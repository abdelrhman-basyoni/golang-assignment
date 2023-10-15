Here's an updated version of your README, incorporating the information you provided earlier:

---

# Golang Assignment 💻

This assignment serves as the foundation for your job interview, where you will present your implementation, discuss the choices you made, and overcome any challenges you encountered.

## Prerequisites ✔️

1. Git
2. Golang
3. Docker (optional)

## Assignment 📝

Fork this repository and commit your code to your own repository.

1. Create a simple Golang application that serves a RESTful API on port 3000 with a resource of your choice (pets, books, memes, etc). The endpoints should support Create, Read, Update, and Delete operations.

2. The data should be persisted in a database like SQLite, MySQL, PostgreSQL, etc.

## Running the Project Locally

To run the project locally, follow these steps:

1. Clone the project repository to your local machine:

   ```bash
   git clone https://github.com/your-username/your-project.git
   ```

2. Navigate to the project directory:

   ```bash
   cd your-project
   ```

3. Install any required dependencies.

4. Run the project:

   ```bash
   go run main.go
   ```

   The project will now be accessible at `http://localhost:3000`.

## Running Tests

To run tests for the project, execute the following command in the project directory:

```bash
go test test/book_test.go
```

## Running the Project with Docker

This project is Dockerized, allowing for easy deployment using Docker containers. To run the project with Docker, follow these steps:

1. Make sure you have Docker installed on your system. If not, please refer to the Docker installation instructions for your platform.

2. Build a Docker image from the project directory (replace `<image-name>` with your desired image name):

   ```bash
   docker build -t <image-name> .
   ```

3. Run a Docker container based on the image, and expose port 3000 to the host (replace `<container-name>` with your desired container name):

   ```bash
   docker run --name <container-name> -p 3000:3000 -d <image-name>
   ```

   The project will now be accessible at `http://localhost:3000`.

   To stop and remove the container, use the following commands:

   ```bash
   docker stop <container-name>
   docker rm <container-name>
   ```



## Tips 🧞

### Application Example

**Endpoints:**

```sh
# Get all books
GET /books

# Get a specific book
GET /books/:id

# Create a book
POST /books

# Update a book
PUT /books/:id

# Delete a book
DELETE /books/:id
```

**Model:**

```json
{
    "id": 1,
    "name": "First Book",
    "genre": "novel",
    "price": 20.5
}
```

### Database

this project uses local SQLite database:



### REST API

The Http Server is UJsing Echo:

- [github.com/labstack/echo](https://github.com/labstack/echo)



---
