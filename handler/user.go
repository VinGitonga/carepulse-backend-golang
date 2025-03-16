package handler

import (
	"CarepluseBackend/dto"
	"CarepluseBackend/helpers"
	"CarepluseBackend/services"
	"github.com/gofiber/fiber/v2"
)

func CreateAdminUser(c *fiber.Ctx) error {
	userData := new(dto.UserCreateAdminDTO)

	if err := c.BodyParser(userData); err != nil {
		return helpers.ResponseError(c, &helpers.ResponseErrorOptions{Msg: "Request Body is required", StatusCode: fiber.StatusBadRequest})
	}

	errors := helpers.ValidateStruct(userData)

	if len(errors) > 0 {
		return helpers.ResponseError(c, &helpers.ResponseErrorOptions{
			StatusCode: fiber.StatusBadRequest,
			Msg:        "Validation Error",
			Errors:     errors,
		})
	}

	user, err := services.CreateAdminUser(userData)

	if err != nil {
		return helpers.ResponseError(c, &helpers.ResponseErrorOptions{
			StatusCode: fiber.StatusBadRequest,
			Msg:        err.Error(),
			Errors:     err,
		})
	}

	// Send an email if successful

	return helpers.ResponseSuccess(c, &helpers.ResponseSuccessOptions{
		StatusCode: fiber.StatusCreated,
		Data:       user,
		Msg:        "User Created",
	})

}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := services.GetUser(id)

	if err != nil {
		return helpers.ResponseError(c, &helpers.ResponseErrorOptions{
			StatusCode: fiber.StatusBadRequest,
			Msg:        err.Error(),
			Errors:     err,
		})
	}

	return helpers.ResponseSuccess(c, &helpers.ResponseSuccessOptions{
		StatusCode: fiber.StatusOK,
		Data:       user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()

	if err != nil {
		return helpers.ResponseError(c, &helpers.ResponseErrorOptions{
			StatusCode: fiber.StatusBadRequest,
			Msg:        err.Error(),
			Errors:     err,
		})
	}

	return helpers.ResponseSuccess(c, &helpers.ResponseSuccessOptions{
		StatusCode: fiber.StatusOK,
		Data:       users,
	})
}

func CreateNewPatient(c *fiber.Ctx) error {
	infoData := new(dto.UserCreatePatientDTO)

	if err := c.BodyParser(infoData); err != nil {
		return helpers.ResponseError(c, &helpers.ResponseErrorOptions{
			StatusCode: fiber.StatusBadRequest,
			Msg:        err.Error(),
			Errors:     err,
		})
	}

	validationErrors := helpers.ValidateStruct(infoData)

	if len(validationErrors) > 0 {
		return helpers.ResponseError(c, &helpers.ResponseErrorOptions{
			StatusCode: fiber.StatusBadRequest,
			Msg:        "Validation Error",
			Errors:     validationErrors,
		})
	}

	new_user_acc, err := services.CreatePatientAccount(&infoData.UserData, &infoData.PatientData)

	if err != nil {
		return helpers.ResponseError(c, &helpers.ResponseErrorOptions{
			StatusCode: fiber.StatusBadRequest,
			Msg:        err.Error(),
			Errors:     err,
		})
	}

	return helpers.ResponseSuccess(c, &helpers.ResponseSuccessOptions{
		StatusCode: fiber.StatusCreated,
		Data:       new_user_acc,
		Msg:        "User Created",
	})
}
