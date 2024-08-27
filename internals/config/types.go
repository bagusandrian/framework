package config

type Server struct {
	HostName string
	HostIP   string
	HTTP     HTTP `yaml:"http"`
	GRPC     GRPC `yaml:"grpc"`
	NSQ      NSQ  `yaml:"nsq"`
	Cron     Cron `yaml:"cron"`
}

// HTTP defines server config for http server
type HTTP struct {
	Address        string `yaml:"address"`
	WriteTimeout   string `yaml:"write_timeout"`
	ReadTimeout    string `yaml:"read_timeout"`
	MaxHeaderBytes int    `yaml:"max_header_bytes"`
}

type GRPC struct {
	Address        string `yaml:"address"`
	MetricsAddress string `yaml:"metrics_address"`
}

type Cron struct {
	MetricsAddress string `yaml:"metrics_address"`
}

type NSQ struct {
	MetricsAddress string `yaml:"metrics_address"`
}

type ElasticSearch struct {
	Address string `yaml:"address"`

	// ES Proxies
	ServiceID    string `yaml:"service_id"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}
