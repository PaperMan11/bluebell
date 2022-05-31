package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"
)

// 投票功能
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
// VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
