table "players" {
    schema = schema.playtics
    column "id" {
        null = false
        type = uuid
    }
    column "name" {
        null = false
        type = varchar(255)
    }
    column "email" {
        null = false
        type = varchar(255)
    }
    column "image_url" {
        null = true
        type = varchar(255)
    }
    column "created_at" {
        null = false
        type = timestamptz
        default = sql("NOW()")
    }
    column "updated_at" {
        null = false
        type = timestamptz
        default = sql("NOW()")
    }
    primary_key {
        columns = [column.id]
    }
    index "idx_players_email" {
        columns = [column.email]
        unique  = true
    }
}

table "matches" {
    schema = schema.playtics
    column "id" {
        null = false
        type = uuid
    }
    column "duration_seconds" {
        null = false
        type = int
    }
    column "created_at" {
        null = false
        type = timestamptz
        default = sql("NOW()")
    }
    primary_key {
        columns = [column.id]
    }
}

table "match_results" {
    schema = schema.playtics
    column "player_id" {
        null = false
        type = uuid
    }
    column "match_id" {
        null = false
        type = uuid
    }
    column "kill_count" {
        null = false
        type = int
    }
    column "death_count" {
        null = false
        type = int
    }
    column "score" {
        null = false
        type = int
    }
    column "created_at" {
        null = false
        type = timestamptz
        default = sql("NOW()")
    }
    column "updated_at" {
        null = false
        type = timestamptz
        default = sql("NOW()")
    }
    primary_key {
        columns = [column.player_id, column.match_id]
    }
    foreign_key "player_id" {
        columns = [column.player_id]
        ref_columns = [table.players.column.id]
        on_delete = CASCADE
        on_update = NO_ACTION
    }
    foreign_key "match_id" {
        columns = [column.match_id]
        ref_columns = [table.matches.column.id]
        on_delete = CASCADE
        on_update = NO_ACTION
    }
}

schema "playtics" {}