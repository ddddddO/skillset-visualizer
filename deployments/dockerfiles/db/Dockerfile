FROM postgres:9.6-alpine
ENV POSTGRES_DB skillsets

COPY /db/init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
