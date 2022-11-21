package quest5

import "gis-project-backend/pkg/module/quest5/model"

type Quest5UseCase interface {
	QuestA(year string) (*model.Quest5AResponse, error)
	QuestB() (*model.Quest5BResponse, error)
	QuestC() (*model.Quest5CResponse, error)
	QuestD() (*model.Quest5DResponse, error)
	QuestE() (*model.Quest5EResponse, error)
	QuestF(year string) (*model.Quest5FResponse, error)
}
