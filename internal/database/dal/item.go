package dal

import (
	"database/sql"
	"fmt"

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

	if item.Name == "" {
		return nil, mokuerrors.InvalidRequest(mokuerrors.ItemEmptyNameMessage)
	}

	// Check if item exists
	itm, itemGetErr := GetItemById(db, item.ID)
	if nil == itemGetErr && nil != itm {
		return nil, mokuerrors.New(mokuerrors.ItemAlreadyExistsErrorCode, mokuerrors.ItemAlreadyExistsMessage)
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
			stock
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

	rows.Scan(&itemId, &name, &description, &stock)

	return &models.Item{
		ID:          itemId,
		Name:        name,
		Description: description,
		Stock:       stock,
	}, nil

}

func GetAllItems(db *sql.DB) ([]*models.Item, *mokuerrors.MokuError) {
	if nil == db {
		return nil, mokuerrors.Unexpected(mokuerrors.DatabaseConnectionEmptyMessage)
	}

	statement, err := db.Prepare(`
		SELECT id,
			name,
			description,
			stock
		FROM items
	`)

	if nil != err {
		return nil, mokuerrors.DatabaseError(err.Error())
	}

	rows, err := statement.Query()

	if nil != err {
		return nil, mokuerrors.DatabaseError(err.Error())
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Stock)
		if err != nil {
			return nil, mokuerrors.DatabaseError(err.Error())
		}
		items = append(items, &item)
	}

	// Check for an error encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, mokuerrors.DatabaseError(err.Error())
	}

	fmt.Println(items)

	return items, nil

}
