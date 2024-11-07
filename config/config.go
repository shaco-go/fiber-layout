package config

type DSN string

type Config struct {
	Env      string         `json:"env" yaml:"env"`
	Log      Log            `json:"log" yaml:"log"`
	Database map[string]DSN `json:"database" yaml:"database"`
	Redis    Redis          `json:"redis" yaml:"redis"`
}
