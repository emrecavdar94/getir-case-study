package record

import (
	"encoding/json"
	"getir-assignment/pkg/response"
	"net/http"
	"regexp"

	"go.uber.org/zap"
)

var (
	getRecordsRegex = regexp.MustCompile(`^\/record[\/]*$`)
)

type RecordService interface {
	GetRecordsByDateAndCount(recordRequest RecordRequest) ([]*Record, error)
}

type Handler struct {
	logger  *zap.Logger
	service RecordService
}

func NewHandler(service RecordService, logger *zap.Logger) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) GetRecordsByDateAndCount(w http.ResponseWriter, r *http.Request) {
	recordRequest := RecordRequest{}

	if err := json.NewDecoder(r.Body).Decode(&recordRequest); err != nil {
		response.InternalServerError(err.Error(), w, r)
		return
	}
	h.logger.Debug("Getting records request arrived", zap.String("startDate", recordRequest.StartDate),
		zap.String("endDate", recordRequest.EndDate),
		zap.Int("minCount", recordRequest.MinCount),
		zap.Int("maxCount", recordRequest.MaxCount))

	records, err := h.service.GetRecordsByDateAndCount(recordRequest)
	if err != nil {
		response.InternalServerError(err.Error(), w, r)
		return
	}
	response.Success(records, w, r)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost && getRecordsRegex.MatchString(r.URL.Path):
		h.GetRecordsByDateAndCount(w, r)
		return
	default:
		response.NotFound("", w, r)
	}
}
