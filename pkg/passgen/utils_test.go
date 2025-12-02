package passgen

import (
	"strings"
	"testing"
)

func TestGenerateRandomString_DefaultLength(t *testing.T) {
	result, err := GenerateRandomString(0)
	if err != nil {
		t.Fatalf("GenerateRandomString(0) error = %v", err)
	}

	if len(result) != DefaultSaltLength {
		t.Errorf("GenerateRandomString(0) length = %d, want %d", len(result), DefaultSaltLength)
	}
}

func TestGenerateRandomString_NegativeLength(t *testing.T) {
	result, err := GenerateRandomString(-1)
	if err != nil {
		t.Fatalf("GenerateRandomString(-1) error = %v", err)
	}

	if len(result) != DefaultSaltLength {
		t.Errorf("GenerateRandomString(-1) length = %d, want %d", len(result), DefaultSaltLength)
	}
}

func TestGenerateRandomString_CustomLength(t *testing.T) {
	tests := []int{1, 10, 50, 100, 500, 4096}

	for _, length := range tests {
		t.Run(string(rune(length)), func(t *testing.T) {
			result, err := GenerateRandomString(length)
			if err != nil {
				t.Fatalf("GenerateRandomString(%d) error = %v", length, err)
			}

			if len(result) != length {
				t.Errorf("GenerateRandomString(%d) length = %d, want %d", length, len(result), length)
			}
		})
	}
}

func TestGenerateRandomString_TooLong(t *testing.T) {
	_, err := GenerateRandomString(4097)
	if err == nil {
		t.Error("GenerateRandomString(4097) should return error")
	}

	expectedErr := "random string length too large (max 4096)"
	if err.Error() != expectedErr {
		t.Errorf("error = %q, want %q", err.Error(), expectedErr)
	}
}

func TestGenerateRandomString_MaxLength(t *testing.T) {
	result, err := GenerateRandomString(4096)
	if err != nil {
		t.Fatalf("GenerateRandomString(4096) error = %v", err)
	}

	if len(result) != 4096 {
		t.Errorf("GenerateRandomString(4096) length = %d, want 4096", len(result))
	}
}

func TestGenerateRandomString_CharacterSet(t *testing.T) {
	result, err := GenerateRandomString(4096)
	if err != nil {
		t.Fatalf("GenerateRandomString(4096) error = %v", err)
	}

	for i, r := range result {
		if !strings.ContainsRune(randomCharset, r) {
			t.Errorf("Character %d (%c) not in randomCharset", i, r)
		}
	}
}

func TestGenerateRandomString_Randomness(t *testing.T) {
	results := make([]string, 10)
	for i := 0; i < 10; i++ {
		var err error
		results[i], err = GenerateRandomString(32)
		if err != nil {
			t.Fatalf("GenerateRandomString(32) error = %v", err)
		}
	}

	seen := make(map[string]bool)
	for _, r := range results {
		if seen[r] {
			t.Error("GenerateRandomString() produced duplicate result")
		}
		seen[r] = true
	}
}

func TestGenerateRandomString_Distribution(t *testing.T) {
	result, err := GenerateRandomString(4096)
	if err != nil {
		t.Fatalf("GenerateRandomString(4096) error = %v", err)
	}

	lowerCount := 0
	upperCount := 0
	digitCount := 0

	for _, r := range result {
		switch {
		case r >= 'a' && r <= 'z':
			lowerCount++
		case r >= 'A' && r <= 'Z':
			upperCount++
		case r >= '0' && r <= '9':
			digitCount++
		}
	}

	if lowerCount == 0 {
		t.Error("Expected some lowercase letters")
	}
	if upperCount == 0 {
		t.Error("Expected some uppercase letters")
	}
	if digitCount == 0 {
		t.Error("Expected some digits")
	}
}

func TestDefaultSaltLength(t *testing.T) {
	if DefaultSaltLength != 32 {
		t.Errorf("DefaultSaltLength = %d, want 32", DefaultSaltLength)
	}
}

func TestRandomCharset(t *testing.T) {
	expectedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if randomCharset != expectedChars {
		t.Errorf("randomCharset = %q, want %q", randomCharset, expectedChars)
	}

	if len(randomCharset) != 62 {
		t.Errorf("randomCharset length = %d, want 62", len(randomCharset))
	}
}

func TestGenerateRandomString_SingleChar(t *testing.T) {
	result, err := GenerateRandomString(1)
	if err != nil {
		t.Fatalf("GenerateRandomString(1) error = %v", err)
	}

	if len(result) != 1 {
		t.Errorf("GenerateRandomString(1) length = %d, want 1", len(result))
	}

	if !strings.ContainsRune(randomCharset, rune(result[0])) {
		t.Errorf("Character %c not in randomCharset", result[0])
	}
}

func TestGenerateRandomString_Uniqueness(t *testing.T) {
	iterations := 100
	seen := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		result, err := GenerateRandomString(DefaultSaltLength)
		if err != nil {
			t.Fatalf("Iteration %d: error = %v", i, err)
		}

		if seen[result] {
			t.Errorf("Iteration %d: duplicate result detected", i)
		}
		seen[result] = true
	}
}
