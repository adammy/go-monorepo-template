package httpapi

import (
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

func NewLogger(cfg LoggerConfig) zerolog.Logger {
	return httplog.NewLogger(cfg.AppName, httplog.Options{
		LogLevel: string(cfg.Level),
		JSON:     cfg.JSON,
		Concise:  false,
		Tags: map[string]string{
			"version": cfg.AppVersion,
			"env":     cfg.EnvName,
		},
	})
}
