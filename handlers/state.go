package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurbekchymbaev/sanctions-go-app/database"
)

func State(c *fiber.Ctx) error {

	var result int64
	database.DB.Db.Table("entries").Count(&result)
	if result == 0 {
		return c.Status(200).JSON(&fiber.Map{
			"result": false,
			"info":   "empty",
		})
	}

	var count int64
	database.DB.Db.Table("pg_locks").
		Where("locktype = 'advisory'").
		Where("objid = ?", database.LockKey).
		Where("pid=pg_backend_pid()").
		Count(&count)
	if count > 0 {
		return c.Status(200).JSON(&fiber.Map{
			"result": false,
			"info":   "updating",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"result": true,
		"info":   "ok",
	})
}
