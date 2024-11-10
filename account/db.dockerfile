FROM postgres:10.3

COPY account/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]