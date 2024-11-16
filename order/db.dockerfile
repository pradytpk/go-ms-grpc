FROM postgres:10.3

COPY catalog/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]