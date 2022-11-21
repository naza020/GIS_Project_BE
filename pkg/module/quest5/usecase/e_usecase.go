package usecase

import (
	"fmt"
	"gis-project-backend/pkg/module/quest5/model"
)

func (service *quest5UseCase) QuestE() (*model.Quest5EResponse, error) {
	fQuery := `SELECT TOP 1  Country,COUNT(City) as noOfCity
	FROM [SpatialDB].[dbo].[AirPollutionPM25] WHERE Year=2011 GROUP BY Country ORDER BY noOfCity DESC;
	`
	results, err := service.CoreRegistry.DB.Query(fQuery)
	if err != nil {
		return nil, err
	}
	results.Next()
	var tempCountry string
	var tempNo int
	err = results.Scan(&tempCountry, &tempNo)
	if err != nil {
		return nil, err
	}

	secQuery := fmt.Sprintf(`SELECT ID,Country,City,Latitude,Longitude
	FROM [dbo].[AirPollutionPM25] WHERE Country ='%s' AND Year=2011`, tempCountry)
	fmt.Println("sQuery=", secQuery)
	secResults, err := service.CoreRegistry.DB.Query(secQuery)
	if err != nil {
		return nil, err
	}
	var resultQuest model.Quest5EResponse
	for secResults.Next() {
		var tempData model.Quest5EData
		err = secResults.Scan(&tempData.ID, &tempData.Country, &tempData.City, &tempData.Latitude, &tempData.Longitude)
		if err != nil {
			return nil, err
		}
		resultQuest.Data = append(resultQuest.Data, tempData)
	}
	fmt.Println(secResults)
	fmt.Println(len(resultQuest.Data))
	return &resultQuest, nil
}
