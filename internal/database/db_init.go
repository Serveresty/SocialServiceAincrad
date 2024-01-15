package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func DB_Init(dbUrl string) error {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return fmt.Errorf("Error while connecting to database: %v", err)
	}

	DB = conn
	return nil
}

func CreateBaseTables() error {
	_, err := DB.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS "users_data" (user_id bigserial PRIMARY KEY, first_name VARCHAR(50) NOT NULL, last_name VARCHAR(50) NOT NULL, sex VARCHAR(10) NOT NULL, username VARCHAR(50) UNIQUE, email VARCHAR(255) UNIQUE NOT NULL, phone VARCHAR(255) UNIQUE NOT NULL, password VARCHAR(255) NOT NULL);
	CREATE TABLE IF NOT EXISTS "roles" (role_id serial PRIMARY KEY, role_name VARCHAR(20) UNIQUE NOT NULL);
	CREATE TABLE IF NOT EXISTS "users_roles" (user_id bigint references users_data (user_id) on delete cascade, role_id int references roles (role_id) on delete cascade);
	CREATE TABLE IF NOT EXISTS "social" (social_id bigserial PRIMARY KEY, personal_web VARCHAR(255), instagram VARCHAR(50), steam VARCHAR(50));
	CREATE TABLE IF NOT EXISTS "info" (info_id bigserial PRIMARY KEY, short_info VARCHAR(255), family_state VARCHAR(20), born_city VARCHAR(170), current_resident VARCHAR(170), social_info bigint references social (social_id) on delete cascade);
	CREATE TABLE IF NOT EXISTS "users_info" (user_id bigint references users_data (user_id) on delete cascade, info_id bigint references info (info_id) on delete cascade);
	INSERT INTO "roles" ("role_name") VALUES ('user'), ('support'), ('moderator'), ('admin');
	`)
	if err != nil {
		return fmt.Errorf("Error while creating base tables: %v", err)
	}

	return nil
}
