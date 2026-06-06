package domain

import (
	"errors"
	"strings"
)

type RiskLevel string

const (
	RiskLow    RiskLevel = "low"
	RiskMedium RiskLevel = "medium"
	RiskHigh   RiskLevel = "high"
)

type Capability struct {
	Key                 string
	Domain              string
	Operation           string
	Method              string
	Path                string
	DefaultEnabled      bool
	Enabled             bool
	RequiredPermission  string
	RiskLevel           RiskLevel
	AuditRequired       bool
	IdempotencyRequired bool
	OwnerPackage        string
	TestProof           string
	DisabledReason      string
}

func NewCapability(input Capability) (Capability, error) {
	capability := Capability{
		Key:                 normalize(input.Key),
		Domain:              normalize(input.Domain),
		Operation:           normalize(input.Operation),
		Method:              strings.ToUpper(normalize(input.Method)),
		Path:                normalize(input.Path),
		DefaultEnabled:      input.DefaultEnabled,
		Enabled:             input.Enabled,
		RequiredPermission:  normalize(input.RequiredPermission),
		RiskLevel:           input.RiskLevel,
		AuditRequired:       input.AuditRequired,
		IdempotencyRequired: input.IdempotencyRequired,
		OwnerPackage:        normalize(input.OwnerPackage),
		TestProof:           normalize(input.TestProof),
		DisabledReason:      normalize(input.DisabledReason),
	}

	if err := capability.Validate(); err != nil {
		return Capability{}, err
	}

	if capability.Enabled {
		capability.DisabledReason = ""
	}

	return capability, nil
}

func (c Capability) Validate() error {
	if c.Key == "" {
		return errors.New("capability key is required")
	}
	if c.Domain == "" {
		return errors.New("capability domain is required")
	}
	if c.Operation == "" {
		return errors.New("capability operation is required")
	}
	if c.Method == "" {
		return errors.New("capability method is required")
	}
	if c.Path == "" {
		return errors.New("capability path is required")
	}
	if c.RequiredPermission == "" {
		return errors.New("capability required permission is required")
	}
	if c.OwnerPackage == "" {
		return errors.New("capability owner package is required")
	}
	if c.TestProof == "" {
		return errors.New("capability test proof is required")
	}
	if !c.RiskLevel.Valid() {
		return errors.New("capability risk level is invalid")
	}

	return nil
}

func (c Capability) Disable(reason string) Capability {
	c.Enabled = false
	c.DisabledReason = normalize(reason)

	return c
}

func (c Capability) Enable() Capability {
	c.Enabled = true
	c.DisabledReason = ""

	return c
}

func (r RiskLevel) Valid() bool {
	return r == RiskLow || r == RiskMedium || r == RiskHigh
}

func normalize(value string) string {
	return strings.TrimSpace(value)
}
