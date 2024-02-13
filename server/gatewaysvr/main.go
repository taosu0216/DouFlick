package main

import "gatewaysvr/routes"

func main() {
	r := routes.RouteInit()
	if err := r.Run(); err != nil {
		panic(err)
	}
}
