package main

import (
	"fmt"
	"globaldevtools.bbva.com/bbva_global_secaas_swat_components/go-openapi-bp"
	"globaldevtools.bbva.com/bbva_global_secaas_swat_components/go-utils/logger"
	holaMundo "holaMundo/routes"
	"os"
)

func main() {
	instanceBuilder := func(logger2 logger.Logger, opt interface{}) interface{} {
		return holaMundo.New(opt.(holaMundo.Config))
	}

	specificConfig := struct {
		ConfigFile string           `flag:"configFile;Configuration file path"`
		Config     holaMundo.Config `configFile:"ConfigFile"`
	}{
		ConfigFile: "./config/config.json",
	}

	if server, err := openapi.GetServerBootStrap(&specificConfig, instanceBuilder); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error initializing server: %v", err)
	} else {
		server.Start()
	}

	os.Exit(0)
}
