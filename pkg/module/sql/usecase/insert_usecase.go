package usecase

import (
	"fmt"
	"gis-project-backend/pkg/module/sql/model"
)

func (service *sqlUseCase) Insert(req *model.InsertDataRequest) (*model.InsertDataResponse, error) {
	tx, err := service.CoreRegistry.DB.Begin()
	if err != nil {
		return nil, err
	}
	count := 1
	for _, item := range req.Data {
		_, err := tx.Exec("INSERT INTO dbo.AirPollutionPM25(Country,City,Year,Pm25,Latitude,Longitude,Population,Wbinc16_text,Region,Conc_pm25,Color_pm25) VALUES (?,?,?,?,?,?,?,?,?,?,?)",
			item.Country, item.City, item.Year, item.Pm25, item.Latitude, item.Longitude, item.Population, item.Wbinc16, item.Region, item.Conc, item.Color)
		if err != nil {
			fmt.Println("err =", err)
			tx.Rollback()
			return nil, err
		}
		count++
	}
	fmt.Println("count = ", count)
	if err == nil {
		err := tx.Commit()
		if err != nil {
			fmt.Println("err =", err)
			return nil, err
		}
	}
	return &model.InsertDataResponse{
		Message: "Insert Complete",
		Record:  count,
	}, nil
}
