package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Dialects map[string]ConfigDialect `yaml:"dialects"`
}

func (c Config) FindDialectChain(dialect string, chain string) *ConfigDialectChain {
	if d, ok := c.Dialects[dialect]; ok {
		if c, ok := d.Chains[chain]; ok {
			return &c
		}
	}
	return nil
}

type ConfigDialect struct {
	Chains map[string]ConfigDialectChain `yaml:"chains"`
}

type ConfigDialectChain struct {
	Methods []string `yaml:"methods"`
}

func LoadConfig() (Config, error) {
	fn := filepath.Join(getCurrentDir(), "config.yaml")
	f, err := os.Open(fn)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)
	var c Config
	err = dec.Decode(&c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
