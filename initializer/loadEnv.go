package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

// Penggunaan huruf kapital pada awal nama fungsi menandakan bahwa fungsi tersebut diekspor (public)
// Penggunaan huruf kecil pada awal nama fungsi menandakan bahwa fungsi tersebut tidak diekspor (private)

func LoadEnvVariablesss() { // Diekspor (Public)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
