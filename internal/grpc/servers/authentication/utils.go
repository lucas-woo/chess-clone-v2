package authenticationgrpc

import (
	"crypto/rand"
	"encoding/base64"
)

func generateSessionId() (string, error) {

	temp := make([]byte, 32)

	if _, err := rand.Read(temp); err != nil {
		return "", nil;
	}
	return base64.URLEncoding.EncodeToString(temp), nil
}