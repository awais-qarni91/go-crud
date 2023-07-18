# go-crud 

A crud application developed in golang

1- Articles entity with memory persistence(local memory).

2- Categories entity with text file persistence.

3- A CRUD with another entity(Products) using a mysql database "db_crud".

 

## Docker Setup
Follow these steps to run the project using Docker:

 

### Prerequisites

 

- Docker: [Installation Guide](https://docs.docker.com/get-docker/)

 

### Build the Docker Images

 

1. Clone the project repository:

 

   ```bash
   git clone https://github.com/awais-qarni91/go-crud
   ```

 

2. Change to the project directory:

 

   ```bash
   cd go-crud
   ```

 

3. Build the Docker image:

 
 Build image for api server
   ```bash
   
   1- docker build -t api-image -f Dockerfile.api .
 ```
  Build image for database server
   
 ```bash

   2- docker build -t mysql-image -f Dockerfile.sql .
   ```

 

### Run the Docker Multi-Container 

1- Create docker network 

  ```bash
  docker network create mynetwork
   ```

2. Run a database container:

 ```bash
   docker run -d --name mysql-container --network=mynetwork -e MYSQL_ROOT_PASSWORD=golang456  mysql-image
   ```
3. Run api container:
    
```bash
   docker run  --name api-container --network=mynetwork -p 8080:8080 api-image
   ```

 
## Useful docker commands

### Stop the Docker Container

 

To stop the Docker container, use the following command:

 

```bash
docker stop <container_id>
```

 

Replace `<container_id>` with the actual container ID or name. You can find the container ID or name by running `docker ps`.

### Show all Containers

To show all running containers, use the following command:

```bash
docker ps
```


To show all stopped containers, use the following command:

```bash
docker ps -a
```


 ### Show Docker Images

To show all images, use the following command:

```bash
docker images
```

 ### Remove Docker Image

To remove image, use the following command:

```bash
docker rmi <image_name>
```

 ### Remove Docker Container

To remove a container, use the following command:

```bash
docker rm -f <container_name>/<container_id>
```

 ### Remove Docker Network

To remove a network, use the following command:

```bash
docker network rm <network_name>
```

