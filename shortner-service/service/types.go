package service

import (
	"github.com/MostafaOsama223/shortner-service/repo"
)

type Service interface {
	Repo() repo.Repo
}
