package main

import (
	"encoding/base64"
	"fmt"

	auth2 "git.innoai.tech/idp/auth_tools/apis/auth"
	"git.innoai.tech/idp/auth_tools/pkg/utils/crypto"
	"github.com/go-resty/resty/v2"
)

type RsaPublicKey struct {
	KeyID     string `json:"keyID"`
	PublicKey string `json:"publicKey"`
}

func main() {
	client := resty.New().SetBaseURL("http://srv-transmission-line---sentry.dev.innoai.tech")
	rsaPublic := RsaPublicKey{}
	_, err := client.
		R().SetResult(&rsaPublic).Get("api/transmission-line/sso/v0/auth/rsa-public-key/encrypt?secure=true")
	if err != nil {
		fmt.Printf("Get rsaPublic err:%s\n", err.Error())
		return
	}

	pk, err := base64.StdEncoding.DecodeString(rsaPublic.PublicKey)
	account := "{\"name\": \"liuyanhao\", \"password\": \"Admin@123456\"}"
	pkey, err := crypto.ParseRSAPublicKeyFromPEM(pk)
	data, err := pkey.Encrypt([]byte(account))
	body := fmt.Sprintf("%s %s", rsaPublic.KeyID, base64.StdEncoding.EncodeToString(data))

	tokenResp := auth2.TokenByPasswordResp{}
	_, err = client.R().SetResult(&tokenResp).SetBody(body).Post("api/transmission-line/sso/v0/auth/token/password")
	if err != nil {
		fmt.Printf("Get token err:%s\n", err.Error())
	}
}
