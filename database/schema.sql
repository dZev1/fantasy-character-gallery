-- Activar foreign keys (importante en SQLite, por defecto están desactivadas)
PRAGMA foreign_keys = ON;

-- Tabla de Stats
CREATE TABLE IF NOT EXISTS stats (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    strength     INTEGER NOT NULL,
    dexterity    INTEGER NOT NULL,
    constitution INTEGER NOT NULL,
    intelligence INTEGER NOT NULL,
    wisdom       INTEGER NOT NULL,
    charisma     INTEGER NOT NULL
);

-- Tabla de Customization
CREATE TABLE IF NOT EXISTS customization (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    hair  INTEGER NOT NULL,
    face  INTEGER NOT NULL,
    shirt INTEGER NOT NULL,
    pants INTEGER NOT NULL,
    shoes INTEGER NOT NULL
);

-- Tabla de Characters
CREATE TABLE IF NOT EXISTS characters (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    body_type INTEGER NOT NULL,   -- mapea a tu enum BodyType
    species   INTEGER NOT NULL,   -- mapea a tu enum Species
    class     INTEGER NOT NULL,   -- mapea a tu enum Class

    stats_id INTEGER NOT NULL,
    customization_id INTEGER NOT NULL,

    FOREIGN KEY (stats_id) REFERENCES stats(id) ON DELETE CASCADE,
    FOREIGN KEY (customization_id) REFERENCES customization(id) ON DELETE CASCADE
);
