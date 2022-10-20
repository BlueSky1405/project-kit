package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Option is a token generator option
type Option func(*TokenGenerator)

// WithProduct set product
func WithProduct(product string) Option {
	return func(t *TokenGenerator) {
		t.product = product
	}
}

// WithExpires set expires for token
func WithExpires(expires time.Duration) Option {
	return func(t *TokenGenerator) {
		t.expires = expires
	}
}

// TokenGenerator token generator
type TokenGenerator struct {
	signer  *Signer
	expires time.Duration
	product string
}

// NewTokenGenerator .
func NewTokenGenerator(signingMethod SigningMethod, privateKey []byte, opts ...Option) (*TokenGenerator, error) {
	signer, err := newSigner(signingMethod, privateKey)
	if err != nil {
		return nil, err
	}

	g := &TokenGenerator{
		signer:  signer,
		expires: 12*time.Hour, // default 12 hour
		product: "default", // default product
	}

	for _, opt := range opts {
		opt(g)
	}

	return g, nil
}

// Generate generate token
func (g *TokenGenerator) Generate(info TokenInfo) (string, error) {

	claims := TokenClaims{
		info,
		jwt.RegisteredClaims{
			Issuer:    g.product,
			Audience:  []string{},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(g.expires)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return jwt.NewWithClaims(g.signer.signingMethod, claims).SignedString(g.signer.privateKey)
}
