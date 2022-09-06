package main

import (
	"github.com/boushib/postify/config"
	"github.com/boushib/postify/routes"
)

func main() {
	config.InitDBConnection()
	routes.InitRouter()
}
