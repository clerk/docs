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

{% api-method method="post" host="https://api.clerk.dev" path="/v1/emails" %}
{% api-method-summary %}
Create an email
{% endapi-method-summary %}

{% api-method-description %}
Create and send an email to the supplied email address ID.
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Authorization" type="string" required=true %}
Bearer \[YOUR\_API\_KEY\]
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-form-data-parameters %}
{% api-method-parameter name="email\_address\_id" type="string" required=true %}
The ID of a verified email address
{% endapi-method-parameter %}

{% api-method-parameter name="from\_email\_name" type="string" required=true %}
The name portion of the from email address.  
e.g. "support" will send the email from support@\[YOUR\_DOMAIN\].com
{% endapi-method-parameter %}

{% api-method-parameter name="subject" type="string" required=true %}
The subject of the email being sent.
{% endapi-method-parameter %}

{% api-method-parameter name="body" type="string" required=true %}
The body of the email being sent.
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
        "object": "email",
        "id": "ema_1q93E6N25LCxLpHeR0NekHCsBiY",
        ...
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

