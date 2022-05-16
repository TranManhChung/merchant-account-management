package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	mAccount "main.go/handler/m-account"
	mMember "main.go/handler/m-member"

	"github.com/gorilla/mux"
)

func Register(mAccountHandler mAccount.Service, mMemberHandler mMember.Service) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/v1/merchant/account/create", CreateMAccount(mAccountHandler)).Methods(http.MethodPost)
	router.HandleFunc("/v1/merchant/account/read", ReadMAccount(mAccountHandler)).Methods(http.MethodGet)
	router.HandleFunc("/v1/merchant/account/update", UpdateMAccount(mAccountHandler)).Methods(http.MethodPost)
	router.HandleFunc("/v1/merchant/account/delete", DeleteMAccount(mAccountHandler)).Methods(http.MethodPost)

	router.HandleFunc("/v1/merchant/member/create", CreateMMember(mMemberHandler)).Methods(http.MethodPost)
	router.HandleFunc("/v1/merchant/member/read", ReadMMember(mMemberHandler)).Methods(http.MethodGet)
	router.HandleFunc("/v1/merchant/member/update", UpdateMMember(mMemberHandler)).Methods(http.MethodPost)
	router.HandleFunc("/v1/merchant/member/delete", DeleteMMember(mMemberHandler)).Methods(http.MethodPost)

	return router
}

func CreateMAccount(mAccountHandler mAccount.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request *mAccount.CreateRequest
		if e := json.NewDecoder(r.Body).Decode(request); e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		} else {
			resp := mAccountHandler.Create(context.Background(), request)
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func ReadMAccount(mAccountHandler mAccount.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		resp := mAccountHandler.Read(context.Background(), code)
		json.NewEncoder(w).Encode(resp)
	}
}

func UpdateMAccount(mAccountHandler mAccount.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request *mAccount.UpdateRequest
		if e := json.NewDecoder(r.Body).Decode(request); e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		} else {
			resp := mAccountHandler.Update(context.Background(), request)
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func DeleteMAccount(mAccountHandler mAccount.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request *mAccount.DeleteRequest
		if e := json.NewDecoder(r.Body).Decode(request); e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		} else {
			resp := mAccountHandler.Delete(context.Background(), request)
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func CreateMMember(mMemberHandler mMember.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request *mMember.CreateRequest
		if e := json.NewDecoder(r.Body).Decode(request); e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		} else {
			resp := mMemberHandler.Create(context.Background(), request)
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func ReadMMember(mMemberHandler mMember.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		gmail := r.URL.Query().Get("gmail")
		resp := mMemberHandler.Read(context.Background(), gmail)
		json.NewEncoder(w).Encode(resp)
	}
}

func UpdateMMember(mMemberHandler mMember.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request *mMember.UpdateRequest
		if e := json.NewDecoder(r.Body).Decode(request); e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		} else {
			resp := mMemberHandler.Update(context.Background(), request)
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func DeleteMMember(mMemberHandler mMember.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request *mMember.DeleteRequest
		if e := json.NewDecoder(r.Body).Decode(request); e != nil {
			http.Error(w, e.Error(), http.StatusBadRequest)
		} else {
			resp := mMemberHandler.Delete(context.Background(), request)
			json.NewEncoder(w).Encode(resp)
		}
	}
}
