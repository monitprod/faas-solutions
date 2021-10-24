# Core

## Env Variables

    CORE_MONGO_URI=<MONGO_DB_URI>
    CORE_ROOT_PATH=<ROOT_PATH>

## How to use?

1. Install this package

        go get github.com/monitprod/core


4. Start Env Variables

5. Start core


    ``` go
    import ( 
        "context"
        "github.com/monitprod/core" 
    )

    core.UseCoreSmp(func(ctx context.Context) {
        // Your code
    })
    ```

    Are you free to use ``` pkg/repository ``` and other packages of this project


## You want to test direct core?
Follow this steps

1. Duplicate and rename file .example.env to .env

2. Configure .env file

3. Are you free to run tests

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
