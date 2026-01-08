package helpers

import "hash/fnv"

func GetShard(key string, numShards int) int {
	// Get the shard by hashing the key
	hash := fnv.New64()
	hash.Write([]byte(key))

	// And then getting the remainder of the hash / the number of shards
	shardId := int(hash.Sum64() % uint64(numShards))

	return shardId
}
