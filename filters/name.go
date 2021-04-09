package filters

import (
	"net/http"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// Name apply filter for column name
func Name(r *http.Request, db *gorm.DB) {
	query := r.URL.Query()
	name := cast.ToString(query.Get("name"))

	if name != "" {
		db.Where("name = ?", name)
	}
}

// Done apply filter for column done
func Done(r *http.Request, db *gorm.DB) *gorm.DB {
	query := r.URL.Query()
	done := query.Get("done")

	if done != "" {
		db.Where("done = ?", cast.ToBool(done))
	}

	return db
}
