# Sessions

A **Session** object represents an active session for a user.  Sessions are created when a user successfully goes through the sign in or sign up flows.

### Available requests

* **`GET`** `/v1/client/sessions/:id`
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

{% api-method method="get" host="https://clerk.example.com" path="/v1/client/sessions/:id" %}
{% api-method-summary %}
Retrieve a session
{% endapi-method-summary %}

{% api-method-description %}
Retrieve the details of a session.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

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

{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sessions/:id/end" %}
{% api-method-summary %}
End a session
{% endapi-method-summary %}

{% api-method-description %}
Ends a currently active Session.  Synonymous with "signing out" the user.  The session will still be present on the client object.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

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

{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sessions/:id/remove" %}
{% api-method-summary %}
Remove a session
{% endapi-method-summary %}

{% api-method-description %}
Completely removes the session from client.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

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

