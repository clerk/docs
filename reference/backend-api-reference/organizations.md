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
* `PATCH /v1/organizations/:id`
* `DELETE /v1/organizations/:id`
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

{% swagger method="post" path="/v1/organizations/:id/memberships" baseUrl="https://api.clerk.dev" summary="Create a new organization membership" %}
{% swagger-description %}
Adds a user as a member to the given organization.

Only users in the same instance as the organization can be added as members.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="string" required="true" %}
The id of the organization that the new membership will be created.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="user_id" type="string" required="true" %}
The id of the user that will be added as a member in the organization.

The user needs to exist in the same instance as the organization and must not be a member of the given organization already.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="role" type="string" required="true" %}
The role that the new member will have in the organization.

Valid values for the role are `admin` and `basic_member`.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The new membership was created successfully." %}
```javascript
{
  "object": "organization_membership",
  "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
  "role": "basic_member",
  "organization": {
    "object": "organization",
    "id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
    "logo_url": "https://images.clerk.services/default-logo.png",
    "name": "Acme Inc",
    "private_metadata": {},
    "public_metadata": {},
    "slug": "acme-inc",
    "created_at": 1638000669544,
    "updated_at": 1638000669544
  },
  "created_at": 1638000669544,
  "updated_at": 1638000669544,
  "public_user_data": {
    "first_name": "Sarah",
    "last_name": "Connor",
    "profile_image_url": "https://images.clerk.dev/uploaded/img_jlkkcq2786n0.jpeg",
    "identifier": "sarah@connor.com",
    "user_id": "user_1o4q123qMeCkKKIXcA9h8"
  }
}
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="User is already a member of the organization." %}
```javascript
{
    "errors": [{
      "code": "already_a_member_in_organization",
      "long_message": "User your_id is already a member of the organization.",
      "message": "already a member"
    }]
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="No user was found with given user id." %}
```javascript
{
    "errors": [{
      "code": "resource_not_found",
      "long_message": "No user was found with id your_id",
      "message": "not found"
    }]
}
```
{% endswagger-response %}

{% swagger-response status="422: Unprocessable Entity" description="Invalid value for role." %}
```javascript
{
    "errors": [{
      "code": "form_param_value_invalid",
      "long_message": "your_value does not match the allowed values for parameter role. You can use one of the following: admin or basic_member.",
      "message": "is invalid",
      "meta": {
        "param_name": "role"
      }
    }]
}
```
{% endswagger-response %}
{% endswagger %}
