# Go-Template

Boilerplate for go project

## Features

- [x] Clean Architecture
- [x] Config Initialization (yaml, env)
- [ ] API server
  - [x] HTTP (Gin)
  - [ ] GraphQL
  - [ ] GRPC
- [ ] Database
  - [x] PostgresQL
    - [x] Migration (dbmate)
    - [x] ORM (GORM)
  - [ ] Redis
- [x] Dependency Injection (wire)
- [x] Validator Package (validator/v10)
- [x] Error Structure
- [x] Unit Testing (testify/suite)
- [ ] Logger
- [ ] Documentation
- [ ] CI/CD Workflow
- [ ] Tracing
- [ ] Deployment

## Setup

### Install Docker

We will use docker to install any dependencies that the service uses such as postgres, etc. To install docker, it is recommended that you follow the official [Get Docker](https://docs.docker.com/get-docker/) guideline.

### Install DB Migration Tools (+ Postgresql Driver)

Linux (on other machine, please refer to [dbmate](https://github.com/amacneil/dbmate))

```
sudo curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
sudo chmod +x /usr/local/bin/dbmate
```

### Install Mockgen

For mock generation

```
go install github.com/golang/mock/mockgen@v1.6.0
```

### Install Google Wire

Wire is a code generation tool that automates connecting components using [dependency injection](https://en.wikipedia.org/wiki/Dependency_injection)

```
go install github.com/google/wire/cmd/wire@latest
```

and ensuring that `$GOPATH/bin` is added to your `$PATH`.

### Install pre-commit

For setup github hooks (Linux/Mac)

```
git config core.hooksPath .githooks

chmod +x .githooks/*
```

## Getting Started with the Service Locally

### Making Database Migration

In order to do DDL operations on our Database, we use `dbmate` command which will run a database migration in sequential manner. The step
to make a database migrations are as follows:

1. Create your DB migration folder as `db/migrations`
2. Fill the `files/etc/env/env.sh` using `files/etc/env/env.sample` as starting point.
3. Run `source files/etc/env/env.sh` in terminal
4. Create a DDL for `up` and `down` migration using CLI: `dbmate new <migration_file_name>`. This will give you the skeleton ONLY. You need to fill in the DDL query.

### Configuration

In this app, there are 2 different configurations:

- **Non-secret configuration**: stored in `files/etc/app_config/config.{env}.yaml`
- **Secret configuration**: stored in `files/etc/env/env.sh`

Secret is meant ONLY for configuration that people are not supposed to know. This includes, but not limited to:

- **Username**
- **Credentials**
- **Private Key**

Other config should go to `files/etc/app_config/config.yaml`. `files/etc/env/env.sh` can be constructed using skeleton from `files/etc/env/env.sample` in development.

To run in local using the default postgres docker config, `files/etc/env/env.sh` should be filled with:

```
# Database Config
export DB_HOST=localhost
export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export DB_PORT=5432
export DB_NAME=postgres
```

### Running the Service

Once you have all the required configuration run `cmd/start.sh` from root directory to run the service. Or `cmd/start.sh <env>` for specified environment.
For example `cmd/start.sh staging` to run and use staging configuration

## Project Structure

### Clean Architecture

![image](https://user-images.githubusercontent.com/102520846/172805794-7bc613ec-30d3-4898-8a5f-144ce3bb5b74.png)

We use clean architecture to separate between request receiver, business logic, and data layer logic. The main advantage of this separation of layer is to enable:

- Separation of concern
- Parallel works on each layer
- Adaptability to changing handler or data layer

As such, our core code will be in the form of:

```
|--- internal
|    |--- api
|    |    |--- rest
|    |    |    |--- ...package needed
|    |    |    |--- module A
|    |    |    |--- module B
|    |    |--- gql
|    |    |    |--- module A
|    |    |    |--- module B
|    |    |--- grpc
|    |    |    |--- module A
|    |    |    |--- module B
|    |
|    |--- entity
|    |    |--- module A
|    |    |--- module B
|    |
|    |--- repositories
|    |    |--- storage (pq, redis, mongodb, etc)
|    |    |   |--- models
|    |    |   |--- respositories
|    |    |   |   |--- module A
|    |    |   |   |--- module B
|    |
|    |--- usecases
|    |    |--- module A
|    |    |--- module B
|    |
```

Here are the details on what each component does:

- **API/Handler**: Presenter layer which does conversion of data structure from entity from and to a well-defined format of choosing like gRPC, HTTP, or GQL
- **Entity**: Business object which has its own data structure and methods. Even though it can be the same as DB models but it does not necessarily have to be the same
- **Repository**: Adapter for querying or manipulating data in the data layer. Change business entity into a data layer model and vice versa. Data layer to be accessed can include: in-memory, RDBMS, NoSQL, File System, External dependency to other service (internal or 3rd party)
- **Usecase**: Where business logic lives, this layer orchestrated entity and repository to achieve application specific needs

For better visualization, current application looks like the diagram below:
![image](https://user-images.githubusercontent.com/102520846/178428592-301e1626-f699-4d36-bb4d-269388cded07.png)

### Code Generation

Code generation is used for mundane code task that takes a lot of time. Using code generation reduce the amount of time to create additional boilerplate code. Currently, this project use few code generation library:

1. [Mocking w/ mockgen](https://github.com/golang/mock)

There are special sections for our generated code that SHOULD NOT be edited manually by developers. It should only be edited by another action from the code generator.
