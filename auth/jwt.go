package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
)

var jwtKey = []byte("YV9zZWNyZXRfa2V5")

func ValidateToken(signedToken string) (uint, error) {
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		signedToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("Invalid Token. %v", err.Error())
	}

	userIdFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid user_id")
	}

	return uint(userIdFloat), nil
}

func Auth(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Missing auth token", http.StatusUnauthorized)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	userId, err := ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	w.Header().Set("X-Requester-ID", strconv.Itoa(int(userId)))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authenticated"))
}
