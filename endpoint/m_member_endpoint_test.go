package endpoint

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	mMember "main.go/handler/m-member"

	"github.com/stretchr/testify/assert"
)

func TestMMCreate(t *testing.T) {
	bodyReq, err := json.Marshal(mMember.CreateRequest{
		Email:        "email@gmail.com",
		MerchantCode: "mcode",
		Name:         "chung",
		Address:      "Vietnam",
		DoB:          "01-01-2011",
		Phone:        "123",
		Gender:       "male",
	})
	assert.Nil(t, err)

	responseBody := bytes.NewBuffer(bodyReq)

	resp, err := http.Post("http://localhost:8080/v1/merchant/member/create", "application/json", responseBody)
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	sb := string(body)
	assert.Equal(t, `{"status":"success"}`+"\n", sb)
}

func TestMMRead(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/v1/merchant/member/read?email=email@gmail.com")
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	sb := string(body)
	assert.Equal(t, `{"status":"success"}`+"\n", sb)
}

func TestMMUpdate(t *testing.T) {
	bodyReq, err := json.Marshal(mMember.UpdateRequest{
		Email:   "email@gmail.com",
		Name:    "chung",
		Address: "Vietnam",
		DoB:     "01-01-2011",
		Phone:   "123",
		Gender:  "male",
	})
	assert.Nil(t, err)

	responseBody := bytes.NewBuffer(bodyReq)
	resp, err := http.Post("http://localhost:8080/v1/merchant/member/update", "application/json", responseBody)
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	sb := string(body)
	assert.Equal(t, `{"status":"success"}`+"\n", sb)
}

func TestMMDelete(t *testing.T) {
	bodyReq, err := json.Marshal(mMember.DeleteRequest{
		Email:   "email@gmail.com",
	})
	assert.Nil(t, err)

	responseBody := bytes.NewBuffer(bodyReq)
	resp, err := http.Post("http://localhost:8080/v1/merchant/member/delete", "application/json", responseBody)
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	sb := string(body)
	assert.Equal(t, `{"status":"success"}`+"\n", sb)
}
