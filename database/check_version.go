package database

import (
	"fmt"
	"github.com/bunred/buncms-plugin-team/utils"
	"github.com/bunred/buncms/types"
	buncmsutils "github.com/bunred/buncms/utils"
	"github.com/go-pg/pg/v9"
)

func CheckVersion(db *pg.DB) error {

	// Check database version
	dbVersion := &types.ModelSettings{
		Name: utils.PluginSettingName,
	}
	targetVersion := utils.PluginSettingVersion

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

	fmt.Printf("Plugin: %s-%s is LessThan %s, start megring data...\n", utils.PluginSettingName, dbVersion.Value, targetVersion)
	err := initDatabase(db)
	if err != nil {
		return err
	}

	return nil
}
