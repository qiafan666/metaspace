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
	object.Elem().FieldByName("BaseRequest").Set(reflect.ValueOf(baseRequest))
}

func BindApiBaseRequest(entity interface{}, ctx iris.Context) {
	//set base request parameter
	object := reflect.ValueOf(entity)
	baseRequest, _ := ctx.Values().Get(common.BaseApiRequest).(request.BaseApiRequest)
	object.Elem().FieldByName("BaseApiRequest").Set(reflect.ValueOf(baseRequest))
}
