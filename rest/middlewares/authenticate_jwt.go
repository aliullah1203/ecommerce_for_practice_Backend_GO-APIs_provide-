package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse jwt
		// parse header and payload or claims
		// hmac-sha-256 algorithm -> hash hmac(header, payload, secret key)
		// parse signature part from the jwt
		// if the signature and hash is same => forward to create products
		// otherwise 401 status code with unauthorized

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		heareArr := strings.Split(header, " ")

		if len(heareArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := heareArr[1]
		fmt.Println("------accessToken:------", accessToken)

		tokenPart := strings.Split(accessToken, ".")

		if len(tokenPart) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenPart[0]
		jwtPayload := tokenPart[1]
		Signature := tokenPart[2]

		fmt.Println(jwtHeader)
		fmt.Println(jwtPayload)
		fmt.Println(Signature)

		message := jwtHeader + "." + jwtPayload

		//cnf := config.GetConfig()

		byteArrsecret := []byte(m.cnf.JwtSecretkey)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrsecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := base64UrlEncode(hash)

		if newSignature != Signature {
			http.Error(w, "Unauthorized, tui hacker!", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func base64UrlEncode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}
