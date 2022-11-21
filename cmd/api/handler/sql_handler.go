package handler

import (
	"gis-project-backend/cmd/api/core/api"
	"gis-project-backend/pkg/core"

	sqlQuest "gis-project-backend/pkg/module/sql"
	"gis-project-backend/pkg/module/sql/model"

	"github.com/gofiber/fiber/v2"
)

type SQLAPIHandler struct {
	app          *fiber.App
	coreRegistry *core.CoreRegistry
	SQLUseCase   sqlQuest.SQLUseCase
}

func NewSQLAPIHandler(app *fiber.App, coreRegistry *core.CoreRegistry, SQLUseCase sqlQuest.SQLUseCase) *SQLAPIHandler {
	return &SQLAPIHandler{
		app:          app,
		coreRegistry: coreRegistry,
		SQLUseCase:   SQLUseCase,
	}

}

func (handler *SQLAPIHandler) Init() {
	endpoint := "/api/sql"

	router := handler.app.Group(endpoint)
	router.Post("/insert", handler.InsertSQL)

}

// @Tags SQL
// @Summary  Insert Data
// @Produce json
// @Param body body model.InsertDataRequest true "body"
// @Success 200 {object} model.InsertDataResponse
// @Router /api/sql/insert [post]
func (handler *SQLAPIHandler) InsertSQL(c *fiber.Ctx) error {
	req := &model.InsertDataRequest{}
	return api.HandlerWithBody(c, req, func() (interface{}, error) {
		return handler.SQLUseCase.Insert(req)
	})
}
