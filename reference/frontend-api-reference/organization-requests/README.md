---
description: Manage organizations and memberships with the Frontend API
---

# Organization requests

{% hint style="info" %}
Organizations is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Overview

Organizations are shared accounts for your application users, so that they can collaborate and share information.

Users that belong to an organization are called members. Each member has a role, and Clerk supports a role with elevated privileges, which is the organization administrator.

Organizations requests allow you to create organizations and view information about which organizations your users belong to.

We've built an organization invitation system that you can trigger from our Frontend API and invite users to become organization members. You can also update members' roles and remove members from an organization.

## Requests summary

### [Organization requests](organization-requests.md)

* `POST /v1/organizations`
* `DELETE /v1/organizations/:organization_id`

### [Organization invitation requests](organization-invitations-requests.md)

* `POST /v1/organizations/:organization_id/invitations`
* `POST /v1/organizations/:organization_id/invitations/:invitation_id/revoke`
* `GET /v1/organizations/:organization_id/invitations/pending`

### [Organization membership requests](organization-membership-requests.md)

* `POST /v1/organizations/:organization_id/memberships`
* `GET /v1/organizations/:organization_id/memberships`
* `PATCH /v1/organizations/:organization_id/memberships/:user_id`
* `DELETE /v1/organizations/:organization_id/memberships/:user_id`

