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
	UserDB database.UserRepository
}

func NewUserHandler(db database.UserRepository) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// GetJWT godoc
// @Summary Get a user JWT token
// @Description Generate a JWT token for authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.GetJWTInput true "user credentials"
// @Success 200 {object} map[string]string
// @Failure 400
// @Failure 401
// @Router /users/login [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JWTExpiresIn").(int)
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

	_, tokenString, _ := jwt.Encode(map[string]any{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)

}

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /users [post]
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
