package main

import (
	"github.com/fingo-martpedia/fingo-wallet/cmd"
	"github.com/fingo-martpedia/fingo-wallet/helpers"
)

func main() {
	helpers.SetupLogger()

	helpers.SetupDatabase()

	// go cmd.ServeGRPC()

	cmd.ServeHTTP()
}
