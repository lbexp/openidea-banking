package main

import (
	"flag"
<<<<<<< HEAD
	"openidea-banking/src"

	_ "openidea-banking/configs"
)

func main() {

	var port string
	var prefork bool
=======
	app "openidea-banking/src"
)

func main() {
	var port string
	var prefork bool

>>>>>>> origin/main
	flag.StringVar(&port, "port", "8080", "application port")
	flag.BoolVar(&prefork, "prefork", true, "enable prefork")
	flag.Parse()

<<<<<<< HEAD
	src.StartApplication(port, prefork)
=======
	app.Start(port, prefork)
>>>>>>> origin/main
}
