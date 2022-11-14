package function

import (
	"github.com/qiafan666/fundametality/common"
	"reflect"

	"github.com/kataras/iris/v12"
	"github.com/qiafan666/fundametality/pojo/request"
)

func BindBaseRequest(entity interface{}, ctx iris.Context) {
	//set base request parameter
	object := reflect.ValueOf(entity)

	baseRequest, _ := ctx.Values().Get(common.BaseRequest).(request.BaseRequest)
	elem := object.Elem()
	base := elem.FieldByName("BaseRequest")
	if base.Kind() != reflect.Invalid {
		base.Set(reflect.ValueOf(baseRequest))
	}

	basePortalRequest, _ := ctx.Values().Get(common.BasePortalRequest).(request.BasePortalRequest)
	basePortal := elem.FieldByName("BasePortalRequest")
	if basePortal.Kind() != reflect.Invalid {
		basePortal.Set(reflect.ValueOf(basePortalRequest))
	}

}
