FROM mysql:latest

ADD ./00_schema.sql /docker-entrypoint-initdb.d/00_schema.sql
ADD ./01_data.sql /docker-entrypoint-initdb.d/01_data.sql

ENV MYSQL_USER=chris \
  MYSQL_PASSWORD=1234 \
  MYSQL_ROOT_PASSWORD=12345 \
  MYSQL_DATABASE=widgetsapi \
  DOCKER=true
