package main

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Dialects map[string]ConfigDialect `yaml:"dialects"`
}

type ConfigDialect struct {
	Default ConfigDialectDefault        `yaml:"default"`
	Mods    map[string]ConfigDialectMod `yaml:"mods"`
}

type ConfigDialectDefault struct {
	Chains map[string]ConfigDialectModChain `yaml:"chains"`
}

type ConfigDialectMod struct {
	Skip   []string                         `yaml:"skip"`
	Funcs  []ConfigDialectModFunc           `yaml:"funcs"`
	Chains map[string]ConfigDialectModChain `yaml:"chains"`
}

type ConfigDialectModFunc struct {
	Prefix          string `yaml:"prefix"`
	ReplacePrefix   string `yaml:"replacePrefix"`
	SecondTypeParam string `yaml:"secondTypeParam"`
}

type ConfigDialectModChain struct {
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

func (c Config) FindDialectSkip(dialect string, mod string, def []string) []string {
	if d, ok := c.Dialects[dialect]; ok {
		if m, ok := d.Mods[mod]; ok {
			return append(def, m.Skip...)
		}
	}
	return def
}

func (c Config) FindDialectFunc(dialect string, mod string, fn string) *ConfigDialectModFunc {
	if d, ok := c.Dialects[dialect]; ok {
		if m, ok := d.Mods[mod]; ok {
			for _, f := range m.Funcs {
				if strings.HasPrefix(fn, f.Prefix) {
					return &f
				}
			}
		}
	}
	return nil
}

func (c Config) FindDialectChain(dialect string, mod string, chain string) *ConfigDialectModChain {
	if d, ok := c.Dialects[dialect]; ok {
		if m, ok := d.Mods[mod]; ok {
			if c, ok := m.Chains[chain]; ok {
				return &c
			}
		}
		if c, ok := d.Default.Chains[chain]; ok {
			return &c
		}
	}
	return nil
}
