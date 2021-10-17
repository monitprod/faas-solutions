# DB Repository

## Env Variables

    DB_REPOSITORY_MONGO_URI=<MONGO_DB_URI>

## How to use?

1. install package

    go get github.com/monitprod/db_repository

2. start db_repository

``` go
import ( 
    "context"
    "github.com/monitprod/db_repository" 
)

ctx, _ := context.Background()

db_repository.StartRepository(ctx)

```

3. Are you free to use ``` pkg/repository ``` and other packages of this project

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