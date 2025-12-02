# pass-gen

`pass-gen` is a command-line tool for generating deterministic passwords and random strings. It allows you to create reproducible passwords based on an input string and a salt, or generate secure random strings for various uses.

## Features

- **Deterministic Password Generation**: Generate the same password every time given the same input, salt, and configuration.
- **Random String Generation**: Create cryptographically secure random strings.
- **Configurable Security Levels**: Choose between `low`, `medium`, and `strong` complexity.
- **Salt Support**: Use a custom salt, a random salt, or an environment variable (`PASSGEN_SALT`).
- **Flexible Length**: Generate passwords from 1 to 4096 characters.

## Installation

### From Source

Ensure you have Go installed (1.24+).

```bash
go install github.com/zapsaang/pass-gen/cmd/passgen@latest
```

Or clone the repository and build:

```bash
git clone https://github.com/zapsaang/pass-gen.git
cd pass-gen
go build -o passgen ./cmd/passgen
```

## Usage

### Deterministic Mode (Default)

Generate a password based on an input string. This is useful for creating strong passwords that you don't need to memorize, as long as you remember the input and salt.

```bash
# Basic usage (default level: medium, default length: 64)
passgen -i "my-secret-input"

# Specify length and security level
passgen -i "my-secret-input" -l 16 -L strong

# Use a custom salt
passgen -i "my-secret-input" -s "my-salt"
```

### Random Salt

You can let the tool generate a random salt for you. **Important:** You must save the salt to recover the password later.

```bash
passgen -i "my-secret-input" --random-salt
```

### Random String Mode

Generate a completely random string (not deterministic).

```bash
# Generate a 32-character random string
passgen --gen-random -l 32
```

## Options

| Flag | Shorthand | Description | Default |
|------|-----------|-------------|---------|
| `--input` | `-i` | Input string (required for deterministic mode) | - |
| `--salt` | `-s` | Salt string (can also be set via `PASSGEN_SALT` env var) | - |
| `--random-salt` | | Generate a random salt automatically | `false` |
| `--gen-random` | | Generate a random string and exit | `false` |
| `--length` | `-l` | Password/String length | `64` |
| `--level` | `-L` | Security level (`low`, `medium`, `strong`) | `medium` |
| `--version` | | Print version information | - |
| `--help` | `-h` | Show help message | - |

## Security Levels

- **low**: Lowercase letters only (`a-z`).
- **medium**: Lowercase, uppercase, and digits (`a-z`, `A-Z`, `0-9`).
- **strong**: Lowercase, uppercase, digits, and special characters (`!@#%^&*()_=+[]{}:,.?-`).
