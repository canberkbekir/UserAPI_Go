package main

func main() {
	router := createServer("8080")
	router.Run()
}
