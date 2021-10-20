# Core

## Env Variables

    CORE_MONGO_URI=<MONGO_DB_URI>
    CORE_ROOT_PATH=<ROOT_PATH> # Read-only, don't modify it

## How to use?

1. Install this package

        go get github.com/monitprod/core

2. Duplicate and rename file .example.env to .env

3. Configure .env file

4. Start core

``` go
import ( 
    "context"
    "github.com/monitprod/core" 
)

core.UseCoreSmp(func(ctx context.Context) {
    // Your code
})

```

5. Are you free to use ``` pkg/repository ``` and other packages of this project

## How to use mongo db client?
``` go
import ( 
    "context"
    "github.com/monitprod/core" 
    "github.com/monitprod/core/pkg/loaders/database"
)

core.UseCoreSmp(func(ctx context.Context) {
    client := database.GetClient()

    // Your code
})

```

## How to handler context and error?
``` go
import ( 
    "context"
    "github.com/monitprod/core" 
    "github.com/monitprod/core/pkg/loaders/database"
)

ctx := context.Background()

err := core.UseCore(ctx, func() error {
    client := database.GetClient()

    // Your code

    return nil
})

if (err != nil) {

    // Your error handling
    
}

```
