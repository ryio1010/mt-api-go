package model

import "mt-api-go/domain/model"

type musclePartId string

type MusclePart struct {
	ID   musclePartId `json:"musclepartid"`
	Name string       `json:"musclepartname"`
}

func MusclePartFromDomainModel(m *model.MMusclepart) *MusclePart {
	musclePart := &MusclePart{
		ID:   musclePartId(m.Musclepartid),
		Name: m.Musclepartname,
	}

	return musclePart
}
