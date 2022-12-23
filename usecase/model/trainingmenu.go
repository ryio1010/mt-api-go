package model

import "mt-api-go/domain/model"

type TrainingMenu struct {
	ID             int32  `json:"menuid"`
	Name           string `json:"menuname"`
	MusclePartId   string `json:"musclepartid"`
	MusclePartName string `json:"musclepartname"`
	Status         string `json:"status"`
}

type TrainingMenuRequest struct {
	MenuName     string `json:"menuname"`
	MusclePartId string `json:"musclepartid"`
	UserId       string `json:"userid"`
}

func TrainingMenuFromDomainModel(m *model.MMenu) *TrainingMenu {
	trainingMenu := &TrainingMenu{
		ID:           int32(m.Menuid),
		Name:         m.Menuname,
		MusclePartId: m.Musclepartid,
		Status:       m.Status,
	}

	return trainingMenu
}
