package buncms_plugin_team

import (
	"github.com/bunred/buncms-plugin-team/database"
	"github.com/bunred/buncms/types"
	buncmsutils "github.com/bunred/buncms/utils"
)

func Routes() {
	db := buncmsutils.DbConnect(types.DbConnect{})

	err := database.CheckVersion(db)

	if err != nil {
		panic(err)
	}
}
