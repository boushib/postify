package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boushib/postify/utils"
)

type LoginReqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSuccessRes struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqBody LoginReqBody
	json.NewDecoder(r.Body).Decode(&reqBody)

	if reqBody.Username == "admin" && reqBody.Password == "admin" {
		token, err := utils.GenerateToken()
		if err != nil {
			fmt.Println("Error generating token" + err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		res := LoginSuccessRes{Token: token}
		json.NewEncoder(w).Encode(res)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid credentials")
	}
}
