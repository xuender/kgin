package valid_test

import (
	"fmt"

	"github.com/AgentCosmic/xvalid"
	"github.com/samber/lo"
	"github.com/xuender/kgin/valid"
)

type Model struct {
	Name string
	Age  int
}

func (p Model) ValidationPost() xvalid.Rules {
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
}

func (p Model) ValidatePost() error {
	return p.ValidationPost().Validate(p)
}

func (p Model) ValidationPut() xvalid.Rules {
	return xvalid.New(&p).
		Field(&p.Name,
			xvalid.Required().SetMessage("名称必填"),
			xvalid.MinStr(3).Optional().SetMessage("名称最少3个字"),
			xvalid.MaxStr(10).SetMessage("名称最多10个字"),
		)
}

func (p Model) ValidatePut() error {
	return p.ValidationPut().Validate(p)
}

func ExamplePostValue() {
	mod := &Model{Name: "new name", Age: 18}
	old := &Model{Name: "old name", Age: 28}

	lo.Must1(valid.PostValue(mod, old))
	fmt.Println(old.Name)
	fmt.Println(old.Age)

	// Output:
	// new name
	// 18
}

func ExamplePutValue() {
	mod := &Model{Name: "new name", Age: 18}
	old := &Model{Name: "old name", Age: 28}

	lo.Must1(valid.PutValue(mod, old))
	fmt.Println(old.Name)
	fmt.Println(old.Age)

	// Output:
	// new name
	// 28
}
