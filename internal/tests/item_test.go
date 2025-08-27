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
