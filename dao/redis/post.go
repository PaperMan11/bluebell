package redis

import "bluebell/models"

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
