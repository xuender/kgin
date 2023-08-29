package valid_test

import (
	"fmt"
	"net/http"

	"github.com/AgentCosmic/xvalid"
	"github.com/samber/lo"
	"github.com/xuender/kgin/valid"
)

type Model struct {
	Name string
	Age  int `json:"age,omitempty"`
}

func (p Model) Validation(method string) xvalid.Rules {
	switch method {
	case http.MethodPost:
		return xvalid.New(&p).
			Field(&p.Name,
				xvalid.Required().SetMessage("名称必填"),
				xvalid.MinStr(3).Optional().SetMessage("名称最少3个字"),
				xvalid.MaxStr(10).SetMessage("名称最多10个字"),
			).
			Field(&p.Age,
				xvalid.MinInt(10).SetMessage("最小年龄10岁"),
				xvalid.MaxInt(40).SetMessage("最大年龄40岁"),
			)
	default:
		return xvalid.New(&p).
			Field(&p.Name,
				xvalid.Required().SetMessage("名称必填"),
				xvalid.MinStr(3).Optional().SetMessage("名称最少3个字"),
				xvalid.MaxStr(10).SetMessage("名称最多10个字"),
			)
	}
}

func (p Model) Validate(method string) error {
	return p.Validation(method).Validate(p)
}

func ExampleValue_post() {
	mod := &Model{Name: "new name", Age: 18}
	old := &Model{Name: "old name", Age: 28}

	lo.Must1(valid.Value(mod, old, http.MethodPost))
	fmt.Println(old.Name)
	fmt.Println(old.Age)

	// Output:
	// new name
	// 18
}

func ExampleValue_put() {
	mod := &Model{Name: "new name", Age: 18}
	old := &Model{Name: "old name", Age: 28}

	lo.Must1(valid.Value(mod, old, http.MethodPut))
	fmt.Println(old.Name)
	fmt.Println(old.Age)

	// Output:
	// new name
	// 28
}
