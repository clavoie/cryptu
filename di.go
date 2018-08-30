package cryptu

import "github.com/clavoie/di"

// NewDiDefs returns new dependency injection definitions for this package
func NewDiDefs() []*di.Def {
	return []*di.Def{
		{NewBase64, di.Singleton},
	}
}
