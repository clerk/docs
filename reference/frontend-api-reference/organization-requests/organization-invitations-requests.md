---
description: >-
  Frontend API endpoints to invite users to an organization and manage the
  invitations
---

# Organization invitations requests

{% hint style="info" %}
Organizations is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Available requests

* `POST /v1/organizations/:organization_id/invitations`
* `POST /v1/organizations/:organization_id/invitations/:invitation_id/revoke`
* `GET /v1/organizations/:organization_id/invitations/pending`

## The organization invitation object

```json
{
  "object": "organization_invitation",
  "id": "orginv_21Ufcy98STcA11s3QckIwtwHIES",
  "email_address": "invitee@example.com",
  "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
  "public_metadata": {},
  "role": "basic_member",
  "status": "pending",
  "created_at": 1638000669544,
  "updated_at": 1638000669544
}
```

{% swagger baseUrl="https://clerk.example.com" path="/v1/organizations/:organization_id/invitations" method="post" summary="Create an organization invitation" %}
{% swagger-description %}
Creates a new organization invitation and sends an email to the provided `email_address` with a link to accept the invitation and join the organization.

You can specify the `role` for the invited organization member.&#x20;

New organization invitations get a "pending" status until they are revoked by an organization administrator or accepted by the invitee.

The request body supports passing an optional `redirect_url` parameter. When the invited user clicks the link to accept the invitation, they will be redirected to the URL you provided. Use this parameter to implement a custom invitation acceptance flow.&#x20;

The authenticated user must be a member with administrator privileges in the organization. Only "admin" members can create organization invitations.
{% endswagger-description %}

{% swagger-parameter in="body" name="email_address" type="String" required="true" %}
The email address of the user that is invited to join the organization. An email to accept the invitation will be sent to this email address.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="role" type="String" required="true" %}
The role of the new organization member. Can be one of 

`admin`

 or 

`basic_member`

.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="redirect_url" type="String" %}
Optional URL that the invitee will be redirected to once they accept the invitation by clicking the join link in the invitation email.
{% endswagger-parameter %}

{% swagger-response status="200" description="Organization invitation was created successfully." %}
```javascript
{
  "object": "organization_invitation",
  "id": "orginv_21Ufcy98STcA11s3QckIwtwHIES",
  "email_address": "invitee@example.com",
  "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
  "role": "basic_member",
  "public_metadata": {},
  "status": "pending",
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

{% swagger-response status="403: Forbidden" description="Insufficient permissions to create organization invitations." %}
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

{% swagger baseUrl="https://clerk.example.com" path="/v1/organizations/:organization_id/invitations/:invitation_id/revoke" method="post" summary="Revoke a pending organization invitation" %}
{% swagger-description %}
Use this request to revoke a previously issued organization invitation.&#x20;

Revoking an organization invitation makes it invalid; the invited user will no longer be able to join the organization with the revoked invitation.

Only organization invitations with "pending" status can be revoked.&#x20;

The authenticated user needs to have administrator privileges in the organization. Only users with "admin" role can revoke invitations.
{% endswagger-description %}

{% swagger-response status="200" description="Organization invitation was revoked successfully." %}
```javascript
{
  "object": "organization_invitation",
  "id": "orginv_21Ufcy98STcA11s3QckIwtwHIES",
  "email_address": "invitee@example.com",
  "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
  "public_metadata": {},
  "role": "basic_member",
  "status": "revoked",
  "created_at": 1638000669544,
  "updated_at": 1638000889544
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

{% swagger-response status="403: Forbidden" description="Insufficient permissions to revoke organization invitations." %}
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
      "code": "organization_invitation_not_pending",
      "long_message": "The organization invitation is not in the \"pending\" status.",
      "message": "not pending"
    }
  ]
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://clerk.example.com" path="/v1/organizations/:organization_id/invitations/pending" method="get" summary="Retrieve a list of pending invitations for an organization" %}
{% swagger-description %}
This request returns the list of organization invitations with "pending" status. This means that the organization invitations can still be used to join the organization, but have not been accepted by the invited user.

The authenticated user must have administrator privileges in the organization. Only members with "admin" role can list pending organization invitations.

Results can be paginated using the optional `limit` and `offset` query parameters. The organization invitations are ordered by descending creation date. Most recent invitations will be returned first.
{% endswagger-description %}

{% swagger-parameter in="query" name="limit" type="Number" %}
Applies a limit to the number of organization invitations returned. Can be used for paginating the results together with 

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

{% swagger-response status="200" description="Pending organization invitations were retrieved successfully." %}
```javascript
[
  {
    "object": "organization_invitation",
    "id": "orginv_21Ufcy98STcA11s3QckIwtwHIES",
    "email_address": "invitee@example.com",
    "organization_id": "org_1o4qfak5AdI2qlXSXENGL05iei6",
    "public_metadata": {},
    "role": "basic_member",
    "status": "pending",
    "created_at": 1638000669544,
    "updated_at": 1638000889544
  },
  {
    "object": "organization_invitation",
    "id": "orginv_21Ufcy98STcA11s3QckIwtwyy6D",
    "email_address": "another@toinvite.com",
    "organization_id": "org_1o4qfak5AdI22rtSXENGL05iei6",
    "public_metadata": {},
    "role": "admin",
    "status": "pending",
    "created_at": 1538000669544,
    "updated_at": 1566000889544
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

{% swagger-response status="403: Forbidden" description="Insufficient permissions to list pending organization invitations." %}
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
