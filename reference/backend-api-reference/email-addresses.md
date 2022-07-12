# Email Addresses

This object represents an email address in your instance.

### Available requests

* **`POST`**` ``/v1/email_addresses`
* **`GET`**`  ``/v1/email_addresses/:id`
* **`PATCH`**`/v1/email_addresses/:id`
* **`DEL`**`  ``/v1/email_addresses/:id`

### Example email address schema

```javascript
{
    "id": "idn_2AOf58YAKamiQEdeQjzSVLRRiXx",
    "object": "email_address",
    "email_address": "example@clerk.dev",
    "reserved": false,
    "verification": {
        "status": "verified",
        "strategy": "email_link",
        "expire_at": 1654885965363,
        "verified_at_client": "client_2ACtvuO8Nwjpf33zagezgdpaQy2"
    },
    "linked_to": []
}
```



{% swagger baseUrl="https://api.clerk.dev" path="/v1/email_addresses" method="post" summary="Create an email address" %}
{% swagger-description %}
Create an email address for a user.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="user_id" type="string" required="true" %}
The ID representing the user
{% endswagger-parameter %}

{% swagger-parameter in="body" name="email_address" type="string" required="true" %}
The new email address.  Must adhere 
{% endswagger-parameter %}

{% swagger-parameter in="body" name="verified" type="bool" %}
When created, the email address will be marked verified.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="primary" type="bool" %}
Create this email address as the primary email address for the user.&#x20;



Default: false, unless it is the first email address.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// See example schema
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/email_addresses/:id" method="get" summary="Retrieve an email address" %}
{% swagger-description %}
Retrieve the details of an email address. 
{% endswagger-description %}

{% swagger-parameter in="header" name="Authentication" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// see example schema
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/email_addresses/:id" method="patch" summary="Update an email address" %}
{% swagger-description %}
Update an email address.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="verified" type="bool" %}
Mark the email address as verified.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="primary" type="bool" %}
Mark the email address as primary for the user.
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
// See example schema
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://api.clerk.dev" path="/v1/email_addresses/:id" method="delete" summary="Delete an email address" %}
{% swagger-description %}
Remove an email address from the user.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200" description="User metadata successfully updated." %}
```
// see example schema
```
{% endswagger-response %}
{% endswagger %}
