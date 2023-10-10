package main

import (
	"calendly/common/helpers"
	"calendly/configurator"
	Core "calendly/core"
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var mode string

	flag.StringVar(&mode, "mode", "default", "Config file that should be initialized")
	flag.Parse()

	var (
		wg sync.WaitGroup
	)

	ctxCancel, fnCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer fnCancel()

	// Configurator initialisation
	var cfg configurator.Cfg
	err := cfg.Build(mode)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Run http server if used
	//var httpServer http.Server
	//if cfg.Api.IsUsed {
	//	httpServer.Initial(&cfg)
	//	// TODO - here should be injection of existing controllers, then run server
	//	err = httpServer.Run(&cfg)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//}

	// Initial highway
	//var ghw highway.FXHighway
	//err = ghw.Initial(&cfg)
	//if err != nil {
	//	helpers.HandleError(err)
	//	return
	//}

	//Initial core
	var core Core.FXCore
	err = core.Initial(&cfg)
	if err != nil {
		helpers.HandleError(err)
		return
	}

	wg.Add(1)
	go func() {
		core.Run()
		wg.Done()
	}()

	wg.Wait()

	// Grace shut down
	<-ctxCancel.Done()
	fmt.Println("Going to grace shut down")
	err = core.Shutdown(ctxCancel)
	helpers.HandleError(err)

	//err = cfg.Shutdown(ctxCancel)
	//helpers.HandleError(err)

}
