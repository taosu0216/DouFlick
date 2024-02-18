package middleWare

import (
	"encoding/base64"
	"gatewaysvr/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

var (
	Secret              = []byte("TikTok")
	TokenExpireDuration = time.Hour * 24 //过期时间
)

type JWTClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int64, username string) (string, error) {
	timeExpire := time.Now().Add(TokenExpireDuration)
	timeStr := &jwt.NumericDate{Time: timeExpire}
	claims := JWTClaims{
		Username: username,
		UserID:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "taosu_douflick_gateway",
			ExpiresAt: timeStr,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	secret, _ := base64.URLEncoding.DecodeString("TikTok")
	token, _ := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	//if err != nil {
	//	return nil, err
	//}
	//if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
	//	return claims, nil
	//}
	claim, _ := token.Claims.(*JWTClaims)
	return claim, nil
	//return nil, errors.New("invalid token")
}

func VerifyToken(tokenString string) (int64, error) {
	zap.L().Debug("tokenString", zap.String("tokenString", tokenString))
	if tokenString == "" {
		return int64(0), nil
	}
	claims, err := ParseToken(tokenString)
	if err != nil {
		return int64(0), err
	}
	return claims.UserID, nil
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.PostForm("token")
		if tokenStr == "" {
			tokenStr = c.Query("token")
		}
		userId, err := VerifyToken(tokenStr)
		if err != nil || userId == int64(0) {
			response.Fail(c, "鉴权失败", nil)
			c.Abort()
		}
		c.Set("UserId", userId)
	}
}

// AuthWithOutMiddleware 部分接口不需要用户登录也可访问，如feed，pushlishlist，favList，follow/follower list
func AuthWithOutMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Query("token")
		userId, err := VerifyToken(tokenString)
		if err != nil {
			response.Fail(c, "auth error", nil)
			c.Abort()
		}
		c.Set("UserId", userId)
		c.Next()
	}
}
