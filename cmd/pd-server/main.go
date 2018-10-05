package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hawkingrei/wfs/pkg/logutil"
	"github.com/hawkingrei/wfs/server"
)

func main() {
	cfg := server.NewConfig()
	err := cfg.Parse(os.Args[1:])

	if cfg.Version {
		server.PrintPDInfo()
		os.Exit(0)
	}

	err = logutil.InitLogger(&cfg.Log)
	if err != nil {
		log.Fatalf("initialize logger error: %s\n", fmt.Sprintf("%+v", err))
	}
}
