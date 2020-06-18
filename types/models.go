package types

import "time"

type ModelTeams struct {
	tableName struct{}  `pg:"teams"`
	Id        int64     `pg:",pk,type:bigint,default:shard_1.id_generator()"`
	CreatedAt time.Time `pg:"default: now()"`
	UpdatedAt time.Time `pg:"default: now()"`
}
