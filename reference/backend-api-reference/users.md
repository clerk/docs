# Users

This object represents a verified user in your instance. &#x20;

### Available requests

* **`GET`**`  ``/v1/users/:id`
* **`GET`**`  ``/v1/users`
* **`POST`**` ``/v1/users`
* **`PATCH`**`/v1/users/:id`
* **`POST`**` ``/v1/users/:id/profile_image`
* **`DEL`**`  ``/v1/users/:id`
* **`GET`**`  ``/v1/users/:id/oauth_access_tokens/:provider`

### Example user schema

```javascript
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user",
    "username": null,
    "first_name": "Boss",
    "last_name": "Clerk",
    "profile_image_url": "https://images.clerk.services/clerk/default-profile.svg",
    "primary_email_address_id": "idn_1oBNgISXFbSf5m0uP2Wl0qWtNGX",
    "primary_phone_number_id": null,
    "primary_web3_wallet_id": null,
    "password_enabled": true,
    "two_factor_enabled": false,
    "email_addresses": [
        {
            "id": "idn_1oBNgISXFbSf5m0uP2Wl0qWtNGX",
            "object": "email_address",
            "email_address": "boss@clerk.dev",
            "verification": {
                "status": "verified",
                "strategy": "email_code",
                "attempts": 1,
                "expire_at": 1612756733
            },
            "linked_to": []
        }
    ],
    "phone_numbers": [
        {
            "id": "idn_1q8Uq8Mc4t7WWMy9Z6Og0gNJVui",
            "object": "phone_number",
            "phone_number": "+15555555555",
            "reserved_for_second_factor": false,
            "verification": {
                "status": "verified",
                "strategy": "phone_code",
                "attempts": 1,
                "expire_at": 1616461499
            },
            "linked_to": []
        }
    ],
    "web3_wallets": [
        {
            "id": "idn_23xyKkHEPpQQk9RzTCJVHsKoOTW",
            "object": "web3_wallet",
            "web3_wallet": "0x0000000000000000000000000000000000000000",
            "verification": {
                "status": "verified",
                "strategy": "web3_metamask_signature",
                "attempts": 1,
                "expire_at": 1642690572653,
                "nonce": "foz4bbtrtw283h0bz72r101d9t9uyjdyx0t84krg"
            }
        }
    ],
    "external_accounts": [],
    "public_metadata": {},
    "private_metadata": {},
    "created_at": 1612756155,
    "updated_at": 1612756155
}
```

###

{% swagger baseUrl="https://api.clerk.dev" path="/v1/users/:id" method="get" summary="Retrieve a user" %}
{% swagger-description %}
Retrieve the details of a user. 
{% endswagger-description %}

{% swagger-parameter in="header" name="Authentication" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user"
    ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/users" method="get" summary="List all users" %}
{% swagger-description %}
List all users.  Ordered by creation date, with the newest user first.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="query" name="phone_number" type="string[]" %}
Returns users with the phone numbers specified. Accepts up to 100 phone numbers. Any phone numbers not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="email_address" type="string[]" %}
Returns users with the email addresses specified. Accepts up to 100 email addresses.  Any email addresses not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="user_id" type="string[]" %}
Returns users with the user ids specified.  Accepts up to 100 user ids. Any user ids not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="offset" type="string" %}
Offset allows pagination through all users.  If used, returns users starting after the number supplied.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="limit" type="integer" %}
Puts a limit on the number of users returned.  You may return anywhere from 1 to 100 users at a time.  Defaults to 10.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="order_by" type="string" %}
Allows to return users in a particular order. At the moment, you can order the returned users either by their `created_at` or `updated_at` timestamp.

In order to specify the direction, you can use the `+/-` symbols prepended in the order property. For example, if you want users to be returned in descending order according to their `created_at` property, you can use `-created_at`.

If you don't use `+` or `-`, then `+` is implied.

Defaults to `-created_at`.
{% endswagger-parameter %}

{% swagger-response status="200" description="Cake successfully retrieved." %}
```
[
    // see example schema
    {
        "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
        "object": "user"
        ...
    },
    {
        "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
        "object": "user"
        ...
    },
    ]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/users" method="post" summary="Create a user" %}
{% swagger-description %}
Create a user.  Your user management settings determine how you should setup your user model.  Any email address and phone number created using this method will be created as verified.

\




\


Note: If you're performing a migration, checkout our guide on zero downtime migrations
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="external_id" type="string" %}
The ID of the user you use in in your external systems.  Must be unique across your instance.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="first_name" type="string" %}
The first name to give to the user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="last_name" type="string" %}
The last name to give to the user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="email_address[]" type="string" %}
Email addresses to add to the user.  Must be unique across your instance.  The first email address will be set as the users primary email address.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="phone_number[]" type="string" %}
Phone numbers that will be added to the user.  Must be unique across your instance.  The first phone number will be set as the users primary phone number.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="username" type="string" %}
The username to give to the user.  It must be unique across your instance.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" type="string" %}
The plaintext password to give the user.  Must be at least 8 characters long, and can not be found in any list of hacked passwords.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password_digest" type="string" %}
In case you already have the password digests and not the passwords, you can use them for the newly created user via this property.&#x20;

The digests should be generated with one of the algorithms we support. The hashing algorithm can be specified using the `password_hasher` property.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password_hasher" type="string" %}
The hashing algorithm that was used to generate the password digest.

The algorithms we support at the moment are [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) and [pbkdf2\_sha256\_django](https://docs.djangoproject.com/en/4.0/topics/auth/passwords/).

If you need support for any particular hashing algorithm, [let us know](https://clerk.dev/support).
{% endswagger-parameter %}

{% swagger-parameter in="body" name="skip_password_checks" type="boolean" %}
When set to true all password checks are skipped. It is recommended to use this method only when migrating plaintext passwords to Clerk. Upon migration the user base should be prompted to pick stronger password.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="public_metadata" type="object" %}
Metadata saved on the user, that is visible to both your Frontend and Backend APIs.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="private_metadata" type="object" %}
Metadata saved on the user, that is only visible to your Backend API
{% endswagger-parameter %}

{% swagger-parameter in="body" name="unsafe_metadata" type="object" %}
Metadata saved on the user, that can be updated from both the Frontend and Backend APIs.  

\




\


Note: Since this data can be modified from the frontend, it is not guaranteed to be safe.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/users/:id" method="patch" summary="Update a user" %}
{% swagger-description %}
Update a user.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="first_name" %}
The first name to give to the user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="last_name" %}
The first name to give to the user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="username" %}
The username to give to the user.  It must be unique across your instance.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" %}
The plaintext password to give the user.  Must be at least 8 characters long, and can not be found in any list of hacked passwords.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="primary_email_address_idd" %}
The ID of the email address to set as primary. It must be verified, and present on the current user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="primary_phone_number_id" %}
The ID of the phone number to set as primary. It must be verified, and present on the current user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="profile_image_id" %}
The ID of the image to set as the users profile image.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="public_metadata" type="object" %}
Metadata saved on the user, that is visible to both your frontend and backend.

\




\


Note:  Object passed in will replace previous value.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="private_metadata" type="object" %}
Metadata saved on the user that is only visible to your backend.

\




\


Note: Object passed in will replace previous value.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="unsafe_metadata" type="object" %}
Metadata saved on the user, that can be updated from both the Frontend and Backend APIs.  

\




\


Note: Since this data can be modified from the frontend, it is not guaranteed to be safe.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user"
    ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/users/:id/profile_image" method="post" summary="Set a user's profile image" %}
{% swagger-description %}
Upload a new profile image for a user. Must use multipart/form-data with one image file.  It must be a jpg, png, gif, or webp image smaller than 10 MB.
{% endswagger-description %}

{% swagger-parameter in="body" name="file" type="object" %}
The image to upload.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/users/:id" method="delete" summary="Delete a user" %}
{% swagger-description %}
Delete a user.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user",
    "deleted": true
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev/" path="v1/users/:id/oauth_access_tokens/:provider" method="get" summary="Retrieve the OAuth access token of a user" %}
{% swagger-description %}
Fetch the corresponding OAuth access token for a user that has previously authenticated with a particular OAuth provider.

For OAuth 2.0, if the access token has expired and we have a corresponding refresh token, the access token will be refreshed transparently the new one will be returned.&#x20;
{% endswagger-description %}

{% swagger-parameter in="path" name="provider" type="string" required="true" %}
The ID of the OAuth provider (e.g. oauth_google)
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}
The user ID
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Response for an OAuth 2.0 token" %}
```javascript
{
  "provider": "oauth_google",
  "token": "xxxxxxxxxxxxxxxxxxxxx",
  "scopes": [
    "openid",
    "https://www.googleapis.com/auth/userinfo.email"
    "https://www.googleapis.com/auth/userinfo.profile"
  ]
}
```
{% endswagger-response %}

{% swagger-response status="200: OK" description="Response for an OAuth 1.0 token" %}
```javascript
{
  "provider": "oauth_twitter",
  "token": "xxxxxxxxxxxxxxxxxxxxx",
  "secret": "yyyyyyyyyyyyyyyyyyyy"
}
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="Unknown OAuth provider" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="No token exists for this particular user/provider combination" %}

{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="The access token has expired but the provider hasn't provided us with a refresh token and so we cannot fetch a new access token." %}

{% endswagger-response %}
{% endswagger %}
