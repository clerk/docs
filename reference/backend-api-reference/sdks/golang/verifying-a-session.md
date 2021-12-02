# Verifying a session

## Protect Your Backend APIs

Go makes it really easy to create a simple HTTP server,  and Clerk makes it really easy to authenticate any request. In the following example you can learn how to verify a session and retrieve the corresponding user.

{% tabs %}
{% tab title="Auth v2" %}
```go
package main

import (
	"net/http"
	"strings"
	
	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func main() {
	client, _ := clerk.NewClient("CLERK_API_KEY")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// get session token from Authorization header
		sessionToken := r.Header.Get("Authorization")
		sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")
		
		// verify the session
		sessClaims, err := client.VerifyToken(sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		// get the user, and say welcome!
		user, err := client.Users().Read(sessClaims.Subject)
		if err != nil {
			panic(err)
		}

		w.Write([]byte("Welcome " + *user.FirstName))
	})

	http.ListenAndServe(":8080", nil)
}

```
{% endtab %}

{% tab title="Auth v1 (deprecated)" %}
```go
package main

import (
	"net/http"
	
	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func main() {
	client, _ := clerk.NewClient("CLERK_API_KEY")	

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// verify the session
		sess, err := client.Verification().Verify(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

                // Optionally get the user, and say welcome!
		user, err := client.Users().Read(sess.UserID)
		if err != nil {
			panic(err)
		}

		w.Write([]byte("Welcome " + *user.FirstName))
	})

	http.ListenAndServe(":8080", nil)
}

```
{% endtab %}
{% endtabs %}

## Using middleware

The Clerk SDK also provides a simple middleware that adds the active session to the request's context.

{% tabs %}
{% tab title="Auth v2" %}
```go
package main

import (
	"net/http"
	
	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func main() {
	client, _ := clerk.NewClient("CLERK_API_KEY")
	
	mux := http.NewServeMux()

	injectActiveSession := clerk.WithSession(client)
	mux.Handle("/hello", injectActiveSession(helloUserHandler(client)))

	http.ListenAndServe(":8080", mux)
}

func helloUserHandler(client clerk.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		
		sessClaims, ok := ctx.Value(clerk.ActiveSessionClaims).(*clerk.SessionClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		user, err := client.Users().Read(sessClaims.Subject)
		if err != nil {
			panic(err)
		}

		w.Write([]byte("Welcome " + *user.FirstName))
	}
}

```
{% endtab %}

{% tab title="Auth v1 (deprecated)" %}
```go
package main

import (
	"net/http"
	
	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func main() {
	client, _ := clerk.NewClient("CLERK_API_KEY")
	
	mux := http.NewServeMux()

	injectActiveSession := clerk.WithSession(client)
	mux.Handle("/hello", injectActiveSession(helloUserHandler(client)))

	http.ListenAndServe(":8080", mux)
}

func helloUserHandler(client clerk.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		
		sess, ok := ctx.Value(clerk.ActiveSession).(*clerk.Session)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

                // Optionally get the user, and say welcome!
		user, err := client.Users().Read(sess.UserID)
		if err != nil {
			panic(err)
		}

		w.Write([]byte("Welcome " + *user.FirstName))
	}
}

```
{% endtab %}
{% endtabs %}
