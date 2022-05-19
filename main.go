package main

import (
	"log"
	"main.go/db"
	mAccountDB "main.go/db/m-account"
	mMemberDB "main.go/db/m-member"
	"main.go/endpoint"
	mAccount "main.go/handler/m-account"
	mMember "main.go/handler/m-member"
	"main.go/infra"
)

func main() {
	bunDB, err := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	if err != nil {
		log.Fatalln(err)
	}
	accountRepo, err := mAccountDB.NewMAccountRepoRepo(bunDB, map[string]int{mAccountDB.MerchantCode: 10, mAccountDB.MerchantName: 10, mAccountDB.Username: 10})
	if err != nil {
		log.Fatalln(err)
	}
	memberRepo, err := mMemberDB.NewMMemberRepoRepo(bunDB, map[string]int{mMemberDB.Email: 20, mMemberDB.Name: 20, mMemberDB.Address: 20, mMemberDB.Phone: 20, mMemberDB.MerchantID: 20})
	if err != nil {
		log.Fatalln(err)
	}
	mAccountHandler := mAccount.New(accountRepo,memberRepo)
	mMemberHandler := mMember.New(memberRepo,accountRepo)
	app, err := infra.New(endpoint.Register(mAccountHandler, mMemberHandler))
	if err != nil {
		log.Fatalln(err)
	}
	if err = app.Start(); err != nil {
		log.Fatalln(err)
	}
}
