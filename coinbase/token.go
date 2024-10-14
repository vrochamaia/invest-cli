package coinbase

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"investcli/utils"
	"math"
	"math/big"
	"time"

	"gopkg.in/go-jose/go-jose.v2"
	"gopkg.in/go-jose/go-jose.v2/jwt"
)

type APIKeyClaims struct {
	*jwt.Claims
	URI string `json:"uri"`
}

type coingbaseApiKey struct {
	ApiKeys utils.ApiKeys `json:"coinbase"`
}

func buildJWT(uri string) (string, error) {
	apiKey := utils.GetDataFromJson[coingbaseApiKey]("./secrets.json").ApiKeys

	block, _ := pem.Decode([]byte(apiKey.PrivateKey))
	if block == nil {
		return "", fmt.Errorf("jwt: Could not decode private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	sig, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.ES256, Key: key},
		(&jose.SignerOptions{NonceSource: nonceSource{}}).WithType("JWT").WithHeader("kid", apiKey.Key),
	)
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}

	cl := &APIKeyClaims{
		Claims: &jwt.Claims{
			Subject:   apiKey.Key,
			Issuer:    "coinbase-cloud",
			NotBefore: jwt.NewNumericDate(time.Now()),
			Expiry:    jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
		},
		URI: uri,
	}
	jwtString, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", fmt.Errorf("jwt: %w", err)
	}
	return jwtString, nil
}

var max = big.NewInt(math.MaxInt64)

type nonceSource struct{}

func (n nonceSource) Nonce() (string, error) {
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return r.String(), nil
}
