package main

import "github.com/vvelikodny/ff-go-test/api/services"

func main() {
	NewApp(services.NewInMemSessionService()).Run()
}
