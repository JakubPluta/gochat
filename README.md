# gochat
Go Chat Application


To start application

```bash
EXPORT POSTGRES_PASSWORD
# to start docker container
make up 

# to create database
make createdb

# to migrate db
make migrate-up # make migrate down

# start backend
make rbe

# start frontend
make rfe
```

Now you can visit:
```bash
http://localhost:3000/
```


