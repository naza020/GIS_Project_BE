package api

import (
	"errors"
	"fmt"
	"gis-project-backend/pkg/core/model"
	coreMode "gis-project-backend/pkg/core/model"
	"gis-project-backend/pkg/core/utils"
	"log"
	"os"
	"strconv"

	reqId "github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/fiber/v2"
	errs "github.com/pkg/errors"
)

type HandlerFn func() (interface{}, error)
type HandlerUploadFileFn func(file *os.File, fileName string) (interface{}, error)
type HandlerDownloadFileFn func() (*model.FileDownloadResponse, error)

func Handler(c *fiber.Ctx, handlerFn HandlerFn) error {
	data, err := handlerFn()
	if err != nil {
		return ResponseError(c, err)
	}
	return ResponseSuccess(c, data)
}

func HandlerWithBody(c *fiber.Ctx, req interface{}, handlerFn HandlerFn) error {
	if err := c.BodyParser(req); err != nil {
		return ResponseError(c, err)
	}
	data, err := handlerFn()
	if err != nil {
		return ResponseError(c, err)
	}
	return ResponseSuccess(c, data)
}

func HandlerPageQueryWithBody(c *fiber.Ctx, req interface{}, page *coreMode.PageLimit, sort *coreMode.Sort, handlerFn HandlerFn) error {
	if err := c.QueryParser(page); err != nil {
		return err
	}

	pageNumber, err := strconv.ParseInt(c.Query("pageNumber", "1"), 10, 64)
	if err != nil {
		return err
	}
	pageSize, err := strconv.ParseInt(c.Query("pageSize", "10"), 10, 64)
	if err != nil {
		return err
	}
	page.PageNumber = pageNumber
	page.PageSize = pageSize

	if err := c.QueryParser(sort); err != nil {
		return err
	}

	if err := c.BodyParser(req); err != nil {
		return ResponseError(c, err)
	}
	data, err := handlerFn()
	if err != nil {
		return ResponseError(c, err)
	}
	return ResponseSuccess(c, data)
}

func HandlerWithFileQueryData(c *fiber.Ctx, fileKey string, req interface{}, handlerUploadFileFn HandlerUploadFileFn) error {
	uploadFile, err := c.FormFile(fileKey)
	if err != nil {
		return ResponseError(c, err)
	}

	if err := c.QueryParser(req); err != nil {
		return err
	}
	log.Println(req)

	fileNameOriginal := uploadFile.Filename

	tmpFilePath := os.TempDir() + utils.NewSID() + fileNameOriginal
	err = c.SaveFile(uploadFile, tmpFilePath)
	if err != nil {
		return ResponseError(c, err)
	}
	defer os.Remove(tmpFilePath)
	file, err := os.Open(tmpFilePath)
	if err != nil {
		return ResponseError(c, err)
	}
	defer file.Close()

	data, err := handlerUploadFileFn(file, fileNameOriginal)

	if err != nil {
		return ResponseError(c, err)
	}
	return ResponseSuccess(c, data)
}

func HandlerWithFileFormData(c *fiber.Ctx, fileKey string, req interface{}, handlerUploadFileFn HandlerUploadFileFn) error {
	uploadFile, err := c.FormFile(fileKey)
	if err != nil {
		return ResponseError(c, err)
	}

	if err := c.BodyParser(req); err != nil {
		return err
	}
	log.Println(req)

	fileNameOriginal := uploadFile.Filename

	tmpFilePath := os.TempDir() + utils.NewSID() + fileNameOriginal
	err = c.SaveFile(uploadFile, tmpFilePath)
	if err != nil {
		return ResponseError(c, err)
	}
	defer os.Remove(tmpFilePath)
	file, err := os.Open(tmpFilePath)
	if err != nil {
		return ResponseError(c, err)
	}
	defer file.Close()

	data, err := handlerUploadFileFn(file, fileNameOriginal)

	if err != nil {
		return ResponseError(c, err)
	}
	return ResponseSuccess(c, data)
}

// ResponseSuccess ...
func ResponseSuccess(c *fiber.Ctx, payload interface{}) error {
	return c.JSON(payload)
}

// ResponseError ...
func ResponseError(c *fiber.Ctx, err error) error {
	errorStatus := MapErrorStatus(err)
	return ResponseErrorWithCode(c, errorStatus, err)
}

// ResponseErrorWithCode ...
func ResponseErrorWithCode(c *fiber.Ctx, errorStatus *fiber.Error, err error) error {
	fmt.Println(err)
	requestId := utils.InterfaceToString(c.Locals(reqId.ConfigDefault.ContextKey))
	return c.Status(errorStatus.Code).JSON(&model.ApiErrorResponse{
		Status:     false,
		StatusCode: errorStatus.Code,
		Message:    errorStatus.Message + " : " + err.Error(),
		Detail:     "uri:" + c.OriginalURL() + "|x-request-id:" + requestId,
		Cause:      err,
	})
}

// MapErrorStatus ...
func MapErrorStatus(err error) *fiber.Error {
	errCause := errs.Cause(err)
	if errors.Is(errCause, model.ErrUnauthorized) {
		return fiber.ErrUnauthorized
	}
	if errors.Is(errCause, model.ErrValidationFailed) {
		return fiber.ErrBadRequest
	}
	if errors.Is(errCause, model.ErrCallExternalServiceFailed) {
		return fiber.ErrBadRequest
	}
	if errors.Is(errCause, model.ErrNotFound) {
		return fiber.ErrNotFound
	}
	return fiber.ErrInternalServerError
}
