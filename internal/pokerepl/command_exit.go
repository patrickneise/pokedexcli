package pokerepl

import "os"

func CommandExit(cfg *Config, args ...string) error {
	os.Exit(0)
	return nil
}
