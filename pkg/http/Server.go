package http

import (
	"calendly/configurator"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	cfg *configurator.Cfg
}

func (srv *Server) Initial(cfg *configurator.Cfg) {
	srv.cfg = cfg
	fmt.Println("Http server initialized")
}

func (srv *Server) Run(cfg *configurator.Cfg) error {
	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})
	http.Handle("/", r)
	fmt.Println(srv.cfg.Api.GetUrl())
	log.Fatal(http.ListenAndServe(srv.cfg.Api.GetUrl(), nil))
	return nil
}
