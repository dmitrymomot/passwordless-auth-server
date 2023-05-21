package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Bind binds json request body to struct.
//
// Example:
//
//	var u user.User
//	if err := BindJSON(r, &u); err != nil {
//		return err
//	}
func BindJSON(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("failed to decode request body: %w", err)
	}

	return nil
}
