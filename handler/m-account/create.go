package m_account

import (
	"fmt"
	"net/http"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("you are calling create api")
}