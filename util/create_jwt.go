package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         string `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := base64UrlEncode(byteArrHeader)
	byteArrdata, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := base64UrlEncode(byteArrdata)
	message := headerB64 + "." + payloadB64

	byteArrsecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrsecret)
	h.Write(byteArrMessage)
	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)
	jwt := message + "." + signatureB64

	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}
