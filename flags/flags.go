package flags

import (
	goFlags "github.com/jessevdk/go-flags"
)

// Flags is the container for command line flag data
type Flags struct {
	Config  string `short:"c" long:"config" optional:"yes" description:"Path to config file"`
	Module  string `short:"m" long:"module" optional:"yes" description:"Display info about a specific module, i.e.: 'wtfutil -m=todo'"`
	Profile bool   `short:"p" long:"profile" optional:"yes" description:"Profile application memory usage"`
	Version bool   `short:"v" long:"version" description:"Show version info"`
	// Work-around go-flags misfeatures. If any sub-command is defined
	// then `wtf` (no sub-commands, the common usage), is warned about.
	Opt struct {
		Cmd  string   `positional-arg-name:"command"`
		Args []string `positional-arg-name:"args"`
	} `positional-args:"yes"`
}

var EXTRA = `
Commands:
  save-secret <service>
    service      Service URL or module name of secret.
  Save a secret into the secret store. The secret will be prompted for.
  Requires wtf.secretStore to be configured.  See individual modules for
  information on what service and secret means for their configuration,
  not all modules use secrets.
`

// NewFlags creates an instance of Flags
func NewFlags() *Flags {
	flags := Flags{}
	return &flags
}

// Parse parses the incoming flags
func (flags *Flags) Parse() {
	parser := goFlags.NewParser(flags, goFlags.Default)
}
