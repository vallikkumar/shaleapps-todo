# todo

# Prerequisite tools

- Docker https://hub.docker.com/editions/community/docker-ce-desktop-windows/

```
  - sudo yum update -y
  - sudo yum install -y docker
  - sudo service docker start
  - sudo usermod -a -G docker ec2-user
```

- Golang Version 1.14 https://golang.org/dl/

```
  - wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
  - tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz
  - vim ~/.bashrc
  - export GOPATH=$HOME/work
  - export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
  - source ~/.bashrc
  - go version

```

- Node Version v13.12.0

  > You can install using nvm (https://github.com/nvm-sh/nvm)

  > nvm install 13.12.0

  > nvm use 13.12.0

- Angular-cli

  > npm install -g @angular/cli

# how to run

- Initiating api and ui with

```
make init

or you can do

 - cd ui && npm install

 - cd api && go mod tidy

```

- run test

```
 make test

 or you can do
   - cd ui && ng test
   - cd api && go test ./...
```

- build

```
make build

or you can do
   - cd ui && ng build --prod
   - docker build -t todo .
```

- run

```
make run

or you can do
  - docker run -p 8080:8080 todo
```

- run withoud docker

```
 - cd api
 - go run main.go
```

# how to pull the code

- git clone `gitlink`

# how to deploy

- Pull the code to the server
- Make sure Prerequisite tools already exist
- then run command
  ```
  make init
  make build
  make deploy
  ```

---

**Golang**

1. Using fiber webframework (https://github.com/gofiber/fiber)
2. Database sqlite3
3. Gorm ORM

**Angular**

1. Using Angular-Cli
2. Command `ng build --prod` will generate UI files and stored the file to folder `/api/public`.
   This folder will be load as a static files in `server.go`

```
s.app.Static("/", "./public")
```

**Dockerfile**

1. Use multi-stage build
2. Use image `golang:1.14-alpine` as a builder
3. Use image `alpine:3.10` for production

multi-stage build will create smaller image size
