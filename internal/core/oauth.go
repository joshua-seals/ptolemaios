package core

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/microsoft"
	"golang.org/x/oauth2/slack"
)

type Provider int

const (
	_               = iota
	Github Provider = iota
	Microsoft
	Slack
)

// Default to github
func NewOauthConfig(p Provider) *oauth2.Config {
	switch p {
	case Github:
		return &oauth2.Config{
			ClientID:     getEnv("CLIENT_ID", "need client ID"),
			ClientSecret: getEnv("CLIENT_SECRET", "need client secret"),
			Scopes:       []string{"user:email"},
			RedirectURL:  getEnv("REDIRECT_URL", "http://localhost:8585/login/oauth2/callback/"),
			Endpoint:     github.Endpoint,
		}
	case Microsoft:
		return &oauth2.Config{
			ClientID:     getEnv("CLIENT_ID", "need client ID"),
			ClientSecret: getEnv("CLIENT_SECRET", "need client secret"),
			Scopes:       []string{"user:email"},
			RedirectURL:  getEnv("REDIRECT_URL", "http://localhost:8585/login/oauth2/callback/"),
			Endpoint:     microsoft.LiveConnectEndpoint,
		}
	case Slack:
		return &oauth2.Config{
			ClientID:     getEnv("CLIENT_ID", "need client ID"),
			ClientSecret: getEnv("CLIENT_SECRET", "need client secret"),
			Scopes:       []string{"user:email"},
			RedirectURL:  getEnv("REDIRECT_URL", ""),
			Endpoint:     slack.Endpoint,
		}
	default:
		// Default to Github
		return &oauth2.Config{
			ClientID:     getEnv("CLIENT_ID", "need client ID"),
			ClientSecret: getEnv("CLIENT_SECRET", "need client secret"),
			Scopes:       []string{"user:email"},
			RedirectURL:  getEnv("REDIRECT_URL", "http://localhost:8585/login/oauth2/callback/"),
			Endpoint:     github.Endpoint,
		}
	}

}

// Ref: https://sharmarajdaksh.github.io/blog/github-oauth-with-go
