package main

import (
	"encoding/base64"
	"fmt"
	"testing"

	"git.innoai.tech/idp/auth_tools/apis/auth"
	"git.innoai.tech/idp/auth_tools/pkg/utils/crypto"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

type PublicRsa struct {
	KeyID     string `json:"keyID"`
	PublicKey string `json:"publicKey"`
}

func TestDataEncrypt(t *testing.T) {
	client := resty.New().SetBaseURL("http://srv-transmission-line---sentry.dev.innoai.tech")
	publicRes := PublicRsa{}
	_, err := client.
		R().SetResult(&publicRes).Get("api/transmission-line/sso/v0/auth/rsa-public-key/encrypt?secure=true")
	if err != nil {
		fmt.Printf("client.SetBaseURL err:%s\n", err.Error())
		return
	}

	pk, err := base64.StdEncoding.DecodeString(publicRes.PublicKey)
	str := "{\"name\": \"liuyanhao\", \"password\": \"Admin@123456\"}"
	pkey, err := crypto.ParseRSAPublicKeyFromPEM(pk)
	fmt.Println(pkey)
	require.NoError(t, err)
	data, err := pkey.Encrypt([]byte(str))
	require.NoError(t, err)
	body := fmt.Sprintf("%s %s", publicRes.KeyID, base64.StdEncoding.EncodeToString(data))
	fmt.Println(body)

	tokenResp := auth.TokenByPasswordResp{}
	resp, err := client.R().SetResult(&tokenResp).SetBody(body).Post("api/transmission-line/sso/v0/auth/token/password")
	fmt.Printf("%v", resp.RawResponse)
	if err != nil {

	}
}
