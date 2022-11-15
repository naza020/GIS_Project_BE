package quest4

import "gis-project-backend/pkg/module/quest4/model"

type Quest4UseCase interface {
	GetString() (string, error)
	QuestA() (*model.Quest4AResponse, error)
	QuestB() (*model.Quest4BResponse, error)
	QuestC(country string) (*model.Quest4CResponse, error)
	QuestD(year string, color string) (*model.Quest4DResponse, error)
}
