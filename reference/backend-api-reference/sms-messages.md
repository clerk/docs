# SMS Messages

This object represents an SMS message that has been sent to a phone number.  You can only send an SMS message to phone numbers that have been verified by a user.  Only recommended for transactional SMS messages.

### Available requests

* **`POST`**`/v1/sms_messages`

### Example SMS message schema

```javascript
{
    "object": "sms_message",
    "id": "sms_1q8uR4AxV2G8l6dgbtuwh4Gorws",
    "from_phone_number": "9738321914",
    "to_phone_number": "+12019613198",
    "phone_number_id": "idn_1oDEYpGZ9m7XbIyqaML55jEaETE",
    "message": "Welcome to Clerk!",
    "status": "queued"
}
```

{% api-method method="post" host="https://api.clerk.dev" path="/v1/sms\_messages" %}
{% api-method-summary %}
Create an SMS message
{% endapi-method-summary %}

{% api-method-description %}
Create and send an SMS message to the supplied phone number ID.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-form-data-parameters %}
{% api-method-parameter name="phone\_number\_id" type="string" required=true %}
The ID of a verified phone number.
{% endapi-method-parameter %}

{% api-method-parameter name="message" type="string" required=true %}
The message you would like to send.
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
    "object": "sms_message",
    "id": "sms_1q8uR4AxV2G8l6dgbtuwh4Gorws"
    ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

