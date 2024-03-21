package main

import (
	"fmt"
	"os"
	"time"

	"github.com/elhardian/go-clean-architecture/libs/manager"
	"github.com/elhardian/go-clean-architecture/libs/server"

	customerRoute "github.com/elhardian/go-clean-architecture/api/customer/delivery"
)
 
func run() error {
	mgr, err := manager.NewInit()
	if err != nil {
		return err
	}
	// app config
	tzLocation, err := time.LoadLocation(mgr.GetConfig().AppTz)
	if err != nil {
		return err
	}
	time.Local = tzLocation
	// server config
	server := server.NewServer(mgr.GetConfig())

	// start routes
	customerRoute.NewRoutes(server.Router, mgr)
	// end routes

	server.RegisterRouter(server.Router)

	return server.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
