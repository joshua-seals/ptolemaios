package v1

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/joshua-seals/ptolemaios/internal/core"
)

// Begin Google OAuth2
// Ref: https://www.kungfudev.com/blog/2018/07/10/oauth2-example-with-go

func (m *Mux) oauthLogin(w http.ResponseWriter, r *http.Request) {
	// Create oauthState cookie
	oauthState := m.generateStateOauthCookie(w)
	conf := core.NewOauthConfig()
	u := conf.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (m *Mux) generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

// func (m *Mux) oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
// 	// Read oauthState from Cookie
// 	oauthState, _ := r.Cookie("oauthstate")

// 	if r.FormValue("state") != oauthState.Value {
// 		log.Println("invalid oauth google state")
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	data, err := getUserDataFromGoogle(r.FormValue("code"))
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	// GetOrCreate User in your db.
// 	// Redirect or response with a token.
// 	// More code .....
// 	fmt.Fprintf(w, "UserInfo: %s\n", data)
// }

// func (m *Mux) getUserDataFromGoogle(code string) ([]byte, error) {
// 	// Use code to get token and get user info from Google.

// 	token, err := googleOauthConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
// 	}
// 	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
// 	}
// 	defer response.Body.Close()
// 	contents, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed read response: %s", err.Error())
// 	}
// 	return contents, nil
// }

// End Google OAuth2

func (m *Mux) listAuthProviders(w http.ResponseWriter, r *http.Request) {}
