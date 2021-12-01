# Email addresses

## Available requests

* **`GET`**` ``/v1/me/email_addresses`
* **`POST`**`/v1/me/email_addresses`
* **`GET`**` ``/v1/me/email_addresses/:id`
* **`POST`**`/v1/me/email_addresses/:id/send_verification_email`
* **`POST`**`/v1/me/email_addresses/:id/verify`
* **`DEL`**` ``/v1/me/email_addresses/:id`

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

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/email_addresses" method="get" summary="List all email addresses" %}
{% swagger-description %}
Returns a list of the current user's email addresses.
{% endswagger-description %}

{% swagger-response status="200" description="" %}
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
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/email_addresses/:id" method="get" summary="Retrieve an email address" %}
{% swagger-description %}
Returns the current user's email address with the given ID.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the email address.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the email address object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "email_address",
  ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/email_addresses" method="post" summary="Create an email address" %}
{% swagger-description %}
Create an email address for the current user.  This email address will need to be verified before it is active on the user.
{% endswagger-description %}

{% swagger-parameter in="body" name="email_address" type="string" %}
The email address that you are adding to the current user.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the email address object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "email_address",
  ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/email_addresses/:id/send_verification_email" method="post" summary="Send a verification email" %}
{% swagger-description %}
Send a verification email to the email address with the given ID.  Can only send to email addresses that have not been verified.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the email address.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the email address object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "email_address",
  ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/email_addresses/:id/verify" method="post" summary="Verify an email address" %}
{% swagger-description %}
Verify the current user's email address.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the email address.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="code" type="string" %}
The 6 digit code that was sent via email.  You only get 3 chances to correctly verify, before a new code must be generated.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the email address object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "email_address",
  ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/email_addresses/:id" method="delete" summary="Delete an email address" %}
{% swagger-description %}
Delete the current user's email address.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the email address.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
{
    "id": "idn_1wKRXU2bsrzOxiVPJnUK2d757ex",
    "object": "email_address",
    "deleted": true
}
```
{% endswagger-response %}
{% endswagger %}
