## How to Run

```sh
go run main.go
```

## How to Run via Docker

```sh
docker compose up -d --build
```

## Endpoints

### `POST /books`
- **Description:** Create a new book (ISBN must be unique)
- **Request Body (JSON):**
  - `title` (string, required)
  - `author` (string, required)
  - `isbn` (string, required, unique)
  - `release_date` (string, required, format: YYYY-MM-DD)
- **Response:**  
  - `201 Created`  
  - `400 Bad Request` 
  - `409 Conflict`

---

### `GET /books?page=1&limit=10`
- **Description:** List books with pagination
- **Query Parameters:**
  - `page` (integer, optional, default: 1)
  - `limit` (integer, optional, default: 10)
- **Response:**  
  - `200 OK`  
  - JSON:
    ```json
    {
      "data": [ ... ],
      "pagination": {
        "page": 1,
        "limit": 10,
        "total": 25,
        "total_pages": 3
      }
    }
    ```

---

### `GET /books/{isbn}`
- **Description:** Get book by ISBN
- **Path Parameter:**
  - `isbn` (string, required)
- **Response:**  
  - `200 OK`
  - `404 Not Found` 

---

### `PUT /books/{isbn}`
- **Description:** Update book by ISBN
- **Path Parameter:**
  - `isbn` (string, required)
- **Request Body (JSON):**
  - `title` (string, required)
  - `author` (string, required)
  - `release_date` (string, required, format: YYYY-MM-DD)
- **Response:**  
  - `200 OK` 
  - `400 Bad Request`  
  - `404 Not Found`

---

### `DELETE /books/{isbn}`
- **Description:** Delete book by ISBN
- **Path Parameter:**
  - `isbn` (string, required)
- **Response:**  
  - `200 OK`  
  - `404 Not Found`

## Swagger Access
If you want access via Swagger, please run this url in your browser: `http://localhost:8080/swagger-ui/dist`

## Log

- Log application: `storage/logs/app.log`

