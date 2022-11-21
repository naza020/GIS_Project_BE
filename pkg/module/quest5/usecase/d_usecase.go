package usecase

import (
	"gis-project-backend/pkg/module/quest5/model"
	"strconv"
	"strings"
)

func (service *quest5UseCase) QuestD() (*model.Quest5DResponse, error) {
	query := `DECLARE @Thai geometry = 'POLYGON EMPTY'
	SELECT @Thai = geom
	FROM [dbo].world
	WHERE NAME='Thailand'
	SELECT @Thai.STEnvelope().ToString();
	`
	results, err := service.CoreRegistry.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var resultQuest model.Quest5DResponse
	var tempData string
	results.Next()
	err = results.Scan(&tempData)
	if err != nil {
		return nil, err
	}
	sTrim := strings.ReplaceAll(tempData, "((", " ")
	sTrim = strings.ReplaceAll(sTrim, "))", " ")
	sTrim = strings.ReplaceAll(sTrim, ",", "")
	// fmt.Println("sTrim := ", sTrim)
	sSplit := strings.Split(sTrim, " ")
	// for col, row := range sSplit {
	// 	// fmt.Println(col, " =", row)
	// }
	for i := 0; i < 4; i++ {
		lat, err := strconv.ParseFloat(sSplit[2+(i*2)], 64)
		if err != nil {
			return nil, err
		}
		long, err := strconv.ParseFloat(sSplit[3+(i*2)], 64)
		if err != nil {
			return nil, err
		}
		resultQuest.Data = append(resultQuest.Data, model.Quest5DMBR{Latitude: lat, Longtitude: long})
	}
	// fmt.Println("Data := ", sSplit)
	// fmt.Println(len(resultQuest.Data))
	return &resultQuest, nil
}
