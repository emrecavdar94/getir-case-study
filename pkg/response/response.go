package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Records interface{} `json:"records"`
}

func Success(data interface{}, w http.ResponseWriter, r *http.Request) {
	response := Response{
		Code:    0,
		Message: "success",
		Records: data,
	}
	w.WriteHeader(http.StatusOK)
	jsonBytes, _ := json.Marshal(response)
	w.Write(jsonBytes)
}

func InternalServerError(msg string, w http.ResponseWriter, r *http.Request) {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	response := Response{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Records: nil,
	}
	w.WriteHeader(http.StatusInternalServerError)
	jsonBytes, _ := json.Marshal(response)
	w.Write(jsonBytes)
}

func MethodNotAllowedError(msg string, w http.ResponseWriter, r *http.Request) {
	if msg == "" {
		msg = "Method Not Allowed."
	}
	response := Response{
		Code:    http.StatusMethodNotAllowed,
		Message: msg,
		Records: nil,
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	jsonBytes, _ := json.Marshal(response)
	w.Write(jsonBytes)
}

func NotFound(msg string, w http.ResponseWriter, r *http.Request) {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	response := Response{
		Code:    http.StatusNotFound,
		Message: msg,
		Records: nil,
	}
	w.WriteHeader(http.StatusNotFound)
	jsonBytes, _ := json.Marshal(response)
	w.Write(jsonBytes)
}

func BadRequest(msg string, w http.ResponseWriter, r *http.Request) {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	response := Response{
		Code:    http.StatusBadRequest,
		Message: msg,
		Records: nil,
	}
	w.WriteHeader(http.StatusBadRequest)
	jsonBytes, _ := json.Marshal(response)
	w.Write(jsonBytes)
}
