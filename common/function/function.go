package function

import (
	"context"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/kataras/iris/v12"
	"reflect"
)

func BindBaseRequest(entity interface{}, ctx iris.Context) {
	//set base request parameter
	object := reflect.ValueOf(entity)
	object.Elem().FieldByName("Ctx").Set(reflect.ValueOf(ctx.Values().Get("ctx").(context.Context)))
	baseRequest := ctx.Values().Get("base_request").(request.BaseRequest)
	object.Elem().FieldByName("BaseUUID").Set(reflect.ValueOf(baseRequest.BaseUUID))
	object.Elem().FieldByName("BaseEmail").Set(reflect.ValueOf(baseRequest.BaseEmail))
	object.Elem().FieldByName("Language").Set(reflect.ValueOf(baseRequest.Language))
}
