
CREATE TABLE
IF NOT EXISTS characters
(
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    body_type   SMALLINT NOT NULL,
    species     SMALLINT NOT NULL,
    class       SMALLINT NOT NULL
);


CREATE TABLE
IF NOT EXISTS stats
(
    id           BIGINT PRIMARY KEY REFERENCES characters
(id) ON
DELETE CASCADE,
    strength     SMALLINT
NOT NULL CHECK
(strength BETWEEN 0 AND 255),
    dexterity    SMALLINT NOT NULL CHECK
(dexterity BETWEEN 0 AND 255),
    constitution SMALLINT NOT NULL CHECK
(constitution BETWEEN 0 AND 255),
    intelligence SMALLINT NOT NULL CHECK
(intelligence BETWEEN 0 AND 255),
    wisdom       SMALLINT NOT NULL CHECK
(wisdom BETWEEN 0 AND 255),
    charisma     SMALLINT NOT NULL CHECK
(charisma BETWEEN 0 AND 255)
);


CREATE TABLE
IF NOT EXISTS customizations
(
    id    BIGINT PRIMARY KEY REFERENCES characters
(id) ON
DELETE CASCADE,
    hair  SMALLINT
NOT NULL CHECK
(hair BETWEEN 0 AND 255),
    face  SMALLINT NOT NULL CHECK
(face BETWEEN 0 AND 255),
    shirt SMALLINT NOT NULL CHECK
(shirt BETWEEN 0 AND 255),
    pants SMALLINT NOT NULL CHECK
(pants BETWEEN 0 AND 255),
    shoes SMALLINT NOT NULL CHECK
(shoes BETWEEN 0 AND 255)
);
