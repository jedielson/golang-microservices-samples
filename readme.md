# Sample Microservice with Golang


## Docker

To execute the environment, you need to run `docker-compose up -d`

It will setup the services:
 - Postgres
 - PG Admin

### Connecting with PgAdmin

After up your environment with docker-compose, go to http://localhost:8001.
Then Setup the email and password.

```
email:admin@admin.com
password:postgres
```

After that, you need to configure a new server. Click in `Add New Server`
At general tab, put some name and go to Connection tab and put this params:

```
Host: host.docker.internal
Port: 5432
Maintence Db: postgres
Username: pguser
Password: postgres
```

And that's it =]

### How to remove all local images from project

`docker ps -a | grep jedi | awk '{print $1}' | xargs docker rm -f && docker images -a | grep jedims | awk '{print $3}' | xargs docker rmi -f`
