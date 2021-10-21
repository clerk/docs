# Sessions

## Available requests

* **`GET `**`/v1/me/sessions`
* **`GET `**`/v1/me/sessions/active`
* **`POST`**`/v1/me/sessions/:id/revoke`

## The session object

```javascript
{
    "object": "session",
    "id": "sess_1w1Nn26zzA47J1hYfdO26m4Rxdo",
    "status": "active",
    "expire_at": 1628227097160,
    "abandon_at": 1630214297160,
    "last_active_at": 1627622297160,
    "user": {
        "id": "user_1o4qfqMeCkK4KD9fiCPKIXcA9h8",
        "object": "user",
        "username": "boss-clerk",
        "first_name": "Boss",
        "last_name": "Clerk",
        "gender": "",
        "birthday": "",
        "profile_image_url": "https://images.clerk.dev/uploaded/img_1ux177vBXVeUGn4wQXfnr3lSwIg.jpeg",
        "primary_email_address_id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
        "primary_phone_number_id": null,
        "password_enabled": true,
        "two_factor_enabled": false,
        "email_addresses": [],
        "phone_numbers": [],
        "external_accounts": [],
        "public_metadata": {},
        "unsafe_metadata": {},
        "created_at": 1612556316784,
        "updated_at": 1628205483345
    },
    "created_at": 1627622297160,
    "updated_at": 1627622297165
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/sessions" method="get" summary="List all sessions" %}
{% swagger-description %}
Returns a list of the current user's sessions.
{% endswagger-description %}

{% swagger-response status="200" description="" %}
```javascript
// see the session object above.
[
    {
        "object": "session",
        "id": "sess_1w1Nn26zzA47J1hYfdO26m4Rxdo",
        ...
    },
    ...
]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/sesssions/active" method="get" summary="List all active sessions" %}
{% swagger-description %}
Returns a list of the current user's active sessions.
{% endswagger-description %}

{% swagger-response status="200" description="" %}
```javascript
// see the session object above.
[
    {
        "object": "session",
        "id": "sess_1w1Nn26zzA47J1hYfdO26m4Rxdo",
        "status": "active",
        ...
    },
    ...
]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/sessions/:id/revoke" method="post" summary="Revoke a session" %}
{% swagger-description %}
Revoke the session with the given ID.  It will be entirely removed from it's client.  You can not revoke the session used to authenticate this request.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the session.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the session object above.
{
    "object": "session",
    "id": "sess_1w1Nn26zzA47J1hYfdO26m4Rxdo",
    "status": "revoked",
    ...
}
```
{% endswagger-response %}
{% endswagger %}

