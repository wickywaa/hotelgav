package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {

	// AppConfig holds the application config
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
