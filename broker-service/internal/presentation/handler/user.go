package handler

import (
	"encoding/json"
	"net/http"

	"github.com/brandon-a-pinto/nebula/broker-service/internal/application/usecase"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/domain/dto"
)

type UserHandler struct {
	CreateUserUsecase usecase.CreateUserUsecase
}

func NewUserHandler(createUserUsecase usecase.CreateUserUsecase) *UserHandler {
	return &UserHandler{
		CreateUserUsecase: createUserUsecase,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	dto := new(dto.CreateUserInput)
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := h.CreateUserUsecase.Exec(r.Context(), *dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
