package main

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the template cache as a global variable (app wide use)
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Infolog       *log.Logger
	InProduction  bool //this variable was made to handle security/inproduction  status
	Session       *scs.SessionManager
}
