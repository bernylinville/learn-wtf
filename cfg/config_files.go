package cfg

import "os"

const (
	// WtfConfigDirV2 defines the path to the second version of the configuration. Use this.
	WtfConfigDirV2 = "~/.config/wtf"
)

// WtfConfigDir returns the absolute path to the configuration directory
func WtfConfigDir() (string, error) {
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir = WtfConfigDirV2
	} else {
		configDir += "/wtf"
	}

	return configDir, nil
}
