// +build integration_test

package endpoint

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"main.go/common/status"
	mAccount "main.go/handler/m-account"
	mMember "main.go/handler/m-member"
	"net/http"
	"testing"
)

func TestMACreate(t *testing.T) {
	bodyReq, err := json.Marshal(mAccount.CreateRequest{
		Code:     "11112",
		Name:     "nike",
		UserName: "user_name",
		Password: "password",
	})
	assert.Nil(t, err)

	resp, err := http.Post("http://localhost:8080/v1/merchant/account?action=new", "application/json", bytes.NewBuffer(bodyReq))
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	respModel := mAccount.CreateResponse{}
	err = json.Unmarshal(body, &respModel)
	assert.Nil(t, err)
	assert.Equal(t, status.Success, respModel.Status)
}

func TestMARead(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/v1/merchant/account?action=get&id=1652924346853519000", "application/json", bytes.NewBuffer([]byte{}))
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	respModel := mAccount.ReadResponse{}
	err = json.Unmarshal(body, &respModel)
	assert.Nil(t, err)
	assert.Equal(t, status.Success, respModel.Status)
}

func TestMAUpdate(t *testing.T) {
	bodyReq, err := json.Marshal(mAccount.UpdateRequest{
		MerchantID: "1652924346853519000",
		Name:       "nikenike",
		Password:   "password",
	})
	assert.Nil(t, err)

	resp, err := http.Post("http://localhost:8080/v1/merchant/account?action=update", "application/json", bytes.NewBuffer(bodyReq))
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	respModel := mAccount.UpdateResponse{}
	err = json.Unmarshal(body, &respModel)
	assert.Nil(t, err)
	assert.Equal(t, status.Success, respModel.Status)
}

func TestMADelete(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/v1/merchant/account?action=delete&id=1652924346853519000", "application/json", bytes.NewBuffer([]byte{}))
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	respModel := mAccount.DeleteResponse{}
	err = json.Unmarshal(body, &respModel)
	assert.Nil(t, err)
	assert.Equal(t, status.Success, respModel.Status)
}

//------------------------------------------

func TestMMCreate(t *testing.T) {
	bodyReq, err := json.Marshal(mMember.CreateRequest{
		Email:      "email8@gmail.com",
		Name:       "chung",
		Address:    "Vietnam",
		Phone:      "123",
		MerchantID: "1652924313363295000",
	})
	assert.Nil(t, err)

	resp, err := http.Post("http://localhost:8080/v1/merchant/account?source=member&action=new", "application/json", bytes.NewBuffer(bodyReq))
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	respModel := mMember.CreateResponse{}
	err = json.Unmarshal(body, &respModel)
	assert.Nil(t, err)
	assert.Equal(t, status.Success, respModel.Status)
}

func TestMMRead(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/v1/merchant/account?source=member&action=get&email=email8@gmail.com", "application/json", bytes.NewBuffer([]byte{}))
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	respModel := mMember.ReadResponse{}
	err = json.Unmarshal(body, &respModel)
	assert.Nil(t, err)
	assert.Equal(t, status.Success, respModel.Status)
}

func TestMMReads(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/v1/merchant/account?source=member&action=gets&merchant_id=1652924313363295000&offset=2&limit=2", "application/json", bytes.NewBuffer([]byte{}))
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	respModel := mMember.ReadsResponse{}
	err = json.Unmarshal(body, &respModel)
	assert.Nil(t, err)
	assert.Equal(t, status.Success, respModel.Status)
}

func TestMMUpdate(t *testing.T) {
	bodyReq, err := json.Marshal(mMember.UpdateRequest{
		Email:   "email8@gmail.com",
		Name:    "chung",
		Address: "Vietnam1",
		Phone:   "123",
	})
	assert.Nil(t, err)

	resp, err := http.Post("http://localhost:8080/v1/merchant/account?source=member&action=update", "application/json", bytes.NewBuffer(bodyReq))
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	respModel := mMember.UpdateResponse{}
	err = json.Unmarshal(body, &respModel)
	assert.Nil(t, err)
	assert.Equal(t, status.Success, respModel.Status)
}

func TestMMDelete(t *testing.T) {

	resp, err := http.Post("http://localhost:8080/v1/merchant/account?source=member&action=delete&email=email8@gmail.com", "application/json", bytes.NewBuffer([]byte{}))
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	respModel := mMember.DeleteResponse{}
	err = json.Unmarshal(body, &respModel)
	assert.Nil(t, err)
	assert.Equal(t, status.Success, respModel.Status)
}
