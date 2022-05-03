---
description: The Organization object describes a user created organization structure.
---

# Organization

## Overview

The `Organization` object is the model around the organization entity.



## Attributes

| Name          | Description                                                               |
| ------------- | ------------------------------------------------------------------------- |
| **id**        | <p><em>string</em></p><p>A unique identifier for this organization.</p>   |
| **name**      | <p><em>string</em></p><p>The organization name.</p>                       |
| **createdAt** | <p><em>Date</em></p><p>Date of the time the organization was created.</p> |
| **updatedAt** | <p><em>Date</em></p><p>Date of the last time the user was updated.</p>    |

## Methods

### getMemberships(params)

`getMemberships(params?:` GetMembershipsParams`)=>Promise<`OrganizationMembershipResource\[]`>`

Retrieve the members of the selected organization.

{% tabs %}
{% tab title="Parameters" %}
| Name        | Description                                                                     |
| ----------- | ------------------------------------------------------------------------------- |
| **limit?**  | <p><em>number</em></p><p>Limit of the results returned.</p>                     |
| **offset?** | <p><em>number</em><br><em></em>The offset of results to start the retrieval</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise<_OrganizationMembershipResource_>_

This method returns a `Promise` which resolves with a list of  `OrganizationMembershipResource` objects.
{% endtab %}
{% endtabs %}

### getPendingInvitations()

`getPendingInvitations()=>Promise<`OrganizationInvitationResource\[]`>`

Retrieve the members of the selected organization.

{% tabs %}
{% tab title="Returns" %}
_Promise<_OrganizationInvitationResource\[]_>_

This method returns a `Promise` which resolves with a list of `OrganizationInvitationResource` objects.
{% endtab %}
{% endtabs %}

### inviteMember(params)

`inviteMember(params: InviteMemberParams) =>Promise<`OrganizationInvitationResource`>`

Creates and sends an invitation to the target email address for becoming a member with the role passed on the function parameters.

{% tabs %}
{% tab title="Parameters" %}
| Name             | Description                                                                                   |
| ---------------- | --------------------------------------------------------------------------------------------- |
| **emailAddress** | <p><em>string</em></p><p>The email address to invite.</p>                                     |
| **role**         | <p><em>MembershipRole</em><br><em></em>The role of the new member.</p>                        |
| **redirectUrl?** | <p><em>string</em><br><em></em>A URL to redirect the user after accepting the invitation.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise<_OrganizationInvitationResource_>_

This method returns a `Promise` which resolves to the  `OrganizationInvitationResource` for the created invitation.
{% endtab %}
{% endtabs %}

### update(params)

`update(params: UpdateOrganizationParams) => Promise<OrganizationResource>`

Updates an organization's attributes.

{% tabs %}
{% tab title="Parameters" %}


| Name     | Description                                      |
| -------- | ------------------------------------------------ |
| **name** | <p><em>string</em><br>The organization name.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<OrganizationResource>_\
This method returns a `Promise` which resolves to an [Organization](organization.md).
{% endtab %}
{% endtabs %}

### updateMember(params)

`updateMember(params: UpdateMembershipParams) =>Promise<`OrganizationMembershipResource`>`

Updates a member based on the _userId_. For now only role can be changed.

{% tabs %}
{% tab title="Parameters" %}
| Name       | Description                                                            |
| ---------- | ---------------------------------------------------------------------- |
| **role**   | <p><em>MembershipRole</em><br><em></em>The role of the new member.</p> |
| **userId** | <p><em>string</em><br><em></em>The user identifier.</p>                |
{% endtab %}

{% tab title="Returns" %}
_Promise<_OrganizationInvitationResource_>_

This method returns a `Promise` which resolves to the  `OrganizationMembershipResource` for the updated membership.
{% endtab %}
{% endtabs %}

### removeMember(params)

`removeMember(userId: string) =>Promise<`OrganizationMembershipResource`>`

Removes a member from the organization based on the userId.

{% tabs %}
{% tab title="Parameters" %}
| Name       | Description                                             |
| ---------- | ------------------------------------------------------- |
| **userId** | <p><em>string</em><br><em></em>The user identifier.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise<_OrganizationInvitationResource_>_

This method returns a `Promise` which resolves to the removed  `OrganizationMembershipResource`.
{% endtab %}
{% endtabs %}

###

