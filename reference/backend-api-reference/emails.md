# Emails

This object represents an email that has been sent to an email address.  You can only send an email to email addresses that have been verified by a user.  Only recommended for transactional emails.

### Available requests

* **`POST`**`/v1/emails`

### Example email schema

```javascript
{
    "id": "ema_1q93E6N25LCxLpHeR0NekHCsBiY",
    "object": "email",
    "from_email_name": "support",
    "to_email_address": "boss@clerk.dev",
    "email_address_id": "idn_1n5AVtYAHNQY73d6dY6GVJt8gmD",
    "subject": "green test",
    "body": "this is a green test",
    "status": "queued"
}
```

{% swagger baseUrl="https://api.clerk.dev" path="/v1/emails" method="post" summary="Create an email" %}
{% swagger-description %}
Create and send an email to the supplied email address ID.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="email_address_id" type="string" %}
The ID of a verified email address
{% endswagger-parameter %}

{% swagger-parameter in="body" name="from_email_name" type="string" %}
The name portion of the from email address.

\


e.g. "support" will send the email from support@[YOUR_DOMAIN].com
{% endswagger-parameter %}

{% swagger-parameter in="body" name="subject" type="string" %}
The subject of the email being sent.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="body" type="string" %}
The body of the email being sent.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
{
        "object": "email",
        "id": "ema_1q93E6N25LCxLpHeR0NekHCsBiY",
        ...
}
```
{% endswagger-response %}
{% endswagger %}
