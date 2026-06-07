package http

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

type testEnvelope struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

type capabilityResponseForTest struct {
	Key            string `json:"key"`
	RiskLevel      string `json:"risk_level"`
	Enabled        bool   `json:"enabled"`
	DisabledReason string `json:"disabled_reason"`
}

func decodeEnvelope(t *testing.T, rec *httptest.ResponseRecorder) testEnvelope {
	t.Helper()

	var envelope testEnvelope
	if err := json.Unmarshal(rec.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("json.Unmarshal(envelope) error = %v", err)
	}

	return envelope
}

func decodeCapability(t *testing.T, raw json.RawMessage) capabilityResponseForTest {
	t.Helper()

	var data capabilityResponseForTest
	decodeRawData(t, raw, &data)

	return data
}

func decodeCapabilityList(t *testing.T, raw json.RawMessage) []capabilityResponseForTest {
	t.Helper()

	var data []capabilityResponseForTest
	decodeRawData(t, raw, &data)

	return data
}

func decodeRawData(t *testing.T, raw json.RawMessage, out any) {
	t.Helper()

	if len(raw) == 0 {
		t.Fatal("data is empty")
	}
	if err := json.Unmarshal(raw, out); err != nil {
		t.Fatalf("json.Unmarshal(data) error = %v", err)
	}
}
