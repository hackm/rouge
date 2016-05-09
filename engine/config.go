package engine

type Config struct {
  Port int
  Plugins []string
  Proxy map[string]string
  Theme string
}

func LoadConfig() *Config {
  return &Config {
    Port: 8080,
    Plugins: []string {
      "user@hackm",
      "basic-auth@hackm",
      "paper@hackm",
    },
    Proxy: map[string]string{},
    Theme: "simple",
  }
}
