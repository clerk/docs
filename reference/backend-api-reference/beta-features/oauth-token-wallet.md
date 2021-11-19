# OAuth Token Wallet

The OAuth Token Wallet allows developers to easily get an active access token for a user's connected OAuth account, e.g. _Sign In with Google._

If an access token has expired, a new one will be issued transparently and returned in the response. Note that the refresh functionality is not supported for every OAuth provider, in which case an error will be returned.

### Available requests

* **`GET`**` /v1/users/:id/oauth_access_tokens/:provider`

{% swagger baseUrl="https://api.clerk.dev" path="/v1/users/:id/oauth_access_tokens/:provider" method="get" summary="Retrieve token" %}
{% swagger-description %}
Retrieve a valid (i.e. non-expired) OAuth access token for a user's connected OAuth account
{% endswagger-description %}

{% swagger-parameter in="path" name="provider" type="string" required="true" %}
The ID of the OAuth provider (e.g. 

`facebook`

,

`google`

).
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}
The Clerk user ID.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
{
  "token": "xxxxxxxxxxxxxxxxxxxxx",
  "provider": "google",
  "scopes": [
    "openid",
    "https://www.googleapis.com/auth/userinfo.email"
    "https://www.googleapis.com/auth/userinfo.profile"
  ]
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="The access token has expired but the provider hasn't provided us with a refresh token and so we cannot fetch a new access token." %}

{% endswagger-response %}
{% endswagger %}



