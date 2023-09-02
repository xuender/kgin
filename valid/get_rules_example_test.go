package valid_test

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/samber/lo"
	"github.com/xuender/kgin/valid"
)

// nolint: lll
func ExampleGetRules_put() {
	data := lo.Must1(json.Marshal(valid.GetRules(http.MethodPut, &Model{})))
	fmt.Println(string(data))

	// Output:
	// {"Model":{"Name":[{"rule":"required","msg":"名称必填"},{"rule":"minStr","min":3,"msg":"名称最少3个字"},{"rule":"maxStr","max":10,"msg":"名称最多10个字"}]}}
}

// nolint: lll
func ExampleGetRules_post() {
	data := lo.Must1(json.Marshal(valid.GetRules(http.MethodPost, &Model{})))
	fmt.Println(string(data))

	// Output:
	// {"Model":{"Name":[{"rule":"required","msg":"名称必填"},{"rule":"minStr","min":3,"msg":"名称最少3个字"},{"rule":"maxStr","max":10,"msg":"名称最多10个字"}],"age":[{"rule":"minNum","min":10,"msg":"最小年龄10岁"},{"rule":"maxNum","max":40,"msg":"最大年龄40岁"}]}}
}
