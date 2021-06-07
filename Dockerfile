FROM mysql:5.7.22
ENV MYSQL_ROOT_PASSWORD admin
ENV MYSQL_DATABASE lp
ENV MYSQL_USER admin
ENV MYSQL_PASSWORD admin
ADD script.sql /docker-entrypoint-initdb.d
EXPOSE 3306