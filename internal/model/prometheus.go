package model

// Prometheus holds data necessary for prometheus configuration
type Prometheus struct {
	Namespace string `yaml:"namespace"`
	Subsystem string `yaml:"subsystem"`
	Name      string `yaml:"name"`
}
