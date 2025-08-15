# Character API Documentation

Esta API permite gestionar personajes con estadísticas, especie, clase y personalización.

## Endpoints

---

### **1. Crear un personaje**

**POST** `/characters`

**Request Body (JSON)**:

```json
{
  "creator": "Diego",
  "name": "Thalor",
  "species": 3,
  "class": 4,
  "stats": {
    "strength": 15,
    "dexterity": 12,
    "constitution": 14,
    "intelligence": 10,
    "wisdom": 8,
    "charisma": 13
  },
  "customization": {
    "hair": 1,
    "face": 2,
    "shirt": 3,
    "pants": 1,
    "shoes": 4
  }
}
```

**Response (201 Created)**:

```json
{
  "id": "64f1a9b8",
  "message": "Character created successfully"
}
```

---

### **2. Obtener todos los personajes**

**GET** `/characters`

**Response (200 OK)**:

```json
[
  {
    "id": "64f1a9b8",
    "creator": "Diego",
    "name": "Thalor",
    "species": 3,
    "class": 4,
    "stats": {
      "strength": 15,
      "dexterity": 12,
      "constitution": 14,
      "intelligence": 10,
      "wisdom": 8,
      "charisma": 13
    },
    "customization": {
      "hair": 1,
      "face": 2,
      "shirt": 3,
      "pants": 1,
      "shoes": 4
    }
  }
]
```

---

### **3. Obtener un personaje por ID**

**GET** `/characters/{id}`

**Response (200 OK)**:

```json
{
  "id": "64f1a9b8",
  "creator": "Diego",
  "name": "Thalor",
  "species": 3,
  "class": 4,
  "stats": {
    "strength": 15,
    "dexterity": 12,
    "constitution": 14,
    "intelligence": 10,
    "wisdom": 8,
    "charisma": 13
  },
  "customization": {
    "hair": 1,
    "face": 2,
    "shirt": 3,
    "pants": 1,
    "shoes": 4
  }
}
```

---

### **4. Actualizar un personaje**

**PUT** `/characters/{id}`

**Request Body (JSON)**:

```json
{
  "name": "Thalor the Brave",
  "species": 3,
  "class": 5,
  "stats": {
    "strength": 16,
    "dexterity": 12,
    "constitution": 15,
    "intelligence": 10,
    "wisdom": 8,
    "charisma": 13
  },
  "customization": {
    "hair": 2,
    "face": 3,
    "shirt": 4,
    "pants": 2,
    "shoes": 5
  }
}
```

**Response (200 OK)**:

```json
{
  "id": "64f1a9b8",
  "message": "Character updated successfully"
}
```

---

### **5. Borrar un personaje**

**DELETE** `/characters/{id}`

**Response (200 OK)**:

```json
{
  "message": "Character deleted successfully"
}
```

---

### **6. Filtrar personajes por creador**

**GET** `/characters?creator={creatorName}`

**Example**:

```
GET /characters?creator=Diego
```

**Response (200 OK)**:

```json
[
  {
    "id": "64f1a9b8",
    "creator": "Diego",
    "name": "Thalor",
    "species": 3,
    "class": 4,
    "stats": {
      "strength": 15,
      "dexterity": 12,
      "constitution": 14,
      "intelligence": 10,
      "wisdom": 8,
      "charisma": 13
    },
    "customization": {
      "hair": 1,
      "face": 2,
      "shirt": 3,
      "pants": 1,
      "shoes": 4
    }
  }
]
```

