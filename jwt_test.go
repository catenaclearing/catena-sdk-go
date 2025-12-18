package catena

import (
	"testing"
)

func TestDecodeJWT(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		wantErr bool
	}{
		{
			name:    "valid token",
			token:   "header.eyJleHAiOjEyMzQ1Njc4OTAsInN1YiI6InVzZXIxIiwib3JnIjp7Im9yZzEiOnsiaWQiOiJvcmcxIiwib3JnX3R5cGUiOlsiY2FycmllciJdfX19.sig",
			wantErr: false,
		},
		{
			name:    "malformed token (no parts)",
			token:   "invalid",
			wantErr: true,
		},
		{
			name:    "malformed token (bad base64)",
			token:   "header.invalid-base64.sig",
			wantErr: true,
		},
		{
			name:    "malformed token (bad json)",
			token:   "header.bm90LWpzb24.sig", // "not-json"
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := decodeJWT(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeJWT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
