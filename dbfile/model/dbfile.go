package model

import (
	"context"
	"time"

	"github.com/moetang/webapp-scaffold/frmpg"
	"github.com/moetang/webapp-scaffold/utils"
)

type Dbfile struct {
	FileId      string `mx.orm:"file_id"`
	Metadata    string `mx.orm:"metadata"`
	Filecontent []byte `mx.orm:"filecontent"`

	TimeCreated int64 `mx.orm:"time_created"`
	TimeUpdated int64 `mx.orm:"time_updated"`
}

func SaveDbFile(dbfile *Dbfile) error {
	now := utils.UnixTime(time.Now())
	dbfile.TimeCreated = now
	dbfile.TimeUpdated = now
	_, err := db.GetPostgresPool().Exec(context.Background(),
		"insert into moeobject_dbfile(file_id, metadata, filecontent, time_created, time_updated) values($1, $2, $3, $4, $5)",
		dbfile.FileId, dbfile.Metadata, dbfile.Filecontent, dbfile.TimeCreated, dbfile.TimeUpdated)
	if err != nil {
		return err
	}
	return nil
}

func LoadDbFileByFileId(fileId string) (*Dbfile, error) {
	d := new(Dbfile)
	err := frmpg.QuerySingle(db.GetPostgresPool(), d, context.Background(),
		"select * from moeobject_dbfile where file_id = $1",
		fileId)
	return d, err
}
