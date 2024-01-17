package getfromenv

import (
	"SocialServiceAincrad/configs"
	"SocialServiceAincrad/models"
)

func GetDatabaseConData() string {
	db := &models.DBConfig{
		Username: configs.GetEnv("DB_USERNAME"),
		Password: configs.GetEnv("DB_PASSWORD"),
		Host:     configs.GetEnv("DB_HOST"),
		Port:     configs.GetEnv("DB_PORT"),
		DBName:   configs.GetEnv("DB_NAME"),
	}

	dbUrl := "postgres://" + db.Username + ":" + db.Password + "@" + db.Host + ":" + db.Port + "/" + db.DBName

	return dbUrl
}
