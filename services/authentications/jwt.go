package authentications

import (
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/trend-ai/TrendAI_mobile_backend/conf"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"time"
)

// Generate JWT authentication token by user
func GenerateAuthenticationTokenByUser(user models.User) (*models.AuthenticationToken, error) {
	now := time.Now().UTC()
	expTime := now.Add(604800 * time.Second) //Expire for 7 days
	hmacSecret := []byte(conf.Get().AppKey)

	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = expTime.Unix()
	claims["uid"] = user.Id
	accessToken.Claims = claims
	logs.Debug("key: ", conf.Get().AppKey)
	accessTokenStr, err := accessToken.SignedString(hmacSecret)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims = make(jwt.MapClaims)
	claims["uid"] = user.Id
	claims["ref"] = "refresh"
	refreshToken.Claims = claims
	refreshTokenStr, err := refreshToken.SignedString(hmacSecret)
	if err != nil {
		return nil, err
	}

	return &models.AuthenticationToken{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
	}, nil
}
