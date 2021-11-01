# Sign in

A **Sign In** represents an active sign in.  A **Client **can only have one active sign in at a time.

## Available requests

* **`POST`**`/v1/client/sign_ins`
* **`POST`**`/v1/client/sign_ins/:id/prepare_first_factor`
* **`POST`**`/v1/client/sign_ins/:id/attempt_first_factor`
* **`POST`**`/v1/client/sign_ins/:id/attempt_second_factor`
* **`POST`**`/v1/client/sign_ins/:id/attempt_second_factor`

## The Sign in object

```javascript
{
    "object": "sign_in_attempt",
    "status": "needs_identifier",
    "allowed_identifier_types": [
        "email_address",
        "oauth_google"
    ],
    "identifier": null,
    "user_data": null,
    "allowed_factor_one_strategies": [
        {
            "name": "oauth_google"
        },
        {
            "name": "password"
        }
    ],
    "factor_one_verification": null,
    "allowed_factor_two_strategies": [],
    "factor_two_verification": null,
    "created_session_id": null,
    "abandon_at": 1621115506
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ins" method="post" summary="Create a sign in" %}
{% swagger-description %}
**Creates or replaces the current Sign in object.  **

In order to authenticate a 

**Sign in **

in as few requests as possible, you can pass in parameters to this request that can identify and verify the 

**Sign in.**

\




\


Parameter rules:

\




\


If the 

**strategy**

 equals `

**phone_code**

\` then an 

**identifier**

 is required.

\




\


If the 

**strategy**

 equals `

**email_code**

\` then

_ _

an 

**identifier**

 is required.

\




\


If the 

**strategy**

 equals `

**password**

\` then both an 

**identifier**

 and a 

**password**

 is required.

\




\


If the 

**strategy**

 equals `

**oauth_[provider]**

\` then a 

**redirect_url**

 is required, and an 

**action_complete_redirect_url**

 is optional.

\



{% endswagger-description %}

{% swagger-parameter in="body" name="identifier" type="string" %}
An email address, phone number, or username.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="strategy" type="string" %}
 The first step of the strategy to perform, as part of this request. allowed options are:

\


\`phone_code`

\


\`email_code`

\


\`password`

\


\`oauth_[provider]`
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" type="string" %}
Required if the strategy is password.  The password to attempt to sign in with.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="redirect_url" type="string" %}
Required if the strategy is one of the OAuth providers.

** **

This is the URL that the user will be redirected to after the OAuth verification completes.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="action_complete_redirect_url" type="string" %}
Optional if the strategy is one of the OAuth providers.  If the OAuth verification results in a completed Sign in, this is the URL that the user will be redirected to.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
        "object": "session",
        "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
        ...

```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ins/:id/prepare_first_factor" method="post" summary="Prepare first factor" %}
{% swagger-description %}
Prepares the verification object for the identified 

**Sign in**

.

**  **

This step authenticates that the user is 

_who they say they are.  _

Depending on the strategy, this request request will do something different.

\




\


Parameter actions:

\


If the 

**strategy **

equals 

**email_code**

 then this request will send an email with an OTP code.

\




\


If the 

**strategy**

 equals 

**phone_code **

then this request will send an SMS with an OTP code.

\




\


If the 

**strategy**

 equals 

**oauth_[provider]**

 then this request generate a URL that the 

**User**

 needs to visit in order to authenticate.

\




\




\


Parameter rules:

\


If the 

**strategy**

 equals `

**oauth_[provider]**

\` then a 

**redirect_url**

 is required, and an 

**action_complete_redirect_url**

 is optional.

\



{% endswagger-description %}

{% swagger-parameter in="body" name="strategy" type="string" %}
The strategy to prepare.  Allowed options are:

\


\`email_code`

\


\`phone_code`

\


\`oauth_[provider]`
{% endswagger-parameter %}

{% swagger-parameter in="body" name="redirect_url" type="string" %}
Required if the strategy is one of the oauth providers. This is the URL the user will be directed to  after the Oauth verification completes
{% endswagger-parameter %}

{% swagger-parameter in="body" name="action_complete_redirect_url" type="string" %}
Optional if the strategy is one of the OAuth providers.  If the OAuth verification results in a completed Sign in, this is the URL the user will be redirected to.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}

{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}



{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ins/:id/attempt_first_factor" method="post" summary="Attempt first factor" %}
{% swagger-description %}
Attempt the first verification.  Requires the sign in attempt to be identified, and the first factor verification to be prepared, unless you're using a password.

\




\


Parameter rules:

\




\


If the 

**strategy**

 equals `email_code`

** **

or `phone_code` then a 

**code**

 is required.

\


If the 

**strategy**

 equals `password` then a 

**password**

 is required. Additionally

\



{% endswagger-description %}

{% swagger-parameter in="body" name="strategy" type="string" %}
The strategy to attempt.  Allowed options are: 

\


\`email_code`

\


\`phone_code`

\


\`password`
{% endswagger-parameter %}

{% swagger-parameter in="body" name="code" type="string" %}
The code to attempt the verification with.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" type="string" %}
The password to attempt the verification with.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}

{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}



{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ins/:id/prepare_second_factor" method="post" summary="Prepare second factor" %}
{% swagger-description %}
Prepare the second verification.  Requires the sign in attempt `status` to be equal to `needs_factor_two`.

\



{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}

{% endswagger-parameter %}

{% swagger-parameter in="body" name="strategy" type="string" %}
The strategy to prepare.  Allowed options are:

\


\`phone_code`
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ins/:id/attempt_second_factor" method="post" summary="Attempt factor two" %}
{% swagger-description %}
Attempt the second verification.  Requires the sign in attempt `status` to be equal to `needs_factor_two`, and for the preparation step to have been called.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}

{% endswagger-parameter %}

{% swagger-parameter in="body" name="strategy" type="string" %}
The strategy to attempt.  Allowed options are:

\


\`phone_code`
{% endswagger-parameter %}

{% swagger-parameter in="body" name="code" type="string" %}
The code to attempt the verification with.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}
