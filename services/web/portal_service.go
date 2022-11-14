package web

import (
	"github.com/qiafan666/fundametality/dao"
	"github.com/qiafan666/fundametality/pojo/request"
	"github.com/qiafan666/fundametality/pojo/response"
	cornus "github.com/qiafan666/quickweb"
	"github.com/qiafan666/quickweb/commons"
	"sync"
)

// PortalService service layer interface
type PortalService interface {
	Login(info request.UserLogin) (out response.UserLogin, code commons.ResponseCode, err error)
}

var portalConfig struct {
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&portalConfig)
}

var portalServiceIns *portalServiceImp
var portalServiceInitOnce sync.Once

func NewPortalServiceInstance() PortalService {

	portalServiceInitOnce.Do(func() {
		portalServiceIns = &portalServiceImp{
			dao: dao.Instance(),
		}
	})

	return portalServiceIns
}

type portalServiceImp struct {
	dao dao.Dao
}

func (p portalServiceImp) Login(info request.UserLogin) (out response.UserLogin, code commons.ResponseCode, err error) {

	return
}
