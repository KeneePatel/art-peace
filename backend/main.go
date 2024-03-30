package main

import (
	"flag"

	"art-peace-backend/backend"
	"art-peace-backend/config"
	"art-peace-backend/routes"
)

func main() {
	canvasConfigFilename := flag.String("canvas-config", config.DefaultCanvasConfigPath, "Canvas config file")
	databaseConfigFilename := flag.String("database-config", config.DefaultDatabaseConfigPath, "Database config file")
  port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	canvasConfig, err := config.LoadCanvasConfig(*canvasConfigFilename)
	if err != nil {
		panic(err)
	}

	databaseConfig, err := config.LoadDatabaseConfig(*databaseConfigFilename)
	if err != nil {
		panic(err)
	}

  databases := backend.NewDatabases(databaseConfig)
	defer databases.Close()

  routes.InitRoutes()

  backend.ArtPeaceBackend = backend.NewBackend(databases, canvasConfig, *port)
  backend.ArtPeaceBackend.Start()
}
