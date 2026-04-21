package jwt

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"pos-go/internal/modules/auth/ports"
)

type Verifier struct {
	issuer string
	aud    string
	secret []byte
	nowFn  func() time.Time
}

type verifiedPayload struct {
	Iss string `json:"iss"`
	Aud string `json:"aud"`
	Sub string `json:"sub"`
	Sid string `json:"sid"`
	AAL string `json:"aal"`
	IAT int64  `json:"iat"`
	EXP int64  `json:"exp"`
}

func NewHMACVerifier(issuer, aud, secret string) (*Verifier, error) {
	if strings.TrimSpace(secret) == "" {
		return nil, errors.New("jwt secret empty")
	}

	return &Verifier{
		issuer: issuer,
		aud:    aud,
		secret: []byte(secret),
		nowFn:  time.Now,
	}, nil
}

func (v *Verifier) VerifyAccessToken(ctx context.Context, token string) (ports.AccessTokenClaims, error) {
	_ = ctx

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return ports.AccessTokenClaims{}, errors.New("invalid token format")
	}

	headerPart, payloadPart, sigPart := parts[0], parts[1], parts[2]

	if !verifyHS256(v.secret, headerPart+"."+payloadPart, sigPart) {
		return ports.AccessTokenClaims{}, errors.New("invalid token signature")
	}

	payloadJSON, err := base64.RawURLEncoding.DecodeString(payloadPart)
	if err != nil {
		return ports.AccessTokenClaims{}, err
	}

	var payload verifiedPayload
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return ports.AccessTokenClaims{}, err
	}

	if payload.Iss != v.issuer {
		return ports.AccessTokenClaims{}, errors.New("invalid token issuer")
	}
	if payload.Aud != v.aud {
		return ports.AccessTokenClaims{}, errors.New("invalid token audience")
	}
	if payload.Sub == "" || payload.Sid == "" || payload.AAL == "" {
		return ports.AccessTokenClaims{}, errors.New("missing token claims")
	}

	nowUnix := v.nowFn().Unix()
	if payload.EXP <= nowUnix {
		return ports.AccessTokenClaims{}, errors.New("token expired")
	}

	return ports.AccessTokenClaims{
		AccountID:  payload.Sub,
		SessionID:  payload.Sid,
		TrustLevel: payload.AAL,
	}, nil
}

func verifyHS256(secret []byte, input, givenSig string) bool {
	mac := hmac.New(sha256.New, secret)
	_, _ = mac.Write([]byte(input))
	expected := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(givenSig))
}

var _ ports.AccessTokenVerifier = (*Verifier)(nil)
