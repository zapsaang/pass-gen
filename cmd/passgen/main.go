package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zapsaang/pass-gen/pkg/passgen"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	versionFlag := flag.Bool("version", false, "Print version information")

	inputPtr := flag.String("input", "", "Input string (required)")
	inputShortPtr := flag.String("i", "", "Input string (shorthand)")

	saltPtr := flag.String("salt", "", "Salt string (optional)")
	saltShortPtr := flag.String("s", "", "Salt string (shorthand)")

	randomSaltPtr := flag.Bool("random-salt", false, "Generate a random salt automatically (overrides -s and ENV)")

	genRandomPtr := flag.Bool("gen-random", false, "Generate a standalone random string (Exclusive, but supports -l)")

	lengthPtr := flag.Int("length", 64, "Password/String length (1-4096)")
	lengthShortPtr := flag.Int("l", -1, "Length shorthand")

	levelPtr := flag.String("level", "medium", "Security level: low, medium, strong")
	levelShortPtr := flag.String("L", "", "Security level (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
		fmt.Println("Generate a deterministic password OR a random string")
		fmt.Println("\nModes:")
		fmt.Println("  1. Deterministic Mode (default): Requires -i/--input")
		fmt.Println("  2. Random String Mode: Use --gen-random (Supports -l)")
		fmt.Println("\nOptions:")
		fmt.Println("  -i, --input TEXT    Input string")
		fmt.Println("  --gen-random        Generate a random string and exit")
		fmt.Println("  -s, --salt TEXT     Salt string (optional)")
		fmt.Println("  --random-salt       Generate a random salt for the password")
		fmt.Println("  -l, --length NUM    Length (default: 64)")
		fmt.Println("  -L, --level LEVEL   Security level (default: medium)")
		fmt.Println("  -h, --help          Show this help message")
	}

	flag.Parse()

	if *versionFlag {
		fmt.Printf("passgen %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Built:  %s\n", date)
		os.Exit(0)
	}

	length := *lengthPtr
	if *lengthShortPtr != -1 {
		length = *lengthShortPtr
	}

	if *genRandomPtr {
		conflict := false
		flag.Visit(func(f *flag.Flag) {
			name := f.Name
			if name != "gen-random" && name != "length" && name != "l" {
				conflict = true
			}
		})

		if conflict {
			fmt.Fprintln(os.Stderr, "Error: --gen-random can only be used with -l or --length")
			os.Exit(1)
		}

		randStr, err := passgen.GenerateRandomString(length)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(randStr)
		os.Exit(0)
	}

	input := *inputPtr
	if input == "" {
		input = *inputShortPtr
	}
	if input == "" {
		fmt.Fprintln(os.Stderr, "Error: input is required (-i or --input)")
		os.Exit(1)
	}

	salt := ""
	isRandomSalt := *randomSaltPtr

	if isRandomSalt {
		var err error
		salt, err = passgen.GenerateRandomString(32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating random salt: %v\n", err)
			os.Exit(1)
		}
	} else {
		salt = *saltPtr
		if salt == "" {
			salt = *saltShortPtr
		}
		if salt == "" {
			salt = os.Getenv("PASSGEN_SALT")
		}
	}

	level := *levelPtr
	if *levelShortPtr != "" {
		level = *levelShortPtr
	}

	config := passgen.Config{
		Input:  input,
		Salt:   salt,
		Length: length,
		Level:  passgen.Level(level),
	}

	password, err := passgen.Generate(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if isRandomSalt {
		fmt.Println("--------------------------------------------------")
		fmt.Printf("Salt:     %s\n", salt)
		fmt.Printf("Password: %s\n", password)
		fmt.Println("--------------------------------------------------")
		fmt.Println("IMPORTANT: Save the Salt! It is required to recover this password.")
	} else {
		fmt.Println(password)
	}
}
