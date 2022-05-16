package endpoint

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	mAccount "main.go/handler/m-account"

	"github.com/stretchr/testify/assert"
)

func TestMACreate(t *testing.T) {
	bodyReq, err := json.Marshal(mAccount.CreateRequest{
		Code:     "1",
		Name:     "nike",
		UserName: "user_name",
		Password: "password",
	})
	assert.Nil(t, err)

	responseBody := bytes.NewBuffer(bodyReq)

	resp, err := http.Post("http://localhost:8080/v1/merchant/account/create", "application/json", responseBody)
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	sb := string(body)
	assert.Equal(t, `{"status":"success"}`+"\n", sb)
}

func TestMARead(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/v1/merchant/account/read?id=1")
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	sb := string(body)
	assert.Equal(t, `{"status":"success"}`+"\n", sb)
}

func TestMAUpdate(t *testing.T) {
	bodyReq, err := json.Marshal(mAccount.UpdateRequest{
		Code:     "1",
		Name:     "nike",
		Password: "password",
	})
	assert.Nil(t, err)

	responseBody := bytes.NewBuffer(bodyReq)
	resp, err := http.Post("http://localhost:8080/v1/merchant/account/update", "application/json", responseBody)
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	sb := string(body)
	assert.Equal(t, `{"status":"success"}`+"\n", sb)
}

func TestMADelete(t *testing.T) {
	bodyReq, err := json.Marshal(mAccount.DeleteRequest{
		Code: "1",
	})
	assert.Nil(t, err)

	responseBody := bytes.NewBuffer(bodyReq)
	resp, err := http.Post("http://localhost:8080/v1/merchant/account/delete", "application/json", responseBody)
	assert.Nil(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	sb := string(body)
	assert.Equal(t, `{"status":"success"}`+"\n", sb)
}
