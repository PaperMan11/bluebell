package redis

// redis key
// redis key 注意使用命名空间的方式，方便查询和拆分
const (
	Prefix             = "bluebell:"   // 前缀
	KeyPostTimeZSet    = "post:time"   // zset：贴子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset：帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" // zset：记录用户及投票类型；参数是post_id
)

// 给 redis 加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
