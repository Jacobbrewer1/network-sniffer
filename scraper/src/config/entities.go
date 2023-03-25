package config

var Cfg Root

type (
	Root struct {
		Setup *Setup `json:"setup,omitempty" yaml:"setup,omitempty"`
	}

	Setup struct {
		ApiEndpoint string `json:"api_endpoint,omitempty" yaml:"api_endpoint,omitempty"`
		Username    string `json:"username,omitempty" yaml:"username,omitempty"`
		Password    string `json:"password,omitempty" yaml:"password,omitempty"`
	}
)
