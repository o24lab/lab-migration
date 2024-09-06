package pkg

func CalculateShardingKey(id int64) int64 {
	return id % 1024
}
