# Users

This object represents a verified user in your instance.  

### Available requests

* **`GET`**  `/v1/users/:id`
* **`GET`**  `/v1/users`
* **`PATCH`**`/v1/users/:id`
* **`DEL`**  `/v1/users/:id`
* **`GET`**  `/v1/users/:id/oauth_access_token/`

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
    "external_accounts": [],
    "public_metadata": {},
    "private_metadata": {},
    "created_at": 1612756155,
    "updated_at": 1612756155
}
```

### 

{% api-method method="get" host="https://api.clerk.dev" path="/v1/users/:id" %}
{% api-method-summary %}
Retrieve a user
{% endapi-method-summary %}

{% api-method-description %}
Retrieve the details of a user. 
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authentication" type="string" required=true %}
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
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user"
    ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.clerk.dev" path="/v1/users" %}
{% api-method-summary %}
List all users
{% endapi-method-summary %}

{% api-method-description %}
List all users.  Ordered by creation date, with the newest user first.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-query-parameters %}
{% api-method-parameter name="phone\_number\[\]" type="string" required=false %}
Returns users with the phone numbers specified. Accepts up to 100 phone numbers. Any phone numbers not found are ignored.
{% endapi-method-parameter %}

{% api-method-parameter name="email\_address\[\]" type="string" required=false %}
Returns users with the email addresses specified. Accepts up to 100 email addresses.  Any email addresses not found are ignored.
{% endapi-method-parameter %}

{% api-method-parameter name="user\_id\[\]" type="string" required=false %}
Returns users with the user ids specified.  Accepts up to 100 user ids. Any user ids not found are ignored.
{% endapi-method-parameter %}

{% api-method-parameter name="offset" type="string" required=false %}
Offset allows pagination through all users.  If used, returns users starting after the number supplied.
{% endapi-method-parameter %}

{% api-method-parameter name="limit" type="integer" required=false %}
Puts a limit on the number of users returned.  You may return anywhere from 1 to 100 users at a time.  Defaults to 10.
{% endapi-method-parameter %}
{% endapi-method-query-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Cake successfully retrieved.
{% endapi-method-response-example-description %}

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
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="patch" host="https://api.clerk.dev" path="/v1/users/:id" %}
{% api-method-summary %}
Update a user
{% endapi-method-summary %}

{% api-method-description %}
Update a user.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-form-data-parameters %}
{% api-method-parameter name="public\_metadata" type="object" required=false %}
Metadata saved on the user, that is visible to both your frontend and backend.  
  
Note:  Object passed in will replace previous value.
{% endapi-method-parameter %}

{% api-method-parameter name="private\_metadata" type="object" required=false %}
Metadata saved on the user that is only visible to your backend.  
  
Note: Object passed in will replace previous value.
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
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user"
    ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="delete" host="https://api.clerk.dev" path="/v1/users/:id" %}
{% api-method-summary %}
Delete a user
{% endapi-method-summary %}

{% api-method-description %}
Delete a user.
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
{
    "id": "user_1oBNj55jOjSK9rOYrT5QHqj7eaK",
    "object": "user",
    "deleted": true
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://api.clerk.dev/" path="v1/users/:id/oauth\_access\_tokens/:provider" %}
{% api-method-summary %}
Retrieve OAuth access token for a user
{% endapi-method-summary %}

{% api-method-description %}
Retrieve a valid \(i.e. non-expired\) OAuth access token for a user that has previously authenticated with a particular OAuth provider.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="provider" type="string" required=true %}
The ID of the OAuth provider \(e.g. oauth\_google\).
{% endapi-method-parameter %}

{% api-method-parameter name="id" type="string" required=true %}
The user ID.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
{
  "token": "xxxxxxxxxxxxxxxxxxxxx",
  "provider": "oauth_google",
  "scopes": [
    "openid",
    "https://www.googleapis.com/auth/userinfo.email"
    "https://www.googleapis.com/auth/userinfo.profile"
  ]
}

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

