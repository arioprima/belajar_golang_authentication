package test

import (
	"database/sql"
	"fmt"
	"github.com/arioprima/belajar_golang_authentication/controller"
	"github.com/arioprima/belajar_golang_authentication/helper"
	"github.com/arioprima/belajar_golang_authentication/repository"
	"github.com/arioprima/belajar_golang_authentication/route"
	"github.com/arioprima/belajar_golang_authentication/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func ConnectionTestDB() (*sql.DB, error) {
	err := godotenv.Load("../.env")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME_TEST"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db, nil
}

func setupRouter() *gin.Engine {
	db, err := ConnectionTestDB()
	helper.PanicIfError(err)

	validate := validator.New()

	customersRespository := repository.NewCustomersRepositoryImpl(db)

	customersService := service.NewCustomersServiceImpl(customersRespository, db, validate)
	customersServiceAuth := service.NewCustomersServiceAuth(customersRespository, db, validate)

	customersController := controller.NewCustomersControllerImpl(customersService)
	customersAuthController := controller.NewCustomersAuthController(customersServiceAuth)

	router := route.NewRouter(customersController, customersAuthController)

	return router
}

func TestRegisterSuccess(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{
	"firstname": "Itadori",
	"lastname": "Yuji",
	"username": "itadori",
	"email": "testing@test.com",
	"password": "test1234"
}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/customers/register", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestRegisterFailed(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{
	"firstname": "Itadori",
	"lastname": "Yuji",
	"username": "itadori",
	"email": "asdasdad",
	"password": "test1234"
}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/customers/register", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 500, response.StatusCode)
}

func TestLoginSuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{
	"username": "itadori",
	"password": "test1234"
}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/customers/login", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestLoginFailed(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{
	"username": "itadoasdasdasdadri",
	"password": "test"
}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/customers/login", requestBody)

	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 500, response.StatusCode)
}
