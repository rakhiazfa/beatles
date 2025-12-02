package beatles

import "time"

type ApplicationConfig struct {
	// Default: ""
	Name string `json:"name"`

	// Default: 256 * 1024
	Concurrency int `json:"concurrency"`

	// Default: 4 * 1024 (4 KB)
	ReadBufferSize int `json:"read_buffer_size"`

	// Default: 4 * 1024 (4 KB)
	WriteBufferSize int `json:"write_buffer_size"`

	// Default: 10 Seconds
	ReadTimeout time.Duration `json:"read_timeout"`

	// Default: 10 Seconds)
	WriteTimeout time.Duration `json:"write_timeout"`

	// Default: 60 Seconds
	IdleTimeout time.Duration `json:"idle_timeout"`

	// Default: 4 * 1024 * 1024
	MaxRequestBodySize int `json:"max_request_body_size"`

	// Default: false
	DisableKeepalive bool `json:"disable_keepalive"`

	// Default: false
	ReduceMemoryUsage bool `json:"reduce_memory_usage"`

	// Default: false
	DisablePreParseMultipartForm bool `json:"disable_pre_parse_multipart_form"`

	// Default: false
	StreamRequestBody bool `json:"stream_request_body"`
}

var DefaultApplicationConfig = ApplicationConfig{
	Name:                         "",
	Concurrency:                  256 * 1024,
	ReadBufferSize:               4 * 1024,
	WriteBufferSize:              4 * 1024,
	ReadTimeout:                  10 * time.Second,
	WriteTimeout:                 10 * time.Second,
	IdleTimeout:                  60 * time.Second,
	MaxRequestBodySize:           4 * 1024 * 1024,
	DisableKeepalive:             false,
	ReduceMemoryUsage:            false,
	DisablePreParseMultipartForm: false,
	StreamRequestBody:            false,
}

func mergeApplicationConfig(def ApplicationConfig, config ...ApplicationConfig) ApplicationConfig {
	if len(config) == 0 {
		return def
	}

	conf := config[0]
	merged := def

	if conf.Name != "" {
		merged.Name = conf.Name
	}

	if conf.Concurrency > 0 {
		merged.Concurrency = conf.Concurrency
	}

	if conf.ReadBufferSize > 0 {
		merged.ReadBufferSize = conf.ReadBufferSize
	}
	if conf.WriteBufferSize > 0 {
		merged.WriteBufferSize = conf.WriteBufferSize
	}

	if conf.ReadTimeout > 0 {
		merged.ReadTimeout = conf.ReadTimeout
	}
	if conf.WriteTimeout > 0 {
		merged.WriteTimeout = conf.WriteTimeout
	}
	if conf.IdleTimeout > 0 {
		merged.IdleTimeout = conf.IdleTimeout
	}

	if conf.MaxRequestBodySize > 0 {
		merged.MaxRequestBodySize = conf.MaxRequestBodySize
	}

	merged.DisableKeepalive = conf.DisableKeepalive
	merged.ReduceMemoryUsage = conf.ReduceMemoryUsage
	merged.DisablePreParseMultipartForm = conf.DisablePreParseMultipartForm
	merged.StreamRequestBody = conf.StreamRequestBody

	return merged
}
