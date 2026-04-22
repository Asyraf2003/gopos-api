package google

import (
	"pos-go/internal/modules/auth/ports"

	"golang.org/x/oauth2"
)

func (o *OIDC) AuthCodeURL(p ports.OIDCAuthURLParams) string {
	cfg := o.oauthConfigFor(p.RedirectURL)

	opts := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("nonce", p.Nonce),
		oauth2.SetAuthURLParam("code_challenge", p.CodeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("prompt", "select_account"),
	}

	return cfg.AuthCodeURL(p.State, opts...)
}
