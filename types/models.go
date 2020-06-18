package types

import "time"

type ModelTeams struct {
	tableName struct{}  `pg:"teams"`
	Id        int64     `pg:",pk,type:bigint,default:shard_1.id_generator()"`
	CreatedAt time.Time `pg:"default: now()"`
	UpdatedAt time.Time `pg:"default: now()"`
	UserId    int64     `pg:",notnull"`
	Name      string    `pg:"type:varchar(20)"`
}

type ModelRelTeamUser struct {
	tableName struct{} `pg:"rel_team_user"`
	Id        int64    `pg:",pk,type:bigint,default:shard_1.id_generator()"`
	TeamId    int64    `pg:",notnull"`
	UserId    int64    `pg:",notnull"`
	Read      bool
	Update    bool
	Create    bool
	Delete    bool
}
