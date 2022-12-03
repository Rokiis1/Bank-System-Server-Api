package main

func main() {
	server := newAPIServer(":4000")
	server.Run()
}