package valid_test

import (
	"fmt"
	"net/http"

	"github.com/samber/lo"
	"github.com/xuender/kvalid"
)

type Model struct {
	Name string
	Age  int `json:"age,omitempty"`
}

func (p *Model) Validation(method string) *kvalid.Rules[*Model] {
	switch method {
	case http.MethodPost:
		return kvalid.New(p).
			Field(&p.Name,
				kvalid.Required().SetMessage("名称必填"),
				kvalid.MinStr(3).Optional().SetMessage("名称最少3个字"),
				kvalid.MaxStr(10).SetMessage("名称最多10个字"),
			).
			Field(&p.Age,
				kvalid.MinNum(10).SetMessage("最小年龄10岁"),
				kvalid.MaxNum(40).SetMessage("最大年龄40岁"),
			)
	default:
		return kvalid.New(p).
			Field(&p.Name,
				kvalid.Required().SetMessage("名称必填"),
				kvalid.MinStr(3).Optional().SetMessage("名称最少3个字"),
				kvalid.MaxStr(10).SetMessage("名称最多10个字"),
			)
	}
}

func (p *Model) Validate(method string) error {
	return p.Validation(method).Validate(p)
}

func ExampleBind_post() {
	mod := &Model{Name: "new name", Age: 18}
	old := &Model{Name: "old name", Age: 28}
	rule := mod.Validation(http.MethodPost)

	lo.Must0(rule.Bind(mod, old))
	fmt.Println(old.Name)
	fmt.Println(old.Age)

	// Output:
	// new name
	// 18
}

func ExampleBind_put() {
	mod := &Model{Name: "new name", Age: 18}
	old := &Model{Name: "old name", Age: 28}
	rule := mod.Validation(http.MethodPut)

	lo.Must0(rule.Bind(mod, old))
	fmt.Println(old.Name)
	fmt.Println(old.Age)

	// Output:
	// new name
	// 28
}
