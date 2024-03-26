package main

import (
	"flag"
	app "openidea-banking/src"
)

func main() {
	var port string
	var prefork bool

	flag.StringVar(&port, "port", "8080", "application port")
	flag.BoolVar(&prefork, "prefork", true, "enable prefork")
	flag.Parse()

	app.Start(port, prefork)
}
