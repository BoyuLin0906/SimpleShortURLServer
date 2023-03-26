// deck_test.go
package controllers

import (
	"testing"
)

func TestURL(t *testing.T) {
	test_url := "https://www.google.com/"
	short_url := SetShortURL(test_url)
	original_url := GetOriginalURL(short_url)

	if original_url != test_url {
		t.Errorf("Expected URL")
	}
}
