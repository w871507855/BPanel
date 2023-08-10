package main

import "github.com/w871507855/BPanel/server/initialize"

func main() {
	r := initialize.Routers()
	r.Run(":8888")
}
