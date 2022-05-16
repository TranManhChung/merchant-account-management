package endpoint

import (
	"net/http"

	mAccount "main.go/handler/m-account"
	mMember "main.go/handler/m-member"

	"github.com/gorilla/mux"
)

func Register(mAccountHandler mAccount.IHandler, mMemberHandler mMember.IHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/v1/merchant/account/create", mAccountHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/v1/merchant/account/read", mAccountHandler.Read).Methods(http.MethodGet)
	router.HandleFunc("/v1/merchant/account/update", mAccountHandler.Update).Methods(http.MethodPost)
	router.HandleFunc("/v1/merchant/account/delete", mAccountHandler.Delete).Methods(http.MethodPost)

	router.HandleFunc("/v1/merchant/member/create", mMemberHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/v1/merchant/member/read", mMemberHandler.Read).Methods(http.MethodGet)
	router.HandleFunc("/v1/merchant/member/update", mMemberHandler.Update).Methods(http.MethodPost)
	router.HandleFunc("/v1/merchant/member/delete", mMemberHandler.Delete).Methods(http.MethodPost)

	return router
}
