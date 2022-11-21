package usecase

import (
	"gis-project-backend/pkg/core"

	sqlQuest "gis-project-backend/pkg/module/sql"
)

type sqlUseCase struct {
	CoreRegistry *core.CoreRegistry
}

func NewSQLUseCase(coreRegistry *core.CoreRegistry) sqlQuest.SQLUseCase {
	return &sqlUseCase{
		CoreRegistry: coreRegistry,
	}
}
