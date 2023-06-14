package http

import (
	"fmt"
	http2 "github.com/nourbol/zhandos/pkg/tools/http"
	"net/http"
)

func (d *Delivery) logError(r *http.Request, err error) {
	d.logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (d *Delivery) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	envelopedError := http2.Envelope{"error": message}
	err := http2.WriteJSON(w, status, envelopedError, nil)
	if err != nil {
		d.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (d *Delivery) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	d.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (d *Delivery) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	d.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	d.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (d *Delivery) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	d.errorResponse(w, r, http.StatusNotFound, message)
}

func (d *Delivery) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	d.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (d *Delivery) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	d.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (d *Delivery) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	d.errorResponse(w, r, http.StatusConflict, message)
}

func (d *Delivery) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	d.errorResponse(w, r, http.StatusTooManyRequests, message)
}
