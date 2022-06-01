package redis

import (
	"bluebell/models"

	"github.com/go-redis/redis"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 确定查询的索引起始点
	start := (p.Offset - 1) * p.Limit
	end := start + p.Limit - 1
	// ZRevRange查询 按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

// GetPostVoteData 根据ids查询每篇帖子的赞成票数
func GetPostVoteData(ids []string) (data []int64, err error) {
	// data = make([]int64, 0, len(ids))
	// for _, id := range ids {
	// 	key := getRedisKey(KeyPostVotedZSetPF + id)
	// 	// 查找key中分数是1的元素的数量->统计每篇帖子的赞成票的数量
	// 	v := client.ZCount(key, "1", "1").Val()
	// 	data = append(data, v)
	// }
	// 使用 pipeline 一次发送多命令，减少 RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}
