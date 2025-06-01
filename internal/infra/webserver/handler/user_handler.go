package handler

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/go-chi/jwtauth"
	"github.com/thenopholo/back_commerce/internal/dto"
	"github.com/thenopholo/back_commerce/internal/entity"
	"github.com/thenopholo/back_commerce/internal/infra/database"
)

type UserHandler struct {
	UserDB        database.UserRepository
	JWT           *jwtauth.JWTAuth
	JWTExpiriesIn int
}

func NewUserHandler(db database.UserRepository, JWT *jwtauth.JWTAuth, JWTExpiriesIn int) *UserHandler {
	return &UserHandler{
		UserDB: db,
		JWT: JWT,
		JWTExpiriesIn: JWTExpiriesIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindUserByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_,tokenString, _ :=h.JWT.Encode(map[string]interface{} {
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JWTExpiriesIn)).Unix(),

	})

	accessToken := struct {
		AccessToken string`json:"access_token"`
	}{
		AccessToken: tokenString,
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
