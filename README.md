# RESTful-API with Golang

This is the RESTful-API use Golang. With Gin and Gorm technology to make this project more efficient and clean. Also, use S3 bucket to implement the file upload feature in this project.

## ℹ️ Directory/Layers

- **cmd**: handle the arguments in the command line. In this case, to handle --migrate and --seed for the database

- **internal/config**: setup the project needed. In this case, to setup the database and all the requirements for the project to go up🆙.

- **internal/dto**: Contains data transfer objects that define how data is structured when transferred between layers or exposed through APIs.

- **internal/handlers**: Handles HTTP requests and responses, acting as the entry point for the API.

- **internal/middleware**: Provides reusable logic executed before or after the main route handlers to handle cross-cutting concerns. Also handle the cors problem if exist.

- **internal/migration**: Handles database migration and seeding operations to initialize or populate the database.

- **internal/models**: Represents the database structure and defines ORM mappings.

- **internal/pkg**: Contains reusable utility functions and helper modules that are independent of the main application logic.

- **internal/repository**: Handles database queries and abstracts data access logic.

- **internal/router**: Defines and registers routes for the API.

- **internal/service**: Implements business logic and orchestrates interactions between repositories and handlers.

## Project Structure 

```mermaid
graph TD
    A[Project Root] --> B[cmd]
    B --> BA[command.go]
    
    A --> C[go.mod]
    A --> D[go.sum]
    
    A --> E[internal]
    E --> F[config]
    F --> FA[app_config.go]
    F --> FB[database.go]
    
    E --> G[dto]
    G --> GA[user_dto.go]
    
    E --> H[handlers]
    H --> HA[user_handler.go]
    
    E --> I[middleware]
    I --> IA[auth.go]
    I --> IB[cors.go]
    I --> IC[only_admin.go]
    
    E --> J[migration]
    J --> JA[data]
    JA --> JAA[users.json]
    J --> JB[migration.go]
    J --> JC[seed]
    JC --> JCA[user.go]
    J --> JD[seeder.go]
    
    E --> K[models]
    K --> KA[user_model.go]
    
    E --> L[pkg]
    L --> LA[jwtutils.go]
    L --> LB[password.go]
    L --> LC[s3aws.go]
    
    E --> M[repository]
    M --> MA[user_repository.go]
    
    E --> N[router]
    N --> NA[testing.go]
    N --> NB[user_route.go]
    
    E --> O[services]
    O --> OA[user_service.go]
    
    A --> P[main.go]

    classDef default fill:#f9f9f9,stroke:#333,stroke-width:1px;
    classDef file fill:#e3f2fd,stroke:#1976d2,stroke-width:1px;
    classDef folder fill:#fff3e0,stroke:#f57c00,stroke-width:1px;

    class A,B,E,F,G,H,I,J,JA,JC,K,L,M,N,O folder;
    class BA,FA,FB,GA,HA,IA,IB,IC,JAA,JB,JCA,JD,KA,LA,LB,LC,MA,NA,NB,OA,C,D,P file;
```

## 🌟 Features

- Authentication use JWT
- Authorization
- Uploading use AWS S3

## Prerequisite 🧰

- Golang installed (im using v1.23.4)
- PostgreSQL (im using v16.4 server)
- AWS Knowledge and Account

for the AWS S3 Setup, you can refer to this [blog](https://medium.com/geekculture/go-cafe-creating-and-adding-files-to-aws-s3-using-golang-b92eaa5f2081) 


## 🚀 How To Use

Simple, understandable installation instruction, go to your terminal and paste this!

1. Clone this repository

```bash
git clone https://github.com/paundraP/RESTful-API-with-Golang.git
```

2. Go to the project folder

```bash 
cd RESTful-API-with-Golang
```

3. Install package dependency for this project

```bash
go mod tidy
```

4. Copy the .env.example to .env and configure with your credentials

```bash
DBHOST=
DBUSER=
DBPASSWORD=
DBNAME=
DBPORT=

JWT_SECRET=

bucket=
AWS_REGION=
AWS_ACCESS_KEY=
AWS_SECRET_KEY=
```

- If you dont configure the posgresql before, open the postgresql you installed before.

```bash
psql -U postgres
CREATE DATABASE name_of_db;
\c name_of_db
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
\q
```

5. To add the table and seed them, run this

```bash 
go run main.go --migrate --migrate
```

or if you just want to run the program (dont forget to migrate the table first)
```bash 
go run main.go
```

And be sure to specify any other minimum requirements like Prerequisite above.

# Time to Explore!