package core

import (
	"database/sql"
	"gis-project-backend/pkg/core/client"

	"gis-project-backend/pkg/core/logger"
)

type CoreRegistry struct {
	Logger     logger.Logger
	RestClient client.RestClient
	DB         *sql.DB
}
