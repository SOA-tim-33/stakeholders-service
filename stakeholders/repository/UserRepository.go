package repository

import (
	"errors"
	"stakeholders-service/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func CheckDBConnection(db *gorm.DB) error {
	if db == nil {
		return errors.New("database connection is nil")
	}
	return nil
}

func (repo *UserRepository) Get(id string) (model.User, error) {
	user := model.User{}
	dbResult := repo.DatabaseConnection.First(&user, "id = ?", id)

	if dbResult.Error != nil {
		return user, dbResult.Error
	}

	return user, nil
}

func (repo *UserRepository) Exists(username string) (bool, error) {
	var user model.User
	dbResult := repo.DatabaseConnection.First(&user, "username = ?", username)
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, dbResult.Error
	}
	return true, nil
}

func (repo *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	dbResult := repo.DatabaseConnection.Find(&users)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return users, nil
}

func (repo *UserRepository) Save(user *model.User) (*model.User, error) {
	dbResult := repo.DatabaseConnection.Create(user)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return user, nil
}

func (repo *UserRepository) Update(user *model.User) error {
	dbResult := repo.DatabaseConnection.Save(user)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *UserRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.User{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *UserRepository) GetActiveByName(username string) (*model.User, error) {
	var user model.User
	dbResult := repo.DatabaseConnection.Where("username = ? AND is_active = ?", username, true).First(&user)
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, dbResult.Error
	}
	return &user, nil
}

func (repo *UserRepository) GetPersonId(userID uuid.UUID) (uuid.UUID, error) {
	var person model.Profile
	dbResult := repo.DatabaseConnection.Where("user_id = ?", userID).First(&person)
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return uuid.UUID{}, nil
		}
		return uuid.UUID{}, dbResult.Error
	}
	return person.Id, nil
}
