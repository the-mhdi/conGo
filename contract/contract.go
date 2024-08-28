package contract

import (
	"github.com/ethereum/go-ethereum/common"
)

type Contract struct {
	version  string //0.y.z versioning
	deployer common.Address
	owner    common.Address
	modifier map[string]any //any == any function, i'll use reflection
}

type contract struct {
	version  string //0.y.z versioning
	deployer common.Address
	owner    common.Address
	modifier map[string]any //any == any function, i'll use reflection
}
