package valid

import "github.com/AgentCosmic/xvalid"

type Valid interface {
	Validation(string) xvalid.Rules
	Validate(string) error
}
