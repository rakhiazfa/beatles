package beatles

type ListenConfig struct {
	// Default: NetworkTCP4
	Network string `json:"network"`

	// Default: ""
	CertFile string `json:"cert_file"`

	// Default: ""
	KeyFile string `json:"key_file"`
}

var DefaultListenConfig = ListenConfig{
	Network:  NetworkTCP4,
	CertFile: "",
	KeyFile:  "",
}

func mergeListenConfig(def ListenConfig, config ...ListenConfig) ListenConfig {
	if len(config) == 0 {
		return def
	}

	conf := config[0]
	merged := def

	if conf.Network != "" {
		merged.Network = conf.Network
	}

	if conf.CertFile != "" {
		merged.CertFile = conf.CertFile
	}
	if conf.KeyFile != "" {
		merged.KeyFile = conf.KeyFile
	}

	return merged
}
