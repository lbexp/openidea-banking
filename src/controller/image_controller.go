package controller

import (
	"openidea-banking/src/config"
	"openidea-banking/src/service"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type ImageController struct {
	AuthService  service.AuthService
	ImageService service.ImageService
}

func NewImageController(
	authService service.AuthService,
	imageService service.ImageService,
) ImageController {
	return ImageController{
		AuthService:  authService,
		ImageService: imageService,
	}
}

func (controller *ImageController) UploadImage(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "no image found")
	}

	files := form.File["file"]

	if len(files) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "no image found")
	}

	file := files[0]

	if file.Size > config.MaxUploadSize || file.Size < config.MinUploadSize {
		return fiber.NewError(fiber.StatusBadRequest, "image more than 2MB or less than 10KB")
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" {
		return fiber.NewError(fiber.StatusBadRequest, "not *.jpg | *.jpeg")
	}

	src, err := file.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	url, err := controller.ImageService.UploadImage(ctx.UserContext(), src, file.Filename)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "something error")
	}

	return ctx.JSON(map[string]any{
		"message": "File uploaded sucessfully",
		"data": map[string]string{
			"imageUrl": url,
		},
	})
}
