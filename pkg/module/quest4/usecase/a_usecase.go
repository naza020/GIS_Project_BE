package usecase

import "gis-project-backend/pkg/module/quest4/model"

func (service *quest4UseCase) QuestA() (*model.Quest4AResponse, error) {
	results, err := service.CoreRegistry.DB.Query("SELECT [ID],[Country],[City],[Year],[Pm25],[Latitude],[Longtitude],[Population],[Wbinc16_text],[Region],[Conc_pm25],[Color_pm25] FROM dbo.AirPollutionPM25 WHERE Year=2015 AND Pm25>=50")
	if err != nil {
		return nil, err
	}
	var resultQuest model.Quest4AResponse
	for results.Next() {
		var tempData model.Quest4AData
		err = results.Scan(&tempData.ID, &tempData.Country, &tempData.City, &tempData.Year, &tempData.Pm25, &tempData.Latitude, &tempData.Longtitude,
			&tempData.Population, &tempData.Wbinc16, &tempData.Region, &tempData.ConcPm25, &tempData.ColorPm25)
		if err != nil {
			return nil, err
		}
		resultQuest.Data = append(resultQuest.Data, tempData)
	}
	// fmt.Println(len(resultQuest.Data))
	return &resultQuest, nil
}
