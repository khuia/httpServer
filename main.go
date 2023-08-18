// 服务端

package main

var c Config

func main() {
	getConfig("./config.json", &c)
	go socks5Start(&c)
	server_advPortForward(&c)
}
