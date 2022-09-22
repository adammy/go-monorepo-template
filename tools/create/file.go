package create

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Dir creates a directory from a full path, i.e, "pkg/server" or "cmd".
func Dir(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("os make dir all failed: %w", err)
	}
	return nil
}

// File creates a file from a FileConfig.
func File(cfg *TmplCfg) error {
	var (
		buf bytes.Buffer
	)

	if err := Dir(cfg.Dir); err != nil {
		return fmt.Errorf("create dir failed: %w", err)
	}

	tmpl := template.Must(template.New(cfg.Name).Funcs(getTemplateFuncMap()).ParseFiles(cfg.Tmpl))
	if err := tmpl.Execute(&buf, cfg.Data); err != nil {
		return fmt.Errorf("template execute failed: %w", err)
	}

	file, err := os.Create(path.Join(cfg.Dir, cfg.File))
	if err != nil {
		return fmt.Errorf("os create failed, %w", err)
	}
	defer file.Close()

	if _, err = file.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	}
	return nil
}

// Files creates n number of files from FileConfig.
func Files(cfgs []*TmplCfg) error {
	for _, cfg := range cfgs {
		if err := File(cfg); err != nil {
			return err
		}
	}
	return nil
}

// getLine returns line text from a reader.
func getLine(r io.Reader, line int) (string, error) {
	sc := bufio.NewScanner(r)
	currLine := 1

	for sc.Scan() {
		if currLine == line {
			return sc.Text(), nil
		}
		currLine++
	}
	return "", ErrLineNotFound
}

// getTemplateFuncMap returns a function mapping that can be used
// when executing a template.
func getTemplateFuncMap() template.FuncMap {
	caser := cases.Title(language.English)
	return template.FuncMap{
		"Title": caser.String,
	}
}
