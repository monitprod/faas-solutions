# DB Repository

## Env Variables

    DB_REPOSITORY_MONGO_URI=<MONGO_DB_URI>
    DB_REPOSITORY_ROOT_PATH=<ROOT_PATH> # Read-only, don't modify it

## How to use?

1. Install this package

        go get github.com/monitprod/db_repository

2. Duplicate and rename file .example.env to .env

3. Configure .env file

4. Start db_repository

``` go
import ( 
    "context"
    "github.com/monitprod/db_repository" 
)

ctx, _ := context.Background()

db_repository.StartRepository(ctx)

```

5. Are you free to use ``` pkg/repository ``` and other packages of this project

## How to use mongo db client?
``` go
import ( 
    "context"
    "github.com/monitprod/db_repository" 
    "github.com/monitprod/db_repository/pkg/loaders/database"
)

ctx, _ := context.Background()

db_repository.StartRepository(ctx)

client := database.GetClient()

```