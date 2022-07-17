package main

import (
	apiroute "github.com/shengchaohua/red-packet-backend/server_gin/route/api"
)

func main() {
	router := apiroute.InitRouter()

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
