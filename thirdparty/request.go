package thirdparty

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/qiafan666/quickweb/commons/log"
	"github.com/qiafan666/quickweb/commons/utils"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
)

type Service interface {
	Request(context.Context, Uri, BaseThirdParty, interface{}, interface{}) (ResponseCode, error)
}
type serviceImp struct {
	dao         dao.Dao
	requestPool sync.Pool
}

func NewThirdPartyService() Service {
	imp := serviceImp{dao: dao.Instance()}
	imp.requestPool.New = func() any {
		return BaseRequest{}
	}
	return &imp
}
func (s *serviceImp) Request(ctx context.Context, uri Uri, third BaseThirdParty, input interface{}, output interface{}) (code ResponseCode, err error) {
	//find third party public key
	var thirdPartySystem model.ThirdPartySystem
	err = s.dao.First([]string{model.ThirdPartySystemColumns.ThirdPartyPublicKey, model.ThirdPartySystemColumns.APIkey, model.ThirdPartySystemColumns.CallbackAddress}, map[string]interface{}{
		model.ThirdPartySystemColumns.ID: third.ThirdPartyID,
	}, nil, &thirdPartySystem)

	if err != nil {
		log.Slog.ErrorF(ctx, "serviceImp ThirdPartyLogin thirdPartySystem First error %s", err.Error())
		return 0, err
	}
	marshal, err := json.Marshal(input)
	if err != nil {
		return 0, err
	}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetBodyRaw(marshal)
	encrypt, err := utils.RsaEncrypt([]byte(thirdPartySystem.APIkey), []byte(thirdPartySystem.ThirdPartyPublicKey))
	if err != nil {
		log.Slog.ErrorF(ctx, "serviceImp ThirdPartyLogin rsaEncrypt error %s", err.Error())
		return 0, err
	}
	req.Header.SetMethod(http.MethodPost)
	//set authorization header
	req.Header.Set(HeaderAuthorization, base64.URLEncoding.EncodeToString(encrypt))
	req.Header.SetContentType("application/json")
	requestUri := fasthttp.AcquireURI()
	defer fasthttp.ReleaseURI(requestUri)
	err = requestUri.Parse([]byte(thirdPartySystem.CallbackAddress), []byte(uri))
	if err != nil {
		return 0, err
	}
	req.SetURI(requestUri)

	if err != nil {
		return 0, err
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = fasthttp.Do(req, resp)
	if err != nil {
		return 0, err
	}
	body := resp.Body()
	base := s.requestPool.Get().(BaseRequest)
	err = json.Unmarshal(body, &base)
	if err != nil {
		return 0, err
	}
	if base.Code != 0 {
		return base.Code, errors.New(base.Msg)
	}
	err = json.Unmarshal(base.Data, output)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
