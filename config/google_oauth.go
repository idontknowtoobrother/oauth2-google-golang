package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/auth/google/callback",
	ClientID:     "1089444531537-v32frg73fnuj1dsse50mollf3l1odstb.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/photoslibrary.readonly",
	},
	Endpoint: google.Endpoint,
}

const OAuthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
