package function

import (
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/kataras/iris/v12"
	"reflect"
)

func BindBaseRequest(entity interface{}, ctx iris.Context) {
	//set base request parameter
	object := reflect.ValueOf(entity)
	baseRequest, _ := ctx.Values().Get("base_request").(request.BaseRequest)
	object.Elem().FieldByName("BaseRequest").Set(reflect.ValueOf(baseRequest))
}
