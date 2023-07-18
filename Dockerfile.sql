#Use the official MySQL image as the base image
FROM mysql:latest
EXPOSE 3306

#Set the MySQL root password
ENV MYSQL_ROOT_PASSWORD=golang456

#Copy the SQL script to initialize the database

COPY init.sql /docker-entrypoint-initdb.d/



