package config

var host hostConfig

type hostConfig struct {
	Name string
	Port string
}

func (h *hostConfig) check() {
	if h.Port == "" {
		h.Port = ":3001"
	}
}
