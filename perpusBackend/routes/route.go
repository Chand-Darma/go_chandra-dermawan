package routes

import (
	"project/controllers"
	m "project/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	m.LogMiddlewares(e)

	//admin route
	e.POST("/admins", controllers.CreateAdminController)
	e.GET("/admins", controllers.GetAdminsController)
	e.GET("/admins/:id", controllers.GetAdminsIDController)
	e.PUT("/admins/:id", controllers.UpdateAdminController)
	e.DELETE("/admins/:id", controllers.DeleteAdminController)

	// petugas route
	e.POST("/petugas", controllers.CreatePetugasController)
	e.GET("/petugas", controllers.GetPetugasController)
	e.GET("/petugas/:id", controllers.GetPetugasIDController)
	e.PUT("/petugas/:id", controllers.UpdatePetugasIDController)
	e.DELETE("/petugas/:id", controllers.DeletePetugasIDController)

	// anggota route
	e.POST("/anggota", controllers.CreateAnggotaController)
	e.GET("/anggota", controllers.GetAnggotaController)
	e.GET("/anggota/:id", controllers.GetAnggotaIDController)
	e.PUT("/anggota/:id", controllers.UpdateAnggotaIDController)
	e.DELETE("/anggota/:id", controllers.DeleteAnggotaIDController)

	// buku route
	e.POST("/buku", controllers.CreateBukuController)
	e.GET("/buku", controllers.GetBukuController)
	e.GET("/buku/:id", controllers.GetBukuIDController)
	e.PUT("/buku/:id", controllers.UpdateBukuIDController)
	e.DELETE("/buku/:id", controllers.DeleteBukuIDController)

	// peminjam route
	e.POST("/peminjaman", controllers.CreatePeminjamController)
	e.GET("/peminjaman", controllers.GetPeminjamanController)
	e.GET("/peminjaman/:id", controllers.GetPeminjamanIDController)
	e.PUT("/peminjaman/:id", controllers.UpdatePeminjamIDController)
	e.DELETE("/peminjaman/:id", controllers.DeletePeminjamIDController)

	// peminjam route
	e.POST("/pengembalian", controllers.CreatePengembalianController)
	e.GET("/pengembalian", controllers.GetPengembalianController)
	e.GET("/pengembalian/:id", controllers.GetPengembalianIDController)
	e.PUT("/pengembalian/:id", controllers.UpdatePengembalianIDController)
	e.DELETE("/pengembalian/:id", controllers.DeletePengembalianIDController)

	//auth
	//eBasicAuth := e.Group("/auth")
	//eBasicAuth.Use(echomid.BasicAuth(m.BasicAuthDB))
	//eBasicAuth.GET("/users", controllers.GetUserController)

	//jwt auth user
	//jwtAuth := e.Group("/jwt")
	//jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	//jwtAuth.GET("/users", controllers.GetUsersController)
	//jwtAuth.GET("/users/:id", controllers.GetUserController)
	//jwtAuth.DELETE("/users/:id", controllers.DeleteUserController)
	//jwtAuth.PUT("/users/:id", controllers.UpdateUserController)

	//jwt auth book
	//jwtAuth.GET("/books", controllers.GetBooksController)
	//jwtAuth.GET("/books/:id", controllers.GetBookController)
	//jwtAuth.POST("/books", controllers.CreateBookController)
	//jwtAuth.DELETE("/books/:id", controllers.DeleteBookController)
	//jwtAuth.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
