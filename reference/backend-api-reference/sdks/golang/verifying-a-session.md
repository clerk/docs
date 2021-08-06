# Verifying a session

## Simple server

Go makes it really easy to create a simple HTTP server,  and Clerk makes it really easy to authenticate any request.  The verify function checks for the cookie Clerk set, and will return the session object if it's both valid and active.

```go
package main

import (
	"net/http"
	"github.com/clerkinc/clerk-sdk-go/clerk"
)

var client, _ = clerk.NewClient("CLERK_API_KEY")

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// verify the session
		sess, err := client.Verification().Verify(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		// get the user, and say welcome!
		user, err := client.Users().Read(sess.UserID)
		if err != nil {
			panic(err)
		}

		w.Write([]byte("Welcome " + *user.FirstName))
	})

	http.ListenAndServe(":8080", nil)
}

```

## Using middleware

The Clerk SDK also provides a simple middleware that adds the active session to the request's context.

```go
package main

import (
	"net/http"
	"github.com/clerkinc/clerk-sdk-go/clerk"
)

var client, _ = clerk.NewClient("CLERK_API_KEY")

func main() {
	mux := http.NewServeMux()

	deps := helloDeps{client: client}

	injectActiveSession := clerk.WithSession(client)
	mux.Handle("/hello", injectActiveSession(http.HandlerFunc(deps.helloUser)))

	http.ListenAndServe(":8080", mux)
}

type helloDeps struct{ client clerk.Client }

func (hd *helloDeps) helloUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sess, ok := ctx.Value(clerk.ActiveSession).(*clerk.Session)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	user, err := hd.client.Users().Read(sess.UserID)
	if err != nil {
		panic(err)
	}

	w.Write([]byte("Welcome " + *user.FirstName))
}

```

