package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/dmitrymomot/go-utils"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

type Permission string

// Permission constants.
const (
	PermissionReadOnly       Permission = "ro"
	PermissionReadAndComment Permission = "rc"
	PermissionReadWrite      Permission = "rw"
	PermissionRWDelete       Permission = "rwd"
	PermissionFull           Permission = "full"
)

// TokenClaims struct with user_id and standard claims.
type (
	TokenClaims struct {
		ProjectID string `json:"project_id,omitempty"`
		IdeaID    string `json:"idea_id,omitempty"`
		FeatureID string `json:"feature_id,omitempty"`
		TaskID    string `json:"task_id,omitempty"`
		ReleaseID string `json:"release_id,omitempty"`

		UserID string `json:"user_id,omitempty"`
		Email  string `json:"email,omitempty"`

		Permissions Permission             `json:"perm,omitempty"`
		ExpiresIn   time.Duration          `json:"-"`
		Meta        map[string]interface{} `json:"meta,omitempty"`
	}

	// Token struct with token, expires_in and expires_at.
	Token struct {
		Token     string `json:"token"`
		ExpiresAt int64  `json:"expires_at"`
	}

	// TokenEncoder function.
	TokenEncoder func(ctx context.Context, claims TokenClaims) (Token, error)
)

// GenerateToken generates a new JWT token with the given claims.
func GenerateToken(tokenAuth *jwtauth.JWTAuth) TokenEncoder {
	return func(ctx context.Context, claims TokenClaims) (Token, error) {
		claimsMap, err := utils.StructToMap(claims, "json")
		if err != nil {
			return Token{}, fmt.Errorf("failed to add claims to token: %w", err)
		}
		claimsMap["exp"] = time.Now().Add(claims.ExpiresIn).Unix()
		claimsMap["iat"] = time.Now().Unix()
		claimsMap["nbf"] = time.Now().Unix()
		claimsMap["jti"] = uuid.New().String()

		token, tokenString, err := tokenAuth.Encode(claimsMap)
		if err != nil {
			return Token{}, err
		}

		return Token{
			Token:     tokenString,
			ExpiresAt: token.Expiration().Unix(),
		}, nil
	}
}
