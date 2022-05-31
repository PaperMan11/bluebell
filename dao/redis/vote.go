package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

// 投一票加 432 分 86400/200 -> 需要200张赞成票可以给你的帖子续一天
/*
direction=1时，有两种情况：
	1、之间没有投过票，现在投赞成票
	2、之前投反对票，现在投赞成票

direction=0时，有两种情况：
	1、之前投赞成票，现在要取消投票
	2、之前投反对票，现在要取消投票

direction=-1时，有两种情况：
	1、之前没投过票，现在投反对票
	2、之前投赞成票，现在改投反对票

投票的限制：
	每个帖子发表之日起一个星期之内允许用户投票
	1、到期之后将 redis 中保存的赞成票及反对票存储到 MySQL中
	2、到期之后删除 KeyPostVotedZSetPF
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePreVote     = 432 // 每一票的分数
)

var (
	ErrorVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postID int64) error {
	pipeline := client.TxPipeline()
	// 帖子时间
	client.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()

	// 帖子分数
	client.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1、判断投票的限制
	// 去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeExpire
	}
	// 2 和 3需要放到一个事务中（pipeline）
	// 2、更新分数
	// 先查之前的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	diff := value - ov // 计算两次投票的差值
	pipeline := client.TxPipeline()
	_, err := client.ZIncrBy(getRedisKey(KeyPostScoreZSet), diff*scorePreVote, postID).Result()
	if err != nil {
		return err
	}
	// 3、记录用户为该帖子投票的数据
	if value == 0 {
		client.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID).Result()
	} else {
		client.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		}).Result()
	}
	_, err = pipeline.Exec()
	return err
}
