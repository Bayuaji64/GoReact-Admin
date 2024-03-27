package controllers

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	url := os.Getenv("URL_UPLOAD")

	if url == "" {
		return errors.New("URL_UPLOAD is not set in the environment variables")
	}
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["image"]
	filename := ""

	for _, file := range files {

		filename = file.Filename

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}

	}

	return c.JSON(fiber.Map{

		"url": url + filename,
	})

}
