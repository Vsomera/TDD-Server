-- create players table if it does not exist
DO $$ BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_schema = 'public'
        AND table_name = 'players'
    ) THEN
        CREATE TABLE players (
            id SERIAL PRIMARY KEY,
            player_name VARCHAR(255) NOT NULL,
            player_score INTEGER NOT NULL
        );
    END IF;
END $$;
