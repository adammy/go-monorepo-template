package createcli

import (
	"fmt"

	"github.com/adammy/go-monorepo-template/tools/create"
	"github.com/manifoldco/promptui"
)

func PromptCreateType() (string, error) {
	types := []createType{
		{
			Name: "OpenAPI spec",
			Creates: []string{
				"OpenAPI YML spec file, where you can define the API you are creating",
			},
		},
		{
			Name: "Generated types, client, and server interface",
			Creates: []string{
				"Generated types from the components in the OpenAPI spec",
				"Generated HTTP client from the OpenAPI spec",
				"Generated HTTP server interface from the OpenAPI spec",
			},
		},
		{
			Name: "Server implementation stub",
			Creates: []string{
				"HTTP Chi server that implements the server interface",
				"Doc file to document the package",
				"Main file that starts the implemented server",
				"Initial types file that defines config for the implemented server",
				"Default config YML file",
				"Dockerfile",
			},
		},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | bold }}",
		Active:   `{{ "▸" | cyan }} {{ .Name | cyan }}`,
		Inactive: "  {{ .Name }}",
		Selected: `{{ "✔" | green }} {{ "Type:" | green }} {{ .Name }}`,
		Details: `
{{ "This type creates:" | faint }}
{{ range .Creates }}{{ "•" | faint }} {{ . | faint }}
{{ end }}`,
	}

	prompt := promptui.Select{
		Label:     "What are we creating?",
		Items:     types,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("prompt run failed: %w", err)
	}

	return types[i].Name, nil
}

func PromptServiceName() (string, error) {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "✔ {{ . }}: ",
		Invalid: "✗ {{ . }}: ",
		Success: `{{ "✔" | green }} {{ . | green }}: `,
		// Valid: ,
		// ✗, ✔, ▸, •
	}

	prompt := promptui.Prompt{
		Label:     "Service name",
		Validate:  create.ValidateName,
		Templates: templates,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("prompt run failed: %w", err)
	}

	return result, nil
}
