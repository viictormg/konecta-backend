package feature

import "konecta/repository"

type Feature struct {
	db repository.Repository
}

func NewFeature(db repository.Repository) Feature {
	return Feature{
		db: db,
	}
}
