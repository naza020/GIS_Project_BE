package usecase

import (
	"fmt"
	"gis-project-backend/pkg/module/quest5/model"
)

func (service *quest5UseCase) QuestB() (*model.Quest5BResponse, error) {
	// query := fmt.Sprintf("SELECT [ID],[Country],[City],[Year],[Pm25],[Latitude],[Longitude],[Population],[Wbinc16_text],[Region],[Conc_pm25],[Color_pm25] FROM dbo.AirPollutionPM25 WHERE Year=%s", year)
	country := "Bangkok"
	firstQuery := fmt.Sprintf("DECLARE @%s geometry = 'POLYGON EMPTY'", country)
	secQuery := fmt.Sprintf("SELECT @%s = geom FROM [dbo].[AirPollutionPM25] WHERE City = '%s'", country, country)
	thirdQuery := fmt.Sprintf("SELECT TOP 51 [ID],[Country],[City],[Year],[Pm25],[Latitude],[Longitude],[Population],[Wbinc16_text],[Region],[Conc_pm25],[Color_pm25], geom.MakeValid().STDistance(@%s.MakeValid()) AS Distance FROM [dbo].[AirPollutionPM25] ORDER BY Distance", country)
	query := fmt.Sprintf("%s %s %s", firstQuery, secQuery, thirdQuery)
	results, err := service.CoreRegistry.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var resultQuest model.Quest5BResponse
	for results.Next() {
		var tempData model.Quest5BData
		err = results.Scan(&tempData.ID, &tempData.Country, &tempData.City, &tempData.Year, &tempData.Pm25, &tempData.Latitude, &tempData.Longitude,
			&tempData.Population, &tempData.Wbinc16, &tempData.Region, &tempData.ConcPm25, &tempData.ColorPm25, &tempData.Distance)
		if err != nil {
			return nil, err
		}
		resultQuest.Data = append(resultQuest.Data, tempData)
	}
	fmt.Println(query)
	fmt.Println(len(resultQuest.Data))
	return &resultQuest, nil
}
