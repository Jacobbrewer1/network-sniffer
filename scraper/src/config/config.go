package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

const FlagName = "config"

var Location string

func CreateConfig() error {
	if abs, ok := findFile(Location); ok {
		c, err := os.ReadFile(abs)
		if err != nil {
			return err
		}

		if err = yaml.Unmarshal(c, &Cfg); err != nil {
			return err
		}

		return nil
	}

	return errors.New("no config found")
}

func findFile(path string) (string, bool) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", false
	}

	file, err := os.Open(abs)
	if err != nil {
		return "", false
	}

	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	return abs, true
}
