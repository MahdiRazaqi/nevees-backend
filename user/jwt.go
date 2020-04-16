package user

import (
	"net/http"
	"strings"
	"time"

	"github.com/MahdiRazaqi/nevees-backend/config"
	"github.com/dgrijalva/jwt-go"
)

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var signature = config.CFG.JWT.SigningKey

// CreateToken generate new token
func (u *User) CreateToken() (string, error) {
	claims := new(customClaims)
	claims.Username = u.Username
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signature))
}

// GetToken get token from header
func GetToken(req *http.Request) string {
	token := req.Header.Get("Authorization")
	return strings.Replace(token, "Bearer ", "", -1)
}

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signature), nil
	})
}

// // LoadByToken load user from token
// func LoadByToken(token string) (*User, error) {
// 	t, _ := parseToken(token)

// 	claims, ok := t.Claims.(*customClaims)
// 	if !ok {
// 		return nil, errors.New("converting token failed")
// 	}
// 	return LoadByUsername(claims.Username)
// }
