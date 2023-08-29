package valid_test

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/samber/lo"
	"github.com/xuender/kgin/valid"
)

// nolint: lll
func ExampleValidation_put() {
	data := lo.Must1(json.Marshal(valid.Validation(http.MethodPut, &Model{})))
	fmt.Println(string(data))

	// Output:
	// {"Model":{"Name":[{"rule":"required"},{"rule":"minStr","min":3},{"rule":"maxStr","max":10}]}}
}

// nolint: lll
func ExampleValidation_post() {
	data := lo.Must1(json.Marshal(valid.Validation(http.MethodPost, &Model{})))
	fmt.Println(string(data))

	// Output:
	// {"Model":{"Name":[{"rule":"required"},{"rule":"minStr","min":3},{"rule":"maxStr","max":10}],"age":[{"rule":"minInt","min":10},{"rule":"maxInt","max":40}]}}
}
