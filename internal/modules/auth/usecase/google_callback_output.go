package usecase

import (
	"context"

	"pos-go/internal/modules/auth/ports"
)

func (u *GoogleFlow) issueGoogleCallbackOutput(
	ctx context.Context,
	accountID string,
	seed googleCallbackSeed,
) (GoogleCallbackOutput, error) {
	accessToken, accessExp, err := u.tokens.IssueAccessToken(ctx, ports.AccessTokenRequest{
		AccountID:  accountID,
		SessionID:  seed.sessionID,
		TrustLevel: "aal1",
	})
	if err != nil {
		return GoogleCallbackOutput{}, err
	}

	return GoogleCallbackOutput{
		AccessToken:    accessToken,
		AccessExp:      accessExp,
		RefreshToken:   seed.refreshToken,
		RefreshExp:     seed.refreshExp,
		TrustLevel:     "aal1",
		StepUpRequired: false,
	}, nil
}
