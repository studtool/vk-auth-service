package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

func main() {
	_ = &oauth2.Config{
		Endpoint: vk.Endpoint,
	}
}
