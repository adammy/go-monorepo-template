package main

import (
	"fmt"
	"log"

	"github.com/adammy/go-monorepo-template/tools/create"
	"github.com/adammy/go-monorepo-template/tools/createcli"
)

func main() {
	createType, err := createcli.PromptCreateType()
	if err != nil {
		return
	}

	name, err := createcli.PromptServiceName()
	if err != nil {
		return
	}

	svc := create.NewService(
		create.WithName(name),
		create.WithFetchedSuiteName(),
		create.WithFetchedModuleName(),
		create.WithFetchedRootURL(),
	)

	templateFuncs := map[string]func() error{
		"OpenAPI spec": func() error {
			if err := svc.OpenAPI(); err != nil {
				return fmt.Errorf("create open api failed: %w", err)
			}
			return nil
		},
		"Generated types, client, and server interface": func() error {
			if err := svc.Server(); err != nil {
				return fmt.Errorf("create server failed: %w", err)
			}
			return nil
		},
		"Server implementation stub": func() error {
			if err := svc.ServerStub(); err != nil {
				return fmt.Errorf("create server stub failed: %w", err)
			}
			return nil
		},
	}

	if fn, ok := templateFuncs[createType]; ok {
		if err := fn(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("invalid type")
	}
}
