---
description: Manage organizations via the Backend API
---

# Organizations

{% hint style="info" %}
Organizations is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Overview

Organizations related requests allow you to create new organizations for your instance.&#x20;

### Available requests

* `POST /v1/organizations`
* `DELETE /v1/organizations/:id`

### Example organization schema

```json
{
    "object": "organization",
    "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
    "name": "Acme Inc",
    "privateMetadata": {},
    "publicMetadata": {},
    "slug": "acme-inc",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}
```

{% swagger method="post" path="/v1/organizations" baseUrl="https://api.clerk.dev" summary="Create an organization" %}
{% swagger-description %}
Creates a new organization with the given name for an instance.&#x20;

In order to successfully create an organization you need to provide the ID of the [User](users.md) who will become the organization administrator.

You can specify an optional slug for the new organization. If provided, the organization slug can contain only lowercase alphanumeric characters (letters and digits) and the dash "-". Organization slugs must be unique for the instance.

You can provide additional metadata for the organization and set any custom attribute you want. Organizations support private and public metadata. Private metadata can only be accessed from the Backend API. Public metadata can be accessed from the Backend API, and are read-only from the Frontend API.
{% endswagger-description %}

{% swagger-parameter in="body" name="name" type="string" required="true" %}
The name of the new organization.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="created_by" type="string" required="true" %}
The ID of the User who will become the new organization's administrator.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="private_metadata" type="object" %}
Metadata saved on the organization, read-only from the Frontend API and fully accessible (read/write) from the Backend API.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="public_metadata" type="object" %}
Metadata saved on the organization, accessible only from the Backend API.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="slug" type="string" %}
A slug for the new organization. Can contain only lowercase alphanumeric characters and the dash "-". Must be unique for the instance.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The new organization was created." %}
```javascript
{
    "object": "organization",
    "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
    "name": "Acme Inc",
    "private_metadata": {},
    "public_metadata": {},
    "slug": "acme-inc",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="The organization creator cannot be found." %}
```json
{
  "errors": [
    {
      "code": "organization_creator_not_found",
      "long_message": "No users found with id ...",
      "message": "creator not found"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions to create an organization." %}
```javascript
{
  "errors": [
    {
      "code": "organizations_not_enabled_in_instance",
      "long_message": "The organizations feature is not enabled for your instance. If you want to try it out, contact us at support@clerk.dev.",
      "message": "access denied"
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="delete" path="/v1/organizations/:id" baseUrl="https://api.clerk.dev" summary="Delete an organization" %}
{% swagger-description %}
Deletes the given organization.

Please note that deleting an organization will also delete all memberships and invitations. This is not reversible.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}
The id of the organization to be deleted.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Organization was deleted successfully" %}
```javascript
{
  "object": "organization",
  "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
  "deleted": true
}
```
{% endswagger-response %}
{% endswagger %}
