# fiber-gorm

-----

### Pacotes utilizados

- [Fiber](https://gofiber.io/)
- [Gorm](https://gorm.io/docs/connecting_to_the_database.html)
- [CompileDaemon](github.com/githubnemo/CompileDaemon)

### Instalação

1 - Instalar o pacote Fiber

```sh
go get -u github.com/gofiber/fiber/v2
```

2 - Instalar o pacote Gorm

```sh
go get -u gorm.io/gorm
```

3 - Instalar o pacote Gorm para o PostgreSQL

```sh
go get -u gorm.io/driver/postgres
```

4 - Instalar o Dialeto do PostgreSQL

```sh
go get github.com/jinzhu/gorm/dialects/postgres
```

5 - Instalar o pacote CompileDaemon

```sh
go get github.com/githubnemo/CompileDaemon
```

-----

### Executando o projeto com Hot reload

```sh
CompileDaemon -command="./fiber-gorm"
```
