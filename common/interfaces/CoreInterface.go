package interfaces

import (
	"calendly/configurator"
	"database/sql"
	"github.com/gorilla/mux"
)

type ICore interface {
	GetConfigurator() *configurator.Cfg
	GetRouter() *mux.Router
	GetDB() *sql.DB
}
