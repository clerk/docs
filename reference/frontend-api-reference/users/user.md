# Current user

## Available requests

* **`GET`**`  ``/v1/me`
* **`PATCH`**`/v1/me`

## The user object

```javascript
{
    "id": "user_1o4q123qMeCkKKIXcA9h8",
    "object": "user",
    "external_id" "my_previous_auth_solution_id",
    "username": null,
    "first_name": "Boss",
    "last_name": "Clerk",
    "profile_image_url": "https://images.clerk.dev/uploaded/img_jlkkcq2786n0.jpeg",
    "primary_email_address_id": "idn_1o4qfa1caa235iei6",
    "primary_phone_number_id": null,
    "primary_web3_wallet_id": null,
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
    "web3_wallets": [],
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

{% swagger baseUrl="https://clerk.example.com" path="/v1/me" method="get" summary="Retrieve current user" %}
{% swagger-description %}
Retrieve the details of the current user. 
{% endswagger-description %}

{% swagger-response status="200" description="" %}
```
// User object
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user"
    ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me" method="patch" summary="Update current user" %}
{% swagger-description %}
Updates the current user.

You can set the current user's primary contact identifiers (email address and phone numbers) by updating the `primary_email_address_id` and `primary_phone_number_id` attributes respectively. Both IDs should correspond to verified identifications that belong to the user.

You can remove the current user's username by setting the `username` attribute to `null` or the blank string `""`. This is a destructive action. The identification will be deleted forever.

Usernames can be removed only if they are optional in your instance settings and there's at least one other identifier on the current user which can be used for authentication.
{% endswagger-description %}

{% swagger-parameter in="body" name="profile_image_id" type="string" %}
The ID of the image to set as the users profile image.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="primary_email_address_id" type="string" %}
The ID of the email address to set as primary.  It must be verified, and present on the current user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="primary_phone_number_id" type="string" %}
The ID of the phone number to set as primary.  It must be verified, and present on the current user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" type="string" %}
User's password, must be at least 8 characters, and can't have been in a recent leaked password database.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="username" type="string" %}
User's username, must be unique across the instance.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="first_name" type="string" %}
User's first name.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="last_name" type="string" %}
User's last name.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// User object
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user"
    ...
}
```
{% endswagger-response %}
{% endswagger %}
