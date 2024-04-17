package v1

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (m *Mux) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/v1/apps/", m.listApps)
	router.HandlerFunc(http.MethodGet, "/api/v1/apps/:app_id/", m.appDetails)
	router.HandlerFunc(http.MethodGet, "/api/v1/instances/", m.listInstances)
	router.HandlerFunc(http.MethodPost, "/api/v1/instances/", m.createInstance)
	router.HandlerFunc(http.MethodGet, "/api/v1/instances/:sid", m.instanceDetails)
	router.HandlerFunc(http.MethodDelete, "/api/v1/instances/:sid", m.deleteInstance)
	router.HandlerFunc(http.MethodPatch, "/api/v1/instances/:sid", m.updateInstance)
	router.HandlerFunc(http.MethodDelete, "/api/v1/instances/:sid/is_ready/", m.instanceIsReady)
	router.HandlerFunc(http.MethodGet, "/api/v1/providers/", m.listAuthProviders)
	//  	This endpoint is designed to support
	// 		scenarios where a reverse proxy (like nginx) performs authentication
	// 		before proxying a request.
	router.HandlerFunc(http.MethodDelete, "/api/v1/users/", m.listUsers)
	// Uses private method to retrieve user's access token from session.
	router.HandlerFunc(http.MethodPost, "/api/v1/users/logout/", m.userLogout)
	router.HandlerFunc(http.MethodGet, "/api/v1/context/", m.helxContext)

	// Generate auth token for admin usage here.
	router.Handler(http.MethodPost, "/auth", http.HandlerFunc(m.createAuthToken))
	// Add privileged CRUD items that require auth token
	// These will be to add applications to Workstation-Database
	// This may require a different Database entirely - mongodb perhaps.
	// router.Handler(http.MethodPost, "/app", m.authenticate(http.HandlerFunc(m.newApp)))
	// router.Handler(http.MethodDelete, "/app/:id", m.authenticate(http.HandlerFunc(m.deleteApp)))
	// router.Handler(http.MethodPatch, "/app/:id", m.authenticate(http.HandlerFunc(m.updateApp)))

	router.Handler(http.MethodGet, "/api/v1/metrics", expvar.Handler())

	// Wrap the middleware around router so we run before ServeMux accesses.
	return m.recoverPanic(m.logRequest(secureHeaders(router)))
}
