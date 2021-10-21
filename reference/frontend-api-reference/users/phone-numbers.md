# Phone numbers

## Available requests

* **`GET  `**`/v1/me/phone_numbers`
* **`POST `**`/v1/me/phone_numbers`
* **`GET  `**`/v1/me/phone_numbers/:id`
* **`PATCH`**`/v1/me/phone_numbers/:id`
* **`POST `**`/v1/me/phone_numbers/:id/send_verification_sms`
* **`POST `**`/v1/me/phone_numbers/:id/verify`
* **`DEL  `**`/v1/me/phone_numbers/:id`

## The phone number object

```javascript
{
    "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
    "object": "phone_number",
    "email_address": "+12015555555",
    "verification": {
        "status": "verified",
        "strategy": "phone_code"
    },
    "linked_to": []
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/phone_numbers" method="get" summary="List all phone numbers" %}
{% swagger-description %}
Returns a list of the current user's phone numbers.
{% endswagger-description %}

{% swagger-response status="200" description="" %}
```javascript
// see the phone number object above.
[
  {
    "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
    "object": "phone_number",
    ...
  },
  ...
]
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/phone_numbers/:id" method="get" summary="Retrieve a phone number" %}
{% swagger-description %}
Returns the current user's phone number with the given ID.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the phone number.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/phone_numbers" method="post" summary="Create a phone number" %}
{% swagger-description %}
Create a new phone number for the current user.
{% endswagger-description %}

{% swagger-parameter in="body" name="phone_number" type="string" %}
The phone number that you are adding to the current user.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/phone_numbers/:id/send_verification_sms" method="post" summary="Send a verification SMS" %}
{% swagger-description %}
Send a verification SMS to the phone number with the given ID.  Only valid for unverified phone numbers.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the phone number.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/phone_numbers/:id/verify" method="post" summary="Verify a phone number" %}
{% swagger-description %}
Verify a phone number with the given ID.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the phone number.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="code" type="string" %}
The 6 digit code that was sent via SMS.  You only get 3 chances to correctly verify, before a new code must be generated.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/phone_number/:id" method="patch" summary="Update a phone number" %}
{% swagger-description %}
Update the specified phone number.  Currently all you can do is turn on or off MFA using this phone number.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the phone number.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="reserverd_for_second_factor" type="boolean" %}
Turn on MFA with the specified phone number.  If set to true, sign in will require a second step.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/me/phone_number/:id" method="delete" summary="Delete a phone number" %}
{% swagger-description %}
Delete the current user's phone number.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" %}
ID of the phone number.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endswagger-response %}
{% endswagger %}
