package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidPrivateKey    = errors.New("invalid private key")
	ErrUnknownSigningMethod = errors.New("unknown signing method")
)

// SigningMethod 签名方法
type SigningMethod string

const (
	// EdDSA ed25519
	SigningMethodEdDSA SigningMethod = "EdDSA"
	// RS256 rsa
	SigningMethodRS256 SigningMethod = "RS256"
	// ES256 ecdsa
	SigningMethodES256 SigningMethod = "ES256"
	// HS256 hmac
	SigningMethodHS256 SigningMethod = "HS256"
)

// Signer .
type Signer struct {
	signingMethod jwt.SigningMethod
	privateKey    interface{}
}

func newSigner(signingMethod SigningMethod, pkByte []byte) (*Signer, error) {
	var privateKey interface{}
	var err error
	var jwtSigningMethod jwt.SigningMethod

	switch signingMethod {
	case SigningMethodEdDSA:
		jwtSigningMethod = jwt.SigningMethodEdDSA
		privateKey, err = jwt.ParseEdPrivateKeyFromPEM(pkByte)

	case SigningMethodRS256:
		jwtSigningMethod = jwt.SigningMethodRS256
		privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(pkByte)

	case SigningMethodES256:
		jwtSigningMethod = jwt.SigningMethodES256
		privateKey, err = jwt.ParseECPrivateKeyFromPEM(pkByte)

	case SigningMethodHS256:
		jwtSigningMethod = jwt.SigningMethodHS256
		privateKey = pkByte

	default:
		return nil, ErrUnknownSigningMethod
	}

	if err != nil {
		return nil, ErrInvalidPrivateKey
	}

	return &Signer{
		signingMethod: jwtSigningMethod,
		privateKey:    privateKey,
	}, nil
}
