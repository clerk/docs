---
description: Frontend API endpoints to create organizations
---

# Organization requests

{% hint style="info" %}
Organizations is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Available requests

* `POST /v1/organizations`
* `PATCH /v1/organizations/:id`

## The organization object

```json
{
  "object": "organization",
  "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
  "name": "Acme Inc",
  "slug": "acme-inc",
  "public_metadata": {},
  "created_at": 1638000669544,
  "updated_at": 1638000669544
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/organizations" method="post" summary="Create an organization" %}
{% swagger-description %}
Creates a new organization with the provided `name`.&#x20;

The current user will become an administrator member of the organization, since they are the ones who created the organization.

You can optionally provide a `slug` for the organization. If a slug is provided, it should contain only lowercase alphanumeric characters (letters and digits) and the dash "-". Organization slugs must be unique.

You cannot create more than 100 organizations per account.
{% endswagger-description %}

{% swagger-parameter in="body" name="name" type="String" required="true" %}
The name of the new organization.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="slug" type="String" %}
An optional slug for the new organization. Must contain only lowercase alphanumeric and the dash. Must be unique for your instance.
{% endswagger-parameter %}

{% swagger-response status="200" description="Organization was created successfully." %}
```javascript
{
  "object": "organization",
  "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
  "name": "Acme Inc",
  "public_metadata": {},
  "slug": "acme-inc",
  "created_at": 1638000669544,
  "updated_at": 1638000669544
}
```
{% endswagger-response %}

{% swagger-response status="401: Unauthorized" description="Request not authenticated." %}
```javascript
{
  "errors": [
    {
      "code": "authentication_invalid",
      "long_message": "Unable to authenticate the request, you need to supply an active session",
      "message": "Invalid authentication"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions to create organizations." %}
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

{% swagger-response status="422: Unprocessable Entity" description="Invalid resource." %}
```javascript
{
  "errors": [
    {
      "code": "form_param_nil",
      "long_message": "Enter name.",
      "message": "Enter name.",
      "meta": {
        "param_name": "name"
      }
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/organizations/:id" method="patch" summary="Update an organization" %}
{% swagger-description %}
Updates the organization specified by `id`.

Only administrators can update an organization. Currently, only the organization `name` can be updated.
{% endswagger-description %}

{% swagger-parameter in="body" name="name" type="String" required="true" %}
The new name for the organization.
{% endswagger-parameter %}

{% swagger-response status="200" description="Organization was updated successfully." %}
```javascript
{
  "object": "organization",
  "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
  "name": "New Acme Inc",
  "public_metadata": {},
  "slug": "acme-inc",
  "created_at": 1638000669544,
  "updated_at": 1638000669544
}
```
{% endswagger-response %}

{% swagger-response status="401: Unauthorized" description="Request not authenticated." %}
```javascript
{
  "errors": [
    {
      "code": "authentication_invalid",
      "long_message": "Unable to authenticate the request, you need to supply an active session",
      "message": "Invalid authentication"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions to update organizations." %}
```javascript
{
  "errors": [
    {
      "code": "not_an_admin_in_organization",
      "long_message": "Current user is not an administrator in the organization. Only administrators can perform this action.",
      "message": "not an administrator"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="Invalid resource." %}
```javascript
{
  "errors": [
    {
      "code": "form_param_nil",
      "long_message": "Enter name.",
      "message": "Enter name.",
      "meta": {
        "param_name": "name"
      }
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}
