# Sessions

This object represents an active session for a user.  Sessions are created when a user successfully goes through the sign in or sign up flows. &#x20;

### Available requests

* **`GET`**`/v1/sessions/:id`
* **`GET`**`/v1/sessions`
* **`POST`**`/v1/sessions/:id/revoke`
* **`POST`**`/v1/sessions/:id/verify`
* `POST /v1/sessions/:id/tokens/:templateName`

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

{% swagger method="post" path="/v1/sessions/:id/tokens/:templateName" baseUrl="https://api.clerk.dev" summary="Create a token based on a session and a template name" %}
{% swagger-description %}
Creates a JSON Web Token(JWT) based on a session and a JWT Template name defined for your instance.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="path" required="true" name="id" %}
Session ID
{% endswagger-parameter %}

{% swagger-parameter in="path" name="templateName" required="true" %}
The name of the JWT Template defined in your instance (e.g. 

`custom_hasura`

).
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Success" %}
```javascript
{
    "object": "token",
    "jwt": "eyJhbGciOiJSUzI1NiIsImtpZCI6Imluc18xazVvd0NSWEx3c3huUmt6RGNoVm9XVG1YSTgiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE2NDIwMjk2NDIsImZpcnN0X25hbWUiOiJCb3NzIiwiaWF0IjoxNjQyMDI5NTgyLCJpc3MiOiJodHRwczovL2NsZXJrLmNsZXJrc3RhZ2UuZGV2IiwianRpIjoiYzU2YTUwN2Y3ODU3ZjkyNjBlNzQiLCJuYmYiOjE2NDIwMjk1NzcsInN1YiI6InVzZXJfMW41QVkyUlhKeUJHaEdMQWhkRWZmYU9DcEZKIn0.jMZjFadJiRsFS9uh0JqzmepAmbij52qKL0nEIzXWquTecRlWLgU2dmFt66jRLEN32SdV0ERyD4eyEKfN8L_judgVR_p38qXnKvgnE_pLmZpwjLjSDm05Ow081VH6Kd2PZuvMbW7DU6xAl5x2Fqj6QqlQGdicuEYcTKxzU5Bo3yJ8Heqnjaa-91VyRritF1hz6J3AVaePEAFS0aU2vdNvXABbup_Om2IfbT1Exr76k8btJKhQvGdkQPDd88V_5DIemmDvV76d4SKR9LE5fYXpGAM74Q1aSI5XNmvjf0aTR_hdoNFNAxjss_0geq3xdCCtLpKN8NMOh7beL4MWjBjifA"
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="Session not found or inactive, JWT template not found for the provided name" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="500: Internal Server Error" description="Something went wrong" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}
