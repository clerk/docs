---
description: Frontend API endpoints to manage organization memberships
---

# Organization membership requests

{% hint style="info" %}
Organizations is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Available requests

* `GET /v1/organizations/:organization_id/memberships`
* `PATCH /v1/organizations/:organization_id/memberships/:user_id`
* `DELETE /v1/organizations/:organization_id/memberships/:user_id`

## The organization membership object

```json5
{
  "object": "organization_membership",
  "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
  "role": "basic_member",
  "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
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

{% swagger baseUrl="https://clerk.example.com" path="/v1/organizations/:organization_id/memberships" method="get" summary="Retrieve a list of organization memberships" %}
{% swagger-description %}
This request returns a list of memberships for a given organization.

Each membership provides information about the organization and the member user's public data, along with the membership role.

The authenticated user must be an organization member.&#x20;

Results can be paginated using the optional `limit` and `offset` query parameters. The organization memberships are ordered by descending creation date. Most recent members will be returned first.
{% endswagger-description %}

{% swagger-parameter in="query" name="limit" type="Number" %}
Applies a limit to the number of organization memberships returned. Can be used for paginating the results together with 

`offset`

. Must be an integer greater than zero.
{% endswagger-parameter %}

{% swagger-parameter in="query" name="offset" type="Number" %}
Skip the first 

`offset`

 results when paginating. Needs to be an integer greater than zero. Use it together with 

`limit`

 to skip the last retrieved organization invitations.
{% endswagger-parameter %}

{% swagger-response status="200" description="Organization memberships were retrieved successfully." %}
```javascript
[
  {
    "object": "organization_membership",
    "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
    "role": "basic_member",
    "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
    "created_at": 1638000669544,
    "updated_at": 1638000669544,
    "public_user_data": {
      "first_name": "Sarah",
      "last_name": "Connor",
      "profile_image_url": "https://images.clerk.dev/uploaded/img_jlkkcq2786n0.jpeg",
      "identifier": "sarah@connor.com",
      "user_id": "user_1o4q123qMeCkKKIXcA9h8"
    }
  },
  {
    "object": "organization_membership",
    "id": "orgmem_21Ufcy98STcA11s3QckIwtwJJy6",
    "role": "admin",
    "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
    "created_at": 1438000669544,
    "updated_at": 1438000669544,
    "public_user_data": {
      "first_name": "John",
      "last_name": "Connor",
      "profile_image_url": "https://images.clerk.dev/uploaded/img_kkjkcq2789iu.jpeg",
      "identifier": "john@connor.com",
      "user_id": "user_1oi8U23qM99OKKIXcAft5"
    }
  }
]
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

{% swagger-response status="403: Forbidden" description="Insufficient permissions to list organization memberships." %}
```javascript
{
  "errors": [
    {
      "code": "not_a_member_in_organization",
      "long_message": "Current user is not a member of the organization. Only organization members can perform this action.",
      "message": "not a member"
    }
  ]
}
```
{% endswagger-response %}

{% swagger-response status="404: Not Found" description="Resource not found." %}
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

{% swagger baseUrl="https://clerk.example.com" path="/v1/organizations/:organization_id/memberships" method="patch" summary="Update an organization membership" %}
{% swagger-description %}
This request updates an organization membership.&#x20;

Use it to update a member's role in the organization. You can also update the current user's role in the organization.

The authenticated user must have administrator privileges in the organization. Only members with "admin" role can update other organization members roles.

When updating a member's role we check that at least one member with administrator privileges will remain in the organization. You cannot change the last "admin" member's role.
{% endswagger-description %}

{% swagger-parameter in="body" name="role" type="String" %}
The organization membership role. Can be one of "admin" or "basic_member".
{% endswagger-parameter %}

{% swagger-response status="200" description="Organization membership updated successfully." %}
```javascript
{
  "object": "organization_membership",
  "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
  "role": "admin",
  "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
  "created_at": 1638000669544,
  "updated_at": 1734200669544,
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

{% swagger-response status="400: Bad Request" description="Request cannot be fulfilled." %}
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

{% swagger-response status="403: Forbidden" description="Insufficient permissions to update organization membership." %}
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

{% swagger-response status="404: Not Found" description="Resource not found." %}
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

{% swagger-response status="422: Unprocessable Entity" description="Invalid resource." %}
```javascript
{
  "errors": [
    {
      "code": "form_param_value_invalid",
      "long_message": "... does not match the allowed values for parameter role. You can use one of the following: admin or basic_member",
      "message": "is invalid",
      "meta": {
        "param_name": "role"
      }
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/organizations/:organization_id/memberships" method="delete" summary="Delete an organization membership" %}
{% swagger-description %}
This request deletes an organization membership.&#x20;

Use it to remove a user from the organization.&#x20;

The authenticated user must have administrator privileges in the organization. Only members with "admin" role can remove other organization members.

When removing a member we check that at least one other member with administrator privileges will remain in the organization. You cannot remove the last "admin" member.
{% endswagger-description %}

{% swagger-response status="200" description="Organization membership deleted successfully." %}
```javascript
{
  "object": "organization_membership",
  "id": "orgmem_21Ufcy98STcA11s3QckIwtwHIES",
  "role": "basic_member",
  "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
  "created_at": 1638000669544,
  "updated_at": 1734200669544,
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

{% swagger-response status="400: Bad Request" description="Request cannot be fulfilled." %}
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

{% swagger-response status="403: Forbidden" description="Insufficient permissions to delete the organization membership." %}
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

{% swagger-response status="404: Not Found" description="Resource not found." %}
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
