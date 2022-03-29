package middleware

import (
	"context"
	"github.com/blockfishio/metaspace-backend/pojo/request"
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
}

func CheckAuth(ctx iris.Context) {
	var language, uuid, email string
	//get language
	language = ctx.Request().Header.Get("Language")
	if language == "" {
		language = commons.DefualtLanguage
	}

	//check white list
	if _, ok := witheList[ctx.Request().RequestURI]; !ok {
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
	}
	ctx.Values().Set("base_request", request.BaseRequest{
		Ctx:       ctx.Values().Get("ctx").(context.Context),
		Language:  language,
		BaseUUID:  uuid,
		BaseEmail: email,
	})
	ctx.Next()

}
