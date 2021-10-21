# Sessions

A **Session** object represents an active session for a user.  Sessions are created when a user successfully goes through the sign in or sign up flows.

### Available requests

* **`GET `**`/v1/client/sessions/:id`
* **`POST`**`/v1/client/sessions/:id/end`
* **`POST`**`/v1/client/sessions/:id/remove`

## The Session object

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

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sessions/:id" method="get" summary="Retrieve a session" %}
{% swagger-description %}
Retrieve the details of a session.
{% endswagger-description %}

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

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sessions/:id/end" method="post" summary="End a session" %}
{% swagger-description %}
Ends a currently active Session.  Synonymous with "signing out" the user.  The session will still be present on the client object.
{% endswagger-description %}

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

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sessions/:id/remove" method="post" summary="Remove a session" %}
{% swagger-description %}
Completely removes the session from client.
{% endswagger-description %}

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
