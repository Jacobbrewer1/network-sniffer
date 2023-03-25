package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	IResponse interface {
		Ok(data any)
		Custom(status int, data any)
		Message(status int, message string)
		BadRequest(err error)
		ServerError(err error)
		NotFound()
		MethodNotAllowed()
	}

	response struct {
		w http.ResponseWriter
		r *http.Request
	}
)

func (r response) NotFound() {
	r.Message(http.StatusNotFound, NotFound.Error())
}

func (r response) MethodNotAllowed() {
	r.Message(http.StatusMethodNotAllowed, MethodNotAllowed.Error())
}

func (r response) ServerError(err error) {
	r.Message(http.StatusInternalServerError, err.Error())
}

func (r response) BadRequest(err error) {
	r.Message(http.StatusBadRequest, err.Error())
}

func (r response) Ok(data any) {
	sendResponse(r.w, http.StatusOK, data)
}

func (r response) Message(status int, message string) {
	sendResponse(r.w, status, struct {
		Message *string `json:"message"`
	}{
		Message: &message,
	})
}

func (r response) Custom(status int, data any) {
	sendResponse(r.w, status, data)
}

func NewResponse(w http.ResponseWriter, r *http.Request) IResponse {
	return &response{
		w: w,
		r: r,
	}
}

// sendResponse This will respond to the client with a response code and a body of data
func sendResponse(w http.ResponseWriter, statusCode int, body any) {
	data, err := json.Marshal(body)
	if err != nil {
		log.Print(err)
		statusCode = http.StatusInternalServerError
		data = nil
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(data); err != nil {
		log.Println(err)
		return
	}
}
