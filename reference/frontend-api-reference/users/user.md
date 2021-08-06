# Current user

## Available requests

* **`GET`**  `/v1/me`
* **`PATCH`**`/v1/me`

## The user object

```javascript
{
    "id": "user_1o4q123qMeCkKKIXcA9h8",
    "object": "user",
    "username": null,
    "first_name": "Boss",
    "last_name": "Clerk",
    "profile_image_url": "https://images.clerk.dev/uploaded/img_jlkkcq2786n0.jpeg",
    "primary_email_address_id": "idn_1o4qfa1caa235iei6",
    "primary_phone_number_id": null,
    "password_enabled": true,
    "two_factor_enabled": false,
    "email_addresses": [
        {
            "id": "idn_1o4dI2qlcaXSXE6",
            "object": "email_address",
            "email_address": "boss@clerk.dev",
            "verification": {
                "status": "verified",
                "strategy": "from_oauth_google"
            },
            "linked_to": [
                {
                    "type": "oauth_google",
                    "id": "idn_1o4qfbU8Y2345234acClb"
                }
            ]
        }
    ],
    "phone_numbers": [],
    "external_accounts": [
        {
            "object": "google_account",
            "id": "idn_1o4qfbU8YjpJgClb",
            "google_id": "10086712340624",
            "approved_scopes": "email https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile openid profile",
            "email_address": "boss@clerk.dev",
            "given_name": "Boss",
            "family_name": "Clerk",
            "picture": "https://lh3.googleusercontent.com/a-/AOh14Gl5hWBSGDUdas2341vtMfY6_NUbPx=s1000-c"
        }
    ],
    "public_metadata": {},
    "created_at": 1612556316784,
    "updated_at": 1627622297187
}
```

{% api-method method="get" host="https://clerk.example.com" path="/v1/me" %}
{% api-method-summary %}
Retrieve current user
{% endapi-method-summary %}

{% api-method-description %}
Retrieve the details of the current user. 
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
// User object
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user"
    ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="patch" host="https://clerk.example.com" path="/v1/me" %}
{% api-method-summary %}
Update current user
{% endapi-method-summary %}

{% api-method-description %}
Updates the current user.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="profile\_image\_id" type="string" required=false %}
The ID of the image to set as the users profile image.
{% endapi-method-parameter %}

{% api-method-parameter name="primary\_email\_address\_id" type="string" required=false %}
The ID of the email address to set as primary.  It must be verified, and present on the current user.
{% endapi-method-parameter %}

{% api-method-parameter name="primary\_phone\_number\_id" type="string" required=false %}
The ID of the phone number to set as primary.  It must be verified, and present on the current user.
{% endapi-method-parameter %}

{% api-method-parameter name="password" type="string" required=false %}
User's password, must be at least 8 characters, and can't have been in a recent leaked password database.
{% endapi-method-parameter %}

{% api-method-parameter name="username" type="string" required=false %}
User's username, must be unique across the instance.
{% endapi-method-parameter %}

{% api-method-parameter name="first\_name" type="string" required=false %}
User's first name.
{% endapi-method-parameter %}

{% api-method-parameter name="last\_name" type="string" required=false %}
User's last name.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
// User object
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user"
    ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

