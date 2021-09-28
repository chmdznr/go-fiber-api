package controllers

import (
	"fiber-api/config"
	"fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
	"strconv"
)

func GetUsers(c *fiber.Ctx) error  {
	var users []models.Users

	config.DB.Preload(clause.Associations).Find(&users)

	return c.JSON(&users)
}

func GetUsersById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.Users{Id: uint(id)}

	config.DB.Preload(clause.Associations).Find(&user)

	return c.JSON(&user)
}

func CreateUsers(c *fiber.Ctx) error {
	var newRecord models.Users

	err := c.BodyParser(&newRecord)
	if err != nil {
		return err
	}

	res := config.DB.Create(&newRecord)
	if res.Error != nil || res.RowsAffected <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "error",
			"message": res.Error,
		})
	}

	return c.JSON(&fiber.Map{
		"message": "Users Created",
		"data": newRecord,
	})
}

func UpdateUsers(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.Users{Id: uint(id)}

	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	config.DB.Preload(clause.Associations).Updates(&user).Find(&user)

	return c.JSON(&fiber.Map{
		"message": "Users Updated",
		"data": user,
	})
}

func DeleteUsers(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.Users{Id: uint(id)}

	config.DB.Delete(&user)

	return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
		"message": "Users Deleted",
	})
}
