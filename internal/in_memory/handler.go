package inmemory

import (
	"encoding/json"
	"getir-assignment/pkg/response"
	"net/http"
	"regexp"

	"go.uber.org/zap"
)

var (
	getRegex = regexp.MustCompile(`^\/in-memory[\/]*$`)
	setRegex = regexp.MustCompile(`^\/in-memory[\/]*$`)
)

type MemoryService interface {
	Get(key string) (*InMemory, error)
	Set(key string, value string) (*InMemory, error)
}

type Handler struct {
	logger  *zap.Logger
	service MemoryService
}

func NewHandler(service MemoryService, logger *zap.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	h.logger.Debug("Get key value pair request arrived", zap.String("key", key))

	if key == "" {
		h.logger.Sugar().Error("key is required")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("key is required"))
		return
	}

	inMemory, err := h.service.Get(key)
	if err != nil {
		h.logger.Sugar().Errorf("Error while getting in-memory", zap.Error(err))
		h.logger.Error("Error while getting in-memory", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonBytes, err := json.Marshal(inMemory)
	if err != nil {
		h.logger.Error("Error while marshaling data", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	h.logger.Debug("Get key value pair handler executed successfully")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
	var data InMemory
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.logger.Sugar().Errorf("Json decode error", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your request is in a bad format."))
		return
	}
	h.logger.Debug("Set key value pair request arrived", zap.String("key", data.Key), zap.String("value", data.Value))

	inMemory, err := h.service.Set(data.Key, data.Value)
	if err != nil {
		h.logger.Error("Set in-memory error", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonBytes, err := json.Marshal(inMemory)
	if err != nil {
		h.logger.Error("Error while marshaling data", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	h.logger.Debug("Get key value pair handler executed successfully")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost && setRegex.MatchString(r.URL.Path):
		h.Set(w, r)
		return
	case r.Method == http.MethodGet && getRegex.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	default:
		response.NotFound("", w, r)
	}
}
