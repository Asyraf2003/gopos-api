package domain

import "errors"

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
