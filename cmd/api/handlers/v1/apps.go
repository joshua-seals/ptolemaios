package v1

import (
	"net/http"
)

func (m *Mux) listApps(w http.ResponseWriter, r *http.Request) {
	m.logger.Info("Listing Apps")
	panic("whoops")

}

func (m *Mux) appDetails(w http.ResponseWriter, r *http.Request) {
	m.logger.Info("Get app details")
}
