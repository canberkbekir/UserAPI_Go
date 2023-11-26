package main

import (
	"os"

	"github.com/canberkbekir/UserAPI_Go/initializers"
)

func main() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()

	router := createServer(os.Getenv("PORT"))
	router.Run()
}
