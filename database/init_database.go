package database

import (
	"github.com/bunred/buncms-plugin-team/types"
	"github.com/bunred/buncms-plugin-team/utils"
	buncmstypes "github.com/bunred/buncms/types"
	buncmsutils "github.com/bunred/buncms/utils"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

func initDatabase(db *pg.DB) error {

	// Create table
	model := &types.ModelTeams{}

	err := db.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		return err
	}

	// plugin-user settings data
	teamSettings := []buncmstypes.ModelSettings{
		{
			Name:  utils.PluginSettingName,
			Value: "1.00", // default version
		},
	}

	// Init buncms_settings table/data
	err = buncmsutils.InitBuncms(db, teamSettings)
	if err != nil {
		return err
	}

	// Update db-version
	err = buncmsutils.UpdateSetting(db, utils.PluginSettingName, utils.PluginSettingVersion)

	if err != nil {
		return err
	}

	return nil
}
