package usecase

import (
	"gis-project-backend/pkg/core"

	quest5 "gis-project-backend/pkg/module/quest5"
)

type quest5UseCase struct {
	CoreRegistry *core.CoreRegistry
}

func NewQuest5UseCase(coreRegistry *core.CoreRegistry) quest5.Quest5UseCase {
	return &quest5UseCase{
		CoreRegistry: coreRegistry,
	}
}
