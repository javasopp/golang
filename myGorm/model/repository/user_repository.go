package repository

import "myGorm/model"

type UserRepository struct {
	model.BaseRepository[model.User]
}
