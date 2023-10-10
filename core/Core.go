package core

import (
	"calendly/common/helpers"
	"calendly/configurator"
	"calendly/repositories/Admin"
	"calendly/repositories/Customer"
	"calendly/repositories/SystemDomain"
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type FXCore struct {
	cfg *configurator.Cfg

	// http server
	router  *mux.Router
	srv     *http.Server
	handler http.Handler
}

func (fxc *FXCore) Initial(cfg *configurator.Cfg) error {
	fmt.Println("Start core initialisation...")
	fxc.cfg = cfg
	fmt.Println("Http server on: ", fmt.Sprintf("%s:%s", cfg.Api.Host, cfg.Api.Port))
	fxc.router = mux.NewRouter()
	//start and listen to requests
	fxc.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Api.Host, cfg.Api.Port),
		Handler: fxc.router,
	}

	return nil
}

func (fxc *FXCore) Run() {
	// Domains should be added here
	var adm Admin.Entry
	err := adm.Build(fxc)
	helpers.HandleError(err)

	var td Customer.Entry
	err = td.Build(fxc)
	helpers.HandleError(err)

	var sc SystemDomain.Entry
	err = sc.Build(fxc)
	helpers.HandleError(err)

	go func() {
		err := fxc.srv.ListenAndServe()
		if err != nil {
			helpers.HandleError(err)
		}
	}()
}

func (fxc *FXCore) Shutdown(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			err := fxc.srv.Shutdown(ctx)
			if err != nil {
				helpers.HandleError(err)
				return err
			}
			// TODO - write shut down logic here
			return nil
		default:
			fmt.Println("Wait for 5 second")
			time.Sleep(5 * time.Second)
		}
	}
}

func (fxc *FXCore) GetConfigurator() *configurator.Cfg {
	return fxc.cfg
}

func (fxc *FXCore) GetRouter() *mux.Router {
	return fxc.router
}

func (fxc *FXCore) GetDB() *sql.DB {
	return fxc.cfg.DB.Db
}
