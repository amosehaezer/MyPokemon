package db

import (
	"fmt"
	"log"
	"phincon/pkg/common/config"
	"phincon/pkg/common/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(c *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.PokemonMoves{})
	db.AutoMigrate(&models.PokemonType{})
	db.AutoMigrate(&models.Sprites{})
	db.AutoMigrate(&models.Pokemon{})
	db.AutoMigrate(&models.MyPokemon{})

	return db
}
