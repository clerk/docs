---
description: Frontend API endpoints to create organizations
---

# Organization requests

## Available requests

* `POST /v1/organizations`
* `PATCH /v1/organizations/:id`
* `DELETE /v1/organizations/:id`
* `DELETE /v1/organizations/:id/logo`
* `PUT /v1/organizations/:id/logo`

## The organization object

```json
{
  "object": "organization",
  "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
  "logo_url": "https://images.clerk.services/default-logo.png",
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

{% swagger method="delete" path="/v1/organizations/:id" baseUrl="https://clerk.example.com" summary="Delete an organization" %}
{% swagger-description %}
Deletes the organization specified by `id`.

Only administrators can delete an organization.
{% endswagger-description %}

{% swagger-response status="200: OK" description="Organization was deleted successfully." %}
```javascript
{
  "object": "organization",
  "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
  "deleted": true
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions to delete the organization." %}
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
{% endswagger %}

{% swagger method="delete" path="/v1/organizations/:id/logo" baseUrl="https://api.clerk.dev" summary="Delete an organization's logo" %}
{% swagger-description %}
Remove an organization's logo. The current user must be an organization administrator.
{% endswagger-description %}

{% swagger-response status="200: OK" description="The organization logo was successfully deleted." %}
```javascript
{
    "object": "image",
    "id": "img_21Ufcy98STcA11s3QckIwtwHIES",
    "slug": "",
    "deleted": true
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions to delete an organization logo." %}
```javascript
{
  "errors": [
    {
      "code": "not_an_admin_in_organization",
      "long_message":  "Current user is not an administrator in the organization. Only administrators can perform this action.",
      "message": "not an administrator"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="The organization specified by ID cannot be found or does not have a logo." %}
```javascript
{
  "errors": [
    {
      "code": "resource_not_found",
      "long_message": "Resource not found",
      "message": "not found"
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="put" path="/v1/organizations/:id/logo" baseUrl="https://api.clerk.dev" summary="Update an organization's logo" %}
{% swagger-description %}
Set or replace an organization's logo, by uploading an image file.&#x20;

This endpoint uses the `multipart/form-data` request content type and accepts a file of image type. The file size cannot exceed 10MB. Only the following file content types are supported: `image/jpeg, image/png, image/gif, image/webp, image/x-icon, image/vnd.microsoft.icon`.

The current user will be the logo uploader and must be an organization administrator.
{% endswagger-description %}

{% swagger-parameter in="body" name="file" type="file" required="true" %}
The multipart image file that will be used as the organization's logo. Cannot exceed 10MB.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Content-Type" type="string" required="true" %}
multipart/form-data
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The organization logo was successfully uploaded." %}
```javascript
{
    "object": "organization",
    "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
    "logo_url": "https://images.clerk.services/default-logo.png",
    "name": "Acme Inc",
    "public_metadata": {},
    "slug": "acme-inc",
    "created_at": 1638000669544,
    "updated_at": 1738011669221
}
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="The request body is invalid, file cannot be uploaded." %}
```json
{
  "errors": [
    {
      "code": "request_body_invalid",
      "message": "Request body invalid"
    },
    {
      "code": "form_param_missing",
      "long_message": "There was no image file present in the request",
      "message": "Image file missing"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions to upload an organization logo." %}
```javascript
{
  "errors": [
    {
      "code": "not_an_admin_in_organization",
      "long_message":  "Current user is not an administrator in the organization. Only administrators can perform this action.",
      "message": "not an administrator"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="The organization specified by ID cannot be found." %}
```javascript
{
  "errors": [
    {
      "code": "resource_not_found",
      "long_message": "Resource not found",
      "message": "not found"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="413: Payload Too Large" description="The supplied file exceeds the size limit" %}
```javascript
{
  "errors": [
    {
      "code": "image_too_large",
      "long_message":  "The image being uploaded is more than 10MB. Please choose a smaller one.",
      "message": "Image too large"
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}
