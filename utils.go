package zoom

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HMAC256(text, salt string) (string, error) {
	h := hmac.New(sha256.New, []byte(salt))
	if _, err := h.Write([]byte(text)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
