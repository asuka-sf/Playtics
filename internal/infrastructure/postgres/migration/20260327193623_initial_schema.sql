-- Add new schema named "playtics"
CREATE SCHEMA "playtics";
-- Create "matches" table
CREATE TABLE "playtics"."matches" (
  "id" uuid NOT NULL,
  "duration_seconds" integer NOT NULL,
  "created_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create "players" table
CREATE TABLE "playtics"."players" (
  "id" uuid NOT NULL,
  "name" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "image_url" character varying(255) NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_players_email" to table: "players"
CREATE UNIQUE INDEX "idx_players_email" ON "playtics"."players" ("email");
-- Create "match_results" table
CREATE TABLE "playtics"."match_results" (
  "player_id" uuid NOT NULL,
  "match_id" uuid NOT NULL,
  "kill_count" integer NOT NULL,
  "death_count" integer NOT NULL,
  "score" integer NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("player_id", "match_id"),
  CONSTRAINT "match_id" FOREIGN KEY ("match_id") REFERENCES "playtics"."matches" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "player_id" FOREIGN KEY ("player_id") REFERENCES "playtics"."players" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Drop schema named "public"
DROP SCHEMA "public" CASCADE;
