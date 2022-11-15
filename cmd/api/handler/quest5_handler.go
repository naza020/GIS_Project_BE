package handler

import (
	"gis-project-backend/cmd/api/core/api"
	"gis-project-backend/pkg/core"
	quest5 "gis-project-backend/pkg/module/quest5"

	"github.com/gofiber/fiber/v2"
)

type Quest5APIHandler struct {
	app           *fiber.App
	coreRegistry  *core.CoreRegistry
	Quest5UseCase quest5.Quest5UseCase
}

func NewQuest5APIHandler(app *fiber.App, coreRegistry *core.CoreRegistry, Quest5UseCase quest5.Quest5UseCase) *Quest5APIHandler {
	return &Quest5APIHandler{
		app:           app,
		coreRegistry:  coreRegistry,
		Quest5UseCase: Quest5UseCase,
	}

}

func (handler *Quest5APIHandler) Init() {
	endpoint := "/api/quest5"

	router := handler.app.Group(endpoint)
	router.Get("/a/:year", handler.GetQuestA)
	router.Get("/b", handler.GetQuestB)
	// router.Get("/c/:country", handler.GetQuestC)
	// router.Get("/d/:year/:color", handler.GetQuestD)
	router.Get("/f/:year", handler.GetQuestF)
}

// @Tags quest5-api
// @Summary Get Quest 5A
// @Produce json
// @Param year path string true "year"
// @Success 200 {object} model.Quest5AResponse
// @Router /api/quest5/a/{year} [get]
func (handler *Quest5APIHandler) GetQuestA(c *fiber.Ctx) error {
	year := c.Params("year")
	return api.Handler(c, func() (interface{}, error) {
		return handler.Quest5UseCase.QuestA(year)
	})
}

// @Tags quest5-api
// @Summary Get Quest 5B
// @Produce json
// @Success 200 {object} model.Quest5BResponse
// @Router /api/quest5/b [get]
func (handler *Quest5APIHandler) GetQuestB(c *fiber.Ctx) error {
	return api.Handler(c, func() (interface{}, error) {
		return handler.Quest5UseCase.QuestB()
	})
}

// @Tags quest5-api
// @Summary Get Quest 5F
// @Produce json
// @Param year path string true "year"
// @Success 200 {object} model.Quest5FResponse
// @Router /api/quest5/f/{year} [get]
func (handler *Quest5APIHandler) GetQuestF(c *fiber.Ctx) error {
	year := c.Params("year")
	return api.Handler(c, func() (interface{}, error) {
		return handler.Quest5UseCase.QuestF(year)
	})
}
