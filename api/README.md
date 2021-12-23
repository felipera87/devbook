This is the backend API.

## To run this app:

### Start the database
To start a simple test database: `docker run --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=superpass -d mysql:8.0`

To run the existing container: `docker start mysql-test`

Configure the `.env` file like `.env.example`

Run `sql/sql.sql` on a mysql cli or ide to create database and tables.