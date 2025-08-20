package auth

import (
	"net/http"
	"errors"
	"testing"
)

func TestAuth(t *testing.T) {
	type test struct {
		input http.Header
		wantKey string
		wantErr error
	}

	tests := []test{
	{input: http.Header{"Authorization": []string{"ApiKey test-key-123"}}, wantKey: "test-key-123", wantErr: nil},
	{input: http.Header{}, wantKey: "", wantErr: ErrNoAuthHeaderIncluded},
	{input: http.Header{"Authorization": []string{"test-key-123"}}, wantKey: "", wantErr: errors.New("malformed authorization header")},
	}

	for _, tc := range tests {
		gotKey, gotErr := GetAPIKey(tc.input)
		if gotKey != tc.wantKey {
    			t.Errorf("expected key %s, got %s", tc.wantKey, gotKey)
		}

		// For error comparison, you could do:
			if (gotErr == nil) != (tc.wantErr == nil) {
    			t.Errorf("expected error %v, got %v", tc.wantErr, gotErr)
		} else if gotErr != nil && gotErr.Error() != tc.wantErr.Error() {
    			t.Errorf("expected error %v, got %v", tc.wantErr, gotErr)
		}
	}

}
