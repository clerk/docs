---
description: Learn how to handle invitations via the backend API
---

# Invitations

## Overview

This resource provides the necessary methods to create, revoke and list your invitations. For more information on how invitations work, refer to our [Invitations](broken-reference) guide.

### Available requests

* **`POST`**`/v1/invitations`
* **`GET`**`/v1/invitations`
* **`POST`**`/v1/invitations/:id/revoke`

### Example invitation schema

```javascript
{
        "object": "invitation",
        "id": "inv_21Ufcy98STcA11s3QckIwtwHIES",
        "email_address": "email@example.com",
        "revoked": true
        "created_at": 1638000669544,
        "updated_at": 1638000669544,
}
```

{% swagger method="post" path="/v1/invitations" baseUrl="https://api.clerk.dev" summary="Create an invitation" %}
{% swagger-description %}
Creates a new invitation for the given email address and sends the invitation email.&#x20;

Keep in mind that you cannot create an invitation if there is already one for the given email address. Also, trying to create an invitation for an email address that already exists in your application will result to an error.
{% endswagger-description %}

{% swagger-parameter in="body" name="email_address" type="string" required="true" %}
The email address the invitation will be sent to.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="public_metadata" type="object" %}
Metadata that will be attached to the newly created invitation. The value of this property should be a well-formed JSON object.

Once the user accepts the invitation and signs up, these metadata will end up in the user's public metadata.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="redirect_url" type="string" %}
Optional URL which specifies where to redirect the user once they click the invitation link.

This is only required if you have implemented a [custom flow](broken-reference) and you're not using Clerk Hosted Pages or Clerk Components.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The new invitation was created." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="Something went wrong and the invitation was not created." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="Invalid resource" %}
```javascript
{
  "errors": [
    {
      "code": "form_param_exceeds_allowed_size",
      "long_message": "The given public_metadata exceeds the maximum allowed size of 4096 bytes (4 KB).",
      "message": "The given public_metadata exceeds the maximum allowed size of 4096 bytes (4 KB).",
      "meta": {
        "param_name": "public_metadata"
      }
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="get" path="/v1/invitations" baseUrl="https://api.clerk.dev" summary="List all invitations" %}
{% swagger-description %}
Returns all non-revoked invitations for your application, sorted by creation date.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorizaton" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="A list of all non-revoked invitations." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="post" path="/v1/invitations/:id/revoke" baseUrl="https://api.clerk.dev" summary="Revokes an invitation" %}
{% swagger-description %}
Revokes the given invitation. Revoking an invitation will prevent the user to use the invitation link that was sent to him. However, it doesn't prevent the user from signing up if they follow the sign up flow.

Only active (i.e. non-revoked) invitations can be revoked.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}
The id of the invitation to be revoked.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The invitation was successfully revoked." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="Something went wrong and the invitation has not been revoked." %}
```javascript
{
    // Response
}
```
{% endswagger-response %}
{% endswagger %}
