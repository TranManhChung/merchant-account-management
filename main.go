package main

import (
	"log"

	"main.go/endpoint"
	mAccount "main.go/handler/m-account"
	mMember "main.go/handler/m-member"
	"main.go/infra"
)

func main() {
	app, err := infra.New(endpoint.Register(mAccount.Handler{}, mMember.Handler{}))
	if err != nil {
		log.Fatalln(err)
	}
	if err = app.Start(); err != nil {
		log.Fatalln(err)
	}
}
