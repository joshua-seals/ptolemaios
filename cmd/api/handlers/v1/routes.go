package v1

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (c *Mux) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/v1/apps/", c.listApps)
	router.HandlerFunc(http.MethodGet, "/api/v1/apps/:app_id/", c.appDetails)
	router.HandlerFunc(http.MethodGet, "/api/v1/instances/", c.listInstances)
	router.HandlerFunc(http.MethodPost, "/api/v1/instances/", c.createInstance)
	router.HandlerFunc(http.MethodGet, "/api/v1/instances/:sid", c.instanceDetails)
	router.HandlerFunc(http.MethodDelete, "/api/v1/instances/:sid", c.deleteInstance)
	router.HandlerFunc(http.MethodPatch, "/api/v1/instances/:sid", c.updateInstance)
	router.HandlerFunc(http.MethodDelete, "/api/v1/instances/:sid/is_ready/", c.instanceIsReady)
	router.HandlerFunc(http.MethodGet, "/api/v1/providers/", c.listAuthProviders)
	//  	This endpoint is designed to support
	// 		scenarios where a reverse proxy (like nginx) performs authentication
	// 		before proxying a request.
	router.HandlerFunc(http.MethodDelete, "/api/v1/users/", c.listUsers)
	// Uses private method to retrieve user's access token from session.
	router.HandlerFunc(http.MethodPost, "/api/v1/users/logout/", c.userLogout)
	router.HandlerFunc(http.MethodGet, "/api/v1/context/", c.helxContext)

	router.Handler(http.MethodGet, "/api/v1/metrics", expvar.Handler())

	// Wrap the middleware around router so we run before ServeMux accesses.
	return c.recoverPanic(c.logRequest(secureHeaders(router)))
}
