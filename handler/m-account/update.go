package m_account

import (
	"fmt"
	"net/http"
)

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("you are calling update api")
}