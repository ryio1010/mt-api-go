package model

import "mt-api-go/domain/model"

type TrainingLog struct {
	LogId          int     `json:"logid"`
	MenuId         int     `json:"menuid"`
	MenuName       string  `json:"menuname"`
	MusclePart     string  `json:"musclepart"`
	TrainingWeight float64 `json:"trainingweight"`
	TrainingCount  int     `json:"trainingcount"`
	TrainingDate   string  `json:"trainingdate"`
	TrainingMemo   string  `json:"trainingmemo"`
}

func TrainingLogFromDomainModel(t *model.TTraininglog) *TrainingLog {
	trainingLog := &TrainingLog{
		LogId:          t.Logid,
		MenuId:         t.Menuid,
		TrainingWeight: t.Trainingweight,
		TrainingCount:  t.Trainingcount,
		TrainingDate:   t.Trainingdate,
		TrainingMemo:   t.Trainingmemo,
	}

	return trainingLog
}
