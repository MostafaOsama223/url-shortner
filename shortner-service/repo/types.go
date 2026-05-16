package repo

import (
	"github.com/MostafaOsama223/shortner-service/database"
)

type Repo interface {
	DB() database.Database
}
