package core

import (
	"golang.org/x/oauth2"
)

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

// Default to github
func NewGithubOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     getEnv("CLIENT_ID", "blah"),
		ClientSecret: getEnv("CLIENT_SECRET", ""),
		Scopes:       []string{"user:email"},
		RedirectURL:  getEnv("REDIRECT_URL", "http://localhost:8585/login/github/callback/"),
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://github.com/login/oauth/access_token",
			AuthURL:  "https://github.com/login/oauth/authorize",
		},
	}
}

// Ref: https://sharmarajdaksh.github.io/blog/github-oauth-with-go
