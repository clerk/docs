# Sign in

A **Sign In** represents an active sign in.  A **Client** can only have one active sign in at a time.

## Available requests

* **`POST`**`/v1/client/sign_in_attempt`
* **`POST`**`/v1/client/sign_in_attempt/prepare_factor_one`
* **`POST`**`/v1/client/sign_in_attempt/attempt_factor_one`
* **`POST`**`/v1/client/sign_in_attempt/prepare_factor_two`
* **`POST`**`/v1/client/sign_in_attempt/attempt_factor_two`

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

{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sign\_in\_attempt" %}
{% api-method-summary %}
Create or replace
{% endapi-method-summary %}

{% api-method-description %}
**Creates or replaces the current Sign in object.**  In order to authenticate a **Sign in** in as few requests as possible, you can pass in parameters to this request that can identify and verify the **Sign in.**  
  
Parameter rules:  
  
If the **strategy** equals \`**phone\_code**\` then an **identifier** is required.  
  
If the **strategy** equals \`**email\_code**\` then __an **identifier** is required.  
  
If the **strategy** equals \`**password**\` then both an **identifier** and a **password** is required.  
  
If the **strategy** equals \`**oauth\_\[provider\]**\` then a **redirect\_url** is required, and an **action\_complete\_redirect\_url** is optional.  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="identifier" type="string" required=false %}
An email address, phone number, or username.
{% endapi-method-parameter %}

{% api-method-parameter name="strategy" type="string" required=false %}
 The first step of the strategy to perform, as part of this request. allowed options are:  
\`phone\_code\`  
\`email\_code\`  
\`password\`  
\`oauth\_\[provider\]\`
{% endapi-method-parameter %}

{% api-method-parameter name="password" type="string" required=false %}
Required if the strategy is password.  The password to attempt to sign in with.
{% endapi-method-parameter %}

{% api-method-parameter name="redirect\_url" type="string" required=false %}
Required if the strategy is one of the OAuth providers. ****This is the URL that the user will be redirected to after the OAuth verification completes.
{% endapi-method-parameter %}

{% api-method-parameter name="action\_complete\_redirect\_url" type="string" required=false %}
Optional if the strategy is one of the OAuth providers.  If the OAuth verification results in a completed Sign in, this is the URL that the user will be redirected to.
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
        "object": "session",
        "id": "sess_1q8uCbeJSMgTJnTQMgpAK1Ff0ER",
        ...

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sign\_in\_attempt/prepare\_factor\_one" %}
{% api-method-summary %}
Prepare factor one
{% endapi-method-summary %}

{% api-method-description %}
Prepares the verification object for the identified **Sign in**.  ****This step authenticates that the user is _who they say they are._  Depending on the strategy, this request request will do something different.  
  
Parameter actions:  
If the **strategy** equals **email\_code** then this request will send an email with an OTP code.  
  
If the **strategy** equals **phone\_code** then this request will send an SMS with an OTP code.  
  
If the **strategy** equals **oauth\_\[provider\]** then this request generate a URL that the **User** needs to visit in order to authenticate.  
  
  
Parameter rules:  
If the **strategy** equals \`**oauth\_\[provider\]**\` then a **redirect\_url** is required, and an **action\_complete\_redirect\_url** is optional.  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="strategy" type="string" required=true %}
The strategy to prepare.  Allowed options are:  
\`email\_code\`  
\`phone\_code\`  
\`oauth\_\[provider\]\`
{% endapi-method-parameter %}

{% api-method-parameter name="redirect\_url" type="string" required=false %}
Required if the strategy is one of the oauth providers. This is the URL the user will be directed to  after the Oauth verification completes
{% endapi-method-parameter %}

{% api-method-parameter name="action\_complete\_redirect\_url" type="string" required=false %}
Optional if the strategy is one of the OAuth providers.  If the OAuth verification results in a completed Sign in, this is the URL the user will be redirected to.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sign\_in\_attempt/attempt\_factor\_one" %}
{% api-method-summary %}
Attempt factor one
{% endapi-method-summary %}

{% api-method-description %}
Attempt the first verification.  Requires the sign in attempt to be identified, and the first factor verification to be prepared, unless you're using a password.  
  
Parameter rules:  
  
If the **strategy** equals \`email\_code\` ****or \`phone\_code\` then a **code** is required.  
If the **strategy** equals \`password\` then a **password** is required. Additionally  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="strategy" type="string" required=true %}
The strategy to attempt.  Allowed options are:   
\`email\_code\`  
\`phone\_code\`  
\`password\`
{% endapi-method-parameter %}

{% api-method-parameter name="code" type="string" required=false %}
The code to attempt the verification with.
{% endapi-method-parameter %}

{% api-method-parameter name="password" type="string" required=false %}
The password to attempt the verification with.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}



{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sign\_in\_attempt/prepare\_factor\_two" %}
{% api-method-summary %}
Prepare factor two
{% endapi-method-summary %}

{% api-method-description %}
Prepare the second verification.  Requires the sign in attempt \`status\` to be equal to \`needs\_factor\_two\`.  
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="strategy" type="string" required=true %}
The strategy to prepare.  Allowed options are:  
\`phone\_code\`
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://clerk.example.com" path="/v1/client/sign\_in\_attempt/attempt\_factor\_two" %}
{% api-method-summary %}
Attempt factor two
{% endapi-method-summary %}

{% api-method-description %}
Attempt the second verification.  Requires the sign in attempt \`status\` to be equal to \`needs\_factor\_two\`, and for the preparation step to have been called.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="strategy" type="string" required=true %}
The strategy to attempt.  Allowed options are:  
\`phone\_code\`
{% endapi-method-parameter %}

{% api-method-parameter name="code" type="string" required=true %}
The code to attempt the verification with.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

