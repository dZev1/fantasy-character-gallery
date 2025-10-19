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

- **Succesful Response (201 Created)**: Returns the object of the created character, including their new `id`.
