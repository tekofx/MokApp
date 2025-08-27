package services

import (
	"database/sql"

	"github.com/Itros97/MokApp/internal/database/dal"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
	"github.com/Itros97/MokApp/internal/models"
)

func InsertItem(db *sql.DB, item models.Item) (*int64, *mokuerrors.MokuError) {
	tx, err := db.Begin()
	if err != nil {
		return nil, mokuerrors.DatabaseError(err.Error())
	}

	defer tx.Rollback()

	itemID, rerr := dal.CreateItem(db, &item)
	if nil != rerr {
		return nil, rerr
	}

	return itemID, nil
}

func GetItemById(db *sql.DB, id int64) (*models.Item, *mokuerrors.MokuError) {

	tx, err := db.Begin()
	if err != nil {
		return nil, mokuerrors.DatabaseError(err.Error())
	}

	defer tx.Rollback()

	itemID, rerr := dal.GetItemById(db, id)
	if nil != rerr {
		return nil, rerr
	}

	return itemID, nil

}
