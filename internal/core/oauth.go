package core

import "golang.org/x/oauth2"

// Mainly OAuth2 Centers around the client config
// Ref: https://pkg.go.dev/golang.org/x/oauth2@v0.19.0#Config

//	var googleOauthConfig = &oauth2.Config{
//		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
//		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
//		Endpoint:     google.Endpoint,
//		RedirectURL:  "http://localhost:8585/auth/google/callback",
//		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"}
//	}

type Provider int

const (
	_               = iota
	Google Provider = iota
	Github
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func NewOauthConfig() *oauth2.Config {
	var c oauth2.Config
	c.ClientID = getEnv("CLIENT_ID", "blah")
	c.ClientSecret = getEnv("CLIENT_SECRET", "")
	// c.Endpoint = getEnv("ENDPOINTS", "")
	c.RedirectURL = getEnv("REDIRECT_URL", "")
	// c.Scopes = getEnv("SCOPES", "")
	return &c
}

// Ref: https://sharmarajdaksh.github.io/blog/github-oauth-with-go
