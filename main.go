package main

import "goarch/router"

func main() {
	api := router.InitGet()

	api.Run(":8000")
}