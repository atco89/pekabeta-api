package fixtures

import (
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
)

func Fixtures(db *gorm.DB) error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db.DB()),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("internal/fixtures/data"),
	)

	if err != nil {
		return err
	}
	return fixtures.Load()
}
