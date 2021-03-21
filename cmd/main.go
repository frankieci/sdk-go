package main

import "gitlab.com/sdk-go/server"

func main() {
	svr := server.NewSDKServer()
	svr.Run()
}
