package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

type Country struct {
	Code           string  `json:"code,omitempty" db:"Code"`
	Name           string  `json:"name,omitempty" db:"Name"`
	Continent      string  `json:"continent,omitempty" db:"Continent"`
	Region         string  `json:"region,omitempty" db:"Region"`
	SurfaceArea    float64 `json:"surfaceArea,omitempty" db:"SurfaceArea"`
	IndepYear      int     `json:"indepYear,omitempty" db:"IndepYear"`
	Population     int     `json:"population,omitempty" db:"Population"`
	LifeExpectancy float64 `json:"lifeExpectancy,omitempty" db:"LifeExpectancy"`
	GNP            float64 `json:"gnp,omitempty" db:"GNP"`
	GNPOld         float64 `json:"gnpOld,omitempty" db:"GNPOld"`
	LocalName      string  `json:"localName,omitempty" db:"LocalName"`
	GovernmentForm string  `json:"governmentForm,omitempty" db:"GovernmentForm"`
	HeadOfState    string  `json:"headOfState,omitempty" db:"HeadOfState"`
	Capital        int     `json:"capital,omitempty" db:"Capital"`
	Code2          string  `json:"code2,omitempty" db:"Code2"`
}

var (
	db *sqlx.DB
)

func main() {
	_db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	db = _db

	e := echo.New()

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.POST("/post", addCityInfoHandler)

	e.Start(":4000")

}

func getCityInfoHandler(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	city := City{}
	db.Get(&city, "SELECT * FROM city WHERE Name=?", cityName)
	if city.Name == "" {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, city)
}

func addCityInfoHandler(c echo.Context) error {
	var city City

	if err := c.Bind(&city); err != nil {
		return c.JSON(http.StatusBadRequest, city)
	}
	db.Exec(fmt.Sprintf("INSERT INTO city (Name, CountryCode, District, Population) VALUES ('%s', '%s', '%s', %d);", city.Name, city.CountryCode, city.District, city.Population))
	return c.String(http.StatusOK, "Success")
}

// fmt.Println("Connected!")

// city := City{}
// country := Country{}
// // db.Get(&city, "SELECT * FROM city WHERE Name='Tokyo'")
// db.Get(&city, fmt.Sprintf("SELECT * FROM city WHERE Name='%s'", os.Args[1]))
// db.Get(&country, fmt.Sprintf("SELECT * FROM country WHERE Code='%s'", city.CountryCode))

// // fmt.Printf("%sの人口は%d人です\n", os.Args[1], city.Population)
// fmt.Printf("%sの人口は%sの人口の%f%%です\n", city.Name, country.Name, float64(city.Population)/float64(country.Population)*100)

// cities := []City{}
// db.Select(&cities, "SELECT * FROM city WHERE CountryCode='JPN'")

// fmt.Println("日本の都市一覧")
// for _, city := range cities {
// 	fmt.Printf("都市名: %s, 人口: %d人\n", city.Name, city.Population)
// }

// db.Exec("INSERT INTO city (Name, CountryCode, District, Population) VALUES ('oookayama', 'JPN', 'Tokyo', 2147483647);")
