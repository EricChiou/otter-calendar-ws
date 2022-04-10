# otter-calendar
Golang Architecture base on Clean Architecture.  

## Making config file
Making config.ini with content below:  
```
# Server
SERVER_NAME=otter-calendar framework
SERVER_PORT=7000
SSL_CERT_FILE_PATH=
SSL_KEY_FILE_PATH=

# DB


# JWT
JWT_KEY=your jwt key
# JWT expire time, set 1 for one day, set 2 for two days, ...
JWT_EXPIRE=1

# Environment
ENV=dev
```
