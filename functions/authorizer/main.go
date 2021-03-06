package main

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	jwt "github.com/gbrlsnchs/jwt/v2"
)

func generatePolicy(principalID string, effect string, isGuest bool, context map[string]interface{}) events.APIGatewayCustomAuthorizerResponse {
	authReponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}

	if effect != "" {
		if isGuest {
			authReponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
				Version: "2012-10-17",
				Statement: []events.IAMPolicyStatement{
					{
						Action: []string{"execute-api:Invoke"},
						Effect: effect,
						Resource: []string{
							"arn:aws:execute-api:*:*:*/*/GET/collections/*",
						},
					},
				},
			}
		} else {
			authReponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
				Version: "2012-10-17",
				Statement: []events.IAMPolicyStatement{
					{
						Action:   []string{"execute-api:Invoke"},
						Effect:   effect,
						Resource: []string{"arn:aws:execute-api:*:*:*/*/*"},
					},
				},
			}
		}
	}

	authReponse.Context = context

	return authReponse
}

func verify(token string, keyEncoded string) (string, error) {
	fmt.Println(token, keyEncoded)
	block, _ := pem.Decode([]byte(keyEncoded))
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	es256 := jwt.NewES256(privateKey, &privateKey.PublicKey)

	decoded, sig, err := jwt.Parse(token)

	if err != nil {
		return "", err
	}
	if err = es256.Verify(decoded, sig); err != nil {
		return "", err
	}

	payloadEncoded := strings.Split(string(decoded), ".")[1]
	decoded64, err := base64.StdEncoding.DecodeString(payloadEncoded)
	if err != nil {
		return "", err
	}

	return string(decoded64), nil
}

func handler(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	fmt.Println(event)
	// Guest authorization
	if !strings.HasPrefix(event.AuthorizationToken, "Bearer ") || event.AuthorizationToken == "Bearer null" {
		return generatePolicy("guest", "Allow", true, map[string]interface{}{"id": "guest"}), nil
	}

	token := strings.Replace(event.AuthorizationToken, "Bearer ", "", 1)
	payload, err := verify(token, os.Getenv("JwtPrivate"))
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	var data interface{}
	json.Unmarshal([]byte(payload), &data)

	isAllowed := true

	var effect string
	if isAllowed {
		effect = "Allow"
	} else {
		effect = "Deny"
	}
	userID := data.(map[string]interface{})["id"].(string)

	return generatePolicy(userID, effect, false, data.(map[string]interface{})), nil
}

func main() {
	lambda.Start(handler)
}
