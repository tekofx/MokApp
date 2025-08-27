package tests

import (
	"database/sql"
	"testing"

	"github.com/Itros97/MokApp/internal/database/dal"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
	"github.com/Itros97/MokApp/internal/models"
)

func TestItemCrud(t *testing.T) {
	db, err := NewTestDatabase()
	AssertMokuErrorDoesNotExist(t, err)
	defer db.Close()

	expectedItem := models.Item{
		Name:        "Item1",
		Description: "Item 1 description",
		Stock:       3,
	}

	t.Run("Create item validations", func(t *testing.T) { testCreateItemValidations(t, db) })
	t.Run("Create item", func(t *testing.T) { testCreateItem(t, db, &expectedItem) })
	t.Run("Get item by id", func(t *testing.T) { testGetItemById(t, db, &expectedItem) })
	t.Run("Get all items", func(t *testing.T) { testGetItemById(t, db, &expectedItem) })
}

func testCreateItemValidations(t *testing.T, db *sql.DB) {
	_, err := dal.CreateItem(nil, nil)
	AssertMokuError(t, err, mokuerrors.UnexpectedErrorCode, mokuerrors.DatabaseConnectionEmptyMessage)

	_, err = dal.CreateItem(db, nil)
	AssertMokuError(t, err, mokuerrors.InvalidRequestErrorCode, mokuerrors.ItemInvalidMessage)

	_, err = dal.CreateItem(db, &models.Item{})
	AssertMokuError(t, err, mokuerrors.InvalidRequestErrorCode, mokuerrors.ItemEmptyNameMessage)
}

func testCreateItem(t *testing.T, db *sql.DB, item *models.Item) {
	itemId, err := dal.CreateItem(db, item)
	AssertMokuErrorDoesNotExist(t, err)

	obtainedItem, err := dal.GetItemById(db, *itemId)
	AssertMokuErrorDoesNotExist(t, err)
	Assert(t,
		nil != obtainedItem &&
			item.Name == obtainedItem.Name &&
			item.Description == obtainedItem.Description &&
			item.Stock == obtainedItem.Stock,
		"expected item and obtained item mismatch",
	)
}

func testGetItemById(t *testing.T, db *sql.DB, item *models.Item) {
	_, err := dal.GetItemById(nil, 0)
	AssertMokuError(t, err, mokuerrors.UnexpectedErrorCode, mokuerrors.DatabaseConnectionEmptyMessage)

	_, err = dal.GetItemById(db, -1)
	AssertMokuError(t, err, mokuerrors.InvalidRequestErrorCode, mokuerrors.ItemIdNegativeMessage)

	_, err = dal.GetItemById(db, 900)
	AssertMokuError(t, err, mokuerrors.NotFoundErrorCode, mokuerrors.ItemNotFoundMessage)

	obtainedItem, err := dal.GetItemById(db, 1)
	AssertMokuErrorDoesNotExist(t, err)
	Assert(t,
		obtainedItem != nil &&
			obtainedItem.Description == item.Description &&
			obtainedItem.Name == item.Name &&
			obtainedItem.Stock == item.Stock,
		"expected item and obtained item mismatch",
	)

}

func testGetAllItems(t *testing.T, db *sql.DB) {

	itemId, err := dal.CreateItem(db, &models.Item{
		Name:        "Item2",
		Description: "Item 2 description",
		Stock:       5,
	})

	AssertMokuErrorDoesNotExist(t, err)
	items, err := dal.GetAllItems(db)
	Assert(t, len(items) == 2, "Items length != 2")
	Assert(t, items[0].ID == 1 &&
		items[0].Name == "Item 1" &&
		items[1].ID == *itemId &&
		items[1].Name == "Item 2",
		"expected list of items and obtained list of items mismatch")

}
