package main

import (
	"github.com/boushib/go-blog/config"
	"github.com/boushib/go-blog/routes"
)

func main() {
	config.InitDBConnection()
	routes.InitRouter()
}
