package main

//"fmt"

var c Config

func main() {
	getConfig("./config.json", &c)
	httpServer()
	//advPortForward(&c)
	//start(&c)
}
