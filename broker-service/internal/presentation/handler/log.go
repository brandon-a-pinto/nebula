package handler

import (
	"encoding/json"
	"net/http"

	"github.com/brandon-a-pinto/nebula/broker-service/internal/application/usecase"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/domain/dto"
)

type LogHandler struct {
	CreateLogUsecase usecase.CreateLogUsecase
}

func NewLogHandler(createLogUsecase usecase.CreateLogUsecase) *LogHandler {
	return &LogHandler{
		CreateLogUsecase: createLogUsecase,
	}
}

func (h *LogHandler) CreateLog(w http.ResponseWriter, r *http.Request) {
	dto := new(dto.CreateLogInput)
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.CreateLogUsecase.Exec(r.Context(), *dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
