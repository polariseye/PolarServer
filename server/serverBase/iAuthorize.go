package serverBase

import (
	"github.com/polariseye/polarserver/common"
)

type IAuthorize interface {
	Authorize(request *common.RequestModel) bool
}
