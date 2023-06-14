package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (d *Delivery) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(d.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(d.methodNotAllowedResponse)

	//router.HandlerFunc(http.MethodGet, "/v1/healthcheck", d.healthcheckHandler)
	//
	//router.HandlerFunc(http.MethodPost, "/v1/shelters", d.createShelterHandler)
	//router.HandlerFunc(http.MethodGet, "/v1/shelters/:id", d.showShelterHandler)
	//router.HandlerFunc(http.MethodGet, "/v1/shelters", d.listSheltersHandler)
	//router.HandlerFunc(http.MethodPatch, "/v1/shelters/:id", d.updateShelterHandler)
	//router.HandlerFunc(http.MethodDelete, "/v1/shelters/:id", d.deleteShelterHandler)
	//
	//router.HandlerFunc(http.MethodPost, "/v1/users", d.registerUserHandler)

	return d.rateLimiter(d.recoverPanic(router))
}
