package valid

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kvalid"
)

type Service struct {
	valids map[string]map[string]*kvalid.Rules
}

func NewService() *Service {
	return &Service{
		valids: make(map[string]map[string]*kvalid.Rules),
	}
}

func (p *Service) Add(method string, holders ...kvalid.RuleHolder) {
	valids, has := p.valids[method]
	if !has {
		valids = map[string]*kvalid.Rules{}
		p.valids[method] = valids
	}

	for key, value := range Validation(method, holders...) {
		valids[key] = value
	}
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
