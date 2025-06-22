package ctxdata

import "github.com/golang-jwt/jwt"

// Identify is the custom claim key for storing user ID (issuer identity)
// You can customize this key as needed.
const Identify = "dailaoer"

// GetJwtToken generates a signed JWT token using HS256 algorithm.
//
// Parameters:
// - secretKey: The secret key used to sign the token.
// - iat: The issued-at timestamp (Unix time).
// - seconds: The duration in seconds that the token is valid (used to calculate exp).
// - uid: The user ID or identifier to embed in the token (custom claim).
//
// Returns:
// - A signed JWT token string.
// - An error if signing fails.
func GetJwtToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	// Create the claims map with standard and custom claims
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds // Expiration time
	claims["iat"] = iat           // Issued-at time
	claims[Identify] = uid        // Custom claim: user ID

	// Create a new JWT token with HS256 signing method
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(secretKey))
}
