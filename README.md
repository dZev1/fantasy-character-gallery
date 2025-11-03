# Character Gallery

## Description

A RESTful API built in Go to create, manage and see a gallery of Role Playing Game characters. The API allows CRUD complete operations, including management of base stats and character appearance customization.

## Getting Started

1. Clone repository

    ```(bash)
    git clone https://github.com/dZev1/fantasy-character-gallery.git
    cd character-gallery
    ```

2. Install dependencies

    ```(bash)
    go mod tidy
    ```

3. Configure the database

    - Make sure PostgreSQL is installed and running.
    - Create a database for the project.
    - The database schema is found on `schema.sql` and it will apply automatically when starting the application.

4. Configure environment variables

    - Create a `.env` file in the root of the project.
    - Add the following variable with the URL to your database:

        ```(.env)
        DATABASE_URL="postgres://user:password@localhost:XXXX/database_name?sslmode=disable"
        ```

5. Execute the application:

    - Build the application:

        ```(bash)
        go build ./cmd/
        ```

    - Run `cmd.exe`.

    - Server will be listening in `http://localhost:8080`.

## Character Enumeration References

- **Body Types**:
  - 0: Type A
  - 1: Type B

- **Species**:
  - 0: Aasimar
  - 1: Dragonborn
  - 2: Dwarf
  - 3: Elf
  - 4: Gnome
  - 5: Goliath
  - 6: Halfling
  - 7: Human
  - 8: Orc
  - 9: Tiefling

- **Classes**:
  - 0: Barbarian
  - 1: Bard
  - 2: Cleric
  - 3: Druid
  - 4: Fighter
  - 5: Monk
  - 6: Paladin
  - 7: Ranger
  - 8: Rogue
  - 9: Sorcerer
  - 10: Warlock
  - 11: Wizard

## API References

### Create a character

- **Endpoint**: `POST /characters`
- **Description**: Creates a new character with their stats and customization.
- **Request Body**:

```(JSON)
{
    "name": "Arwen",
    "body_type": 1,
    "species": 3,
    "class": 11,
    "stats": {
        "strength": 10,
        "dexterity": 5,
        "constitution": 10,
        "intelligence": 5,
        "wisdom": 7,
        "charisma": 3
    },
    "customization": {
        "hair": 0,
        "face": 3,
        "shirt": 4,
        "pants": 2,
        "shoes": 1
    }
}
```

- **Succesful Response (`201 Created`)**: Returns the object of the created character, including their new `id`.

### Get all characters

- **Endpoint**: `GET /characters`
- **Description**: Returns an array of all characters in the database.

- **Succesful Response (`200 OK`)**: Returns an array of character objects, each including stats and customization fields:

```(JSON)
[
    {
        "id": 1,
        "name": "Shallan",
        "body_type": 1,
        "species": 7,
        "class": 1,
        "stats": {
            ...
        },
        "customization": {
            ...
        }
    },
    {
        "id": 2,
        "name": "Dalinar",
        "body_type": 0,
        "species": 3,
        "class": 11,
        "stats": {
            ...
        },
        "customization": {
            ...
        }
    }
]
```

### Get a character

- **Endpoint**: `GET /characters/{id}`
- **Description**: Returns a single character by their `id`.
- **Succesful Response (`200 OK`)**: Returns the object of the character with the specified `id`, including stats and customization fields:

```(JSON)
{
    "id": 1,
    "name": "Shallan",
    "body_type": 1,
    "species": 7,
    "class": 1,
    "stats": {
        ...
    },
    "customization": {
        ...
    }
}
```

### Edit a character

- **Endpoint**: `PUT /characters/{id}`
- **Description**: Updates an existing character by their `id`.
- **Request Body**: A character object with the updated fields.

```(JSON)
{
    "name": "Shallan Davar",
    "body_type": 1,
    "species": 7,
    "class": 1,
    "stats": {
        "strength": 8,
        "dexterity": 7,
        "constitution": 9,
        "intelligence": 6,
        "wisdom": 8,
        "charisma": 5
    },
    "customization": {
        "hair": 1,
        "face": 2,
        "shirt": 3,
        "pants": 4,
        "shoes": 0
    }
}
```

- **Succesful Response (`200 OK`)**: Returns the object of the updated character, including their `id`.

### Delete a character

- **Endpoint**: `DELETE /characters/{id}`
- **Description**: Deletes an existing character by their `id`.
- **Succesful Response (`200 OK`)**.
