# GoLang bindings for the Telegraph API

> This project is just to provide a wrapper around the API without any additional features.

All methods and types available and this library (possibly) is ready for use in production.

## Start using telegraph

Download and install it:

```shell
go get -u github.com/Z00mZE/telegraph-api
```

Import it in your code:
```go
import "github.com/Z00mZE/telegraph-api"

func main(){
	
}
```

### Google Wire

Codegen [Google Wire](https://github.com/google/wire)

```shell
wire gen
```

OR use [Taskfile](https://taskfile.dev)

```shell
task wg
```


##  Examples

```go
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Z00mZE/telegraph-api"
	"github.com/Z00mZE/telegraph-api/model"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	c := telegraph.NewTelegraphAPI("")
	account, accountError := c.CreateAccount(context.Background(), model.Account{
		ShortName:  "whoame",
		AuthorName: "Who are me",
	})

	if accountError != nil {
		logger.Error("occured some error", slog.String("error", accountError.Error()))
		return
	}

	logger.Info(
		"account created",
		slog.Any("account", account),
	)
}
```