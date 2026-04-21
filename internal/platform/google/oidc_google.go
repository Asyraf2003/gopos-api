package google

import (
	"context"
	"errors"
	"strings"
	"time"

	"pos-go/internal/modules/auth/ports"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type OIDCConfig struct {
	Issuer       string
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

type OIDC struct {
	provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
	oauth    oauth2.Config
}

func NewOIDC(ctx context.Context, cfg OIDCConfig) (*OIDC, error) {
	if strings.TrimSpace(cfg.Issuer) == "" {
		return nil, errors.New("oidc issuer empty")
	}
	if strings.TrimSpace(cfg.ClientID) == "" {
		return nil, errors.New("oidc client id empty")
	}
	if strings.TrimSpace(cfg.ClientSecret) == "" {
		return nil, errors.New("oidc client secret empty")
	}

	provider, err := oidc.NewProvider(ctx, cfg.Issuer)
	if err != nil {
		return nil, err
	}

	return &OIDC{
		provider: provider,
		verifier: provider.Verifier(&oidc.Config{
			ClientID: cfg.ClientID,
		}),
		oauth: oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Endpoint:     provider.Endpoint(),
			RedirectURL:  cfg.RedirectURL,
			Scopes:       []string{oidc.ScopeOpenID, "email", "profile"},
		},
	}, nil
}

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

func (o *OIDC) ExchangeAndVerify(
	ctx context.Context,
	code string,
	codeVerifier string,
	redirectURL string,
	nonce string,
) (ports.OIDCClaims, error) {
	cfg := o.oauthConfigFor(redirectURL)

	token, err := cfg.Exchange(
		ctx,
		code,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier),
	)
	if err != nil {
		return ports.OIDCClaims{}, err
	}

	rawIDToken, _ := token.Extra("id_token").(string)
	if rawIDToken == "" {
		return ports.OIDCClaims{}, errors.New("missing id_token")
	}

	idToken, err := o.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return ports.OIDCClaims{}, err
	}

	var claims struct {
		Sub           string `json:"sub"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		Nonce         string `json:"nonce"`
		AuthTime      int64  `json:"auth_time"`
	}

	if err := idToken.Claims(&claims); err != nil {
		return ports.OIDCClaims{}, err
	}

	if nonce != "" && claims.Nonce != nonce {
		return ports.OIDCClaims{}, errors.New("nonce mismatch")
	}

	authTime := time.Unix(claims.AuthTime, 0)

	return ports.OIDCClaims{
		Provider:      "google",
		Subject:       claims.Sub,
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
		AuthTime:      authTime,
	}, nil
}

func (o *OIDC) oauthConfigFor(redirectURL string) oauth2.Config {
	cfg := o.oauth
	if strings.TrimSpace(redirectURL) != "" {
		cfg.RedirectURL = redirectURL
	}
	return cfg
}

var _ ports.OIDCProvider = (*OIDC)(nil)
