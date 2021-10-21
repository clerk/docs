# Clients

This object represents a client created by a user.  Clients are unauthenticated objects that are used to track the current sign in, sign up, and an array of sessions.

### Available requests

* **`GET`**`/v1/clients/:id`
* **`GET`**`/v1/clients`
* **`POST`**`/v1/clients/verify`

### Example client schema

```javascript
{
    "object": "client",
    "id": "client_1q8u5xxKGAaflDktdY1IXxvK36S",
    "session_ids": [
        "sess_1q8u7QQhKOGRitdFGreHB00443j"
    ],
    "sign_in_attempt_id": null,
    "sign_up_attempt_id": null,
    "last_active_session_id": "sess_1q8u7QQhKOGRitdFGreHB00443j",
    "created_at": 1616473358,
    "updated_at": 1616473358
}
```

{% swagger baseUrl="https://api.clerk.dev" path="/v1/clients/:id" method="get" summary="Retrieve a client" %}
{% swagger-description %}
Retrieve the details of a client.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
    "object": "client",
    "id": "client_1q8u5xxKGAaflDktdY1IXxvK36S"
    ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/clients" method="get" summary="List all clients" %}
{% swagger-description %}
List all clients.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
[
    {
        "object": "client",
        "id": "client_1q8u5xxKGAaflDktdY1IXxvK36S"
        ...
    },
    {
        "object": "client",
        "id": "client_1q8u5xxKGAaflDktdY1IXxvK36S"
        ...
    }
]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/clients/verify" method="post" summary="Verify a client" %}
{% swagger-description %}
Verify the validity of the supplied token.  
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="token" type="string" %}
The JWT that is sent via the `__session` cookie from your frontend. 
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
    "object": "client",
    "id": "client_1q8u5xxKGAaflDktdY1IXxvK36S"
    ...
}
```
{% endswagger-response %}
{% endswagger %}
