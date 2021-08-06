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

{% api-method method="get" host="https://api.clerk.dev" path="/v1/clients/:id" %}
{% api-method-summary %}
Retrieve a client
{% endapi-method-summary %}

{% api-method-description %}
Retrieve the details of a client.
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
    "object": "client",
    "id": "client_1q8u5xxKGAaflDktdY1IXxvK36S"
    ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.clerk.dev" path="/v1/clients" %}
{% api-method-summary %}
List all clients
{% endapi-method-summary %}

{% api-method-description %}
List all clients.
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
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://api.clerk.dev" path="/v1/clients/verify" %}
{% api-method-summary %}
Verify a client
{% endapi-method-summary %}

{% api-method-description %}
Verify the validity of the supplied token.  
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
The JWT that is sent via the \`\_\_session\` cookie from your frontend. 
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
    "object": "client",
    "id": "client_1q8u5xxKGAaflDktdY1IXxvK36S"
    ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

