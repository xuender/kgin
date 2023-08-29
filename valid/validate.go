package valid

import "github.com/AgentCosmic/xvalid"

type Put interface {
	ValidationPut() xvalid.Rules
	ValidatePut() error
}

type Post interface {
	ValidationPost() xvalid.Rules
	ValidatePost() error
}

type (
	Validation func() xvalid.Rules
	Validate   func() error
)
