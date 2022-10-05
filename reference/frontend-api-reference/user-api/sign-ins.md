# Sign ins

A **Sign In** represents an active sign in.  A **Client** can only have one active sign in at a time.

## Available requests

* **`POST`**`/v1/client/sign_ins`
* **`GET`**` ``/v1/client/sign_ins`
* **`POST`**`/v1/client/sign_ins/:id/prepare_first_factor`
* **`POST`**`/v1/client/sign_ins/:id/attempt_first_factor`
* **`POST`**`/v1/client/sign_ins/:id/prepare_second_factor`
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
**Creates or replaces the current Sign in object.**

  In order to authenticate a 

**Sign in**

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

__

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
&#x20;The first step of the strategy to perform, as part of this request. allowed options are:

`email_link`\
`phone_code`\
`email_code`\
`password`

`web3_metamask_signature`\
`oauth_[provider]`

`ticket`
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" type="string" %}
Required if the strategy is password.  The password to attempt to sign in with.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="redirect_url" type="string" %}
Required if the strategy is one of the OAuth providers. 

****

 This is the URL that the user will be redirected to after the OAuth verification completes.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="action_complete_redirect_url" type="string" %}
Optional if the strategy is one of the OAuth providers.  If the OAuth verification results in a completed Sign in, this is the URL that the user will be redirected to.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="ticket" type="string" %}
Required if the strategy is `ticket`.

It should contain the JWT ticket value that will be used to verify the user.

These JWT tickets are generated in the following cases:

`organization invitations`

`sign in tokens`
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

{% swagger baseUrl="https://clerk.example.com" path="/v1/client/sign_ins/:id" method="get" summary="Get a sign in" %}
{% swagger-description %}
**Get the Sign in object.**

\



{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}

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
Prepares the verification object for the identified **Sign in**.  ****  This step authenticates that the user is _who they say they are._  Depending on the strategy, this request request will do something different.\
\
Parameter actions:

If the **strategy** equals **email\_link** then this request will send an email magic link.

If the **strategy** equals **email\_code** then this request will send an email with an OTP code.\
\
If the **strategy** equals **phone\_code** then this request will send an SMS with an OTP code.

If the **strategy** equals **web3\_metamask\_signature** then this request generate a nonce that the **User** needs to sign in the browser using their Web3 Wallet browser extension.\
\
If the **strategy** equals **oauth\_\[provider]** then this request generate a URL that the **User** needs to visit in order to authenticate.\
\
Parameter rules:\
If the **strategy** equals \`**oauth\_\[provider]**\` then a **redirect\_url** is required, and an **action\_complete\_redirect\_url** is optional.\

{% endswagger-description %}

{% swagger-parameter in="body" name="strategy" type="string" %}
The strategy to prepare.  Allowed options are:

\`email\_link\`\
\`email\_code\`\
\`phone\_code\`

\`web3\_metamask\_signature\`\
\`oauth\_\[provider]\`
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
Attempt the first verification.  Requires the sign in attempt to be identified, and the first factor verification to be prepared, unless you're using a password.\
\
Parameter rules:\
\
If the **strategy** equals \`email\_code\` **** or \`phone\_code\` then a **code** is required.

If the **strategy** equals \`password\` then a **password** is required.

If the **strategy** equals \`web3\_metamask\_signature\` then a **signature** is required.\

{% endswagger-description %}

{% swagger-parameter in="body" name="strategy" type="string" %}
The strategy to attempt.  Allowed options are:&#x20;

\`email\_link\`\
\`email\_code\`\
\`phone\_code\`

\`password\`

\`web3\_metamask\_signature\`
{% endswagger-parameter %}

{% swagger-parameter in="body" name="code" type="string" %}
The code to attempt the verification with.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" type="string" %}
The password to attempt the verification with.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="signature" type="string" %}
Web3 wallet generated signature to be verified. This parameter is required only for web3 verification strategies.
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




`phone_code`

, 

`totp`

, 

`backup_code`
{% endswagger-parameter %}

{% swagger-parameter in="body" name="code" type="string" %}
The code to attempt the verification with.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
```
{% endswagger-response %}
{% endswagger %}
