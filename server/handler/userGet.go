package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Params struct {
	Email string `json:"email"`
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Service 3 UserGet")

	defer r.Body.Close()

	var params Params
	err := json.NewDecoder(r.Body).Decode(&params)
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(User{ID: 1, Email: params.Email})
	if nil != err {
		w.WriteHeader(http.StatusOK)
		return
	}
}
