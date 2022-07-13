
FROM mysql:8.0.29

EXPOSE 3306

COPY ./api/sql/migrations/*.sql /docker-entrypoint-initdb.d/
COPY ./api/sql/seeders/*.sql /docker-entrypoint-initdb.d/
