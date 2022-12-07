package config

// The server configuration structure. This struct maps the config.yaml
// structure, so that it can easily be parsed and read when necessary.
// It is recommended to change this struct according to the changes that
// are made to the config.yaml, to keep it consistent.
type ServerConfig struct {
	Orchestrator struct {
		Server struct {
			Banner       string `yaml:"banner"`
			Version      string `yaml:"version"`
			AppName      string `yaml:"application-name"`
			Port         string `yaml:"port"`
			MaxQueueSize int    `yaml:"max-queue-size"`
			MaxWorkers   int    `yaml:"max-workers"`
		} `yaml:"server"`
		ActiveProfile string `yaml:"active-profile"`
	} `yaml:"orchestrator"`
}

type ProfileConfig struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}
