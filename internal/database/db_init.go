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
	CREATE TABLE IF NOT EXISTS "social" (user_id bigint references users_data (user_id) on delete cascade, personal_web VARCHAR(255), instagram VARCHAR(50), steam VARCHAR(50));
	CREATE TABLE IF NOT EXISTS "info" (user_id bigint references users_data (user_id) on delete cascade, short_info VARCHAR(255), family_state VARCHAR(20), born_city VARCHAR(170), current_resident VARCHAR(170), social_info bigint references social (social_id) on delete cascade);
	CREATE TABLE IF NOT EXISTS "friend_status" (status_id serial PRIMARY KEY, status_name VARCHAR(50) UNIQUE);
	CREATE TABLE IF NOT EXISTS "friends" (first bigint references users_data (user_id) on delete cascade, second bigint references users_data (user_id) on delete cascade, status_id int references friend_status (status_id));
	
	CREATE TABLE IF NOT EXISTS "front_styles" (style_id serial PRIMARY KEY, style_name VARCHAR(20) UNIQUE NOT NULL);
	CREATE TABLE IF NOT EXISTS "general_settings" (user_id bigint references users_data (user_id) on delete cascade, front_style int references front_styles (style_id) NOT NULL, censure_filter bool, language VARCHAR(20) NOT NULL);
	
	CREATE TABLE IF NOT EXISTS "privacy_statuses" (status_id serial PRIMARY KEY, status_name VARCHAR(20) UNIQUE NOT NULL)
	CREATE TABLE IF NOT EXISTS "privacy_settings" (user_id bigint references users_data (user_id) on delete cascade, saved_photos int references privacy_statuses (status_id), groups int references privacy_statuses (status_id), audio int references privacy_statuses (status_id), video int references privacy_statuses (status_id), friends int references privacy_statuses (status_id), posts int references privacy_statuses (status_id), comments int references privacy_statuses (status_id), messages int references privacy_statuses (status_id);
	
	CREATE TABLE IF NOT EXISTS "blacklist" (user_id bigint references users_data (user_id) on delete cascade, blocked_user_id bigint references users_data (user_id) on delete cascade);
	`)
	if err != nil {
		return fmt.Errorf("Error while creating base tables: %v", err)
	}

	_, err = DB.Exec(context.Background(), `
	INSERT INTO "roles" ("role_name") VALUES ('user'), ('support'), ('moderator'), ('admin');
	`)
	if err != nil {
	}

	_, err = DB.Exec(context.Background(), `
	INSERT INTO "friend_status" ("status_name") VALUES ('wait'), ('friend'), ('follower');
	`)
	if err != nil {
	}

	_, err = DB.Exec(context.Background(), `
	INSERT INTO "front_styles" ("style_name") VALUES ('default'), ('aincrad');
	`)
	if err != nil {
	}

	_, err = DB.Exec(context.Background(), `
	INSERT INTO "privacy_statuses" ("status_name") VALUES ('all'), ('friends'), ('nobody');
	`)
	if err != nil {
	}

	return nil
}
