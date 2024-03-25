package v1

import "net/http"

func (c *CoreHandler) listInstances(w http.ResponseWriter, r *http.Request)      {}
func (c *CoreHandler) createInstance(w http.ResponseWriter, r *http.Request)     {}
func (c *CoreHandler) instanceDetails(w http.ResponseWriter, r *http.Request) {}
func (c *CoreHandler) deleteInstance(w http.ResponseWriter, r *http.Request)     {}
func (c *CoreHandler) updateInstance(w http.ResponseWriter, r *http.Request)     {}
func (c *CoreHandler) instanceIsReady(w http.ResponseWriter, r *http.Request)    {}
