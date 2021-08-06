# Sign up

A **Sign up** represents an active sign up.  A **Client** can only have one active sign up at a time.

## Available requests

**Create/Update**

* **`POST`** `/v1/client/sign_up`
* **`PATCH`**`/v1/client/sign_up`

**Verification**

* **`POST`**`/v1/client/sign_up/prepare_verification`
* **`POST`**`/v1/client/sign_up/attempt_verification`

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

{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sign\_up" %}
{% api-method-summary %}
Create a Sign Up
{% endapi-method-summary %}

{% api-method-description %}
Creates a new **Sign Up**. The new Sign Up will replace any pending **Sign Up**s on the current **Client**. Use this call to initiate \(or create\) a new **User** registration at any point.   
  
You can optionally pass any of the following parameters as part of the creation process. You don't have to pass all required registration parameters at once, since a **Sign Up** can also be updated at a later stage with a   
`PATCH /v1/client/sign_up` request.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="phone\_number" type="string" required=false %}
The phone number of the new **User**.  Must be a valid phone number, including the country code, and be unique.
{% endapi-method-parameter %}

{% api-method-parameter name="email\_address" type="string" required=false %}
The email address of the new **User**. Must be a valid email address, and be unique.
{% endapi-method-parameter %}

{% api-method-parameter name="username" type="string" required=false %}
The username of the new **User**.  Must be unique, and only contain alphanumeric characters \(letters A-Z, numbers 0-9\) with the exception of underscores.
{% endapi-method-parameter %}

{% api-method-parameter name="last\_name" type="string" required=false %}
The last name of the new **User**.
{% endapi-method-parameter %}

{% api-method-parameter name="first\_name" type="string" required=false %}
The first name of the new **User**.
{% endapi-method-parameter %}

{% api-method-parameter name="password" type="string" required=false %}
The password of the new **User**.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
`POST /v1/client/sign_up` with no request body parameters.
{% endapi-method-response-example-description %}

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
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="patch" host="https://clerk.example.com" path="/v1/client/sign\_up" %}
{% api-method-summary %}
Update
{% endapi-method-summary %}

{% api-method-description %}
Updates the current **Sign Up** object with the supplied parameters. You can trigger this request at any time during the sign up process, but before the **Sign Up** is completed.  
The update sign up request provides an opportunity to update the new **User** as you gather more information.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="phone\_number" type="string" required=false %}
The phone number of the new **User**.  Must be a valid phone number, including the country code, and be unique.
{% endapi-method-parameter %}

{% api-method-parameter name="email\_address" type="string" required=false %}
The email address of the new **User**. Must be a valid email address, and be unique.
{% endapi-method-parameter %}

{% api-method-parameter name="username" type="string" required=false %}
The username of the new **User**.  Must be unique, and only contain alphanumeric characters \(letters A-Z, numbers 0-9\) with the exception of underscores.
{% endapi-method-parameter %}

{% api-method-parameter name="last\_name" type="string" required=false %}
The last name of the new **User**.
{% endapi-method-parameter %}

{% api-method-parameter name="first\_name" type="string" required=false %}
The first name of the new **User**.
{% endapi-method-parameter %}

{% api-method-parameter name="password" type="string" required=false %}
The password of the new **User**.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
`PATCH /v1/client/sign_up`, passing the `email_address`, `first_name` and `last_name` parameters.
{% endapi-method-response-example-description %}

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
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

## Verification

{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sign\_up/prepare\_verification" %}
{% api-method-summary %}
Prepare a verification
{% endapi-method-summary %}

{% api-method-description %}
In order to kick off the verification process, you need to prepare a verification for the current **Sign Up**. Pass the strategy parameter to specify what kind of verification process will be started. Currently, we support verification by email, phone or an external account \(usually via OAuth\).  
  
You can find more information about the state of the verification under the `verifications` field in the response body.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="strategy" type="string" required=true %}
The verification strategy. Can be one of `email_code`, `phone_code` or one of the oauth related strategies like `oauth_google`, `oauth_facebook`, `oauth_<provider>` in general.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
`POST /v1/client/sign_up/prepare_verification` with `email_code` strategy.
{% endapi-method-response-example-description %}

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
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sign\_up/attempt\_verification" %}
{% api-method-summary %}
Attempt a verification
{% endapi-method-summary %}

{% api-method-description %}
Use this request to check a previously prepared verification. Attempting a verification needs the `strategy` and the already shared `code` that needs to be checked.  
  
You can find more information regarding the verification check under the `verifications` response field. 
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="strategy" type="string" required=true %}
The strategy that was used when preparing the verification.
{% endapi-method-parameter %}

{% api-method-parameter name="code" type="string" required=true %}
A character string that needs to be verified. Usually 6 digits.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
`POST /v1/client/sign_up/attempt_verification` with strategy `email_code` and a valid `code`.
{% endapi-method-response-example-description %}

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
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

