package m_account

import (
	"main.go/common/err"
	"main.go/db/m-member"
)

type CreateRequest struct {
	Email      string `json:"email"`
	MerchantID string `json:"merchant_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
}

type CreateResponse struct {
	Status string     `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
}

type UpdateRequest struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UpdateResponse struct {
	Status string     `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
}

type ReadResponse struct {
	Status string                         `json:"status"`
	Error  *err.Error                     `json:"error,omitempty"`
	Data   *m_member.MerchantMemberEntity `json:"data"`
}

type ReadsResponse struct {
	Status string                           `json:"status"`
	Error  *err.Error                       `json:"error,omitempty"`
	Data   *[]m_member.MerchantMemberEntity `json:"data"`
}

type DeleteResponse struct {
	Status string     `json:"status"`
	Error  *err.Error `json:"error,omitempty"`
}
