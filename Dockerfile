FROM postgres:latest

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD 12345
ENV POSTGRES_DB email-db

COPY users_table.sql /docker-entrypoint-initdb.d/
