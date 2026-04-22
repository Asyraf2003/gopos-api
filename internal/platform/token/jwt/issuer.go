package jwt

import (
	"errors"
	"strings"
	"time"

	"pos-go/internal/modules/auth/ports"
)

type Issuer struct {
	issuer string
	aud    string
	kid    string
	ttl    time.Duration
	secret []byte
}

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
	Kid string `json:"kid,omitempty"`
}

type payload struct {
	Iss string `json:"iss"`
	Aud string `json:"aud"`
	Sub string `json:"sub"`
	Sid string `json:"sid"`
	AAL string `json:"aal"`
	IAT int64  `json:"iat"`
	EXP int64  `json:"exp"`
}

func NewHMACIssuer(issuer, aud, kid, secret string, ttl time.Duration) (*Issuer, error) {
	if strings.TrimSpace(secret) == "" {
		return nil, errors.New("jwt secret empty")
	}
	if ttl <= 0 {
		return nil, errors.New("jwt ttl invalid")
	}

	return &Issuer{
		issuer: issuer,
		aud:    aud,
		kid:    kid,
		ttl:    ttl,
		secret: []byte(secret),
	}, nil
}

var _ ports.TokenIssuer = (*Issuer)(nil)
