package valid

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kvalid"
)

type Service struct {
	valids map[string]map[string]BytesMarshaler
}

func NewService() *Service {
	return &Service{
		valids: make(map[string]map[string]BytesMarshaler),
	}
}

func (p *Service) Add(jsoners ...kvalid.ValidJSONer) {
	for _, jsoner := range jsoners {
		var (
			validJSON = jsoner.ValidJSON()
			bytesJSON = make(map[string]BytesMarshaler, len(validJSON))
		)

		for key, value := range validJSON {
			bytesJSON[key] = NewBytesMarshaler(value)
		}

		p.valids[getName(jsoner)] = bytesJSON
	}
}

func getName(model kvalid.ValidJSONer) string {
	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val.Type().Name()
}

func (p *Service) Router(group *gin.RouterGroup) {
	group.GET("/", p.get)
}

func (p *Service) get(ctx *gin.Context) {
	method := ctx.DefaultQuery("method", http.MethodPost)

	if val, has := p.valids[method]; has {
		ctx.JSON(http.StatusOK, val)

		return
	}

	panic(BadRequestError("Bad Method:" + method))
}
