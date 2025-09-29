package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req RequestLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")

		return
	}

	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	accessToken, err := util.CreateJWT(h.cnf.JwtSecretkey, util.Payload{ // jwt secret key from env variable
		Sub:         strconv.Itoa(usr.ID),
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusCreated, accessToken)
}
