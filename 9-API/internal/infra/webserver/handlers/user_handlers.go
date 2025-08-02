package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/sancheschris/goexpert/9-APIS/internal/dto"
	"github.com/sancheschris/goexpert/9-APIS/internal/entity"
	"github.com/sancheschris/goexpert/9-APIS/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
	Jwt *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// accessToken := struct {
	// 	AccessToken string `json:"access_token"`
	// } {
	// 	AccessToken: token,
	// }
	// _, tokenString, _ := h.Jwt.Encode(jwtauth.Claims{"user_id": u.ID})
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewDecoder(w).Encode(accessToken)
	// token := jwauth.New("HS256", []byte("secret"), nil)
	w.Write([]byte(tokenString))
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	w.WriteHeader(http.StatusCreated)
}