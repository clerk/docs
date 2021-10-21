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

{% swagger baseUrl="https://api.clerk.dev" path="/v1/sms_messages" method="post" summary="Create an SMS message" %}
{% swagger-description %}
Create and send an SMS message to the supplied phone number ID.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="phone_number_id" type="string" %}
The ID of a verified phone number.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="message" type="string" %}
The message you would like to send.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
    "object": "sms_message",
    "id": "sms_1q8uR4AxV2G8l6dgbtuwh4Gorws"
    ...
}
```
{% endswagger-response %}
{% endswagger %}
