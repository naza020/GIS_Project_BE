package usecase

import (
	"gis-project-backend/pkg/module/quest4/model"
)

func (service *quest4UseCase) QuestB() (*model.Quest4BResponse, error) {
	results, err := service.CoreRegistry.DB.Query("SELECT Country,AVG(Pm25) as AvgPm25 FROM dbo.AirPollutionPM25 GROUP BY Country ORDER BY AvgPm25 DESC")
	if err != nil {
		return nil, err
	}
	var resultQuest model.Quest4BResponse
	for results.Next() {
		var tempData model.Quest4BData
		err = results.Scan(&tempData.Country, &tempData.AvgPm25)
		if err != nil {
			return nil, err
		}
		resultQuest.Data = append(resultQuest.Data, tempData)
	}
	// fmt.Println(len(resultQuest.Data))
	return &resultQuest, nil
}
