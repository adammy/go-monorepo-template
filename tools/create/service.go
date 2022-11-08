package create

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

// Service provides functionality for creating generated OpenAPI specs,
// types, servers, clients, etc.
type Service struct {
	// Name of the service/server, e.g. "user".
	Name string

	// ModuleName as defined in the go.mod file, e.g., "github.com/adammy/go-monorepo-template".
	ModuleName string

	// SuiteName defines the prefix for the service/server apps, e.g., "adammy".
	SuiteName string

	// RootURL defines the RootURL the service/server will be hosted at, e.g., "adammy.com".
	RootURL string

	// PkgDir defines the pkg directory that will contain files for the
	// service/server, e.g., "pkg/user".
	PkgDir string

	// CmdDir defines the cmd directory that will contain files for the
	// service/server, e.g., "cmd/user".
	CmdDir string

	// TmplData contains the dynamic data to be supplied to executed templates.
	TmplData map[string]string
}

// NewService constructs a Service.
func NewService(options ...func(*Service)) *Service {
	svc := &Service{
		TmplData: make(map[string]string),
	}

	for _, o := range options {
		o(svc)
	}

	if svc.Name != "" {
		svc.PkgDir = fmt.Sprintf("pkg/%s", svc.Name)
		svc.CmdDir = fmt.Sprintf("cmd/%s", svc.Name)
		svc.TmplData["name"] = svc.Name
	}

	if svc.SuiteName != "" {
		svc.TmplData["suiteName"] = svc.SuiteName
	}

	if svc.ModuleName != "" {
		svc.TmplData["moduleName"] = svc.ModuleName
	}

	if svc.RootURL != "" {
		svc.TmplData["rootURL"] = svc.RootURL
	}

	return svc
}

// WithName is a functional option for Service.Name.
func WithName(name string) func(*Service) {
	return func(svc *Service) {
		svc.Name = name
	}
}

// WithSuiteName is a functional option for Service.SuiteName.
func WithSuiteName(name string) func(*Service) {
	return func(svc *Service) {
		svc.SuiteName = name
	}
}

// WithFetchedSuiteName is a functional option for Service.SuiteName from .create.yml.
func WithFetchedSuiteName() func(*Service) {
	return func(svc *Service) {
		cfg, err := fetchCfg()
		if err == nil {
			svc.SuiteName = cfg.SuiteName
		}
	}
}

// WithModuleName is a functional option for Service.ModuleName.
func WithModuleName(name string) func(*Service) {
	return func(svc *Service) {
		svc.ModuleName = name
	}
}

// WithFetchedModuleName is a functional option that fetches Service.ModuleName from go.mod.
func WithFetchedModuleName() func(*Service) {
	return func(svc *Service) {
		modName, err := fetchModuleName()
		if err == nil {
			svc.ModuleName = modName
		}
	}
}

// WithRootURL is a functional option for Service.RootURL.
func WithRootURL(url string) func(*Service) {
	return func(svc *Service) {
		svc.RootURL = url
	}
}

// WithRootURL is a functional option for Service.RootURL from .create.yml.
func WithFetchedRootURL() func(*Service) {
	return func(svc *Service) {
		cfg, err := fetchCfg()
		if err == nil {
			svc.RootURL = cfg.RootURL
		}
	}
}

// OpenAPI generates an OpenAPI spec file.
func (svc *Service) OpenAPI() error {
	cfg := &TmplCfg{
		Name: "openapi.yml.tmpl",
		Tmpl: "tools/create/templates/openapi.yml.tmpl",
		Dir:  "api/openapi",
		File: fmt.Sprintf("%s.yml", svc.Name),
		Data: svc.TmplData,
	}

	if err := File(cfg); err != nil {
		return err
	}

	return nil
}

// Server generates types, a server, and a client from an OpenAPI spec.
func (svc *Service) Server() error {
	cfgs := []*TmplCfg{
		{
			Name: "generate.go.tmpl",
			Tmpl: "tools/create/templates/generate.go.tmpl",
			Dir:  svc.PkgDir,
			File: "generate.go",
			Data: svc.TmplData,
		},
	}

	if err := Files(cfgs); err != nil {
		return err
	}

	cmd := exec.Command("go", "generate")
	cmd.Dir = fmt.Sprintf("pkg/%s", svc.Name)
	if _, err := cmd.Output(); err != nil {
		return fmt.Errorf("cmd output command failed: %w", err)
	}

	return nil
}

// ServerStub generates an initial implementation for a server from an OpenAPI spec.
func (svc *Service) ServerStub() error {
	cfgs := []*TmplCfg{
		{
			Name: "doc.go.tmpl",
			Tmpl: "tools/create/templates/doc.go.tmpl",
			Dir:  svc.PkgDir,
			File: "doc.go",
			Data: svc.TmplData,
		},
		{
			Name: "main.go.tmpl",
			Tmpl: "tools/create/templates/main.go.tmpl",
			Dir:  svc.CmdDir,
			File: "main.go",
			Data: svc.TmplData,
		},
		{
			Name: "generate.stub.go.tmpl",
			Tmpl: "tools/create/templates/generate.stub.go.tmpl",
			Dir:  svc.PkgDir,
			File: "generate.stub.go",
			Data: svc.TmplData,
		},
		{
			Name: "types.go.tmpl",
			Tmpl: "tools/create/templates/types.go.tmpl",
			Dir:  svc.PkgDir,
			File: "types.go",
			Data: svc.TmplData,
		},
		{
			Name: "dockerfile.tmpl",
			Tmpl: "tools/create/templates/dockerfile.tmpl",
			Dir:  svc.CmdDir,
			File: "Dockerfile",
			Data: svc.TmplData,
		},
		{
			Name: "config.yml.tmpl",
			Tmpl: "tools/create/templates/config.yml.tmpl",
			Dir:  "configs",
			File: fmt.Sprintf("%s.yml", svc.Name),
			Data: svc.TmplData,
		},
	}

	if err := Files(cfgs); err != nil {
		return err
	}

	cmd := exec.Command("go", "generate", "-tags", "stub")
	cmd.Dir = fmt.Sprintf("pkg/%s", svc.Name)
	if _, err := cmd.Output(); err != nil {
		return fmt.Errorf("cmd output command failed: %w", err)
	}

	return nil
}

// Delete will delete all associated files with a service.
func (svc *Service) Delete() error {
	cmd := exec.Command("rm", fmt.Sprintf("api/openapi/%s.yml", svc.Name))
	if _, err := cmd.Output(); err != nil {
		return fmt.Errorf("cmd output command failed: %w", err)
	}

	return nil
}

// fetchModuleName gets the root module name from the go.mod file, i.e.,
// "github.com/adammy/go-monorepo-template".
func fetchModuleName() (string, error) {
	r, err := os.Open("go.mod")
	if err != nil {
		return "", fmt.Errorf("os open failed: %w", err)
	}

	line, err := getLine(r, 1)
	if err != nil {
		return "", fmt.Errorf("get line failed: %w", err)
	}

	split := strings.Split(line, " ")
	if len(split) <= 1 {
		return "", ErrUnableToParseModFile
	}

	return split[1], nil
}

// fetchCfg gets the root URL for the project's servers.
func fetchCfg() (*Cfg, error) {
	cfg := &Cfg{}

	yml, err := os.ReadFile(".create.yml")
	if err != nil {
		return nil, fmt.Errorf("os read file failed: %w", err)
	}

	if err := yaml.Unmarshal(yml, cfg); err != nil {
		return nil, fmt.Errorf("yaml unmarshal failed: %w", err)
	}

	return cfg, nil
}
