package endpoint

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	mAccount "main.go/handler/m-account"
	mMember "main.go/handler/m-member"

	"github.com/gorilla/mux"
)

func Register(mAccountHandler mAccount.Service, mMemberHandler mMember.Service) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1/merchant/account", Handle(mMemberHandler, mAccountHandler)).Methods(http.MethodPost)
	return router
}

func Handle(mMemberHandler mMember.Service, mAccountHandler mAccount.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		source := r.URL.Query().Get("source")
		action := r.URL.Query().Get("action")

		if source == "member" {
			switch action {
			case "new":
				request := &mMember.CreateRequest{}
				if e := json.NewDecoder(r.Body).Decode(request); e != nil {
					http.Error(w, e.Error(), http.StatusBadRequest)
				} else {
					resp := mMemberHandler.Create(context.Background(), request)
					json.NewEncoder(w).Encode(resp)
				}
			case "get":
				gmail := r.URL.Query().Get("email")
				resp := mMemberHandler.Read(context.Background(), gmail)
				json.NewEncoder(w).Encode(resp)
			case "gets":
				merchantID := r.URL.Query().Get("merchant_id")
				offsetStr := r.URL.Query().Get("offset")
				offset, e := strconv.Atoi(offsetStr)
				if e != nil {
					http.Error(w, e.Error(), http.StatusBadRequest)
				}
				limitStr := r.URL.Query().Get("limit")
				limit, e := strconv.Atoi(limitStr)
				if e != nil {
					http.Error(w, e.Error(), http.StatusBadRequest)
				}
				resp := mMemberHandler.Reads(context.Background(), merchantID, offset, limit)
				json.NewEncoder(w).Encode(resp)
			case "update":
				request := &mMember.UpdateRequest{}
				if e := json.NewDecoder(r.Body).Decode(request); e != nil {
					http.Error(w, e.Error(), http.StatusBadRequest)
				} else {
					resp := mMemberHandler.Update(context.Background(), request)
					json.NewEncoder(w).Encode(resp)
				}
			case "delete":
				email := r.URL.Query().Get("email")
				resp := mMemberHandler.Delete(context.Background(), email)
				json.NewEncoder(w).Encode(resp)
			default:
				http.Error(w, "action not found", http.StatusBadRequest)
			}
		} else {
			switch action {
			case "new":
				request := &mAccount.CreateRequest{}
				if e := json.NewDecoder(r.Body).Decode(request); e != nil {
					http.Error(w, e.Error(), http.StatusBadRequest)
				} else {
					resp := mAccountHandler.Create(context.Background(), request)
					json.NewEncoder(w).Encode(resp)
				}
			case "get":
				id := r.URL.Query().Get("id")
				resp := mAccountHandler.Read(context.Background(), id)
				json.NewEncoder(w).Encode(resp)
			case "update":
				request := &mAccount.UpdateRequest{}
				if e := json.NewDecoder(r.Body).Decode(request); e != nil {
					http.Error(w, e.Error(), http.StatusBadRequest)
				} else {
					resp := mAccountHandler.Update(context.Background(), request)
					json.NewEncoder(w).Encode(resp)
				}
			case "delete":
				merchantID := r.URL.Query().Get("id")
				resp := mAccountHandler.Delete(context.Background(), merchantID)
				json.NewEncoder(w).Encode(resp)
			default:
				http.Error(w, "action not found", http.StatusBadRequest)
			}
		}
	}
}
