-- +goose Up
-- +goose StatementBegin

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(50) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password_hash" varchar(255) NOT NULL,
  "avatar_url" text,
  "cpf" varchar(11) UNIQUE,
  "role" varchar(20) NOT NULL DEFAULT 'user',
  "points" int DEFAULT 0,
  "created_at" timestamp WITH time zone DEFAULT (now()),
  "updated_at" timestamp WITH time zone DEFAULT (now())
);

CREATE TABLE "artist_profiles" (
  "id" bigserial PRIMARY KEY,
  "user_id" int UNIQUE NOT NULL,
  "public_name" varchar(255) NOT NULL,
  "bio" text,
  "banner_url" text,
  "created_at" timestamp WITH time zone DEFAULT (now()),
  "updated_at" timestamp WITH time zone DEFAULT (now())
);

CREATE TABLE "albums" (
  "id" bigserial PRIMARY KEY,
  "artist_id" int NOT NULL,
  "title" varchar(255) NOT NULL,
  "released_at" date,
  "created_at" timestamp WITH time zone DEFAULT (now()),
  "updated_at" timestamp WITH time zone DEFAULT (now())
);

CREATE TABLE "musics" (
  "id" bigserial PRIMARY KEY,
  "album_id" int NOT NULL,
  "title" varchar(255) NOT NULL,
  "length_sec" int,
  "track_num" int,
  "audio_url" text NOT NULL,
  "created_at" timestamp WITH time zone DEFAULT (now())
);

CREATE TABLE "discos" (
  "id" bigserial PRIMARY KEY,
  "album_id" int NOT NULL,
  "format" varchar(50),
  "edition" varchar(255),
  "cover_url" text,
  "price_points" int NOT NULL,
  "created_at" timestamp WITH time zone DEFAULT (now()),
  "updated_at" timestamp WITH time zone DEFAULT (now())
);

CREATE TABLE "user_discos" (
  "user_id" int,
  "disco_id" int,
  "unlocked_at" timestamp WITH time zone DEFAULT (now()),
  PRIMARY KEY(user_id, disco_id)
);

CREATE TABLE "listening_history" (
  "id" bigserial PRIMARY KEY,
  "user_id" int,
  "music_id" int,
  "listened_at" timestamp WITH time zone DEFAULT (now()),
  "points_earned" int DEFAULT 0
);

CREATE TABLE "artist_analytics" (
  "id" bigserial PRIMARY KEY,
  "artist_id" int,
  "music_id" int,
  "listens" int DEFAULT 0,
  "points_from_listens" int DEFAULT 0
);

--------------------------------------------------------------------------------
-- 2. Add Foreign Key Constraints (NOT VALID)
--------------------------------------------------------------------------------

ALTER TABLE "artist_profiles"
  ADD CONSTRAINT fk_artist_profiles_user_id
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") NOT VALID;

ALTER TABLE "albums"
  ADD CONSTRAINT fk_albums_artist_id
  FOREIGN KEY ("artist_id") REFERENCES "users" ("id") NOT VALID;

ALTER TABLE "musics"
  ADD CONSTRAINT fk_musics_album_id
  FOREIGN KEY ("album_id") REFERENCES "albums" ("id") NOT VALID;

ALTER TABLE "discos"
  ADD CONSTRAINT fk_discos_album_id
  FOREIGN KEY ("album_id") REFERENCES "albums" ("id") NOT VALID;

ALTER TABLE "user_discos"
  ADD CONSTRAINT fk_user_discos_user_id FOREIGN KEY ("user_id") REFERENCES "users" ("id") NOT VALID,
  ADD CONSTRAINT fk_user_discos_disco_id FOREIGN KEY ("disco_id") REFERENCES "discos" ("id") NOT VALID;

ALTER TABLE "listening_history"
  ADD CONSTRAINT fk_history_user_id FOREIGN KEY ("user_id") REFERENCES "users" ("id") NOT VALID,
  ADD CONSTRAINT fk_history_music_id FOREIGN KEY ("music_id") REFERENCES "musics" ("id") NOT VALID;

ALTER TABLE "artist_analytics"
  ADD CONSTRAINT fk_analytics_artist_id FOREIGN KEY ("artist_id") REFERENCES "users" ("id") NOT VALID,
  ADD CONSTRAINT fk_analytics_music_id FOREIGN KEY ("music_id") REFERENCES "musics" ("id") NOT VALID;

--------------------------------------------------------------------------------
-- 3. Validate Constraints
--------------------------------------------------------------------------------

ALTER TABLE "artist_profiles" VALIDATE CONSTRAINT fk_artist_profiles_user_id;
ALTER TABLE "albums" VALIDATE CONSTRAINT fk_albums_artist_id;
ALTER TABLE "musics" VALIDATE CONSTRAINT fk_musics_album_id;
ALTER TABLE "discos" VALIDATE CONSTRAINT fk_discos_album_id;

ALTER TABLE "user_discos"
  VALIDATE CONSTRAINT fk_user_discos_user_id,
  VALIDATE CONSTRAINT fk_user_discos_disco_id;

ALTER TABLE "listening_history"
  VALIDATE CONSTRAINT fk_history_user_id,
  VALIDATE CONSTRAINT fk_history_music_id;

ALTER TABLE "artist_analytics"
  VALIDATE CONSTRAINT fk_analytics_artist_id,
  VALIDATE CONSTRAINT fk_analytics_music_id;

--------------------------------------------------------------------------------
-- 4. Indexes for Foreign Keys
--------------------------------------------------------------------------------

CREATE INDEX idx_artist_profiles_user_id ON artist_profiles(user_id);
CREATE INDEX idx_albums_artist_id ON albums(artist_id);
CREATE INDEX idx_musics_album_id ON musics(album_id);
CREATE INDEX idx_discos_album_id ON discos(album_id);

CREATE INDEX idx_user_discos_user_id ON user_discos(user_id);
CREATE INDEX idx_user_discos_disco_id ON user_discos(disco_id);

CREATE INDEX idx_listening_history_user_id ON listening_history(user_id);
CREATE INDEX idx_listening_history_music_id ON listening_history(music_id);

CREATE INDEX idx_artist_analytics_artist_id ON artist_analytics(artist_id);
CREATE INDEX idx_artist_analytics_music_id ON artist_analytics(music_id);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "artist_analytics";
DROP TABLE IF EXISTS "listening_history";
DROP TABLE IF EXISTS "user_discos";
DROP TABLE IF EXISTS "discos";
DROP TABLE IF EXISTS "musics";
DROP TABLE IF EXISTS "albums";
DROP TABLE IF EXISTS "artist_profiles";
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
