package middleware

import (
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
	"/metaspace/web/register": "",
	"/metaspace/web/login":    "",
}

func CheckAuth(ctx iris.Context) {
	if _, ok := witheList[ctx.Request().RequestURI]; ok {
		ctx.Next()
		return
	}

	parseToken, err := jwt.Parse(ctx.Request().Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.JWT.Secret), nil
	})
	if err != nil {
		_, _ = ctx.JSON(commons.BuildFailed(commons.TokenError))
		return
	}
	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		ctx.Values().Set("claims", claims)
		ctx.Next()
		return
	} else {
		_, _ = ctx.JSON(commons.BuildFailed(commons.UnKnowError))
		return
	}

}
