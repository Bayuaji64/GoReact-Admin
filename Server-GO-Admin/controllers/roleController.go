package controllers

import (
	"strconv"

	"example.com/go-admin/db"
	"example.com/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func Allrole(c *fiber.Ctx) error {
	var role []models.Role

	db.DB.Find(&role)

	return c.JSON(role)
}

func CreateRoleSTRING(c *fiber.Ctx) error {
	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {

		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	if err := db.DB.Create(&role).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create role"})
	}

	return c.JSON(role)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	list, ok := roleDTO["permissions"].([]interface{})
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Permissions format is invalid"})
	}

	permissions := make([]models.Permission, len(list))
	for i, v := range list {
		switch permissionId := v.(type) {
		case float64: // JSON numbers are decoded as float64
			permissions[i] = models.Permission{Id: uint(permissionId)}
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid permission ID format, expected a number"})
		}
	}

	role := models.Role{
		Name:        roleDTO["name"].(string),
		Permissions: permissions, // Ensure this matches your struct field name for permissions
	}

	if err := db.DB.Create(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create role"})
	}

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	db.DB.Find(&role)

	return c.JSON(role)

}

func UpdateRole(c *fiber.Ctx) error {
	roleID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid role ID"})
	}

	var roleDTO fiber.Map
	if err := c.BodyParser(&roleDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	permissionsInterface, ok := roleDTO["permissions"].([]interface{})
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Permissions format is invalid"})
	}

	var permissions []models.Permission
	for _, v := range permissionsInterface {
		permissionID, ok := v.(float64) // Permissions are received as float64
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid permission ID format"})
		}
		permissions = append(permissions, models.Permission{Id: uint(permissionID)})
	}

	tx := db.DB.Begin()

	// Update the role with new name
	if err := tx.Model(&models.Role{}).Where("id = ?", roleID).Updates(models.Role{Name: roleDTO["name"].(string)}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update role"})
	}

	var role models.Role
	if err := tx.First(&role, roleID).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Role not found"})
	}

	// Replace permissions
	if err := tx.Model(&role).Association("Permission").Replace(permissions); err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update permissions"})
	}

	tx.Commit()
	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	db.DB.Delete(&role)
	return c.JSON(fiber.Map{"message": "Delete user success"})

}
