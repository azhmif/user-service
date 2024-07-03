package integrations

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	UserInfoURL          = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	Provider             = "https://accounts.google.com"
	ScopesURLUserInfo    = "https://www.googleapis.com/auth/userinfo.email"
	ScopesURLUserProfile = "https://www.googleapis.com/auth/userinfo.profile"
	RandomString         = "123qwerty"
)

var (
	SSOSignup *oauth2.Config
	SSOSignin *oauth2.Config

	ClientIDSignup     = "2089389193-q3la82qdnhkcdpddmkuhvbb9mctktf7r.apps.googleusercontent.com"
	ClientSecretSignup = "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"
	RedirectURLSignup  = "http://localhost:9990/users/signup/callback"

	ClientIDSignin     = "2089389193-hh9d0b6iorrlmhtgoqdsh9fe7iopabft.apps.googleusercontent.com"
	ClientSecretSignin = "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"
	RedirectURLSignin  = "http://localhost:9990/users/signin/callback"
)

func init() {
	SSOSignup = initOAuthConfig(ClientIDSignup, ClientSecretSignup, RedirectURLSignup)
	SSOSignin = initOAuthConfig(ClientIDSignin, ClientSecretSignin, RedirectURLSignin)
}

func initOAuthConfig(clientID, clientSecret, redirectURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			ScopesURLUserInfo,
			ScopesURLUserProfile,
		},
		Endpoint: google.Endpoint,
	}
}
