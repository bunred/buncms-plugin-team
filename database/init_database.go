package database

import (
	"buncms-plugin-team/types"
	buncmstypes "github.com/bunred/buncms/types"
	buncmsutils "github.com/bunred/buncms/utils"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

func initDatabase(db *pg.DB) error {

	// Create table
	model := []interface{}{
		(*types.ModelTeams)(nil),
	}

	err := db.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		return err
	}

	// plugin-user settings data
	teamSettings := []buncmstypes.ModelSettings{
		{
			Name:  "team-db-version",
			Value: "1.00",
		},
	}

	// Init buncms_settings table/data
	err = buncmsutils.InitBuncms(db, teamSettings)
	if err != nil {
		return err
	}

	// Update db-version
	err = buncmsutils.UpdateSetting(db, "team-db-version", "1.00")

	if err != nil {
		return err
	}

	return nil
}
