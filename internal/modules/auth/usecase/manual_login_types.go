package usecase

import "time"

type ManualLoginInput struct {
	Email string
}

type ManualLoginOutput struct {
	AccessToken    string    `json:"access_token"`
	AccessExp      time.Time `json:"access_exp"`
	RefreshToken   string    `json:"refresh_token"`
	RefreshExp     time.Time `json:"refresh_exp"`
	TrustLevel     string    `json:"trust_level"`
	StepUpRequired bool      `json:"step_up_required"`
}
