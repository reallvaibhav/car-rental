package auth

import (
  "github.com/golang-jwt/jwt/v4"
  "time"
)

type JWTManager struct {
  secret string
}

func New(secret string) *JWTManager {
  return &JWTManager{secret: secret}
}

func (j *JWTManager) Generate(userID string) (string, error) {
  claims := jwt.MapClaims{
    "user_id": userID,
    "exp":     time.Now().Add(24 * time.Hour).Unix(),
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString([]byte(j.secret))
}

func (j *JWTManager) Validate(tokenStr string) (string, error) {
  token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
    return []byte(j.secret), nil
  })
  if err != nil || !token.Valid {
    return "", err
  }
  claims := token.Claims.(jwt.MapClaims)
  return claims["user_id"].(string), nil
}
