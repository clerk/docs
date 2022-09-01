# Users

This object represents a verified user in your instance. &#x20;

### Available requests

* **`GET`**`  ``/v1/users/:id`
* **`GET`**`  ``/v1/users`
* **`GET`**`  ``/v1/users/count`
* **`POST`**` ``/v1/users`
* **`PATCH`**`/v1/users/:id`
* **`PATCH`**`/v1/users/:id/metadata`
* **`POST`**` ``/v1/users/:id/profile_image`
* **`POST`**` ``/v1/users/:id/verify_password`
* **`DEL`**`  ``/v1/users/:id`
* **`GET`**`  ``/v1/users/:id/oauth_access_tokens/:provider`
* **`DEL`**`  ``/v1/users/:id/totp`

### Example user schema

```javascript
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user",
    "external_id": "my_previous_auth_solution_id",
    "username": null,
    "first_name": "Boss",
    "last_name": "Clerk",
    "profile_image_url": "https://images.clerk.services/clerk/default-profile.svg",
    "primary_email_address_id": "idn_1oBNgISXFbSf5m0uP2Wl0qWtNGX",
    "primary_phone_number_id": null,
    "primary_web3_wallet_id": null,
    "password_enabled": true,
    "two_factor_enabled": false,
    "totp_enabled": false,
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

{% swagger-parameter in="query" name="email_address" type="string[]" %}
Returns users with the specified email addresses. Accepts up to 100 email addresses.  Any email addresses not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="phone_number" type="string[]" %}
Returns users with the specified phone numbers. Accepts up to 100 phone numbers. Any phone numbers not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="external_id" type="string[]" %}
Returns users with the specified external ids. Accepts up to 100 external ids. Any external ids not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="username" type="string[]" %}
Returns users with the specified usernames. Accepts up to 100 usernames. Any usernames not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="web3_wallet" type="string[]" %}
Returns users with the specified web3 wallet addresses. Accepts up to 100 web3 wallet addresses. Any web3 wallet addressed not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="user_id" type="string[]" %}
Returns users with the user ids specified.  Accepts up to 100 user ids. Any user ids not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="query" type="string" %}
Returns users that match the given query.

For possible matches, we check the email addresses, phone numbers, usernames, web3 wallets, user ids, first and last names.

The query value doesn't need to match the exact value you're looking for, it's capable of partial matches as well.
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

{% swagger method="get" path="/v1/users/count" baseUrl="https://api.clerk.dev" summary="Count users" %}
{% swagger-description %}
Returns a total count of all users that match the given filtering criteria.
{% endswagger-description %}

{% swagger-parameter in="query" name="email_address" type="string[]" %}
Returns users with the specified email addresses. Accepts up to 100 email addresses.  Any email addresses not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="phone_number" type="string[]" %}
Returns users with the specified phone numbers. Accepts up to 100 phone numbers. Any phone numbers not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="username" type="string[]" %}
Returns users with the specified usernames. Accepts up to 100 usernames. Any usernames not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="web3_wallet" type="string[]" %}
Returns users with the specified web3 wallet addresses. Accepts up to 100 web3 wallet addresses. Any web3 wallet addressed not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="user_id" type="string[]" %}
Returns users with the user ids specified.  Accepts up to 100 user ids. Any user ids not found are ignored.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="query" type="string[]" %}
Returns users that match the given query.

For possible matches, we check the email addresses, phone numbers, usernames, web3 wallets, user ids, first and last names.

The query value doesn't need to match the exact value you're looking for, it's capable of partial matches as well.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorizaton" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The total number of users that match the given filter criteria" %}
```javascript
{
    object: "total_count",
    total_count: 1 // number of users
}
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
The ID of the user you use in your external systems or your previous authentication solution.  Must be unique across your instance.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="first_name" type="string" %}
The first name to give to the user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="last_name" type="string" %}
The last name to give to the user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="email_address" type="string[]" %}
Email addresses to add to the user.  Must be unique across your instance.  The first email address will be set as the users primary email address.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="phone_number" type="string[]" %}
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

The algorithms we support at the moment are [bcrypt](https://en.wikipedia.org/wiki/Bcrypt), [pbkdf2\_sha256\_django](https://docs.djangoproject.com/en/4.0/topics/auth/passwords/) and [scrypt\_firebase](https://firebaseopensource.com/projects/firebase/scrypt/).

Each of the above expects the incoming digest to be of a particular format. More specifically:



_**bcrypt**_: The digest should be of the following form:&#x20;

`$<algorithm version>$<cost>$<salt & hash>`\


_**pbkdf2\_sha256\_django**:_ This is the Django-specific variant of PBKDF2 and the digest should have the following format (as exported from Django):

`pbkdf2_sha256$<iterations>$<salt>$<hash>`\


_**scrypt\_firebase**_: The Firebase-specific variant of scrypt.

The value is expected to have 6 segments separated by the `$` character and include the following information:

_hash_: The actual Base64 hash. This can be retrieved when exporting the user from Firebase.

_salt_: The salt used to generate the above hash. Again, this is given when exporting the user.

_signer key_: The base64 encoded signer key.

_salt separator_: The base64 encoded salt separator.

_rounds_: The number of rounds the algorithm needs to run.

_memory cost_: The cost of the algorithm run



The first 2 (hash and salt) are per user and can be retrieved when exporting the user from Firebase.

The other 4 values (signer key, salt separator, rounds and memory cost) are project-wide settings and can be retrieved from the project's password hash parameters.



Once you have all these, you can combine it in the following format and send this as the digest in order for Clerk to accept it:

`<hash>$<salt>$<signer key>$<salt separator>$<rounds>$<memory cost>`



If you need support for any particular hashing algorithm, [let us know](https://clerk.dev/support).
{% endswagger-parameter %}

{% swagger-parameter in="body" name="skip_password_checks" type="boolean" %}
When set to true all password checks are skipped. It is recommended to use this method only when migrating plaintext passwords to Clerk. Upon migration the user base should be prompted to pick stronger password.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="skip_password_requirement" type="boolean" %}
When set to true, `password` is not required anymore when creating the user and can be omitted.

This is useful when you are trying to create a user that doesn't have a password, in an instance that's using passwords.

Please note that you cannot use this flag if password is the only way for a user to sign into your instance.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="totp_secret" type="string" %}
In case the user already have TOTP configured you can provide the secret to enable it in the newly created user without the need to reset it.&#x20;

Please note that currently the supported options are:\
\- Period: 30 seconds\
\- Code length: 6 digits\
\- Algorithm: SHA1
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
Update a user's attributes.

You can set the user's primary contact identifiers (email address and phone numbers) by updating the `primary_email_address_id` and `primary_phone_number_id` attributes respectively. Both IDs should correspond to verified identifications that belong to the user.

You can remove a user's username by setting the `username` attribute to `null` or the blank string `""`. This is a destructive action. The identification will be deleted forever.

Usernames can be removed only if they are optional in your instance settings and there's at least one other identifier which can be used for authentication.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="first_name" %}
The first name to give to the user.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="last_name" %}
The last name to give to the user.
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

{% swagger baseUrl="https://api.clerk.dev" path="/v1/users/:id/metadata" method="patch" summary="Merge and update a user's metadata" %}
{% swagger-description %}
Update a user's metadata attributes by merging existing values with the provided parameters.&#x20;

This endpoint behaves differently than the **Update a user** endpoint. Metadata values will not be replaced entirely. Instead, a deep merge will be performed. Deep means that any nested JSON objects will be merged as well.

You can remove metadata keys at any level by setting their value to `null`.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="public_metadata" type="object" %}
Metadata saved on the user, that is visible to both your frontend and backend.

\


The new object will be merged with the existing value.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="private_metadata" type="object" %}
Metadata saved on the user that is only visible to your backend.

\


The new object will be merged with the existing value.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="unsafe_metadata" type="object" %}
Metadata saved on the user, that can be updated from both the Frontend and Backend APIs.  

\


The new object will be merged with the existing value.

\




\


Note: Since this data can be modified from the frontend, it is not guaranteed to be safe.
{% endswagger-parameter %}

{% swagger-response status="200" description="User metadata successfully updated." %}
```
// see example schema
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user"
    ...}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="User not found on the instance." %}
```javascript
{
  "errors": [
    {
      "code": "resource_not_found",
      "long_message": "No user was found with id user_26oyVupb9nU9hSYPh1yJJNqe7Hl",
      "message": "not found"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="Invalid metadata parameters" %}
```javascript
{
  "errors": [
    {
      "code": "form_param_exceeds_allowed_size",
      "long_message": "The given public_metadata exceeds the maximum allowed size of 4096 bytes (4 KB).",
      "message": "The given public_metadata exceeds the maximum allowed size of 4096 bytes (4 KB).",
      "meta": {
        "param_name": "public_metadata"
      }
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/users/:id/profile_image" method="post" summary="Set a user's profile image" %}
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

{% swagger method="post" path="/v1/users/:id/verify_password" baseUrl="https://api.clerk.dev" summary="Verify a user's password" %}
{% swagger-description %}
Check that the users password matches the supplied input.  Useful for custom auth flows and re-verification.
{% endswagger-description %}

{% swagger-parameter in="body" name="password" type="string" required="true" %}
the password
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Returns if the password matches" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="Returns if there's an error, or if the password doesn't match" %}
```javascript
{
    // Response
}
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
[
  {
    "object": "oauth_access_token",
    "provider": "oauth_google",
    "token": "xxxxxxxxxxxxxxxxxxxxx",
    "public_metadata": {},
    "label": "clerk",
    "scopes": [
      "openid",
      "https://www.googleapis.com/auth/userinfo.email"
      "https://www.googleapis.com/auth/userinfo.profile"
    ]
  }
  ...
]
```
{% endswagger-response %}

{% swagger-response status="200: OK" description="Response for an OAuth 1.0 token" %}
```javascript
[
  {
    "object": "oauth_access_token",
    "provider": "oauth_twitter",
    "token": "xxxxxxxxxxxxxxxxxxxxx",
    "public_metadata": {},
    "label": "clerk",
    "token_secret": "xxxxxxxxxxxxxxxxxxxxx"
  }
  ...
]
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

{% swagger method="delete" path="v1/users/:id/totp" baseUrl="https://api.clerk.dev/" summary="Disable user's TOTP" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" required="true" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Success" %}
```javascript
{
    "id": "totp_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "totp",
    "deleted": true
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="User not found or TOTP not enabled for the specific user" %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}
