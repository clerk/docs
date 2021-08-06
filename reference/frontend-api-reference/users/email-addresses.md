# Email addresses

## Available requests

* **`GET`** `/v1/me/email_addresses`
* **`POST`**`/v1/me/email_addresses`
* **`GET`** `/v1/me/email_addresses/:id`
* **`POST`**`/v1/me/email_addresses/:id/send_verification_email`
* **`POST`**`/v1/me/email_addresses/:id/verify`
* **`DEL`** `/v1/me/email_addresses/:id`

## The email address object

```javascript
{
    "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
    "object": "email_address",
    "email_address": "boss@clerk.dev",
    "verification": {
        "status": "verified",
        "strategy": "from_oauth_google"
    },
    "linked_to": [
        {
            "type": "oauth_google",
            "id": "idn_1o4qfbU8YjpJgfpDza4U4aTbClb"
        }
    ]
}
```

{% api-method method="get" host="https://clerk.example.com" path="/v1/me/email\_addresses" %}
{% api-method-summary %}
List all email addresses
{% endapi-method-summary %}

{% api-method-description %}
Returns a list of the current user's email addresses.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the email address object above.
[
  {
    "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
    "object": "email_address",
    ...
  },
  ...
]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://clerk.example.com" path="/v1/me/email\_addresses/:id" %}
{% api-method-summary %}
Retrieve an email address
{% endapi-method-summary %}

{% api-method-description %}
Returns the current user's email address with the given ID.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="id" type="string" required=true %}
ID of the email address.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the email address object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "email_address",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://clerk.example.com" path="/v1/me/email\_addresses" %}
{% api-method-summary %}
Create an email address
{% endapi-method-summary %}

{% api-method-description %}
Create an email address for the current user.  This email address will need to be verified before it is active on the user.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="email\_address" type="string" required=true %}
The email address that you are adding to the current user.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the email address object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "email_address",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://clerk.example.com" path="/v1/me/email\_addresses/:id/send\_verification\_email" %}
{% api-method-summary %}
Send a verification email
{% endapi-method-summary %}

{% api-method-description %}
Send a verification email to the email address with the given ID.  Can only send to email addresses that have not been verified.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="id" type="string" required=true %}
ID of the email address.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the email address object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "email_address",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://clerk.example.com" path="/v1/me/email\_addresses/:id/verify" %}
{% api-method-summary %}
Verify an email address
{% endapi-method-summary %}

{% api-method-description %}
Verify the current user's email address.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="id" type="string" required=true %}
ID of the email address.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-form-data-parameters %}
{% api-method-parameter name="code" type="string" required=true %}
The 6 digit code that was sent via email.  You only get 3 chances to correctly verify, before a new code must be generated.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the email address object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "email_address",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="delete" host="https://clerk.example.com" path="/v1/me/email\_addresses/:id" %}
{% api-method-summary %}
Delete an email address
{% endapi-method-summary %}

{% api-method-description %}
Delete the current user's email address.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="id" type="string" required=true %}
ID of the email address.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
    "id": "idn_1wKRXU2bsrzOxiVPJnUK2d757ex",
    "object": "email_address",
    "deleted": true
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

