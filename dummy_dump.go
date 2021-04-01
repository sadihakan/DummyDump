package dummy_dump

import (
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/errors"
	"github.com/sadihakan/dummy-dump/internal"
	"github.com/sadihakan/dummy-dump/util"
)

// DummyDump ..
type DummyDump struct {
	c     *config.Config
	dump  internal.Dump
	Error error
}

// New ..
func New(cfg ...*config.Config) (*DummyDump, error) {
	dd := new(DummyDump)

	if len(cfg) > 0 {
		dd.c = cfg[0]

		if err := dd.configParser(); err != nil {
			return nil, err
		}

	} else {
		dd.c = new(config.Config)
	}

	return dd, nil
}

// SetUser ..
func (dd *DummyDump) SetUser(username, password string) {
	dd.c.User = username
	dd.c.Password = password
}

// SetBinaryPath ..
func (dd *DummyDump) SetBinaryPath(binaryPath string) {
	dd.c.BinaryPath = binaryPath
}

func (dd *DummyDump) configParser() (err error) {
	switch dd.c.Source {
	case config.PostgreSQL:
		if err = dd.c.CheckConfigPostgreSQL(); err != nil {
			return err
		}
		dd.dump = internal.Postgres{}
		break
	case config.MySQL:
		if err = dd.c.CheckConfigMySQL(); err != nil {
			return err
		}
		dd.dump = internal.MySQL{}
		break
	case config.MSSQL:
		if err = dd.c.CheckConfigMsSQL(); err != nil {
			return err
		}
		dd.dump = internal.MSSQL{}
	default:
		err = errors.New("not implemented")
	}

	return err
}

func (dd *DummyDump) Import() *DummyDump {
	dumpConfig := dd.c

	if !util.PathExists(dumpConfig.Path) {
		dd.Error = errors.New(errors.ConfigPathNotExist)
	}

	err := dd.dump.Import(*dumpConfig)

	if err != nil {
		dd.Error = err
	}

	return dd
}

func (dd *DummyDump) Export() *DummyDump {
	dumpConfig := dd.c
	err := dd.dump.Export(*dumpConfig)

	if err != nil {
		dd.Error = err
	}

	return dd
}

func (dd *DummyDump) Check() *DummyDump {
	dd.Error = dd.dump.Check()

	return dd
}

func (dd *DummyDump) Run() (*DummyDump, error) {
	if dd.Error != nil {
		return dd, dd.Error
	}

	return dd, nil
}

func (dd *DummyDump) GetBinary() (binaryPath string, version string) {
	dumpConfig := dd.c
	binaryPath, err := internal.CheckBinary("", dumpConfig.Source, dumpConfig.Import, dumpConfig.Export)
	version, err = internal.CheckVersion(binaryPath, dumpConfig.Source)

	if err != nil {
		dd.Error = err
	}

	return binaryPath, version
}
