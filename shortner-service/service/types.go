package service

import (
	"shortner-service/repo"
)

type Service interface {
	Repo() repo.Repo
}
