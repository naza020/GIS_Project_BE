package usecase

import (
	"fmt"
	"gis-project-backend/pkg/module/quest5/model"
)

func (service *quest5UseCase) QuestC() (*model.Quest5CResponse, error) {
	// query := fmt.Sprintf("SELECT [ID],[Country],[City],[Year],[Pm25],[Latitude],[Longitude],[Population],[Wbinc16_text],[Region],[Conc_pm25],[Color_pm25] FROM dbo.AirPollutionPM25 WHERE Year=%s", year)
	var country []string
	fQuery := `DECLARE @Country geometry = 'POLYGON EMPTY'
	SELECT @Country = geom
	FROM dbo.world
	WHERE NAME = 'Thailand'
	SELECT Name
	FROM dbo.world
	WHERE Geom.MakeValid().STTouches(@Country.MakeValid())=1
	`
	results, err := service.CoreRegistry.DB.Query(fQuery)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		var tempData string
		err = results.Scan(&tempData)
		if err != nil {
			return nil, err
		}
		country = append(country, tempData)
	}
	countryString := ""
	for col, row := range country {
		countryString += "'" + row + "'"
		if col != len(country)-1 {
			countryString += ","
		}
	}
	secQuery := fmt.Sprintf(`SELECT Country,City,Latitude,Longitude
	FROM [SpatialDB].[dbo].[AirPollutionPM25] WHERE Country IN (%s) 
	GROUP BY Country,City,Latitude,Longitude
  `, countryString)
	fmt.Println("sQuery=", secQuery)
	secResults, err := service.CoreRegistry.DB.Query(secQuery)
	if err != nil {
		return nil, err
	}
	var resultQuest model.Quest5CResponse
	for secResults.Next() {
		var tempData model.Quest5CData
		err = secResults.Scan(&tempData.Country, &tempData.City, &tempData.Latitude, &tempData.Longitude)
		if err != nil {
			return nil, err
		}
		resultQuest.Data = append(resultQuest.Data, tempData)
	}
	fmt.Println(secResults)
	fmt.Println(len(resultQuest.Data))
	return &resultQuest, nil
}
