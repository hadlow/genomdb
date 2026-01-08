package types

type Config struct {
	Database string  `yaml:"database"`
	Shards   []Shard `yaml:"shards"`
}
