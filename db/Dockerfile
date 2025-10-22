FROM mariadb

ADD ddl.sql /docker-entrypoint-initdb.d/
ADD dml.sql /docker-entrypoint-initdb.d/

EXPOSE 3306