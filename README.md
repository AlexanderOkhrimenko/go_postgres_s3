## Minimal framework of a simple micro service application

### Container interaction diagram :

 --> 8080 --> API <--> POSTGRESQL <--> WORKERS ---> MINIO 
 
 #### Only original images are used. At first Golang application assembled and then copied to the container for use, this leads to a minimum size of the final image. 
 

##### API - accepts requests of the form  ````POST http://localhost:8080/v1/url.insert ```` the query adds text information to the database
##### WORKER - looking for the value "wait” in the database. As an example of a function for writing a file to the repository was given MINIO


-	```` .env ````- contains all environment variables available to all containers (PostgreSQL / MINIO)
-	```` build.sh ```` - run ```` docker-compose.yml```` and five processes Worker
-	```` purge.sh ```` - stop ````docker-compose.yml```` and deletes all images and volumes associated with the assembly; this is useful when you change the code and values ​​in the ````.env ```` file.
-	note that hosts inside the docker are accessible by the names specified in docker-compose.yml, for example ```` minio ```` or ```` postgresql ````


