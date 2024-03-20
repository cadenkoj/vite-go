package controllers_v1

import (
	"net/http"

	"github.com/cadenkoj/vera/backend/db"
	"github.com/cadenkoj/vera/backend/model"
	"github.com/cadenkoj/vera/backend/utils"
	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error {
	tx := db.GetDB().Begin()
	slug := c.Param("slug")
	profile := model.Profile{}

	if err := tx.Where("user_id = ?", slug).First(&profile).Error; err != nil {
		tx.Rollback()
		return c.JSON(utils.NewError(http.StatusNotFound))
	}

	if err := tx.Commit().Error; err != nil {
		return c.JSON(utils.NewError(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, profile)
}

func GetProfiles(c echo.Context) error {
	tx := db.GetDB().Begin()
	profiles := []model.Profile{}

	if err := tx.Find(&profiles).Error; err != nil {
		tx.Rollback()
		return c.JSON(utils.NewError(http.StatusNotFound))
	}

	if err := tx.Commit().Error; err != nil {
		return c.JSON(utils.NewError(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, profiles)
}

func PostProfile(c echo.Context) error {
	tx := db.GetDB().Begin()
	profile := model.Profile{}

	if err := c.Bind(&profile); err != nil {
		return c.JSON(utils.NewError(http.StatusBadRequest))
	}

	if err := tx.Create(&profile).Error; err != nil {
		tx.Rollback()
		return c.JSON(utils.NewError(http.StatusInternalServerError))
	}

	if err := tx.Commit().Error; err != nil {
		return c.JSON(utils.NewError(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusCreated, profile)
}

func PatchProfile(c echo.Context) error {
	tx := db.GetDB().Begin()
	slug := c.Param("slug")
	profile := model.Profile{}

	if err := c.Bind(&profile); err != nil {
		return c.JSON(utils.NewError(http.StatusBadRequest))
	}

	if err := tx.Model(&profile).Where("user_id = ?", slug).Updates(&profile).Error; err != nil {
		tx.Rollback()
		return c.JSON(utils.NewError(http.StatusInternalServerError))
	}

	if err := tx.Commit().Error; err != nil {
		return c.JSON(utils.NewError(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, profile)
}

func DeleteProfile(c echo.Context) error {
	tx := db.GetDB().Begin()
	slug := c.Param("slug")
	profile := model.Profile{}

	if err := tx.Where("user_id = ?", slug).Delete(&profile).Error; err != nil {
		tx.Rollback()
		return c.JSON(utils.NewError(http.StatusInternalServerError))
	}

	if err := tx.Commit().Error; err != nil {
		return c.JSON(utils.NewError(http.StatusInternalServerError))
	}

	return c.NoContent(http.StatusNoContent)
}
