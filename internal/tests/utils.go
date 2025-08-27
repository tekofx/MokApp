package tests

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Itros97/MokApp/internal/database"
	"github.com/Itros97/MokApp/internal/database/tables"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
	"github.com/Itros97/MokApp/internal/logger"
)

func Assert(t *testing.T, predicate bool, failMessage string) {
	if !predicate {
		logger.Log("Test failed:", failMessage)
		t.FailNow()
	}
}

func AssertMokuErrorDoesNotExist(t *testing.T, error *mokuerrors.MokuError) {
	if nil != error {
		logger.Error("Test failed with error:", error.Message)
		t.FailNow()
	}
}

func AssertMokuError(t *testing.T, error *mokuerrors.MokuError, code mokuerrors.MokuErrorCode, message string) {
	if nil == error {
		logger.Error("Test failed because error is empty.")
		t.FailNow()
	}
	Assert(t, error.Code == code && error.Message == message, fmt.Sprintf("\n[%d - %s] \nwas expected but \n[%d - %s] \nwas found\n", code, message, error.Code, error.Message))
}

func NewTestDatabase() (*sql.DB, *mokuerrors.MokuError) {

	db, err := database.Connect(":memory:")
	if err != nil {
		return nil, mokuerrors.New(mokuerrors.DatabaseErrorCode, err.Error())
	}

	err = tables.UpdateDatabaseTablesToLatestVersion("../..", db)
	if err != nil {
		return nil, mokuerrors.New(mokuerrors.DatabaseErrorCode, err.Error())
	}

	return db, nil
}
