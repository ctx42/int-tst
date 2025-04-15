package first

import (
	"github.com/ctx42/int-tst/pkg/mocker/second"
)

// Medium represents test interface.
type Medium interface {
	// Method0 method documentation.
	Method0() error

	// Method1 method documentation.
	Method1(a *second.T0) error
}
