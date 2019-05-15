package main

import (
	"SimpleGameApiPlatform/server/api"
)

func main() {
	server := api.Init()

	server.Run(":9999")
}
