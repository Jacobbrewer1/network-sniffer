package config

var Cfg Root

type (
	Root struct {
		Setup *Setup `json:"setup,omitempty" yaml:"setup,omitempty"`
	}

	Setup struct {
		ListeningPort string `json:"listeningPort,omitempty" yaml:"listeningPort,omitempty"`
		CertPath      string `json:"certificatePath,omitempty" yaml:"certPath,omitempty"`
		KeyPath       string `json:"keyPath,omitempty" yaml:"keyPath,omitempty"`
	}
)
