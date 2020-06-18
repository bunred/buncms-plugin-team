package buncms_plugin_team

import (
	"github.com/bunred/buncms-plugin-team/database"
	"github.com/bunred/buncms/types"
	buncmsutils "github.com/bunred/buncms/utils"
	"os"
)

func Routes() {
	// preset DbConnect
	params := types.DbConnect{}
	if os.Getenv("TEAM_DATABASE_ADDR") != "" {
		params.Options.Addr = os.Getenv("TEAM_DATABASE_ADDR")
	}
	if os.Getenv("TEAM_DATABASE_USER") != "" {
		params.Options.User = os.Getenv("TEAM_DATABASE_USER")
	}
	if os.Getenv("TEAM_DATABASE_PASSWORD") != "" {
		params.Options.Password = os.Getenv("TEAM_DATABASE_PASSWORD")
	}
	if os.Getenv("TEAM_DATABASE_NAME") != "" {
		params.Options.Database = os.Getenv("TEAM_DATABASE_NAME")
	}

	db := buncmsutils.DbConnect(params)

	err := database.CheckVersion(db)

	if err != nil {
		panic(err)
	}
}
