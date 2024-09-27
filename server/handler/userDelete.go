package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UserDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Service 3 UserDelete")

	defer r.Body.Close()

	var params Params
	err := json.NewDecoder(r.Body).Decode(&params)
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
