package dal

import (
	"database/sql"

	mokuerrors "github.com/Itros97/MokApp/internal/errors"
	"github.com/Itros97/MokApp/internal/models"
)

func CreateItem(db *sql.DB, item *models.Item) (*int64, *mokuerrors.MokuError) {
	if nil == db {
		return nil, mokuerrors.Unexpected(mokuerrors.DatabaseConnectionEmptyMessage)
	}

	if nil == item {
		return nil, mokuerrors.InvalidRequest(mokuerrors.ItemInvalidMessage)
	}

	// Check if item exists
	itm, itemGetErr := GetItemById(db, item.ID)
	if nil == itemGetErr && nil != itm {
		return nil, mokuerrors.New(mokuerrors.ItemAlreadyExistsErrorCode, mokuerrors.ItemAlreadyExistsMessage)
	}

	if itemGetErr.Code != mokuerrors.NotFoundErrorCode {
		return nil, itemGetErr
	}

	statement, err := db.Prepare(
		"INSERT INTO items(name, description, stock) VALUES(?,?,?)",
	)

	if nil != err {
		return nil, mokuerrors.DatabaseError(err.Error())
	}

	res, err := statement.Exec(
		item.Name,
		item.Description,
		item.Stock,
	)

	if nil != err {
		return nil, mokuerrors.DatabaseError(err.Error())
	}

	item.ID, err = res.LastInsertId()
	if nil != err {
		return nil, mokuerrors.DatabaseError(err.Error())
	}

	return &item.ID, nil

}

func GetItemById(db *sql.DB, id int64) (*models.Item, *mokuerrors.MokuError) {

	if nil == db {
		return nil, mokuerrors.Unexpected(mokuerrors.DatabaseConnectionEmptyMessage)
	}

	if id <= 0 {
		return nil, mokuerrors.InvalidRequest(mokuerrors.ItemIdNegativeMessage)
	}

	statement, err := db.Prepare(`
		SELECT id,
			name,
			description,
			stock,
		FROM items
		WHERE id = ?
	`)

	if nil != err {
		return nil, mokuerrors.DatabaseError(err.Error())
	}

	rows, err := statement.Query(id)

	if nil != err {
		return nil, mokuerrors.DatabaseError(err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, mokuerrors.NotFound(mokuerrors.ItemNotFoundMessage)
	}

	var itemId int64
	var name string
	var description string
	var stock int64

	rows.Scan(itemId, name, description, stock)

	return &models.Item{
		ID:          itemId,
		Name:        name,
		Description: description,
		Stock:       stock,
	}, nil

}
