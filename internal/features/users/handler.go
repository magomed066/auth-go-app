package users

import (
	"errors"
	"fmt"
	"net/http"

	"charm.land/log/v2"
	"github.com/magomed066/auth-go-app/internal/helpers/jwt"
	"github.com/magomed066/auth-go-app/internal/request"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// ? Methods
func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	var user CreateUserParams

	err := request.Read(r, &user)
	if err != nil {
		log.Error(err)
		request.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := request.Validate(&user, FieldErrorMessages); err != nil {
		request.Error(w, http.StatusBadRequest, err)
		return
	}

	if len([]byte(user.Password)) > 72 {
		request.Error(w, http.StatusBadRequest,
			errors.New("Your password is too long. Please choose a shorter password."))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Error("Could not hash the password", "err", err)
		request.Error(w, http.StatusInternalServerError, err)
		return
	}

	newUser := NewUser{
		CreateUserParams: user,
		PasswordHash:     hashedPassword,
	}

	createdUser, err := h.service.Register(r.Context(), newUser)
	if err != nil {
		log.Error(err)
		if errors.Is(err, ErrUserAlreadyExists) {
			request.Error(w, http.StatusConflict, fmt.Errorf("User with email %s already exists.", user.Email))
			return
		}
		request.Error(w, http.StatusInternalServerError, err)
		return
	}

	token, err := jwt.CreateToken(createdUser.ID)
	if err != nil {
		log.Error("Could not create token", "err", err)
		request.Error(w, http.StatusInternalServerError, err)
		return
	}

	request.Success(w, http.StatusCreated, map[string]any{
		"user":        createdUser,
		"accessToken": token,
	})
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var body LoginUserParams

	err := request.Read(r, &body)
	if err != nil {
		request.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := request.Validate(&body, LoginValidationErrors); err != nil {
		request.Error(w, http.StatusBadRequest, err)
		return
	}
	if len(body.Password) > 72 {
		request.Error(w, http.StatusBadRequest, ErrPasswordTooLong)
		return
	}
	
	user, err := h.service.Login(r.Context(), body)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			request.Error(w, http.StatusUnauthorized, ErrInvalidCredentials)
			return
		}
		log.Error("Could not log in user", "err", err)
		request.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		request.Error(w, http.StatusUnauthorized, ErrUserPassword)
		return
	}

	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		log.Error("Could not create token", "err", err)
		request.Error(w, http.StatusInternalServerError, err)
		return
	}


	request.Success(w, http.StatusOK, map[string]any{
		"user":        user,
		"accessToken": token,
	})
}
