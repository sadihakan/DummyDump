package main

import (
	"context"
	"fmt"
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New()

	if err != nil {
		fmt.Println("DummyDump new error: ", err)
	}

	dd.SetBinaryConfig(config.PostgreSQL, false, true)

	ctx := context.Background()

	binary, version := dd.GetBinary(ctx)
	fmt.Println("Bin: ", binary)
	fmt.Println("Version: ", version)

	dd2, err := dummydump.New(&config.Config{
		Source:         config.PostgreSQL,
		Import:         false,
		Export:         true,
		User:           "hakankosanoglu",
		Password:       "",
		DB:             "hell",
		Host:           "localhost",
		Port:           5432,
		BackupFilePath: "/Users/hakankosanoglu/Desktop",
		BackupName:     "aa.backup",
		BinaryPath:     binary,
	})

	if err != nil {
		fmt.Println("DummyDump error ", err)
	}

	_, err = dd2.CheckPath(ctx).Export(ctx).Run()

	if err != nil {
		fmt.Println("Run error: ", err)
	}

}
