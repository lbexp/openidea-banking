package main

import (
	"flag"
	"openidea-banking/src"

	_ "openidea-banking/configs"
)

func main() {

	var port string
	var prefork bool
	flag.StringVar(&port, "port", "8080", "application port")
	flag.BoolVar(&prefork, "prefork", true, "enable prefork")
	flag.Parse()

	src.StartApplication(port, prefork)
}
