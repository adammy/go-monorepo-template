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
		createcli.OpenAPIPrompt: func() error {
			if err := svc.OpenAPI(); err != nil {
				return fmt.Errorf("create open api failed: %w", err)
			}
			return nil
		},
		createcli.ServerPrompt: func() error {
			if err := svc.Server(); err != nil {
				return fmt.Errorf("create server failed: %w", err)
			}
			return nil
		},
		createcli.ServerStubPrompt: func() error {
			if err := svc.ServerStub(); err != nil {
				return fmt.Errorf("create server stub failed: %w", err)
			}
			return nil
		},
		createcli.DeletePrompt: func() error {
			if err := svc.Delete(); err != nil {
				return fmt.Errorf("delete failed: %w", err)
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
