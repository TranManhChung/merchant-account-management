package m_account

import (
	"fmt"
	"net/http"
)

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("you are calling delete api")
}
