package middlewares

import (
	"errors"
	"strconv"

	"example.com/go-admin/db"
	"example.com/go-admin/models"
	"example.com/go-admin/utility"
	"github.com/gofiber/fiber/v2"
)

func IsAuthorized(c *fiber.Ctx, page string) error {

	cookie := c.Cookies("jwt")

	_, claims, err := utility.ParseJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
	}

	userId, _ := strconv.Atoi(claims.Issuer)

	user := models.User{
		Id: uint(userId),
	}

	db.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	db.DB.Preload("Permissions").Find(&role)

	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")

}
