package controller

import (
	"github.com/labstack/echo/v4"

	"net/http"
	"time"
)

type Status struct {
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Date    time.Time `json:"date"`
}

func NewStatusController(projectName string, projectVersion string) Status {
	return Status{
		Name:    projectName,
		Version: projectVersion,
		Date:    time.Now(),
	}
}

func (s Status) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, s)
}
