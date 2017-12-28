package extend

import (
	"github.com/alibaba/pouch/daemon/mgr"
)

// RegisterSpecSetup registers some function used to setup the spec object.
func RegisterSpecSetup(setups ...mgr.SetupFunc) {
	for _, setup := range setups {
		mgr.Register(setup)
	}
}
