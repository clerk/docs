# Phone numbers

## Available requests

* **`GET`**  `/v1/me/phone_numbers`
* **`POST`** `/v1/me/phone_numbers`
* **`GET`**  `/v1/me/phone_numbers/:id`
* **`PATCH`**`/v1/me/phone_numbers/:id`
* **`POST`** `/v1/me/phone_numbers/:id/send_verification_sms`
* **`POST`** `/v1/me/phone_numbers/:id/verify`
* **`DEL`**  `/v1/me/phone_numbers/:id`

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

{% api-method method="get" host="https://clerk.example.com" path="/v1/me/phone\_numbers" %}
{% api-method-summary %}
List all phone numbers
{% endapi-method-summary %}

{% api-method-description %}
Returns a list of the current user's phone numbers.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

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
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://clerk.example.com" path="/v1/me/phone\_numbers/:id" %}
{% api-method-summary %}
Retrieve a phone number
{% endapi-method-summary %}

{% api-method-description %}
Returns the current user's phone number with the given ID.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="id" type="string" required=true %}
ID of the phone number.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://clerk.example.com" path="/v1/me/phone\_numbers" %}
{% api-method-summary %}
Create a phone number
{% endapi-method-summary %}

{% api-method-description %}
Create a new phone number for the current user.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-form-data-parameters %}
{% api-method-parameter name="phone\_number" type="string" required=true %}
The phone number that you are adding to the current user.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://clerk.example.com" path="/v1/me/phone\_numbers/:id/send\_verification\_sms" %}
{% api-method-summary %}
Send a verification SMS
{% endapi-method-summary %}

{% api-method-description %}
Send a verification SMS to the phone number with the given ID.  Only valid for unverified phone numbers.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="id" type="string" required=true %}
ID of the phone number.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="post" host="https://clerk.example.com" path="/v1/me/phone\_numbers/:id/verify" %}
{% api-method-summary %}
Verify a phone number
{% endapi-method-summary %}

{% api-method-description %}
Verify a phone number with the given ID.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="id" type="string" required=true %}
ID of the phone number.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-form-data-parameters %}
{% api-method-parameter name="code" type="string" required=true %}
The 6 digit code that was sent via SMS.  You only get 3 chances to correctly verify, before a new code must be generated.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="patch" host="https://clerk.example.com" path="/v1/me/phone\_number/:id" %}
{% api-method-summary %}
Update a phone number
{% endapi-method-summary %}

{% api-method-description %}
Update the specified phone number.  Currently all you can do is turn on or off MFA using this phone number.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="id" type="string" required=true %}
ID of the phone number.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-form-data-parameters %}
{% api-method-parameter name="reserverd\_for\_second\_factor" type="boolean" required=false %}
Turn on MFA with the specified phone number.  If set to true, sign in will require a second step.
{% endapi-method-parameter %}
{% endapi-method-form-data-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="delete" host="https://clerk.example.com" path="/v1/me/phone\_number/:id" %}
{% api-method-summary %}
Delete a phone number
{% endapi-method-summary %}

{% api-method-description %}
Delete the current user's phone number.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="id" type="string" required=true %}
ID of the phone number.
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
// see the phone number object above.
{
  "id": "idn_1o4qfak5AdI2qlXSXENGL05iei6",
  "object": "phone_number",
  ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

