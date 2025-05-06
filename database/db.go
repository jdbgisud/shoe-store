package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла: ", err)
	}

	// Собираем строку подключения из переменных окружения
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable TimeZone=Asia/Almaty"

	// Подключаемся к базе данных
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	DB = database

	RunMigrationsFromSQL()
}

func RunMigrationsFromSQL() {
	files, err := filepath.Glob("database/migrations/*.sql")
	if err != nil {
		log.Fatal("Ошибка при чтении миграций: ", err)
	}

	if len(files) == 0 {
		log.Println("Нет SQL миграций для выполнения.")
		return
	}

	sort.Strings(files)

	for _, file := range files {
		sqlBytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Ошибка при чтении файла миграции %s: %v", file, err)
		}

		sql := string(sqlBytes)
		if strings.TrimSpace(sql) != "" {
			log.Printf("Выполняется миграция: %s", file)
			if err := DB.Exec(sql).Error; err != nil {
				log.Fatalf("Ошибка выполнения миграции %s: %v", file, err)
			}
		}
	}

	log.Println("Все SQL миграции успешно выполнены")
}
