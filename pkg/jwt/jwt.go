package jwt

import (
	"errors"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenExpired           = errors.New("token expired")
	ErrTokenExpiredMaxRefresh = errors.New("token expired max refresh")
	ErrTokenMalformed         = errors.New("token malformed")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrHeaderEmpty            = errors.New("header is empty")
	ErrHeaderMalformed        = errors.New("header malformed")
)

type JWT struct {
	SignKey    []byte
	MaxRefresh time.Duration
}

type CustomClaims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwtpkg.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

func (jwt *JWT) ParseToken(ctx *gin.Context) (*CustomClaims, error) {
	tokenString, parseErr := jwt.getTokenFromRequest(ctx)
	if parseErr != nil {
		return nil, parseErr
	}

	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		if errors.Is(err, jwtpkg.ErrTokenMalformed) {
			return nil, ErrTokenMalformed
		} else if errors.Is(err, jwtpkg.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, ErrTokenInvalid
	}

	return claims, nil
}

func (jwt *JWT) RefreshToken(ctx *gin.Context) (string, error) {
	tokenString, parseErr := jwt.getTokenFromRequest(ctx)
	if parseErr != nil {
		return "", parseErr
	}

	token, err := jwt.parseTokenString(tokenString)
	if err != nil && !errors.Is(err, jwtpkg.ErrTokenExpired) {
		return "", err
	}

	claims := token.Claims.(*CustomClaims)
	x := app.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt.Unix() > x {
		claims.RegisteredClaims.ExpiresAt = jwtpkg.NewNumericDate(jwt.expireAtTime())
		return jwt.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

func (jwt *JWT) IssueToken(userID, userName string) string {
	expireAtTime := jwt.expireAtTime()
	claims := CustomClaims{
		userID,
		userName,
		jwtpkg.RegisteredClaims{
			NotBefore: jwtpkg.NewNumericDate(app.TimenowInTimezone()),
			IssuedAt:  jwtpkg.NewNumericDate(app.TimenowInTimezone()),
			ExpiresAt: jwtpkg.NewNumericDate(expireAtTime),
			Issuer:    config.GetString("app.name"),
		},
	}

	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

func (jwt *JWT) createToken(claims CustomClaims) (string, error) {
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

func (jwt *JWT) expireAtTime() time.Time {
	timenow := app.TimenowInTimezone()

	var expireAtTime int64
	if config.GetBool("app.debug") {
		expireAtTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireAtTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireAtTime) * time.Minute

	return timenow.Add(expire)
}

func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

func (jwt *JWT) getTokenFromRequest(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if len(authHeader) == 0 {
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}

	return parts[1], nil
}
