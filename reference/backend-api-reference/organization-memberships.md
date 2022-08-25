---
description: >-
  Backend API endpoints to create, update and delete user memberships in
  organizations
---

# Organization Memberships

{% hint style="info" %}
Organizations is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Overview

Organization memberships related requests which allow you to create new user memberships on an organization, update existing memberships or delete them.

### Available requests

* `POST /v1/organizations/:organization_id/memberships`
* `GET /v1/organizations/:organization_id/memberships`
* `PATCH /v1/organizations/:organization_id/memberships/:user_id`
* `PATCH /v1/organizations/:organization_id/memberships/:user_id/metadata`
* `DELETE /v1/organizations/:organization_id/memberships/:user_id`

### Example organization membership schema

```json
{
  "object": "organization_membership",
  "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
  "public_metadata": {},
  "private_metadata": {},
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

{% swagger method="post" path="/v1/organizations/:organization_id/memberships" baseUrl="https://api.clerk.dev" summary="Create a new organization membership" %}
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
  "public_metadata": {},
  "private_metadata": {},
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

{% swagger method="get" path="/v1/organizations/:organization_id/memberships" baseUrl="https://api.clerk.dev" summary="Retrieve all memberships for organization" %}
{% swagger-description %}
Retrieves all user memberships for the given organization.
{% endswagger-description %}

{% swagger-parameter in="path" name="organization_id" type="string" required="true" %}
The id of the organization whose memberships we want to retrieve.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="query" name="limit" type="int" %}
The number of memberships to be returned.

By default, if not supplied, the response will contain 10 results.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="offset" type="int" %}
The number of entries to skip.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Organization membership returned successfully" %}
```javascript
{
    data: [{
        "object": "organization_membership",
        "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
        "public_metadata": {},
        "private_metadata": {},
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
    }]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="patch" path="/v1/organizations/:organization_id/memberships/:user_id" baseUrl="https://api.clerk.dev" summary="Update organization memberships" %}
{% swagger-description %}
Updates the properties of an existing organization membership.
{% endswagger-description %}

{% swagger-parameter in="path" name="organization_id" type="string" required="true" %}
The id of the organization of the membership.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="path" name="user_id" type="string" required="true" %}
The id of the user that this membership belongs to.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="role" type="string" %}
The new role of the given membership. Valid values for role are: 

`admin`

 and 

`basic_member`

.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Organization membership was updated successfully." %}
```javascript
{
  "object": "organization_membership",
  "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
  "public_metadata": {},
  "private_metadata": {},
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

{% swagger-response status="400: Bad Request" description="There should be at least one admin left after the update." %}
```javascript
{
  "errors": [
    {
      "code": "at_least_one_admin_needed",
      "long_message": "Cannot manage membership. There has to be at least one admin in the organization.",
      "message": "at least one admin needed"
    }
  ]
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

{% swagger method="patch" path="/v1/organizations/:organization_id/memberships/:user_id/metadata" baseUrl="https://api.clerk.dev" summary="Merge and update an organization membership's metadata" %}
{% swagger-description %}
Update an organization membership's metadata attributes by merging existing values with the provided parameters.&#x20;

Metadata values will be updated via a deep merge. Deep means that any nested JSON objects will be merged as well.

You can remove metadata keys at any level by setting their value to `null`.
{% endswagger-description %}

{% swagger-parameter in="path" name="organization_id" type="string" required="true" %}
The id of the organization of the membership.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="path" name="user_id" type="string" required="true" %}
The id of the user that this membership belongs to.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="public_metadata" type="object" %}
Metadata saved on the organization membership, that is visible to both your frontend and backend.

\


The new object will be merged with the existing value.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="private_metadata" type="object" %}
Metadata saved on the organization membership that is only visible to your backend.

\


The new object will be merged with the existing value.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Organization membership metadata were updated successfully." %}
```javascript
{
  "object": "organization_membership",
  "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
  "public_metadata": {
    "hello": "world"
  },
  "private_metadata": {
    "super": "secret"
  },
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

{% swagger-response status="404: Not Found" description="Organization membership not found." %}
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

{% swagger-response status="422: Unprocessable Entity" description="Invalid organization membership metadata." %}
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

{% swagger method="delete" path="/v1/organizations/:organization_id/memberships/:user_id" baseUrl="https://api.clerk.dev" summary="Delete organization membership" %}
{% swagger-description %}
Removes the given membership from the organization.
{% endswagger-description %}

{% swagger-parameter in="path" name="organization_id" type="string" required="true" %}
The id of the organization of the membership.
{% endswagger-parameter %}

{% swagger-parameter in="path" name="user_id" type="string" required="true" %}
The id of the user that this membership belongs to.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Membership was deleted successfully." %}
```javascript
{
  "object": "organization_membership",
  "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
  "public_metadata": {},
  "private_metadata": {},
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

{% swagger-response status="400: Bad Request" description="There should be at least one admin left after the update." %}
```javascript
{
  "errors": [
    {
      "code": "at_least_one_admin_needed",
      "long_message": "Cannot manage membership. There has to be at least one admin in the organization.",
      "message": "at least one admin needed"
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}
