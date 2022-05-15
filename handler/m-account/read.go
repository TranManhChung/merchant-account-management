package m_account

import (
	"fmt"
	"net/http"
)

func (h Handler) Read(w http.ResponseWriter, r *http.Request) {
	fmt.Println("you are calling read api")
}