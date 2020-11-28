package servers

import (
	"encoding/json"
	"net/http"
)

func sendResponse(w http.ResponseWriter, status int, data interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(data)
	return err
}

func OK(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusCreated, data)
}

func BadRequest(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusBadRequest, data)
}

func Unauthorized(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusUnauthorized, data)
}

func Forbidden(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusForbidden, data)
}

func NotFound(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusNotFound, data)
}

func MethodNotAllowed(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusMethodNotAllowed, data)
}

func Conflict(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusConflict, data)
}

func Teapot(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusTeapot, data)
}

func InternalServerError(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusInternalServerError, data)
}

func NotImplemented(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusNotImplemented, data)
}

func BadGateway(w http.ResponseWriter, data interface{}) error {
	return sendResponse(w, http.StatusBadGateway, data)
}
