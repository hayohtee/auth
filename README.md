# auth

## Description
A demo REST API authentication and authorization using JWT in Go

## How to run
1. Clone the repository:

  ```bash
  git@github.com:hayohtee/auth.git

  cd auth
  ```
2. Install dependencies:

  ```bash
  go get ./...
  ```
3. Ensure PostgreSQL is installed
   
4. Create DATABASE test and ROLE test using PostgreSQL (psql)

```bash
   // First connect as postgres user
   psql -U postgres -d postgres

   // Inside psql shell run the following
   CREATE DATABASE test;
   CREATE ROLE test WITH LOGIN PASSWORD 'your password';
   GRANT ALL PRIVILEGES ON DATABASE test TO test;

   \c test postgres
   // You are now connected to database "test" as user "postgres".
   GRANT ALL ON SCHEMA public TO test;
   CREATE EXTENSION citext;
 ```

5. Set up environment variables:
  Create enivronment variable for TEST_DB_DSN

  ```bash
  export TEST_DB_DSN='postgres://test:yourpassword@localhost/test?sslmode=disable';
  ```
  Ensure the environemt variable was setup properply (you might need to close and open the terminal again).

6. Run the sql file under migrations folder

  ```bash
  psql $TEST_DB_DSN
  \i 000001_create_user_table.down.sql
  ```
7. Create .env file with the following contents

```
TEST_DB_DSN='postgres://test:yourpassword@localhost/test?sslmode=disable'
JWT_SECRET='your secret'
```

8. Run the command
   
```bash
go run ./cmd/api
```

  If you want to customize the default configurations, there is commandline flags option. Run the following
  to show the list of all commamdline options.
  
  ```bash
  go run ./cmd/api -help 
  ```

## Endpoint
[documentation](https://app.swaggerhub.com/apis-docs/OlamilekanAkintilebo/test/1.0.0)
