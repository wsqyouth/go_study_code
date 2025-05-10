package go_mock

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetS3File(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name     string
		s3URL    string
		mockResp httpmock.Responder
		wantData string
		wantErr  bool
	}{
		{
			name:  "success",
			s3URL: "http://example.com/test.pdf",
			mockResp: httpmock.NewStringResponder(http.StatusOK, "test data"),
			wantData: "test data",
			wantErr:  false,
		},
		{
			name:  "server error 500",
			s3URL: "http://example.com/error-500.pdf",
			mockResp: httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, map[string]interface{}{
				"error": "internal server error",
			}),
			wantErr: true,
		},
		{
			name:  "unauthorized 401",
			s3URL: "http://example.com/error-401.pdf",
			mockResp: httpmock.NewJsonResponderOrPanic(http.StatusUnauthorized, map[string]interface{}{
				"error": "unauthorized",
			}),
			wantErr: true,
		},
		{
			name:  "forbidden 403",
			s3URL: "http://example.com/error-403.pdf",
			mockResp: httpmock.NewJsonResponderOrPanic(http.StatusForbidden, map[string]interface{}{
				"error": "forbidden",
			}),
			wantErr: true,
		},
		{
			name:  "not found 404",
			s3URL: "http://example.com/error-404.pdf",
			mockResp: httpmock.NewJsonResponderOrPanic(http.StatusNotFound, map[string]interface{}{
				"error": "not found",
			}),
			wantErr: true,
		},
		{
			name:  "empty response",
			s3URL: "http://example.com/empty.pdf",
			mockResp: httpmock.NewStringResponder(http.StatusOK, ""),
			wantData: "",
			wantErr:  false,
		},
		{
			name:  "network error",
			s3URL: "http://example.com/network-error.pdf",
			mockResp: httpmock.NewErrorResponder(io.ErrUnexpectedEOF),
			wantErr:  true,
		},
		{
			name:  "invalid url",
			s3URL: "invalid-url",
			mockResp: httpmock.NewStringResponder(http.StatusOK, ""),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.HasPrefix(tt.s3URL, "http") {
				httpmock.RegisterResponder("GET", tt.s3URL, tt.mockResp)
			}

			got, err := GetS3File(context.Background(), tt.s3URL)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantData, string(got))
		})
	}
}