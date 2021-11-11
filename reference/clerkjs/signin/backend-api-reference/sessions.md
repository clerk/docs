# Sessions

This object represents an active session for a user.  Sessions are created when a user successfully goes through the sign in or sign up flows. &#x20;

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

{% swagger baseUrl="https://api.clerk.dev" path="/v1/sessions/:id" method="get" summary="Retrieve a session" %}
{% swagger-description %}
Retrieve the details of a session.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
        "object": "session",
        "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
        ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/sessions" method="get" summary="List all sessions" %}
{% swagger-description %}
List all sessions.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
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
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/sessions/:id/revoke" method="post" summary="Revoke a session" %}
{% swagger-description %}
Sets the status of a session as "revoked".  Which is an unauthenticated state.  In multi-session mode, a revoked session will still return along with it's client object, however the user will need to sign into it again.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
        "object": "session",
        "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
        ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/sessions/:id/verify" method="post" summary="Verify a session" %}
{% swagger-description %}
Returns the session if it is authenticated, otherwise returns a 404 resource_not_found error.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="token" type="string" %}
The JWT that is sent via the `__session` cookie from your frontend.  Note: this JWT must be associated with the supplied session ID
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
        "object": "session",
        "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
        ...
}
```
{% endswagger-response %}

{% swagger-response status="404" description="The supplied session_id was either not valid, or was not related to the supplied token." %}
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
{% endswagger-response %}
{% endswagger %}
