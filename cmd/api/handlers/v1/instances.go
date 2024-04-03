package v1

import "net/http"

func (c *Mux) listInstances(w http.ResponseWriter, r *http.Request)   {}
func (c *Mux) createInstance(w http.ResponseWriter, r *http.Request)  {}
func (c *Mux) instanceDetails(w http.ResponseWriter, r *http.Request) {}
func (c *Mux) deleteInstance(w http.ResponseWriter, r *http.Request)  {}
func (c *Mux) updateInstance(w http.ResponseWriter, r *http.Request)  {}
func (c *Mux) instanceIsReady(w http.ResponseWriter, r *http.Request) {}
