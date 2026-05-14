package repo

import (
	"shortner-service/database"
)

type Repo interface {
	DB() database.Database
}
