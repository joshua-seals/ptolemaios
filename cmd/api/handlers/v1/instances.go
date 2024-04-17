package v1

import "net/http"

func (m *Mux) listInstances(w http.ResponseWriter, r *http.Request)   {}
func (m *Mux) createInstance(w http.ResponseWriter, r *http.Request)  {}
func (m *Mux) instanceDetails(w http.ResponseWriter, r *http.Request) {}
func (m *Mux) deleteInstance(w http.ResponseWriter, r *http.Request)  {}
func (m *Mux) updateInstance(w http.ResponseWriter, r *http.Request)  {}
func (m *Mux) instanceIsReady(w http.ResponseWriter, r *http.Request) {}
