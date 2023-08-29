package valid_test

import (
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
	"github.com/xuender/kgin/valid"
)

// nolint: lll
func ExampleValidationPut() {
	data := lo.Must1(json.Marshal(valid.ValidationPut(&Model{})))
	fmt.Println(string(data))

	// Output:
	// {"Model":{"Name":[{"rule":"required"},{"rule":"minStr","min":3},{"rule":"maxStr","max":10}]}}
}

// nolint: lll
func ExampleValidationPost() {
	data := lo.Must1(json.Marshal(valid.ValidationPost(&Model{})))
	fmt.Println(string(data))

	// Output:
	// {"Model":{"Name":[{"rule":"required"},{"rule":"minStr","min":3},{"rule":"maxStr","max":10}],"age":[{"rule":"minInt","min":10},{"rule":"maxInt","max":40}]}}
}
