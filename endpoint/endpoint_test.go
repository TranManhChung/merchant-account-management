package endpoint

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	m_account "main.go/handler/m-account"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	bodyReq, err:= json.Marshal(m_account.CreateRequest{
		Code: "1",
		Name: "nike",
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

func TestRead(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/v1/merchant/account/read?id=1")
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	sb := string(body)
	assert.Equal(t, `{"status":"success"}`+"\n", sb)
}

func TestUpdate(t *testing.T) {
	bodyReq, err:= json.Marshal(m_account.UpdateRequest{
		Code: "1",
		Name: "nike",
		UserName: "user_name",
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

func TestDelete(t *testing.T) {
	bodyReq, err:= json.Marshal(m_account.DeleteRequest{
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
