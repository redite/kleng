package config

type (

	// BoltDB hold boltdb configuration
	BoltDB struct {
		BoltDBPath string `json:"db_path"`
	}

	// BleveSearch holt blevesearch index configuration
	BleveSearch struct {
		IndexPath string `json:"index_path"`
	}

	// Config hold application configuration
	Config struct {
		BoltDB      BoltDB      `json:"boltdb"`
		BleveSearch BleveSearch `json:"blevesearch"`
		Port        int         `envconfig:"http_port"`
		Token       string      `envconfig:"github_token"`
		Username    string      `envconfig:"github_username"`
	}

	// IndexMapping ...
	IndexMapping struct {
	}
)
