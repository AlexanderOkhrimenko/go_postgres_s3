
![Docker Compose Actions Workflow](https://github.com/AlexanderOkhrimenko/go_postgres_s3/workflows/Docker%20Compose%20Actions%20Workflow/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/AlexanderOkhrimenko/go_postgres_s3)](https://goreportcard.com/report/github.com/AlexanderOkhrimenko/go_postgres_s3)

## Minimal framework for a simple micro service application

### The scheme of interaction  containers :

--> 8080 --> API <--> POSTGRESQL <--> WORKERS ---> MINIO

#### Only the original images are used  and  the Go application is first collected and then copied to a container for use, this results in the minimum size of the final image

- **.env** - contains all environment variables available to all containers (PostgreSQL / MINIO)
- **build.sh** - launches **docker-compose.yml** with 5 worker processes
- **purge.sh** - stops **docker-compose.yml** and deletes all images and volumes related to the build. 
This is useful when changing the code and values in the **.env file.**
- note that inside the Docker hosts are accessible by the names specified in **docker-compose.yml** for example **minio** or **postgresql**

***

## Минимальный каркас простого микросервисного приложения

### Схема взяимодействия контейнеров :

 --> 8080 --> API <--> POSTGRESQL <--> WORKERS ---> MINIO 
 
 #### Используются только оригиналные образа , а Go приложение сначало собирается а далее копируется в контейнер для использования, это приводит к минимальному размеру финального образа 
 

- **.env** - содержит все переменные окружения доступные всем контейнерам (PostgreSQL / MINIO)
- **build.sh** - запускает **docker-compose.yml** с 5 worker процесамми
- **purge.sh** - останавливает **docker-compose.yml** и удаляет все образы и тома связанные со сборкой. 
Это полезно при изменении кода и значений в **.env** файле.
- обратите внимание что внутри докера хосты доступны по именам указанным в **docker-compose.yml** например **minio** или **postgresql**

