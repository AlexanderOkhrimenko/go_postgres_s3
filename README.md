## Минимальный каркас простого микросервисного приложения

### Схема взяимодействия контейнеров :

 --> 8080 --> API <--> POSTGRESQL <--> WORKERS ---> MINIO 
 
 #### Используются только оригиналные образа а Go приложение сначало собирается а далее копируется в контейнер для использования, это приводит к минимальному размеру финального образа 
 

- .env - содержит все переменные окружения доступные всем контейнерам (PostgreSQL / MINIO)
- build.sh - запускает docker-compose.yml с 5 worker процесамми
- purge.sh - останавливает docker-compose.yml и удаляет все образы и тома связанные со сборкой
это полезно при изменении кода и значений в .env файле.
- обратите внимание что внутри докера хосты доступны по именам указанным в docker-compose.yml например minio или postgresql
