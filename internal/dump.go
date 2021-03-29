package internal

import "github.com/sadihakan/dummy-dump/config"

// Dump ...
type Dump interface {
	Check() error
	Export(dump config.Config) error
	Import(dump config.Config) error
}
