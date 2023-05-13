package jwts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/wtkeqrf0/while.act/pkg/middleware/errs"
	"time"
)

type AuthService interface {
	IDExist(id int) (bool, error)
}

type Auth struct {
	s AuthService
}

func NewAuth(s AuthService) *Auth {
	return &Auth{s: s}
}

// ValidateJWT validates the token and identifies the user with IDExist function. Returns an error in case of unsuccessful validation. Otherwise, returns the user ID.
func (a Auth) ValidateJWT(token string) (int, error) {

	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to decrypt token")
		}

		return []byte(cfg.Token.Secret), nil
	})

	claims, ok := at.Claims.(jwt.MapClaims)
	if !ok || !at.Valid || err != nil {
		return 0, fmt.Errorf("token is not valid: %v", err)
	}

	userId, oki := claims["id"].(float64)
	now := float64(time.Now().Unix())

	if !oki || now >= claims["exp"].(float64) {
		return 0, fmt.Errorf("token fields not found: %v", err)
	}

	id := int(userId)
	exist, err := a.s.IDExist(id)

	if err != nil {
		return 0, err
	} else if !exist {
		return 0, fmt.Errorf("user not found by token")
	}

	return id, nil
}

func (Auth) GenerateJWT(id int) (string, error) {

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(cfg.Token.AccessDuration).Unix(),
	}

	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.Token.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (Auth) GetUserId(c *gin.Context) (int, bool) {
	anyID, ok := c.Get("id")
	if !ok {
		c.Error(errs.UnAuthorized)
		return 0, false
	}

	id, ok := anyID.(int)
	if !ok {
		c.Error(errs.ServerError)
		return 0, false
	}

	return id, true
}
