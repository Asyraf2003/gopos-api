package system

import authdomain "pos-go/internal/modules/auth/domain"

type MeResponse struct {
	AccountID   string   `json:"account_id"`
	SessionID   string   `json:"session_id"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	TrustLevel  string   `json:"trust_level"`
}

func Me(principal authdomain.Principal) MeResponse {
	return MeResponse{
		AccountID:   principal.AccountID,
		SessionID:   principal.SessionID,
		Roles:       principal.Roles,
		Permissions: principal.Permissions,
		TrustLevel:  principal.TrustLevel,
	}
}
