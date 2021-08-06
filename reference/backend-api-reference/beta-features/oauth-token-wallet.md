# OAuth Token Wallet

The OAuth Token Wallet allows developers to easily get an active access token for a user's connected OAuth account, e.g. _Sign In with Google._

If an access token has expired, a new one will be issued transparently and returned in the response. Note that the refresh functionality might not work for every OAuth provider right now, in which case an error will be returned.

### Available requests

* **`GET`** `/v1/users/:id/oauth_access_tokens/:provider`

{% api-method method="get" host="https://api.clerk.dev" path="/v1/users/:id/oauth\_access\_tokens/:provider" %}
{% api-method-summary %}
Retrieve token
{% endapi-method-summary %}

{% api-method-description %}
Retrieve a valid \(i.e. non-expired\) OAuth access token for a user's connected OAuth account
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="provider" type="string" required=true %}
The ID of the OAuth provider \(e.g. `facebook`,`google`\).
{% endapi-method-parameter %}

{% api-method-parameter name="id" type="string" required=true %}
The Clerk user ID.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

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
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}





