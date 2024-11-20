package main

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api"
	userController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user"
	userService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/user"
	userRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/repository/user"

	bukuController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku"
	bukuService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/buku"
	bukuRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/repository/buku"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/configs"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Koneksi ke database MySQL
	db := configs.MySQLConn()

	// Membuat repository dan service untuk user
	userRepo := userRepository.NewUserRepository(db)
	userServices := userService.NewUserService(userRepo)
	userCon := userController.NewUserController(userServices)

	// Membuat repository dan service untuk buku
	bukuRepo := bukuRepository.NewBukuRepository(db)
	bukuServices := bukuService.NewBukuService(bukuRepo)
	bukuCon := bukuController.NewBukuController(bukuServices)

	// Konfigurasi server Fiber
	config := configs.ServerTimeOut()
	app := fiber.New(config)

	// Menambahkan path API untuk user dan buku
	api.RegisterPath(app, userCon, bukuCon)

	// Menjalankan aplikasi di port 8080
	log.Fatal(app.Listen(":8080"))
}
