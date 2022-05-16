package main

import (
	"log"

	"main.go/endpoint"
	mAccount "main.go/handler/m-account"
	mMember "main.go/handler/m-member"
	"main.go/infra"
)

func main() {
	mAccountHandler:= mAccount.New()

	app, err := infra.New(endpoint.Register(mAccountHandler, mMember.Handler{}))
	if err != nil {
		log.Fatalln(err)
	}
	if err = app.Start(); err != nil {
		log.Fatalln(err)
	}
}
