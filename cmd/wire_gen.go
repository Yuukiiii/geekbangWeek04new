// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"fmt"
	"github.com/Yuukiiii/geekbangWeek04/internal/server"
	"net/http"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

// initApp init application.
func initServer(portHelloWorld int, m string) *http.Server {
	httpServer := server.NewHelloWorldServer(portHelloWorld)
	message := server.NewMessage(m)
	greeter := server.NewGreeter(message)
	fmt.Println(greeter.Greet())
	return httpServer
}