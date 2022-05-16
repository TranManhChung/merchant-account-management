package main

import (
	"log"
	"main.go/endpoint"
	m_account "main.go/handler/m-account"
	"main.go/infra"
)

func main() {
	app, err := infra.New(endpoint.Register(m_account.Handler{}))
	if err != nil {
		log.Fatalln(err)
	}
	if err = app.Start(); err!=nil {
		log.Fatalln(err)
	}
}
