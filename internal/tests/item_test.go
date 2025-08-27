package tests

import (
	"database/sql"
	"testing"

	"github.com/Itros97/MokApp/internal/database/dal"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
)

func TestItemCrud(t *testing.T) {
	db, err := NewTestDatabase()
	AssertMokuErrorDoesNotExist(t, err)
	defer db.Close()

	t.Run("Create item", func(t *testing.T) { testCreateItem(t, db) })
}

func testCreateItem(t *testing.T, db *sql.DB) {
	_, err := dal.CreateItem(nil, nil)
	AssertMokuError(t, err, mokuerrors.UnexpectedErrorCode, mokuerrors.DatabaseConnectionEmptyMessage)

	_, err = dal.CreateItem(db, nil)
	AssertMokuError(t, err, mokuerrors.InvalidRequestErrorCode, mokuerrors.ItemInvalidMessage)

}
