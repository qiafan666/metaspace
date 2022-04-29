package function

import (
	"github.com/blockfishio/metaspace-backend/common"
	"reflect"

	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/kataras/iris/v12"
)

func BindBaseRequest(entity interface{}, ctx iris.Context) {
	//set base request parameter
	object := reflect.ValueOf(entity)
	baseRequest, _ := ctx.Values().Get(common.BaseRequest).(request.BaseRequest)
	base := object.Elem().FieldByName("BaseRequest")

	if base.Kind() != reflect.Invalid {
		base.Set(reflect.ValueOf(baseRequest))
	}

	baseApiRequest, _ := ctx.Values().Get(common.BaseApiRequest).(request.BaseRequest)
	baseApi := object.Elem().FieldByName("BaseApiRequest")
	if baseApi.Kind() != reflect.Invalid {
		baseApi.Set(reflect.ValueOf(baseApiRequest))
	}
}
