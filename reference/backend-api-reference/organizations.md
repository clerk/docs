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

* `GET /v1/organizations`
* `POST /v1/organizations`
* `GET /v1/organizations/:id_or_slug`
* `DELETE /v1/organizations/:id`
* `PATCH /v1/organizations/:id`
* `PATCH /v1/organizations/:id/metadata`
* `PUT /v1/organizations/:id/logo`
* `POST /v1/organizations/:id/memberships`

### Example organization schema

```json
{
    "object": "organization",
    "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
    "logo_url": "https://images.clerk.services/default-logo.png",
    "name": "Acme Inc",
    "privateMetadata": {},
    "publicMetadata": {},
    "slug": "acme-inc",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}
```

{% swagger method="get" path="/v1/organizations" baseUrl="https://api.clerk.dev" summary="Get a list of organizations for an instance" %}
{% swagger-description %}
This request returns the list of organizations for an instance. The instance is specified by the API key provided in the `Authorization` header.

Results can be paginated using the optional `limit` and `offset` query parameters. The organizations are ordered by descending creation date. Most recent organizations will be returned first.
{% endswagger-description %}

{% swagger-parameter in="body" name="limit" type="number" required="false" %}
Applies a limit to the number of organizations returned. Can be used for paginating the results together with 

`offset`

. Must be an integer greater than zero.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="offset" type="number" required="false" %}
Skip the first 

`offset`

 results when paginating. Needs to be an integer greater than zero. Use it together with 

`limit`

 to skip the last retrieved organizations.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Organizations were retrieved successfully." %}
```javascript
{
  "data": [
    {
      "object": "organization",
      "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
      "logo_url": "https://images.clerk.services/default-logo.png",
      "name": "Acme Inc",
      "privateMetadata": {},
      "publicMetadata": {},
      "slug": "acme-inc",
      "created_at": 1638000669544,
      "updated_at": 1638000669544
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Organizations are not enabled for the instance." %}
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

{% swagger-response status="422: Unprocessable Entity" description="Invalid pagination parameters." %}
```javascript
{
  {
    "errors": [
      {
        "code": "form_param_value_invalid",
        "long_message": "offset is invalid",
        "message": "is invalid",
        "meta": {
          "param_name": "offset"
        }
      }
    ]
  }
}
```
{% endswagger-response %}
{% endswagger %}

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
    "logo_url": null,
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

{% swagger method="get" path="/v1/organizations/:id_or_slug" baseUrl="https://api.clerk.dev" summary="Retrieve an organization by ID or slug" %}
{% swagger-description %}
Fetches the organization whose ID or slug matches the provided `id_or_slug` URL query parameter.

The organization must belong to the instance whose API key is used for authentication in the `Authorization` request header.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="path" name="id_or_slug" type="string" required="true" %}
The id or slug of the organization.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Organization was retrieved successfully" %}
```javascript
{
    "object": "organization",
    "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
    "logo_url": null,
    "name": "Acme Inc",
    "private_metadata": {},
    "public_metadata": {},
    "slug": "acme-inc",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Organizations are not enabled for the instance" %}
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

{% swagger-response status="404: Not Found" description="Resource not found" %}
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

{% swagger-response status="403: Forbidden" description="Organizations are not enabled for the instance" %}
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

{% swagger-response status="404: Not Found" description="Resource not found" %}
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

{% swagger method="patch" path="/v1/organizations/:id" baseUrl="https://api.clerk.dev" summary="Update an organization" %}
{% swagger-description %}
Updates an existing organization.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}
The id of the organization to be updated.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="name" type="string" %}
The new name of the organization.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The organization was updated successfully." %}
```javascript
{
    "object": "organization",
    "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
    "logo_url": null,
    "name": "Acme Inc",
    "private_metadata": {},
    "public_metadata": {},
    "slug": "acme-inc",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Organizations are not enabled for the instance." %}
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

{% swagger-response status="404: Not Found" description="Resource not found" %}
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

{% swagger baseUrl="https://api.clerk.dev" path="/v1/organizations/:id/metadata" method="patch" summary="Merge and update an organization's metadata" %}
{% swagger-description %}
Update an organization's metadata attributes by merging existing values with the provided parameters.&#x20;

Metadata values will be updated via a deep merge. Deep means that any nested JSON objects will be merged as well.

You can remove metadata keys at any level by setting their value to `null`.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="public_metadata" type="object" %}
Metadata saved on the organization, that is visible to both your frontend and backend.

\


The new object will be merged with the existing value.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="private_metadata" type="object" %}
Metadata saved on the organization that is only visible to your backend.

\


The new object will be merged with the existing value.
{% endswagger-parameter %}

{% swagger-response status="200" description="Organization metadata successfully updated." %}
```json
{
    "object": "organizantion",
    "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
    "logo_url": "https://images.clerk.services/default-logo.png",
    "name": "Acme Inc",
    "privateMetadata": { "private": "metadata" },
    "publicMetadata": { "public": "metadata" },
    "slug": "acme-inc",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Organizations not enabled for the instance." %}
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

{% swagger-response status="404: Not Found" description="Organization not found on the instance." %}
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

You must also provide the uploader user's ID. The uploader user must be an organization administrator.
{% endswagger-description %}

{% swagger-parameter in="body" name="file" type="file" required="true" %}
The multipart image file that will be used as the organization's logo. Cannot exceed 10MB.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Content-Type" type="string" required="true" %}
multipart/form-data
{% endswagger-parameter %}

{% swagger-parameter in="body" name="uploader_user_id" type="string" required="true" %}
The ID of the User who will be the logo uploader. Must be an organization administrator.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The organization logo was successfully uploaded." %}
```javascript
{
    "object": "organization",
    "id": "org_21Ufcy98STcA11s3QckIwtwHIES",
    "logo_url": "https://images.clerk.services/default-logo.png",
    "name": "Acme Inc",
    "private_metadata": {},
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
