package configurator

import (
	"calendly/configurator/Api"
	"calendly/configurator/Databases"
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"time"
)

const cfgFileName = ".env.yml"

type Cfg struct {
	Api Api.Gateway      `yaml:"api"`
	DB  Databases.SQLite `yaml:"sqlite"`
}

func (cfg *Cfg) initial(mode string) error {
	path := mode + cfgFileName
	fmt.Println(path)
	configFile, err := os.Open(path)
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(configFile)
	if err != nil {
		return err
	}

	ymlContent, err := io.ReadAll(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(ymlContent, &cfg)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("%s:%s", cfg.Api.Host, cfg.Api.Port))
	fmt.Println(cfg.DB)
	err = cfg.DB.Initial()
	if err != nil {
		return err
	}
	return nil

}
func (cfg *Cfg) Build(mode string) error {
	err := cfg.initial(mode)
	if err != nil {
		return err
	}
	fmt.Println("Build configurator...")
	return nil
}

func (cfg *Cfg) Shutdown(ctx context.Context) error {
	fmt.Println("In configurator cancel start cancelling")
	for {
		select {
		case <-ctx.Done():
			cfg.DB.Close()
			return nil
		default:
			fmt.Println("Wait for 5 second")
			time.Sleep(5 * time.Second)
		}
	}
}
