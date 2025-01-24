package token

import (
    "crypto/rand"
    "encoding/base64"
    "time"
)

// Token represents an API token
type Token struct {
    Value     string    `json:"value"`
    CreatedAt time.Time `json:"created_at"`
    ExpiresAt time.Time `json:"expires_at"`
}

// GenerateToken creates a new secure token with expiration
func GenerateToken(expirationDays int) (*Token, error) {
    // Generate 32 random bytes
    randomBytes := make([]byte, 32)
    if _, err := rand.Read(randomBytes); err != nil {
        return nil, err
    }

    // Encode to base64
    tokenValue := base64.URLEncoding.EncodeToString(randomBytes)

    now := time.Now()
    return &Token{
        Value:     tokenValue,
        CreatedAt: now,
        ExpiresAt: now.AddDate(0, 0, expirationDays),
    }, nil
}

// IsValid checks if the token is still valid
func (t *Token) IsValid() bool {
    return time.Now().Before(t.ExpiresAt)
}
