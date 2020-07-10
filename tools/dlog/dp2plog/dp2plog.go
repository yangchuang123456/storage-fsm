package dp2plog

import (
	"github.com/filecoin-project/storage-fsm/tools/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("fsm")
}
