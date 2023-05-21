package jwt

import (
	"fmt"
	"net/http"

	"github.com/dmitrymomot/go-utils"
	"github.com/dmitrymomot/itm-api/pkg/httpx"
	"github.com/go-chi/jwtauth"
)

// GetUserID returns user ID from the request context.
func GetUserID(r *http.Request) (string, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return "", fmt.Errorf("%w: failed to get token from context: %s", httpx.ErrUnauthorized, err)
	}

	uid, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("%w: failed to get user ID from token", httpx.ErrUnauthorized)
	}

	return uid, nil
}

// GetClaims returns claims from the request context.
func GetClaims(r *http.Request) (TokenClaims, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return TokenClaims{}, fmt.Errorf("%w: failed to get token from context: %s", httpx.ErrUnauthorized, err)
	}

	var tokenClaims TokenClaims
	if err := utils.MapToStruct(claims, &tokenClaims, "json"); err != nil {
		return TokenClaims{}, fmt.Errorf("%w: failed to map claims to struct: %s", httpx.ErrUnauthorized, err)
	}

	return tokenClaims, nil
}
