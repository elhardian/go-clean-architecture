package main

import (
	"fmt"
	"os"
	"time"

	"github.com/elhardian/go-clean-architecture/libs/manager"
	"github.com/elhardian/go-clean-architecture/libs/server"
)

func run() error {
	mgr, err := manager.NewInit()
	if err != nil {
		return err
	}

	tzLocation, err := time.LoadLocation(mgr.GetConfig().AppTz)
	if err != nil {
		return err
	}
	time.Local = tzLocation

	server := server.NewServer(mgr.GetConfig())

	return server.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
