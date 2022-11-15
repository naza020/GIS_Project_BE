package quest5

import "gis-project-backend/pkg/module/quest5/model"

type Quest5UseCase interface {
	QuestA(year string) (*model.Quest5AResponse, error)
	QuestB() (*model.Quest5BResponse, error)

	QuestF(year string) (*model.Quest5FResponse, error)
}
