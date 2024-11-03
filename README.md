# todo-app

Pet-project REST API

## Description

This project is a REST API for managing todo lists. It is written in Go and uses PostgreSQL as the database. The project includes user authentication, creation, reading, updating, and deletion of todo lists and todo items.

## Project Structure

<details>
  <summary>Entity-relationship diagram</summary>
  <p align="center">
    <img src=https://github.com/user-attachments/assets/25fea2ec-c29e-40fd-92e5-012dd419648b
 width=50% />
  </p>
</details>

## How to Run


### Step 1 
Сreate a .env file in the root of the project with the following content:
``` .env
DB_PASSWORD=your_postgres_password

PGADMIN_DEFAULT_EMAIL=your_pgadmin_email
PGADMIN_DEFAULT_PASSWORD=your_pgadmin_password
```

### Step 2
Create a configs/config.yml file with the following content:

``` .yml
port: "8000"
db:
  username: "postgres"
  host: "db"
  port: "5432"
  dbname: "todo-db"
  sslmode: "disable"
```

### Step 3
Make sure you have Docker and Docker Compose installed. Then run the following command:
```
docker-compose up --build
```
