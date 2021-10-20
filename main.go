package main

import (
	"gin/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8090")
}
