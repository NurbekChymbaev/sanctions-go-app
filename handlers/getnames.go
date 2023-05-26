package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurbekchymbaev/sanctions-go-app/database"
	"strings"
)

type QueryOut struct {
	Name string `query:"name"`
	Type string `query:"type"`
}

func Getnames(c *fiber.Ctx) error {

	var input QueryOut
	c.QueryParser(&input)

	if input.Name == "" {
		return c.JSON(&fiber.Map{
			"result": false,
			"info":   "name parameter is required!",
		})
	}
	firstname := strings.ToLower(input.Name)
	lastname := firstname
	if len(strings.Fields(input.Name)) > 1 {
		parts := strings.Fields(firstname)
		firstname = parts[0]
		lastname = parts[1]
	}

	type response struct {
		Uid        uint
		First_name string
		Last_name  string
	}

	var results []response
	query := database.DB.Db.
		Table("entries").
		Select("entries.id as uid, entries.firstname as first_name, entries.lastname as last_name").
		Joins("LEFT JOIN names ON names.entry_id = entries.id")

	if input.Type == "weak" {
		query.
			Where("entries.firstname ilike '%" + firstname + "%'").
			Or("entries.lastname ilike '%" + lastname + "%'").
			Or("names.firstname ilike '%" + firstname + "%'").
			Or("names.lastname ilike '%" + lastname + "%'")
	} else {
		query.
			Where("lower(entries.firstname) = ?", firstname).
			Or("lower(entries.lastname) = ?", lastname).
			Or("lower(names.firstname) = ?", firstname).
			Or("lower(names.lastname) = ?", lastname)
	}
	query.Group("entries.id").
		Find(&results)

	return c.Status(200).JSON(results)
}
