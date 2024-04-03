package v1

import (
	"net/http"
)

func (c *Mux) listApps(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Listing Apps")
	panic("whoops")

}

func (c *Mux) appDetails(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Get app details")
}
