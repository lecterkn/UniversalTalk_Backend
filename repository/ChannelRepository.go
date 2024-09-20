package repository

import (
	"lecter/goserver/db"
	"lecter/goserver/model"

	"github.com/google/uuid"
)

type ChannelRepository struct{}

func (ChannelRepository) Index() ([]model.ChannelModel, error) {
	var models []model.ChannelModel
	err := db.Database().Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (ChannelRepository) Select(id uuid.UUID) (*model.ChannelModel, error) {
	var model model.ChannelModel
	err := db.Database().Where("id = ?", id[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ChannelRepository) Create(model model.ChannelModel) (*model.ChannelModel, error) {
	err := db.Database().Create(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ChannelRepository) Update(model model.ChannelModel) (*model.ChannelModel, error) {
	err := db.Database().Save(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ChannelRepository) Delete(id uuid.UUID) error {
	err := db.Database().Where("id = ?", id[:]).Delete(&model.ChannelModel{}).Error
	return err
}
