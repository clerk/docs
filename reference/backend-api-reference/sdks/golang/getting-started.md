# Getting started with Go

The source: [https://github.com/clerkinc/clerk-sdk-go](https://github.com/clerkinc/clerk-sdk-go)

## Overview

Clerk's Go SDK is a thin wrapper over our Backend API. Add the following import statement:

```go
import "github.com/clerkinc/clerk-sdk-go/clerk"
```

If Go doesn't automatically add the module, you can explicitly add it with:

```bash
$ go get github.com/clerkinc/clerk-sdk-go
```

## Clerk client

Now, you can create a Clerk client by calling the `clerk.NewClient` function. This function requires your Clerk API key.  You can get this from the dashboard of your Clerk application in **Settings &gt; API Keys**.

Then create a client object to use the various services Clerk provides.

```go

client, err := clerk.NewClient("CLERK_API_KEY")
if err != nil {
    // handle error
}
```

