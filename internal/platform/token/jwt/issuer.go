package jwt

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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

func (i *Issuer) IssueAccessToken(ctx context.Context, req ports.AccessTokenRequest) (string, time.Time, error) {
	_ = ctx

	if req.AccountID == "" || req.SessionID == "" {
		return "", time.Time{}, errors.New("missing account/session")
	}

	now := time.Now()
	exp := now.Add(i.ttl)

	h := header{
		Alg: "HS256",
		Typ: "JWT",
		Kid: i.kid,
	}

	p := payload{
		Iss: i.issuer,
		Aud: i.aud,
		Sub: req.AccountID,
		Sid: req.SessionID,
		AAL: req.TrustLevel,
		IAT: now.Unix(),
		EXP: exp.Unix(),
	}

	hs, err := encodeJSON(h)
	if err != nil {
		return "", time.Time{}, err
	}

	ps, err := encodeJSON(p)
	if err != nil {
		return "", time.Time{}, err
	}

	input := hs + "." + ps
	sig := signHS256(i.secret, input)

	return fmt.Sprintf("%s.%s", input, sig), exp, nil
}

func encodeJSON(v any) (string, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(raw), nil
}

func signHS256(secret []byte, input string) string {
	mac := hmac.New(sha256.New, secret)
	_, _ = mac.Write([]byte(input))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

var _ ports.TokenIssuer = (*Issuer)(nil)
