# Client

The **Client** object contains the current state of the **Sign in**, **Sign up**, **and Session(s)** objects.

### Available requests

* **`PUT`**`/v1/client`
* **`GET`**`/v1/client`
* **`DEL`**`/v1/client`

### Example client object schema

```javascript
{
        "object": "client",
        "id": "client_1sV3b42PUQh1kX7wi26opkRXR4x",
        "sessions": [],
        "sign_in_attempt": null,
        "sign_up_attempt": null,
        "last_active_session_id": null,
        "created_at": 1620943997,
        "updated_at": 1620943997
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/client" method="put" summary="Create or replace the current client" %}
{% swagger-description %}
Replaces the current client object with a new one, if one does not exist it will create one.

\




\


Session Management: This request will replace the existing client cookie. If one does not exist, it will create one.  If there are any in-progress actions on the client, such as an active session, the browser will lose access to them
{% endswagger-description %}

{% swagger-response status="200" description="" %}
```javascript
{
    "response": {
        "object": "client",
        "id": "client_1sV3b42PUQh1kX7wi26opkRXR4x",
        "sessions": [],
        "sign_in_attempt": null,
        "sign_up_attempt": null,
        "last_active_session_id": null,
        "created_at": 1620943997,
        "updated_at": 1620943997
    },
    "client": null
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/client" method="get" summary="Retrieve the current client" %}
{% swagger-description %}
Retrieve the details of the current client object.
{% endswagger-description %}

{% swagger-response status="200" description="" %}
```javascript
{
    "response": {
        "object": "client",
        "id": "client_1sCBzTaRqTlV6sLUtguJPrHQYKf",
        "sessions": [
            {
                "object": "session",
                "id": "sess_1sCBzm5mh3K5DjPmyH4bruWMSi1",
                "status": "active",
                "expire_at": 1620971755,
                "abandon_at": 1622958955,
                "last_active_at": 1620366955,
                "user": {
                    "id": "user_1n5AY2RXJyBGhGLAhdEffaOCpFJ",
                    "object": "user",
                    "username": null,
                    "first_name": "Boss",
                    "last_name": "Clerk",
                    "gender": "",
                    "birthday": "",
                    "profile_image_url": "https://www.gravatar.com/avatar?d=mp",
                    "primary_email_address_id": "idn_1n5AVtYAHNQY73d6dY6GVJt8gmD",
                    "primary_phone_number_id": null,
                    "password_enabled": true,
                    "two_factor_enabled": false,
                    "email_addresses": [
                        {
                            "id": "idn_1n5AVtYAHNQY73d6dY6GVJt8gmD",
                            "object": "email_address",
                            "email_address": "boss@clerk.dev",
                            "verification": {
                                "status": "verified",
                                "strategy": "email_code",
                                "attempts": 1,
                                "expire_at": 1610670206
                            },
                            "linked_to": []
                        }
                    ],
                    "phone_numbers": [],
                    "external_accounts": [],
                    "public_metadata": {
                        "test": {
                            "test": "test"
                        },
                        "onboarding": {
                            "ins_1k5owCRXLwsxnRkzDchVoWTmXI8": {
                                "oa": 1
                            }
                        }
                    },
                    "created_at": 1610669621,
                    "updated_at": 1619787935
                },
                "created_at": 1620366955,
                "updated_at": 1620366955
            }
        ],
        "sign_in_attempt": null,
        "sign_up_attempt": null,
        "last_active_session_id": "sess_1sCBzm5mh3K5DjPmyH4bruWMSi1",
        "created_at": 1620366953,
        "updated_at": 1620366955
    },
    "client": null
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/client" method="delete" summary="Delete the current client" %}
{% swagger-description %}
Deletes the current client object. This request will also delete the existing client cookie.

\




\


Session management: This request will delete the existing client cookie.
{% endswagger-description %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}
