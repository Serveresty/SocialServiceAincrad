package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *pgx.Conn
var Mongo *mongo.Client

func DB_Init(dbUrl string) error {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return fmt.Errorf("error while connecting to database: %v", err)
	}

	DB = conn
	return nil
}

func Mongo_Init(mongoUrl string, ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return fmt.Errorf("error while connecting to mongodb: %v", err)
	}

	Mongo = client
	return nil
}

func CreateBaseTables() error {
	_, err := DB.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS "users_data" (user_id bigserial PRIMARY KEY, first_name VARCHAR(50) NOT NULL, last_name VARCHAR(50) NOT NULL, sex VARCHAR(10) NOT NULL, username VARCHAR(50) UNIQUE DEFAULT '', email VARCHAR(255) UNIQUE NOT NULL, phone VARCHAR(255) UNIQUE NOT NULL, password VARCHAR(255) NOT NULL);
	CREATE TABLE IF NOT EXISTS "roles" (role_id serial PRIMARY KEY, role_name VARCHAR(20) UNIQUE NOT NULL);
	CREATE TABLE IF NOT EXISTS "users_roles" (user_id bigint references users_data (user_id) on delete cascade, role_id int references roles (role_id) on delete cascade);
	CREATE TABLE IF NOT EXISTS "social" (user_id bigint references users_data (user_id) on delete cascade, personal_web VARCHAR(255) DEFAULT '', instagram VARCHAR(50) DEFAULT '', steam VARCHAR(50) DEFAULT '');
	CREATE TABLE IF NOT EXISTS "info" (user_id bigint references users_data (user_id) on delete cascade, short_info VARCHAR(255) DEFAULT '', family_state VARCHAR(20) DEFAULT '', born_city VARCHAR(170) DEFAULT '', current_resident VARCHAR(170) DEFAULT '');
	CREATE TABLE IF NOT EXISTS "friend_status" (status_id serial PRIMARY KEY, status_name VARCHAR(50) UNIQUE);
	CREATE TABLE IF NOT EXISTS "friends" (first bigint references users_data (user_id) on delete cascade, second bigint references users_data (user_id) on delete cascade, status_id int references friend_status (status_id));

	CREATE TABLE IF NOT EXISTS "front_styles" (style_id serial PRIMARY KEY, style_name VARCHAR(20) UNIQUE NOT NULL);
	CREATE TABLE IF NOT EXISTS "general_settings" (user_id bigint references users_data (user_id) on delete cascade, front_style int references front_styles (style_id) NOT NULL, censure_filter bool NOT NULL, language VARCHAR(20) NOT NULL);
	
	CREATE TABLE IF NOT EXISTS "privacy_statuses" (status_id serial PRIMARY KEY, status_name VARCHAR(20) UNIQUE NOT NULL);
	CREATE TABLE IF NOT EXISTS "privacy_settings" (user_id bigint references users_data (user_id) on delete cascade, saved_photos int references privacy_statuses (status_id) NOT NULL, groups int references privacy_statuses (status_id) NOT NULL, audio int references privacy_statuses (status_id) NOT NULL, video int references privacy_statuses (status_id) NOT NULL, friends int references privacy_statuses (status_id) NOT NULL, posts int references privacy_statuses (status_id) NOT NULL, comments int references privacy_statuses (status_id) NOT NULL, messages int references privacy_statuses (status_id) NOT NULL);
	
	CREATE TABLE IF NOT EXISTS "blacklist" (user_id bigint references users_data (user_id) on delete cascade, blocked_user_id bigint references users_data (user_id) on delete cascade);
	
	CREATE TABLE IF NOT EXISTS "users_songs" (user_id bigint references users_data (user_id) on delete cascade, songs_list bigint[]);

	CREATE TABLE IF NOT EXISTS "songs" (song_id bigserial PRIMARY KEY, name VARCHAR(55) NOT NULL, author VARCHAR(55) NOT NULL, filename VARCHAR(55) UNIQUE NOT NULL, hash VARCHAR(70) UNIQUE NOT NULL);

	CREATE TABLE IF NOT EXISTS "users_videos" (user_id bigint references users_data (user_id) on delete cascade, videos_list bigint[]);

	CREATE TABLE IF NOT EXISTS "videos" (video_id bigserial PRIMARY KEY, title VARCHAR(55) NOT NULL, description VARCHAR(255), created_at TIMESTAMP, views int);

	CREATE TABLE IF NOT EXISTS "messages" (id bigserial PRIMARY KEY, sender_id bigint references users_data (user_id) on delete cascade, receiver_id bigint references users_data (user_id) on delete cascade, message VARCHAR(255), created_at TIMESTAMP);
	`)
	if err != nil {
		return fmt.Errorf("error while creating base tables: %v", err)
	}

	_, _ = DB.Exec(context.Background(), `
	INSERT INTO "roles" ("role_name") VALUES ('user'), ('support'), ('moderator'), ('admin');
	`)

	_, _ = DB.Exec(context.Background(), `
	INSERT INTO "friend_status" ("status_name") VALUES ('wait'), ('friend'), ('follower');
	`)

	_, _ = DB.Exec(context.Background(), `
	INSERT INTO "front_styles" ("style_name") VALUES ('default'), ('aincrad');
	`)

	_, _ = DB.Exec(context.Background(), `
	INSERT INTO "privacy_statuses" ("status_name") VALUES ('all'), ('friends'), ('nobody');
	`)

	return nil
}

func CreateDBAndCollectionMongo() error {
	db := Mongo.Database("social")
	collName := "videos"
	collections, err := db.ListCollectionNames(context.Background(), bson.M{"name": collName})
	if err != nil {
		return err
	}
	if len(collections) == 0 {
		log.Printf("Collection %s does not exist. Creating...\n", collName)
		return db.CreateCollection(context.Background(), collName)
	}
	log.Printf("Collection %s already exists.\n", collName)
	return nil
}
