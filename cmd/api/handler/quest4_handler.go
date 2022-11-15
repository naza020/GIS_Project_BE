package handler

import (
	"gis-project-backend/cmd/api/core/api"
	"gis-project-backend/pkg/core"
	quest4 "gis-project-backend/pkg/module/quest4"

	"github.com/gofiber/fiber/v2"
)

type Quest4APIHandler struct {
	app           *fiber.App
	coreRegistry  *core.CoreRegistry
	Quest4UseCase quest4.Quest4UseCase
}

func NewQuest4APIHandler(app *fiber.App, coreRegistry *core.CoreRegistry, Quest4UseCase quest4.Quest4UseCase) *Quest4APIHandler {
	return &Quest4APIHandler{
		app:           app,
		coreRegistry:  coreRegistry,
		Quest4UseCase: Quest4UseCase,
	}

}

func (handler *Quest4APIHandler) Init() {
	endpoint := "/api/quest4"

	router := handler.app.Group(endpoint)
	router.Get("/str", handler.GetEvent)
	router.Get("/a", handler.GetQuestA)
	router.Get("/b", handler.GetQuestB)
	router.Get("/c/:country", handler.GetQuestC)
	router.Get("/d/:year/:color", handler.GetQuestD)
}

// @Tags quest4-api
// @Summary test
// @Produce json
// @Success 200 {object} string
// @Router /api/quest4/str [get]
func (handler *Quest4APIHandler) GetEvent(c *fiber.Ctx) error {
	return api.Handler(c, func() (interface{}, error) {
		return handler.Quest4UseCase.GetString()
	})
}

// @Tags quest4-api
// @Summary Get Quest 4A
// @Produce json
// @Success 200 {object} model.Quest4AResponse
// @Router /api/quest4/a [get]
func (handler *Quest4APIHandler) GetQuestA(c *fiber.Ctx) error {
	return api.Handler(c, func() (interface{}, error) {
		return handler.Quest4UseCase.QuestA()
	})
}

// @Tags quest4-api
// @Summary Get Quest 4B
// @Produce json
// @Success 200 {object} model.Quest4BResponse
// @Router /api/quest4/b [get]
func (handler *Quest4APIHandler) GetQuestB(c *fiber.Ctx) error {
	return api.Handler(c, func() (interface{}, error) {
		return handler.Quest4UseCase.QuestB()
	})
}

// @Tags quest4-api
// @Summary Get Quest 4C
// @Produce json
// @Param country path string true "country"
// @Success 200 {object} model.Quest4CResponse
// @Router /api/quest4/c/{country} [get]
func (handler *Quest4APIHandler) GetQuestC(c *fiber.Ctx) error {
	country := c.Params("country")
	return api.Handler(c, func() (interface{}, error) {
		return handler.Quest4UseCase.QuestC(country)
	})
}

// @Tags quest4-api
// @Summary Get Quest 4D
// @Produce json
// @Param year path string true "year"
// @Param color path string true "color"
// @Success 200 {object} model.Quest4DResponse
// @Router /api/quest4/d/{year}/{color} [get]
func (handler *Quest4APIHandler) GetQuestD(c *fiber.Ctx) error {
	year := c.Params("year")
	color := c.Params("color")
	return api.Handler(c, func() (interface{}, error) {
		return handler.Quest4UseCase.QuestD(year, color)
	})
}
