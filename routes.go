package buncms_plugin_user

import (
	"github.com/bunred/buncms/types"
	buncmsutils "github.com/bunred/buncms/utils"
)

func Routes() {
	buncmsutils.DbConnect(types.DbConnect{})
}
