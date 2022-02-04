package util

import (
	"go-api-rest/src/database"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type TestSuiteEnv struct {
	suite.Suite
	db *gorm.DB
}

func (suite *TestSuiteEnv) SetupSuite() {
	database.Setup()
	suite.db = database.GetDB()
}

func (suite *TestSuiteEnv) TearDownTest() {
	database.ClearTable("books")
}

func (suite *TestSuiteEnv) TearDownSuite() {
	suite.db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteEnv))
}
