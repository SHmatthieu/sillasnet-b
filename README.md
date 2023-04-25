# sillasnet backend

## how to use it
```
# to run it
go run .

# to build it and run it
go build .
./serv
```

# REST API
| method | url | result | authentification required |
|--------|-----|--------|---------------------|
|GET     |/api/ping | check if the server is online | false |
|POST     |/api/connection | check if the server is online |false |
|GET     |/api/software | get all softwares info | true |
|GET     |/api/software/:id | get info about a software by id | true |
|POST     |/api/software | add softwares | true |
|GET     |/api/hardware | get all hardwares info | true |
|GET     |/api/hardware/:id | get info about a hardware by id |  true |
|POST     |/api/hardware | add hardwares | true |
|GET     |/api/supplier | get all supplier info | true |
|GET     |/api/supplier/:id | get info about a supplier by id | true |
|POST     |/api/supplier | add suppliers | true |
|GET     |/api/tips | get all security tips | false |
|POST     |/api/report | send a report | true |

# authentification method 
 Add in request header :
  - name:{your name}
  - token:{your token}

