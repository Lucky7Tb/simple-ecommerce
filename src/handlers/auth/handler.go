package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"simple-ecommerce/src/commons/structs"
	"simple-ecommerce/src/config"
	"simple-ecommerce/src/handlers/auth/request"
	model "simple-ecommerce/src/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewAuthHandler(route *fiber.App, db *gorm.DB, validator *config.Validator) {
	route.Post("/api/v1/auth/register", func(c *fiber.Ctx) error {
		request := new(request.RegisterRequest)
		if err := c.BodyParser(request); err != nil {
			fmt.Println(err)
			return err
		}
		validationError := validator.Validate(request)
		if validationError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(validationError)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.KataSandi), bcrypt.DefaultCost)
		if err != nil {
			errors := make([]string, 1)
			errors[0] = "Error while hashing password"
			return c.Status(fiber.StatusInternalServerError).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}

		bornDate, err := time.Parse("02/01/2006", request.TanggalLahir)
		if err != nil {
			errors := make([]string, 1)
			errors[0] = "tanggal lahir must be format DD/MM/YYYY"
			return c.Status(fiber.StatusBadRequest).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}

		db.Begin()
		idProvinsi, _ := strconv.Atoi(request.IdProvinsi)
		response, err := http.Get(fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%d.json", idProvinsi))
		if err != nil {
			errors := make([]string, 1)
			errors[0] = "Cannot validate provinsi"
			return c.Status(fiber.StatusBadRequest).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}

		if response.StatusCode == 404 {
			errors := make([]string, 1)
			errors[0] = "Provinsi not found"
			return c.Status(fiber.StatusBadRequest).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}

		idKota, _ := strconv.Atoi(request.IdKota)
		response, err = http.Get(fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/districts/%d.json", idKota))
		if err != nil {
			errors := make([]string, 1)
			errors[0] = "Cannot validate kota"
			return c.Status(fiber.StatusBadRequest).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}
		if response.StatusCode == 404 {
			errors := make([]string, 1)
			errors[0] = "Kota not found"
			return c.Status(fiber.StatusBadRequest).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}

		user := model.User{Nama: request.Nama, KataSandi: string(hashedPassword), Notelp: request.NoTelepon, Email: request.Email, TanggalLahir: bornDate, Pekerjaan: request.Pekerjaan, IdKota: uint(idKota), IdProvinsi: uint(idProvinsi), IsAdmin: true}
		result := db.Create(&user)
		if result.Error != nil {
			errors := make([]string, 1)
			errors[0] = result.Error.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}

		if err = db.Create(model.Toko{IdUser: user.ID, NamaToko: fmt.Sprintf("toke-%v", request.Nama)}).Error; err != nil {
			db.Rollback()
			errors := make([]string, 1)
			errors[0] = "Internal server error"
			return c.Status(fiber.StatusInternalServerError).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}
		db.Commit()

		return c.Status(fiber.StatusOK).JSON(structs.Response{Status: true, Message: "Succeed to POST data", Errors: nil, Data: "Register success"})
	})

	route.Post("/api/v1/auth/register", func(c *fiber.Ctx) error {
		request := new(request.LoginRequest)
		if err := c.BodyParser(request); err != nil {
			fmt.Println(err)
			return err
		}
		validationError := validator.Validate(request)
		if validationError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(validationError)
		}

		var user model.User
		if err := db.First(&user).Where("notelp = ?", request.NoTelepon).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				errors := make([]string, 1)
				errors[0] = "No Telp atau kata sandi salah"
				return c.Status(fiber.StatusUnauthorized).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
			}

			errors := make([]string, 1)
			errors[0] = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(request.KataSandi))
		if err != nil {
			errors := make([]string, 1)
			errors[0] = "No Telp atau kata sandi salah"
			return c.Status(fiber.StatusUnauthorized).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}

		response, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces")
		if err != nil {
			errors := make([]string, 1)
			errors[0] = "Cannot validate provinsi"
			return c.Status(fiber.StatusBadRequest).JSON(structs.Response{Status: false, Message: "failed to POST data", Errors: errors, Data: nil})
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var provinces
		json.Unmarshal(responseData, &responseObject)

		return c.Status(fiber.StatusOK).JSON(structs.Response{Status: true, Message: "Succeed to POST data", Errors: nil, Data: "login success"})
	})
}
