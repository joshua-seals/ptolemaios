package v1

import (
	"net/http"
)

func (c *CoreHandler) listApps(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Listing Apps")
	panic("whoops")

}

func (c *CoreHandler) appDetails(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Get app details")
}
