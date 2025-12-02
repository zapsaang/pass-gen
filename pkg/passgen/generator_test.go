package passgen

import (
	"math/rand"
	"strings"
	"testing"
	"unicode"
)

func TestGenerate_ValidLevels(t *testing.T) {
	tests := []struct {
		name     string
		level    Level
		length   int
		wantLow  bool
		wantUp   bool
		wantDig  bool
		wantSpec bool
	}{
		{
			name:     "LevelLow with length 8",
			level:    LevelLow,
			length:   8,
			wantLow:  true,
			wantUp:   false,
			wantDig:  false,
			wantSpec: false,
		},
		{
			name:     "LevelMedium with length 12",
			level:    LevelMedium,
			length:   12,
			wantLow:  true,
			wantUp:   true,
			wantDig:  true,
			wantSpec: false,
		},
		{
			name:     "LevelStrong with length 16",
			level:    LevelStrong,
			length:   16,
			wantLow:  true,
			wantUp:   true,
			wantDig:  true,
			wantSpec: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{
				Input:  "testinput",
				Salt:   "testsalt",
				Length: tt.length,
				Level:  tt.level,
			}

			result, err := Generate(cfg)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			if len(result) != tt.length {
				t.Errorf("Generate() length = %d, want %d", len(result), tt.length)
			}

			hasLower := containsAny(result, charsLower)
			hasUpper := containsAny(result, charsUpper)
			hasDigit := containsAny(result, charsDigits)
			hasSpecial := containsAny(result, charsSpecial)

			if tt.wantLow && !hasLower {
				t.Errorf("Generate() should contain lowercase letters")
			}
			if tt.wantUp && !hasUpper {
				t.Errorf("Generate() should contain uppercase letters")
			}
			if tt.wantDig && !hasDigit {
				t.Errorf("Generate() should contain digits")
			}
			if tt.wantSpec && !hasSpecial {
				t.Errorf("Generate() should contain special characters")
			}

			if tt.level == LevelLow {
				for _, r := range result {
					if !unicode.IsLower(r) {
						t.Errorf("LevelLow should only contain lowercase, got %c", r)
					}
				}
			}
		})
	}
}

func TestGenerate_Deterministic(t *testing.T) {
	cfg := Config{
		Input:  "myinput",
		Salt:   "mysalt",
		Length: 20,
		Level:  LevelStrong,
	}

	result1, err := Generate(cfg)
	if err != nil {
		t.Fatalf("First Generate() error = %v", err)
	}

	result2, err := Generate(cfg)
	if err != nil {
		t.Fatalf("Second Generate() error = %v", err)
	}

	if result1 != result2 {
		t.Errorf("Generate() not deterministic: got %q and %q", result1, result2)
	}
}

func TestGenerate_RandomDeterministic(t *testing.T) {
	cfg := Config{
		Input:  "myinput",
		Salt:   "mysalt",
		Length: 20,
		Level:  LevelStrong,
	}

	levels := []Level{LevelLow, LevelMedium, LevelStrong}
	for range 40960 {
		cfg.Length = rand.Intn(4096) + 1
		cfg.Level = levels[rand.Intn(len(levels))]
		cfg.Salt = "salt" + string(rune(rand.Intn(1000)))
		cfg.Input = "input" + string(rune(rand.Intn(1000)))
		result1, err := Generate(cfg)
		if err != nil {
			t.Fatalf("First Generate() error = %v", err)
		}

		result2, err := Generate(cfg)
		if err != nil {
			t.Fatalf("Second Generate() error = %v", err)
		}

		if result1 != result2 {
			t.Errorf("Generate() not deterministic: got %q and %q", result1, result2)
		}
	}

}

func TestGenerate_DifferentInputsProduceDifferentPasswords(t *testing.T) {
	base := Config{
		Input:  "input1",
		Salt:   "salt1",
		Length: 16,
		Level:  LevelStrong,
	}

	result1, _ := Generate(base)

	cfg2 := base
	cfg2.Input = "input2"
	result2, _ := Generate(cfg2)

	cfg3 := base
	cfg3.Salt = "salt2"
	result3, _ := Generate(cfg3)

	cfg4 := base
	cfg4.Level = LevelMedium
	result4, _ := Generate(cfg4)

	cfg5 := base
	cfg5.Length = 20
	result5, _ := Generate(cfg5)

	if result1 == result2 {
		t.Error("Different inputs should produce different passwords")
	}
	if result1 == result3 {
		t.Error("Different salts should produce different passwords")
	}
	if result1 == result4 {
		t.Error("Different levels should produce different passwords")
	}
	if result1 == result5 {
		t.Error("Different lengths should produce different passwords")
	}
}

func TestGenerate_EmptyInput(t *testing.T) {
	cfg := Config{
		Input:  "",
		Salt:   "salt",
		Length: 16,
		Level:  LevelStrong,
	}

	_, err := Generate(cfg)
	if err == nil {
		t.Error("Generate() should return error for empty input")
	}
	if err.Error() != "input is required" {
		t.Errorf("Generate() error = %v, want 'input is required'", err)
	}
}

func TestGenerate_InputTooLong(t *testing.T) {
	cfg := Config{
		Input:  strings.Repeat("a", 1001),
		Salt:   "salt",
		Length: 16,
		Level:  LevelStrong,
	}

	_, err := Generate(cfg)
	if err == nil {
		t.Error("Generate() should return error for input too long")
	}
	if err.Error() != "input too long" {
		t.Errorf("Generate() error = %v, want 'input too long'", err)
	}
}

func TestGenerate_InputAtMaxLength(t *testing.T) {
	cfg := Config{
		Input:  strings.Repeat("a", 1000),
		Salt:   "salt",
		Length: 16,
		Level:  LevelStrong,
	}

	result, err := Generate(cfg)
	if err != nil {
		t.Errorf("Generate() should not return error for input at max length, got %v", err)
	}
	if len(result) != 16 {
		t.Errorf("Generate() length = %d, want 16", len(result))
	}
}

func TestGenerate_InvalidLength(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"zero length", 0},
		{"negative length", -1},
		{"exceeds max", 4097},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{
				Input:  "input",
				Salt:   "salt",
				Length: tt.length,
				Level:  LevelStrong,
			}

			_, err := Generate(cfg)
			if err == nil {
				t.Error("Generate() should return error for invalid length")
			}
		})
	}
}

func TestGenerate_VariedLength(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"128", 128},
		{"512", 512},
		{"1024", 1024},
		{"2048", 2048},
		{"4096", 4096},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{
				Input:  "input",
				Salt:   "salt",
				Length: tt.length,
				Level:  LevelStrong,
			}

			password, err := Generate(cfg)
			if err != nil {
				t.Error("Generate() should not return error for valid length")
			}
			if len(password) != tt.length {
				t.Errorf("Generate() length = %d, want %d", len(password), tt.length)
			}
		})
	}
}

func TestGenerate_LengthAtMax(t *testing.T) {
	cfg := Config{
		Input:  "input",
		Salt:   "salt",
		Length: 4096,
		Level:  LevelStrong,
	}

	result, err := Generate(cfg)
	if err != nil {
		t.Errorf("Generate() should not return error for max length, got %v", err)
	}
	if len(result) != 4096 {
		t.Errorf("Generate() length = %d, want 4096", len(result))
	}
}

func TestGenerate_InvalidLevel(t *testing.T) {
	cfg := Config{
		Input:  "input",
		Salt:   "salt",
		Length: 16,
		Level:  "invalid",
	}

	_, err := Generate(cfg)
	if err == nil {
		t.Error("Generate() should return error for invalid level")
	}
	if err.Error() != "invalid level" {
		t.Errorf("Generate() error = %v, want 'invalid level'", err)
	}
}

func TestGenerate_EmptySalt(t *testing.T) {
	cfg := Config{
		Input:  "input",
		Salt:   "",
		Length: 16,
		Level:  LevelStrong,
	}

	result, err := Generate(cfg)
	if err != nil {
		t.Errorf("Generate() should not return error for empty salt, got %v", err)
	}
	if len(result) != 16 {
		t.Errorf("Generate() length = %d, want 16", len(result))
	}
}

func TestGenerate_MinimumLength(t *testing.T) {
	tests := []struct {
		name   string
		level  Level
		length int
	}{
		{"LevelLow length 1", LevelLow, 1},
		{"LevelMedium length 1", LevelMedium, 1},
		{"LevelMedium length 3", LevelMedium, 3},
		{"LevelStrong length 1", LevelStrong, 1},
		{"LevelStrong length 4", LevelStrong, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Config{
				Input:  "input",
				Salt:   "salt",
				Length: tt.length,
				Level:  tt.level,
			}

			result, err := Generate(cfg)
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}
			if len(result) != tt.length {
				t.Errorf("Generate() length = %d, want %d", len(result), tt.length)
			}
		})
	}
}

func TestGenerate_CharacterDistribution(t *testing.T) {
	cfg := Config{
		Input:  "distribution_test",
		Salt:   "salt",
		Length: 1000,
		Level:  LevelStrong,
	}

	result, err := Generate(cfg)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	countLower := 0
	countUpper := 0
	countDigit := 0
	countSpecial := 0

	for _, r := range result {
		switch {
		case unicode.IsLower(r):
			countLower++
		case unicode.IsUpper(r):
			countUpper++
		case unicode.IsDigit(r):
			countDigit++
		default:
			countSpecial++
		}
	}

	if countLower == 0 {
		t.Error("Expected some lowercase letters")
	}
	if countUpper == 0 {
		t.Error("Expected some uppercase letters")
	}
	if countDigit == 0 {
		t.Error("Expected some digits")
	}
	if countSpecial == 0 {
		t.Error("Expected some special characters")
	}
}

func TestGenerate_UnicodeSafeInput(t *testing.T) {
	cfg := Config{
		Input:  "测试输入🔐",
		Salt:   "盐值",
		Length: 16,
		Level:  LevelStrong,
	}

	result, err := Generate(cfg)
	if err != nil {
		t.Errorf("Generate() should handle unicode input, got error: %v", err)
	}
	if len(result) != 16 {
		t.Errorf("Generate() length = %d, want 16", len(result))
	}
}

func TestLevelConstants(t *testing.T) {
	if LevelLow != "low" {
		t.Errorf("LevelLow = %q, want 'low'", LevelLow)
	}
	if LevelMedium != "medium" {
		t.Errorf("LevelMedium = %q, want 'medium'", LevelMedium)
	}
	if LevelStrong != "strong" {
		t.Errorf("LevelStrong = %q, want 'strong'", LevelStrong)
	}
}

func TestCharacterSets(t *testing.T) {
	if len(charsLower) != 26 {
		t.Errorf("charsLower length = %d, want 26", len(charsLower))
	}
	if len(charsUpper) != 26 {
		t.Errorf("charsUpper length = %d, want 26", len(charsUpper))
	}
	if len(charsDigits) != 10 {
		t.Errorf("charsDigits length = %d, want 10", len(charsDigits))
	}
	if len(charsSpecial) == 0 {
		t.Error("charsSpecial should not be empty")
	}

	if string(runesLower) != charsLower {
		t.Error("runesLower does not match charsLower")
	}
	if string(runesUpper) != charsUpper {
		t.Error("runesUpper does not match charsUpper")
	}
	if string(runesDigits) != charsDigits {
		t.Error("runesDigits does not match charsDigits")
	}
	if string(runesSpecial) != charsSpecial {
		t.Error("runesSpecial does not match charsSpecial")
	}
}

func containsAny(s, chars string) bool {
	for _, c := range chars {
		if strings.ContainsRune(s, c) {
			return true
		}
	}
	return false
}
