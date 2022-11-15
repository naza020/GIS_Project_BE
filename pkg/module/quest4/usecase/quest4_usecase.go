package usecase

import (
	"gis-project-backend/pkg/core"

	quest4 "gis-project-backend/pkg/module/quest4"
)

type quest4UseCase struct {
	CoreRegistry *core.CoreRegistry
}

func NewQuest4UseCase(coreRegistry *core.CoreRegistry) quest4.Quest4UseCase {
	return &quest4UseCase{
		CoreRegistry: coreRegistry,
	}
}
