package v1

import (
	"fmt"
	"net/http"
)

// The logError() method is a generic helper for logging an error message along // with the current request method and URL as attributes in the log entry.
func (c *CoreHandler) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	c.logger.Error(err.Error(), "method", method, "uri", uri)
}

// The errorResponse() method is a generic helper for sending JSON-formatted error // messages to the client with a given status code. Note that we're using the any // type for the message parameter, rather than just a string type, as this gives us // more flexibility over the values that we can include in the response.
func (c *CoreHandler) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	// Write the response using the writeJSON() helper. If this happens to return an
	// error then log it, and fall back to sending the client an empty response with a
	// 500 Internal Server Error status code.
	err := c.writeJSON(w, status, env, nil)
	if err != nil {
		c.logError(r, err)
		w.WriteHeader(500)
	}
}

// The serverErrorResponse() method will be used when our application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses the
// errorResponse() helper to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func (c *CoreHandler) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	c.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	c.errorResponse(w, r, http.StatusInternalServerError, message)
}

// The notFoundResponse() method will be used to send a 404 Not Found status code and // JSON response to the client.
func (c *CoreHandler) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	c.errorResponse(w, r, http.StatusNotFound, message)
}

// The methodNotAllowedResponse() method will be used to send a 405 Method Not Allowed
// status code and JSON response to the client.
func (c *CoreHandler) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	c.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
