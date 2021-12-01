# Sign ups

A **Sign up** represents an active sign up.  A **Client** can only have one active sign up at a time.

## Available requests

**Create/Read/Update**

* **`POST`**` ``/v1/client/sign_ups`
* **`GET`**`/v1/client/sign_ups/:id`
* **`PATCH`**`/v1/client/sign_ups/:id`

**Verification**

* **`POST`**`/v1/client/sign_ups/:id/prepare_verification`
* **`POST`**`/v1/client/sign_ups/:id/attempt_verification`

## The Sign up object

```javascript
{ 
    "object": "sign_up_attempt", 
    "status": "missing_requirements", 
    "required_fields": [
        "email_address",
        "password",
        "first_name",
        "last_name"
    ], 
    "optional_fields": [],
    "missing_fields": [
        "email_address",
        "password",
        "first_name",
        "last_name"
    ],
    "unverified_fields": [
        "email_address"
    ],
    "supported_external_accounts": [
        "oauth_google"
    ],
    "username": null, 
    "email_address": null, 
    "phone_number": null, 
    "external_account": null, 
    "password_enabled": false, 
    "first_name": null, 
    "last_name": null, 
    "created_session_id": null, 
    "created_user_id": null, 
    "abandon_at": 1622256299572,
    "verifications": {
        "email_address": null,
        "phone_number": null,
        "external_account": null
    }
}
```

## Create/Update

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ups" method="post" summary="Create a Sign up" %}
{% swagger-description %}
Creates a new 

**Sign Up**

. The new Sign Up will replace any pending 

**Sign Up**

s on the current 

**Client**

. Use this call to initiate (or create) a new 

**User**

 registration at any point. 

\




\


You can optionally pass any of the following parameters as part of the creation process. You don't have to pass all required registration parameters at once, since a 

**Sign Up**

 can also be updated at a later stage with a 

\




`PATCH /v1/client/sign_up`

 request.
{% endswagger-description %}

{% swagger-parameter in="body" name="phone_number" type="string" %}
The phone number of the new 

**User**

.  Must be a valid phone number, including the country code, and be unique.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="email_address" type="string" %}
The email address of the new 

**User**

. Must be a valid email address, and be unique.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="username" type="string" %}
The username of the new 

**User**

.  Must be unique, and only contain alphanumeric characters (letters A-Z, numbers 0-9) with the exception of underscores.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="last_name" type="string" %}
The last name of the new 

**User**

.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="first_name" type="string" %}
The first name of the new 

**User**

.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" type="string" %}
The password of the new 

**User**

.
{% endswagger-parameter %}

{% swagger-response status="200" description="POST /v1/client/sign_up with no request body parameters." %}
```
{
  "object": "sign_up_attempt",
  "status": "missing_requirements",
  "required_fields": [
    "first_name",
    "last_name",
    "email_address",
    "password"
  ],
  "optional_fields": [],
  "missing_fields": [
    "first_name",
    "last_name",
    "password"
  ],
  "unverified_fields": [
    "email_address"
  ],
  "supported_external_accounts": [
    "oauth_google"
  ],
  "username": null,
  "email_address": null,
  "phone_number": null,
  "external_account": null,
  "password_enabled": false,
  "first_name": null,
  "last_name": null,
  "created_session_id": null,
  "created_user_id": null,
  "abandon_at": 1626170619989,
  "verifications": {
    "email_address": null,
    "phone_number": null,
    "external_account": null
  }
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/v1/client/sign_ups/:id" baseUrl="https://clerk.example.com" summary="Get a Sign up" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}

{% endswagger-parameter %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ups/:id" method="patch" summary="Update" %}
{% swagger-description %}
Updates the current 

**Sign Up**

 object with the supplied parameters. You can trigger this request at any time during the sign up process, but before the 

**Sign Up**

 is completed.

\


The update sign up request provides an opportunity to update the new 

**User**

 as you gather more information.
{% endswagger-description %}

{% swagger-parameter in="body" name="phone_number" type="string" %}
The phone number of the new 

**User**

.  Must be a valid phone number, including the country code, and be unique.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="email_address" type="string" %}
The email address of the new 

**User**

. Must be a valid email address, and be unique.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="username" type="string" %}
The username of the new 

**User**

.  Must be unique, and only contain alphanumeric characters (letters A-Z, numbers 0-9) with the exception of underscores.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="last_name" type="string" %}
The last name of the new 

**User**

.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="first_name" type="string" %}
The first name of the new 

**User**

.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" type="string" %}
The password of the new 

**User**

.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}

{% endswagger-parameter %}

{% swagger-response status="200" description="PATCH /v1/client/sign_up, passing the email_address, first_name and last_name parameters." %}
```
{
  "object": "sign_up_attempt",
  "status": "missing_requirements",
  "required_fields": [
    "first_name",
    "last_name",
    "email_address",
    "password"
  ],
  "optional_fields": [],
  "missing_fields": [
    "password"
  ],
  "unverified_fields": [
    "email_address"
  ],
  "supported_external_accounts": [
    "oauth_google"
  ],
  "username": null,
  "email_address": "homer@example.com",
  "phone_number": null,
  "external_account": null,
  "password_enabled": false,
  "first_name": "Homer",
  "last_name": "Example",
  "created_session_id": null,
  "created_user_id": null,
  "abandon_at": 1626170619989,
  "verifications": {
    "email_address": null,
    "phone_number": null,
    "external_account": null
  }
}
```
{% endswagger-response %}
{% endswagger %}

## Verification

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ups/:id/prepare_verification" method="post" summary="Prepare a verification" %}
{% swagger-description %}
In order to kick off the verification process, you need to prepare a verification for the current 

**Sign Up**

. Pass the strategy parameter to specify what kind of verification process will be started. Currently, we support verification by email, phone or an external account (usually via OAuth).

\




\


You can find more information about the state of the verification under the 

`verifications`

 field in the response body.
{% endswagger-description %}

{% swagger-parameter in="body" name="strategy" type="string" %}
The verification strategy. Can be one of 

`email_code`

, 

`phone_code`

 or one of the oauth related strategies like 

`oauth_google`

, 

`oauth_facebook`

, 

`oauth_<provider>`

 in general.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}

{% endswagger-parameter %}

{% swagger-response status="200" description="POST /v1/client/sign_up/prepare_verification with email_code strategy." %}
```
{
  "object": "sign_up_attempt",
  "status": "missing_requirements",
  "required_fields": [],
  "optional_fields": [],
  "missing_fields": [],
  "unverified_fields": [
    "email_address"
  ],
  "supported_external_accounts": [
    "oauth_google"
  ],
  "username": null,
  "email_address": "homer@example.com",
  "phone_number": null,
  "external_account": null,
  "password_enabled": false,
  "first_name": "Homer",
  "last_name": "Example",
  "created_session_id": null,
  "created_user_id": null,
  "abandon_at": 1626170619989,
  "verifications": {
    "email_address": {
      "status": "unverified",
      "strategy": "email_code",
      "attempts": 0,
      "expire_at": 1626084832122,
      "next_action": "needs_attempt",
      "supported_strategies": [
        "email_code"
      ]
    },
    "phone_number": null,
    "external_account": null
  }
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ups/:id/attempt_verification" method="post" summary="Attempt a verification" %}
{% swagger-description %}
Use this request to check a previously prepared verification. Attempting a verification needs the 

`strategy`

 and the already shared 

`code`

 that needs to be checked.

\




\


You can find more information regarding the verification check under the 

`verifications`

 response field. 
{% endswagger-description %}

{% swagger-parameter in="body" name="strategy" type="string" %}
The strategy that was used when preparing the verification.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="code" type="string" %}
A character string that needs to be verified. Usually 6 digits.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}

{% endswagger-parameter %}

{% swagger-response status="200" description="POST /v1/client/sign_up/attempt_verification with strategy email_code and a valid code." %}
```
{
  "object": "sign_up_attempt",
  "status": "missing_requirements",
  "required_fields": [
    "password"
  ],
  "optional_fields": [],
  "missing_fields": [],
  "unverified_fields": [],
  "supported_external_accounts": [
    "oauth_google"
  ],
  "username": null,
  "email_address": "homer@example.com",
  "phone_number": null,
  "external_account": null,
  "password_enabled": false,
  "first_name": null,
  "last_name": null,
  "created_session_id": null,
  "created_user_id": null,
  "abandon_at": 1626170619989,
  "verifications": {
    "email_address": {
      "status": "verified",
      "strategy": "email_code",
      "attempts": 2,
      "expire_at": 1626084832122,
      "next_action": "",
      "supported_strategies": [
        "email_code"
      ]
    },
    "phone_number": null,
    "external_account": null
  }
}
```
{% endswagger-response %}
{% endswagger %}
