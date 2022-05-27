package middleware

import (
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/services/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/jau1jz/cornus"
	"github.com/jau1jz/cornus/commons"
	"github.com/kataras/iris/v12"
)

var jwtConfig struct {
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&jwtConfig)
}

var witheList = map[string]string{
	"/metaspace/web/register":                   "",
	"/metaspace/web/login":                      "",
	"/metaspace/web/login/nonce":                "",
	"/metaspace/web/subscribe/newsletter/email": "",
	"/metaspace/web/health":                     "",
	"/metaspace/web/orders":                     "",
	"/metaspace/web/login/third":                "",
}

func CheckPortalAuth(ctx iris.Context) {
	var language, uuid, email string
	var userId inner.UserId
	//check white list
	if _, ok := witheList[ctx.Request().RequestURI]; !ok {

		signMidOnce.Do(func() {
			signService = api.NewSignInstance()
		})
		//check jwt
		parseToken, err := jwt.Parse(ctx.Request().Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtConfig.JWT.Secret), nil
		})
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.TokenError, language))
			return
		}

		if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
			uuid, _ = claims["uuid"].(string)
			email, _ = claims["email"].(string)
		} else {
			_, _ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, language))
			return
		}

		userId, _, err = signService.GetUserId(ctx, uuid)
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefualtLanguage))
			return
		}
	}
	ctx.Values().Set(common.BasePortalRequest, request.BasePortalRequest{
		BaseUserID: userId.UserId,
		BaseUUID:   uuid,
		BaseEmail:  email,
	})
	ctx.Next()

}
