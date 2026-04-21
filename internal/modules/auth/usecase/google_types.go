package usecase

import "time"

type GoogleStartInput struct {
	Purpose     string
	RedirectURL string
}

type GoogleStartOutput struct {
	RedirectTo string `json:"redirect_to"`
	State      string `json:"state"`
}

type ClientInfo struct {
	UserAgent string
	IP        string
}

type GoogleCallbackInput struct {
	Code        string
	State       string
	RedirectURL string
	Client      ClientInfo
}

type GoogleCallbackOutput struct {
	AccessToken    string    `json:"access_token"`
	AccessExp      time.Time `json:"access_exp"`
	RefreshToken   string    `json:"refresh_token"`
	RefreshExp     time.Time `json:"refresh_exp"`
	TrustLevel     string    `json:"trust_level"`
	StepUpRequired bool      `json:"step_up_required"`
}
