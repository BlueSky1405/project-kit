package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWTTokenRS256(t *testing.T) {
	signingMethod := SigningMethodRS256
	privateKey := []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCh5Nk2GLiyQFMIU+h3OEA4UeFbu3dCH5sjd/sLTxxvwjXq7JLq
Jbt2rCIdzpAXOi4jL+FRGQnHaxUlHUBZsojnCcHvhrz2knV6rXNogt0emL7f7ZMR
o8IsQGV8mlKIC9xLnlOQQdRNUssmrROrCG99wpTRRNZjOmLvkcoXdeuaCQIDAQAB
AoGAUTcJ1H6QYTOts9bMHsrERLymzir8R9qtLBzrfp/gRxxpigHGLdph8cWmk8dl
N5HDRXmmkdV6t2S7xdOnzZen31lcWe0bIzg0SrFiUEOtg3Lwxzw2Pz0dKwg4ZUoo
GKpcIU6kEpbC2UkjBV4+2E6P1DXuhdgTyHoUA3ycxOdjCAUCQQCyjTzGPXFoHq5T
miJyVd4VXNyCXGU0ZuQayt6nPN8Gd5CcEb2S4kggzPXQcd90FO0kHfZV6+PGTrc2
ZUuz5uwPAkEA6B3lmEmiZsJS/decLzWR0T1CXaFGwTjBQbHXJ0RziAfkuy+VwSmh
vrW/ipk5xbREr5rKx3jVI2PzVOvLw7NgZwJAbUsvDFnH9WfyZZJPy5TsID97awCL
oovozM2phM0p55eAmUfyttp0ND/BqBpMIY49qoH8q5N9FYJRe6Z9tF2B2QJAQBEo
cw039xcB4zCk2l713YQEEmXWarSomuJkWWFKZiyPlJ8Ava0pCMOPl8jNKmWkY7fc
6ovOgJMw8aqXtm+HVwJAerJeUEDez2djG5pIF6aCV0bP3fhQUq8OQCgGF5Qzo9Cn
qvYreGpYKPJGVixAsEPCiLzJRhy1XfFona6VRXIIxw==
-----END RSA PRIVATE KEY-----
`)
	publicKey := []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCh5Nk2GLiyQFMIU+h3OEA4UeFb
u3dCH5sjd/sLTxxvwjXq7JLqJbt2rCIdzpAXOi4jL+FRGQnHaxUlHUBZsojnCcHv
hrz2knV6rXNogt0emL7f7ZMRo8IsQGV8mlKIC9xLnlOQQdRNUssmrROrCG99wpTR
RNZjOmLvkcoXdeuaCQIDAQAB
-----END PUBLIC KEY-----
`)

	tokenGen, err := NewTokenGenerator(signingMethod, privateKey, WithExpires(time.Hour), WithProduct("financial"))
	require.NoError(t, err, "NewTokenGenerator")

	tokenStr, err := tokenGen.Generate(TokenInfo{UserID: 2, RoleID: 3})
	require.NoError(t, err, "Generate token")
	fmt.Println(tokenStr)

	verifier, err := NewTokenVerifier(signingMethod, publicKey)
	require.NoError(t, err, "NewTokenVerifier")

	info, err := verifier.Verify(tokenStr)
	require.NoError(t, err, "Verify")

	assert.Equal(t, uint(2), info.UserID)
	assert.Equal(t, uint(3), info.RoleID)
}
