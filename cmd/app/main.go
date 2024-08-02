package main

import (
	"github.com/hahaclassic/elysium/config"
	"github.com/hahaclassic/elysium/internal/app"
)

func main() {
	conf := config.MustLoad()

	app.Run(conf)
}
