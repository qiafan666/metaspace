package middleware

import (
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	commonService "github.com/blockfishio/metaspace-backend/services/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/jau1jz/cornus"
	"github.com/jau1jz/cornus/commons"
	"github.com/kataras/iris/v12"
	"sync"
)

var comService commonService.PublicService
var commonMidOnce sync.Once

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

	commonMidOnce.Do(func() {
		comService = commonService.NewPublicInstance()
	})

	var language, uuid, email string
	var user inner.User

	BaseRequest := ctx.Values().Get(common.BaseRequest).(request.BaseRequest)
	language = BaseRequest.Language
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

		user, _, err = comService.GetUser(ctx, uuid)
		if err != nil {
			_, _ = ctx.JSON(commons.BuildFailed(commons.UnKnowError, commons.DefualtLanguage))
			return
		}
	}
	ctx.Values().Set(common.BasePortalRequest, request.BasePortalRequest{
		BaseUserID: user.UserId,
		BaseUUID:   uuid,
		BaseEmail:  email,
		BaseWallet: user.WalletAddress,
	})
	ctx.Next()

}
