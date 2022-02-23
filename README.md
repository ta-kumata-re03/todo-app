Sample todo\_app
===

## Requirements

- [go-task](https://taskfile.dev)
- MySQL
- Docker
- Go 1.18

### Install go-task

```shell
brew install go-task/tap/go-task
```

### Environment value

| Varialble | description |
|:--|:--|
| DATASOURCE\_USER     | e.g. sample-go-task   |
| DATASOURCE\_PASSWORD | e.g. sample-go-task   |
| DATASOURCE\_DATABASE | e.g. sample\_go\_task |
| DATASOURCE\_PORT     | e.g. 3306(default)    |

## Usage

### Setup

```shell
task setup
```

### Run database

```shell
task database:run
```

### Seed database

```shell
task database:seed
```

### Create database migration file

```shell
task migrate:up name=<MIGRATE_NAME>
```
