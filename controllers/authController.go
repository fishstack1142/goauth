package controllers

import (
	"../models"
	"../database"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func Register(c *fiber.Ctx) error {

	var data map[string]string

	//long version
	// err := c.BodyParser(&data)

	// if err != nil {
	// 	return err
	// }

	//short version
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["Name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
