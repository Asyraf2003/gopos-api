package jwt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"pos-go/internal/modules/auth/ports"
)

func (i *Issuer) IssueAccessToken(ctx context.Context, req ports.AccessTokenRequest) (string, time.Time, error) {
	_ = ctx

	if req.AccountID == "" || req.SessionID == "" {
		return "", time.Time{}, errors.New("missing account/session")
	}

	now := time.Now()
	exp := now.Add(i.ttl)

	hs, err := encodeJSON(header{
		Alg: "HS256",
		Typ: "JWT",
		Kid: i.kid,
	})
	if err != nil {
		return "", time.Time{}, err
	}

	ps, err := encodeJSON(payload{
		Iss: i.issuer,
		Aud: i.aud,
		Sub: req.AccountID,
		Sid: req.SessionID,
		AAL: req.TrustLevel,
		IAT: now.Unix(),
		EXP: exp.Unix(),
	})
	if err != nil {
		return "", time.Time{}, err
	}

	input := hs + "." + ps
	sig := signHS256(i.secret, input)

	return fmt.Sprintf("%s.%s", input, sig), exp, nil
}
