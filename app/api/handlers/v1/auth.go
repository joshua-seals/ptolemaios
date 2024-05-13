package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/joshua-seals/ptolemaios/internal/core"
)

// Begin Google OAuth2
// Ref: https://www.kungfudev.com/blog/2018/07/10/oauth2-example-with-go

type UserInfo struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

func (m *Mux) loggedinHandler(w http.ResponseWriter, r *http.Request, githubData string) {
	if githubData == "" {
		// Unauthorized users get an unauthorized message
		fmt.Fprintf(w, "UNAUTHORIZED!")
		return
	}

	// Set return type JSON
	w.Header().Set("Content-type", "application/json")

	// Prettifying the json
	var prettyJSON bytes.Buffer
	// json.indent is a library utility function to prettify JSON indentation
	parserr := json.Indent(&prettyJSON, []byte(githubData), "", "\t")
	if parserr != nil {
		log.Panic("JSON parse error")
	}

	// Return the prettified JSON as a string
	fmt.Fprintf(w, string(prettyJSON.Bytes()))
}

func (m *Mux) splashPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<a href="/login/github/">LOGIN</a>`)
}

func (m *Mux) githubLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Get the environment variable
	conf := core.NewOauthConfig(core.Github)

	// Create the dynamic redirect URL for login
	redirectURL := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s",
		conf.Endpoint.AuthURL, conf.ClientID, conf.RedirectURL)

	http.Redirect(w, r, redirectURL, 301)
}

func (m *Mux) githubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	githubAccessToken := getGithubAccessToken(code)

	githubData := getGithubData(githubAccessToken)

	m.loggedinHandler(w, r, githubData)
}

func getGithubAccessToken(code string) string {
	conf := core.NewOauthConfig(core.Github)
	// Set up the request body as JSON
	requestBodyMap := map[string]string{
		"client_id":     conf.ClientID,
		"client_secret": conf.ClientSecret,
		"code":          code,
	}
	requestJSON, _ := json.Marshal(requestBodyMap)

	// POST request to set URL
	req, reqerr := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if reqerr != nil {
		log.Panic("Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Get the response
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	// Response body converted to stringified JSON
	respbody, _ := io.ReadAll(resp.Body)

	// Represents the response received from Github
	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
	var ghresp githubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	// Return the access token (as the rest of the
	// details are relatively unnecessary for us)
	return ghresp.AccessToken
}

func getGithubData(accessToken string) string {
	// Get request to a set URL
	req, reqerr := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil)
	if reqerr != nil {
		log.Panic("API Request creation failed")
	}

	// Set the Authorization header before sending the request
	// Authorization: token XXXXXXXXXXXXXXXXXXXXXXXXXXX
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	// Make the request
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	// Read the response as a byte slice
	respbody, _ := io.ReadAll(resp.Body)

	// Convert byte slice to string and return
	return string(respbody)
}

func (m *Mux) listAuthProviders(w http.ResponseWriter, r *http.Request) {}
