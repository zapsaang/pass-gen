package passgen

import (
	"errors"
	"strconv"
)

type Level string

const (
	LevelLow    Level = "low"
	LevelMedium Level = "medium"
	LevelStrong Level = "strong"
)

const (
	charsLower   = "abcdefghijklmnopqrstuvwxyz"
	charsUpper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsDigits  = "0123456789"
	charsSpecial = "!@#%^&*()_=+[]{}:,.?-"
)

var (
	runesLower   = []rune(charsLower)
	runesUpper   = []rune(charsUpper)
	runesDigits  = []rune(charsDigits)
	runesSpecial = []rune(charsSpecial)
)

type Config struct {
	Input  string
	Salt   string
	Length int
	Level  Level
}

func Generate(cfg Config) (string, error) {
	if cfg.Input == "" {
		return "", errors.New("input is required")
	}
	if len(cfg.Input) > 1000 {
		return "", errors.New("input too long")
	}
	if cfg.Length <= 0 || cfg.Length > 4096 {
		return "", errors.New("length must be positive and not exceed 4096")
	}

	rng := newDetermRNG(cfg.Salt + cfg.Input + string(cfg.Level) + strconv.Itoa(cfg.Length))

	var requiredPools [][]rune
	var allChars []rune

	switch cfg.Level {
	case LevelLow:
		requiredPools = [][]rune{runesLower}
		allChars = runesLower
	case LevelMedium:
		requiredPools = [][]rune{runesLower, runesUpper, runesDigits}
		allChars = make([]rune, 0, len(runesLower)+len(runesUpper)+len(runesDigits))
		allChars = append(allChars, runesLower...)
		allChars = append(allChars, runesUpper...)
		allChars = append(allChars, runesDigits...)
	case LevelStrong:
		requiredPools = [][]rune{runesLower, runesUpper, runesDigits, runesSpecial}
		allChars = make([]rune, 0, len(runesLower)+len(runesUpper)+len(runesDigits)+len(runesSpecial))
		allChars = append(allChars, runesLower...)
		allChars = append(allChars, runesUpper...)
		allChars = append(allChars, runesDigits...)
		allChars = append(allChars, runesSpecial...)
	default:
		return "", errors.New("invalid level")
	}

	passwordRunes := make([]rune, 0, cfg.Length)

	for _, pool := range requiredPools {
		if len(passwordRunes) >= cfg.Length {
			break
		}
		idx := rng.Intn(len(pool))
		passwordRunes = append(passwordRunes, pool[idx])
	}

	for len(passwordRunes) < cfg.Length {
		idx := rng.Intn(len(allChars))
		passwordRunes = append(passwordRunes, allChars[idx])
	}

	for i := len(passwordRunes) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		passwordRunes[i], passwordRunes[j] = passwordRunes[j], passwordRunes[i]
	}

	return string(passwordRunes), nil
}
