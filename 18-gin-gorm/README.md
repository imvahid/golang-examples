## Useful commands

#### Initial module
```shell
go mod init github.com/imvahid/go-crud
```

#### Install CompileDaemon, for auto build
```shell
go get github.com/githubnemo/CompileDaemon
```

#### Install godotenv
```shell
go install github.com/joho/godotenv/cmd/godotenv@latest
go get github.com/joho/godotenv/cmd/godotenv
```

#### Install gin
```shell
go get -u github.com/gin-gonic/gin
```

#### Install gorm
```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

#### Use CompileDaemon, with module name
```shell
CompileDaemon -command="./go-crud"
```
