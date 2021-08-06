# Sessions

This object represents an active session for a user.  Sessions are created when a user successfully goes through the sign in or sign up flows.  

### Available requests

* **`GET`**`/v1/sessions/:id`
* **`GET`**`/v1/sessions`
* **`POST`**`/v1/sessions/:id/revoke`
* **`POST`**`/v1/sessions/:id/verify`

### Example session schema

```javascript
{
        "object": "session",
        "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
        "client_id": "client_1q8sn8pLHjqTU2g1UheJwJ7YWU6",
        "user_id": "user_1n5BS00ns3t21GWIXiPJ0KcSe02",
        "status": "active",
        "last_active_at": 1616473411,
        "expire_at": 1617078211,
        "abandon_at": 1619065411
}
```

{% api-method method="get" host="https://api.clerk.dev" path="/v1/sessions/:id" %}
{% api-method-summary %}
Retrieve a session
{% endapi-method-summary %}

{% api-method-description %}
Retrieve the details of a session.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
// see example schema
{
        "object": "session",
        "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
        ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.clerk.dev" path="/v1/sessions" %}
{% api-method-summary %}
List all sessions
{% endapi-method-summary %}

{% api-method-description %}
List all sessions.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
// see example schema
[
        {
                "object": "session",
                "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
                ...
        },
        {
                "object": "session",
                "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
                ...
        }        
]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://api.clerk.dev" path="/v1/sessions/:id/revoke" %}
{% api-method-summary %}
Revoke a session
{% endapi-method-summary %}

{% api-method-description %}
Sets the status of a session as "revoked".  Which is an unauthenticated state.  In multi-session mode, a revoked session will still return along with it's client object, however the user will need to sign into it again.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
// see example schema
{
        "object": "session",
        "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
        ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://api.clerk.dev" path="/v1/sessions/:id/verify" %}
{% api-method-summary %}
Verify a session
{% endapi-method-summary %}

{% api-method-description %}
Returns the session if it is authenticated, otherwise returns a 404 resource\_not\_found error.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-form-data-parameters %}
{% api-method-parameter name="token" type="string" required=true %}
The JWT that is sent via the \`\_\_session\` cookie from your frontend.  Note: this JWT must be associated with the supplied session ID
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
// see example schema
{
        "object": "session",
        "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
        ...
}
```
{% endapi-method-response-example %}

{% api-method-response-example httpCode=404 %}
{% api-method-response-example-description %}
The supplied session\_id was either not valid, or was not related to the supplied token.
{% endapi-method-response-example-description %}

```
{
    "errors": [
        {
            "message": "Session not found",
            "long_message": "No session was found with id session_id",
            "code": "resource_not_found"
        }
    ]
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

