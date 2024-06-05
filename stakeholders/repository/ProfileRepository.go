package repository

import (
	"stakeholders-service/model"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *ProfileRepository) Get(id string) (model.Profile, error) {
	profile := model.Profile{}
	dbResult := repo.DatabaseConnection.First(&profile, "id = ?", id)

	if dbResult.Error != nil {
		return profile, dbResult.Error
	}

	return profile, nil
}

func (repo *ProfileRepository) GetAll() ([]model.Profile, error) {
	var profiles []model.Profile
	dbResult := repo.DatabaseConnection.Find(&profiles)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return profiles, nil
}

func (repo *ProfileRepository) Save(profile *model.Profile) (*model.Profile, error) {
	dbResult := repo.DatabaseConnection.Create(profile)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return profile, nil
}

func (repo *ProfileRepository) Update(profile *model.Profile) error {
	dbResult := repo.DatabaseConnection.Save(profile)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *ProfileRepository) Delete(id string) error {
	dbResult := repo.DatabaseConnection.Delete(&model.Profile{}, "id = ?", id)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}
