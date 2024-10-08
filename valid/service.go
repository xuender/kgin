package valid

import (
	"encoding/hex"
	"net/http"
	"reflect"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/hash"
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

	name := val.Type().Name()

	return strings.ToLower(name)
}

func (p *Service) Router(group *gin.RouterGroup) {
	group.GET("/:key", p.get)
}

func (p *Service) Etag(ctx *gin.Context) string {
	key := strings.ToLower(ctx.Param("key"))
	if val, has := p.valids[key]; has {
		keys := make([]string, 0, len(val))
		for key := range val {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		sum := hash.NewSipHash64()

		for _, key := range keys {
			sum.Write([]byte(key))
			sum.Write(val[key])
		}

		return hex.EncodeToString(sum.Sum(nil))
	}

	return ""
}

func (p *Service) get(ctx *gin.Context) {
	key := strings.ToLower(ctx.Param("key"))
	if val, has := p.valids[key]; has {
		ctx.JSON(http.StatusOK, val)

		return
	}

	panic(BadRequestError("bad key:" + key))
}
