package cryptu

import "github.com/clavoie/di"

// NewDiDefs returns new dependency injection definitions for this package
func NewDiDefs() []*di.Def {
	return []*di.Def{
		{NewAes, di.Singleton},
		{NewBase64, di.Singleton},

		// add definitions for Key and Base64Encoding to your
		// project in order to use these encoders out of the box,
		// or add your own definitions.
	}
}
