package usecase

import (
	"fmt"
	"gis-project-backend/pkg/module/quest4/model"
)

func (service *quest4UseCase) QuestD(year string, color string) (*model.Quest4DResponse, error) {

	query := fmt.Sprintf("SELECT SUM([Population]) FROM dbo.AirPollutionPM25 WHERE Year=%s AND Color_pm25='%s'", year, color)
	results, err := service.CoreRegistry.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var resultQuest model.Quest4DResponse
	for results.Next() {
		err = results.Scan(&resultQuest.Data)
		if err != nil {
			return nil, err
		}
	}
	// fmt.Println(query)
	// fmt.Println(len(resultQuest.Data))
	return &resultQuest, nil
}
