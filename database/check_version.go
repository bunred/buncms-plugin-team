package database

import (
	"fmt"
	"github.com/bunred/buncms/types"
	buncmsutils "github.com/bunred/buncms/utils"
	"github.com/go-pg/pg/v9"
)

func CheckVersion(db *pg.DB) error {

	// Check database version
	dbVersion := &types.ModelSettings{
		Name: "team-db-version",
	}
	targetVersion := "1.00"

	selectError := db.Select(dbVersion)
	if selectError == nil {
		vLt, vErr := buncmsutils.VersionLessThan(dbVersion.Value, targetVersion)
		if vErr != nil {
			return vErr
		}
		if !vLt {
			return nil
		}
	}

	fmt.Printf("Plugin: user-%s is LessThan %s, start megring data...\n", dbVersion.Value, targetVersion)
	err := initDatabase(db)
	if err != nil {
		return err
	}

	return nil
}
