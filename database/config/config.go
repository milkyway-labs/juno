package config

type Config struct {
	URL           string `yaml:"url"`
	PartitionSize int64  `yaml:"partition_size"`
}

func NewDatabaseConfig(url string, partitionSize int64) Config {
	return Config{
		URL:           url,
		PartitionSize: partitionSize,
	}
}

func (c Config) WithURL(url string) Config {
	c.URL = url
	return c
}

func (c Config) WithPartitionSize(partitionSize int64) Config {
	c.PartitionSize = partitionSize
	return c
}

func (c Config) GetPartitionSize() int64 {
	if c.PartitionSize > 0 {
		return c.PartitionSize
	}
	return 100_0000
}

// DefaultDatabaseConfig returns the default instance of Config
func DefaultDatabaseConfig() Config {
	return NewDatabaseConfig(
		"postgresql://user:password@localhost:5432/database-name?sslmode=disable&search_path=public",
		100000,
	)
}
