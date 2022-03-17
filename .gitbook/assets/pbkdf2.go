package hash

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	keyLen = 32
)

type pbkdf2Hasher struct{}

func (pbkdf2Hasher) Name() string {
	return PBKDF2SHA256Django
}

func (pbkdf2Hasher) Compare(password, encoded string) (bool, error) {
	parts := strings.SplitN(encoded, "$", 4)
	if len(parts) != 4 {
		return false, errors.New("pbkdf2: invalid encoded value, it must consist of 4 segments")
	}

	iterations, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("pbkdf2: wrong number of iterations %s: %w", parts[1], err)
	}

	salt := []byte(parts[2])

	expectedHash, err := base64.StdEncoding.DecodeString(parts[3])
	if err != nil {
		return false, fmt.Errorf("pbkdf2: wrong hash encoding for %s: %w", parts[3], err)
	}

	actualHash := pbkdf2.Key([]byte(password), salt, iterations, keyLen, sha256.New)
	return bytes.Equal(actualHash, expectedHash), nil
}

// Validate will check whether the given value is a valid
// hash for PBKDF2 SHA256 Django.
// A valid hash for this variant of PBKDF2 has the following
// form:
// <algorithm>$<iterations>$<salt>$<hash>
// where:
// 1. `algorithm` has value `pbkdf2_sha256`
// 2. `iterations` is the number of iterations the algorithm
//    has to run
// 3. `salt` is the salt string that was used to generate the
//    hash
// 4. `hash` is the bas64 encoded hash
//
// For more information, check the official Django documentation:
// https://docs.djangoproject.com/en/4.0/topics/auth/passwords/
func (pbkdf2Hasher) Validate(encoded string) bool {
	parts := strings.Split(encoded, "$")
	if len(parts) != 4 {
		return false
	}

	if parts[0] != "pbkdf2_sha256" {
		return false
	}

	// iterations should be a number
	_, err := strconv.Atoi(parts[1])
	return err == nil
}
