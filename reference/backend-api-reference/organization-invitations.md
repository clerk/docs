---
description: >-
  Backend API endpoints to invite users to an organization and manage
  invitations
---

# Organization Invitations

{% hint style="info" %}
Organizations is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Overview

Organization invitations related requests which allow you to invite new users to an organization and manage invitations.

### Available requests

* `POST /v1/organizations/:id/invitations`
* `GET /v1/organizations/:id/invitations/pending`
* `POST /v1/organizations/:organization_id/invitations/:invitation_id/revoke`

### Example organization invitation schema

```json
{
  "object": "organization_invitation",
  "id": "orginv_21Ufcy98STcA11s3QckIwtwHIES",
  "email_address": "invited@example.com",
  "role": "basic_member",
  "organization_id": "org_A1Ufcy88STc111s3QckIwtw956e",
  "status": "pending",
  "created_at": 1638000669544,
  "updated_at": 1638000669544
}
```

{% swagger method="post" path="/v1/organizations/:id/invitations" baseUrl="https://api.clerk.dev" summary="Create and send an organization invitation" %}
{% swagger-description %}
Creates a new organization invitation and sends an email to the provided `email_address` with a link to accept the invitation and join the organization.

You can specify the `role` for the invited organization member.&#x20;

New organization invitations get a "pending" status until they are revoked by an organization administrator or accepted by the invitee.

The request body supports passing an optional `redirect_url` parameter. When the invited user clicks the link to accept the invitation, they will be redirected to the URL you provided. Use this parameter to implement a custom invitation acceptance flow.&#x20;

You must specify the ID of the user that will send the invitation with the `inviter_user_id` parameter. That user must be a member with administrator privileges in the organization. Only "admin" members can create organization invitations.
{% endswagger-description %}

{% swagger-parameter in="body" name="email_address" type="string" required="true" %}
The email address of the new member that's going to be invited to the organization.
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="role" type="string" required="true" %}
The role of the new member in the organization.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="inviter_user_id" type="string" required="true" %}
The ID of the user that invites the new member to the organization. Must be an administrator in the organization.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="redirect_url" type="object" %}
Optional URL that the invitee will be redirected to once they accept the invitation by clicking the join link in the invitation email.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The new organization invitation was created." %}
```javascript
{
  "object": "organization_invitation",
  "id": "orginv_21Ufcy98STcA11s3QckIwtwHIES",
  "email_address": "invitee@example.com",
  "role": "basic_member",
  "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
  "status": "pending",
  "created_at": 1638000669544,
  "updated_at": 1638000669544
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions to create an organization invitation." %}
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

{% swagger-response status="422: Unprocessable Entity" description="Invalid resource" %}
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

{% swagger method="get" path="/v1/organizations/:id/invitations/pending" baseUrl="https://api.clerk.dev" summary="Get a list of pending organization invitations" %}
{% swagger-description %}
This request returns the list of organization invitations with "pending" status. This means that the organization invitations can still be used to join the organization, but have not been accepted by the invited user.

Results can be paginated using the optional `limit` and `offset` query parameters. The organization invitations are ordered by descending creation date. Most recent invitations will be returned first.
{% endswagger-description %}

{% swagger-parameter in="body" name="limit" type="number" required="false" %}
Applies a limit to the number of organization invitations returned. Can be used for paginating the results together with 

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

 to skip the last retrieved organization invitations.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="Pending organization invitations were retrieved successfully." %}
```javascript
{
  "data": [
    {
      "object": "organization_invitation",
      "id": "orginv_21Ufcy98STcA11s3QckIwtwHIES",
      "email_address": "invitee@example.com",
      "role": "basic_member",
      "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
      "status": "pending",
      "created_at": 1638000669544,
      "updated_at": 1638000669544
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

{% swagger-response status="422: Unprocessable Entity" description="Invalid pagination parameters" %}
```javascript
{
  {
    "errors": [
      {
        "code": "form_param_value_invalid",
        "long_message": "limit is invalid",
        "message": "is invalid",
        "meta": {
          "param_name": "limit"
        }
      }
    ]
  }
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger method="post" path="/v1/organizations/:organization_id/invitations/:invitation_id/revoke" baseUrl="https://api.clerk.dev" summary="Revoke a pending organization invitation" %}
{% swagger-description %}
Use this request to revoke a previously issued organization invitation.&#x20;

Revoking an organization invitation makes it invalid; the invited user will no longer be able to join the organization with the revoked invitation.

Only organization invitations with "pending" status can be revoked.&#x20;

The request needs the `requesting_user_id` parameter to specify the user which revokes the invitation. Only users with "admin" role can revoke invitations.
{% endswagger-description %}

{% swagger-parameter in="header" name="Authorization" type="string" required="true" %}
Bearer [YOUR_API_KEY]
{% endswagger-parameter %}

{% swagger-parameter in="body" name="requesting_user_id" type="string" required="true" %}
The ID of the user that revokes the invitation. Must be an administrator in the organization.
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="The organization invitation was revoked." %}
```javascript
{
  "object": "organization_invitation",
  "id": "orginv_21Ufcy98STcA11s3QckIwtwHIES",
  "email_address": "invitee@example.com",
  "role": "basic_member",
  "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
  "status": "revoked",
  "created_at": 1638000669544,
  "updated_at": 1638000669544
}
```
{% endswagger-response %}

{% swagger-response status="403: Forbidden" description="Insufficient permissions to revoke an organization invitation." %}
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
