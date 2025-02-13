# Project Name

## Introduction
Shangrila is a minimalistic microblogging website (similar to Twitter).

## Packages Used
Below is a table of the packages used in this project:

| Package Name  | Version | Description | Link |
|--------------|---------|-------------|-------------|
| pgx  | 3.6.2   | PostgreSQL Driver and Toolkit | https://github.com/jackc/pgx |
| branca  | 0.0.0   | Secure alternative to JWT | https://github.com/hako/branca |
| way  | 0.0.0   | Simple Router | https://github.com/matryer/way |

## Database

We are using CockroachDB (https://www.cockroachlabs.com/) for this project.

1. Install
```
brew install cockroachdb/tap/cockroach
```

2. Start a Local CockroachDB Cluster
```
cockroach start-single-node --insecure --listen-addr=localhost
```

To create a persistent database:

```
cockroach sql --insecure -e 'CREATE DATABASE mydb;'
```

3. Use SQL schema (schema.sql) and use it to create the database

```
cat schema.sql | cockroach sql --insecure -f schema.sql
```

## How to Use
1. Clone the repository:
   ```sh

   ```
2. Install dependencies:
   ```sh

   ```
3. Run the project:
   ```sh

   ```
4. Access the application at `http://localhost:3000` (or relevant URL).

## Images
Include images to showcase your project:

![Screenshot 1](path/to/image1.png)
![Screenshot 2](path/to/image2.png)

> _Replace `path/to/image.png` with actual image paths._