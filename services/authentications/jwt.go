package authentications

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/trend-ai/TrendAI_mobile_backend/conf"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"net/http"
	"time"
)

// Get hmac secret from app key
var hmacSecret = []byte(conf.Get().AppKey)

// Generate JWT authentication token by user
func GenerateAuthenticationTokenByUser(user models.User) (*models.AuthenticationToken, error) {
	now := time.Now().UTC()
	expTime := now.Add(604800 * time.Second) //Expire for 7 days

	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = expTime.Unix()
	claims["uid"] = user.Id
	accessToken.Claims = claims
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

// Filter for verify jwt authentication token from request
func JwtAuthenticationFilter() beego.FilterFunc {
	return func(ctx *context.Context) {
		// Response has indent
		hasIndent := beego.BConfig.RunMode != beego.PROD

		// Parse token from request
		token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return hmacSecret, nil
			},
		)

		// Respond status unauthorized if parse fail
		if err != nil {
			ctx.Output.SetStatus(http.StatusUnauthorized)
			_ = ctx.Output.JSON(models.NewResponseWithError("unauthorized", "Unauthorized access to this resource"), hasIndent, false)
			return
		}

		// Get uid, exp from token
		uid := token.Claims.(jwt.MapClaims)["uid"]
		exp := token.Claims.(jwt.MapClaims)["exp"]
		if !token.Valid || uid == nil || exp == nil {
			ctx.Output.SetStatus(http.StatusUnauthorized)
			_ = ctx.Output.JSON(models.NewResponseWithError("unauthorized", "Token is not valid"), hasIndent, false)
			return
		}

		// Get user from our database
		userId := fmt.Sprintf("%v", uid)
		userCollection := models.GetUserCollection()
		userSnapshot, err := userCollection.Doc(userId).Get(databases.Context)

		// Account doesn't exists
		if err != nil {
			ctx.Output.SetStatus(http.StatusUnauthorized)
			_ = ctx.Output.JSON(models.NewResponseWithError("unauthorized", "User is not exists"), hasIndent, false)
			return
		}

		// Cast user data to user struct
		var user models.User
		err = userSnapshot.DataTo(&user)
		if err != nil {
			logs.Debug("Parse user data failed", err)
			ctx.Output.SetStatus(http.StatusUnauthorized)
			_ = ctx.Output.JSON(models.NewResponseWithError("unauthorized", "User is not exists"), hasIndent, false)
			return
		}

		// Get user id
		user.Id = userSnapshot.Ref.ID

		// Check if token was expired
		if int64(exp.(float64)) < time.Now().Unix() {
			ctx.Output.SetStatus(http.StatusUnauthorized)
			_ = ctx.Output.JSON(models.NewResponseWithError("unauthorized", "Token is not valid"), hasIndent, false)
			return
		}

		// Pass user data to request's input
		ctx.Input.SetData("uid", userId)
		ctx.Input.SetData("user", user)
		logs.Debug("uid:", userId, ", exp:", exp, ", logged at:", time.Now().Unix())
	}
}
