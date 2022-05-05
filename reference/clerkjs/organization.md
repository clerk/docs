---
description: The Organization object describes a user created organization structure.
---

# Organization

## Overview

The `Organization` object is the model around the organization entity.



## Attributes

| Name               | Description                                                                                                                                                                                                                                |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **id**             | <p><em>string</em></p><p>A unique identifier for this organization.</p>                                                                                                                                                                    |
| **logoUrl**        | <p><em>string</em><br>The full URL for the organization logo.</p>                                                                                                                                                                          |
| **name**           | <p><em>string</em></p><p>The organization name.</p>                                                                                                                                                                                        |
| **publicMetadata** | <p><em>object</em><br>Additional information about the organization that can be read through the <a href="../frontend-api-reference/">Frontend API</a>, but written only from the <a href="../backend-api-reference/">Backend API</a>.</p> |
| **slug**           | <p><em>string</em><br>The organization slug. If supplied, it must be unique for the instance.</p>                                                                                                                                          |
| **createdAt**      | <p><em>Date</em></p><p>Date of the time the organization was created.</p>                                                                                                                                                                  |
| **updatedAt**      | <p><em>Date</em></p><p>Date of the last time the user was updated.</p>                                                                                                                                                                     |

## Methods

### addMember(params)

`addMember(params: AddMemberParams) => Promise<OrganizationMembershipResource>`

Adds a user as a member to an organization. A user can only be added to an organization if they are not already members of it and if they already exist in the same instance as the organization.

Please note that only administrators can add members to an organization.

{% tabs %}
{% tab title="Parameters" %}


| Name       | Description                                                                                                                                        |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| **userId** | <p><em>string</em><br>The id of the user that will be added as a member to the organization.</p>                                                   |
| **role**   | <p><em>string</em><br>The role that the user will have in the organization. Valid values are <code>admin</code> and <code>basic_member</code>.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<OrganizationMembershipResource>_

This method returns a `Promise` which resolves to an `OrganizationMembershipResource` object.
{% endtab %}
{% endtabs %}

### destroy()

`destroy() => Promise<void>`

Deletes the organization. Only administrators can delete an organization.

Please note that deleting an organization will also delete all memberships and invitations. This is not reversible.

{% tabs %}
{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

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

### setLogo(params)

`setLogo(params: SetOrganizationLogoParams) => Promise<OrganizationResource>`

Sets or replaces an organization's logo. Accept the logo as a file or blob. The logo must be an image and its size cannot exceed 10MB.

{% tabs %}
{% tab title="Paramters" %}
| Name     | Description                                                            |
| -------- | ---------------------------------------------------------------------- |
| **file** | <p><em>File | Blob</em><br>An image file which cannot exceed 10MB.</p> |


{% endtab %}

{% tab title="Returns" %}
_Promise\<OrganizationResource>_\
This method returns a `Promise` which resolves to an [Organization](organization.md).
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

