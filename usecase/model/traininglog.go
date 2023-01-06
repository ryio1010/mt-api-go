package model

import (
	"go/types"
	"mt-api-go/domain/model"
)

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

type TrainingLogAddForm struct {
	MenuId              int         `json:"menuid"`
	TrainingDate        string      `json:"trainingdate"`
	UserId              string      `json:"userid"`
	WeightCountMemoList types.Slice `json:"weightcountmemolist"`
}

type TrainingLogUpdateForm struct {
	UserId     string      `json:"userid"`
	UpdateList types.Slice `json:"updatelist"`
	DeleteList types.Slice `json:"deletelist"`
}

type TrainingLogDeleteForm struct {
	MenuId       int    `json:"menuid"`
	TrainingDate string `json:"trainingdate"`
	UserId       string `json:"userid"`
}

type TrainingLogUpdateResponse struct {
	UpdateList types.Slice `json:"updatelist"`
	DeleteList types.Slice `json:"deletelist"`
}

type TrainingLogDeleteResponse struct {
	MenuId       int    `json:"Menuid"`
	TrainingDate string `json:"trainingdate"`
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
