package model

// HTTP holds data necessary for http configuration
type HTTP struct {
	Addr string `yaml:"addr"`
	Port int16  `yaml:"port,omitempty"`
}
