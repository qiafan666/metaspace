package api

import (
	"fmt"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"
	"github.com/jau1jz/cornus/commons/utils"
	"strconv"
	"sync"
	"time"
)

// LogicService service layer interface
type LoginService interface {
	CreateAuthCode(info request.CreateAuthCode) (out response.CreateAuthCode, code commons.ResponseCode, err error)
}

var ThirdLoginConfig struct {
	ThirdLogin struct {
		ThirdLoginAddress string `yaml:"third_login_address"`
	} `yaml:"thirdLogin"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&ThirdLoginConfig)
}

var LoginServiceIns *LoginServiceImp
var LoginServiceInitOnce sync.Once

func NewLoginInstance() LoginService {
	LoginServiceInitOnce.Do(func() {
		LoginServiceIns = &LoginServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return LoginServiceIns
}

type LoginServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (l LoginServiceImp) CreateAuthCode(info request.CreateAuthCode) (out response.CreateAuthCode, code commons.ResponseCode, err error) {

	authCode := utils.GenerateUUID()

	var thirdPartySystem model.ThirdPartySystem
	err = l.dao.First([]string{model.ThirdPartySystemColumns.CallbackAddress}, map[string]interface{}{
		model.ThirdPartySystemColumns.ID: info.BaseThirdPartyId,
	}, nil, &thirdPartySystem)

	if err != nil {
		slog.Slog.ErrorF(nil, "SignServiceImp CreateAuthCode thirdPartySystem First error %s", err.Error())
		return out, 0, err
	}

	err = l.redis.SetAuthCode(info.Ctx, inner.AuthCode{
		ThirdPartyPublicId: strconv.FormatUint(info.BaseThirdPartyId, 10),
		AuthCode:           authCode,
		CallbackUrl:        thirdPartySystem.CallbackAddress,
	}, time.Minute*3)
	if err != nil {
		slog.Slog.ErrorF(nil, "SignServiceImp SetRand error %s", err.Error())
		return out, 0, err
	}

	out.AuthCode = authCode
	out.LoginAddress = ThirdLoginConfig.ThirdLogin.ThirdLoginAddress
	out.LoginUrl = fmt.Sprintf("%s?authCode=%s", ThirdLoginConfig.ThirdLogin.ThirdLoginAddress, authCode)
	return

}
