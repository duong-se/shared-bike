
FROM mysql:8.0.29

EXPOSE 3306

COPY ./sql/migrations/*.sql /docker-entrypoint-initdb.d/
