package sql

import "gis-project-backend/pkg/module/sql/model"

type SQLUseCase interface {
	Insert(req *model.InsertDataRequest) (*model.InsertDataResponse, error)
}
