package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/adammy/go-monorepo-template/tools/create"
)

func main() {
	flag.Parse()
	args := flag.Args()

	cmdType := args[0]

	serverName := args[1]
	if err := create.ValidateName(serverName); err != nil {
		log.Fatal(err)
	}

	svc := create.NewService(
		create.WithName(serverName),
		create.WithFetchedSuiteName(),
		create.WithFetchedModuleName(),
		create.WithFetchedRootURL(),
	)

	templateFuncs := map[string]func() error{
		"api": func() error {
			if err := svc.OpenAPI(); err != nil {
				return fmt.Errorf("create open api failed: %w", err)
			}
			return nil
		},
		"server": func() error {
			if err := svc.Server(); err != nil {
				return fmt.Errorf("create server failed: %w", err)
			}
			return nil
		},
	}

	if fn, ok := templateFuncs[cmdType]; ok {
		if err := fn(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("invalid type")
	}
}
