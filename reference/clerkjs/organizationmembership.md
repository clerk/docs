# OrganizationMembership

## Overview

The `OrganizationMembership` object is the model around an organization membership entity.

## Attributes

| Name               | Description                                                                                                                                         |
| ------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| **id**             | <p><em>string</em></p><p>A unique identifier for this organization membership.</p>                                                                  |
| **role**           | <p><em>MembershipRole</em></p><p>The role of the current user in the organization.</p>                                                              |
| **publicUserData** | <p><em>PublicUserData</em></p><p>Public information about the user that this membership belongs to.</p>                                             |
| **organization**   | <p><em></em><a href="organization.md"><em>OrganizationResource</em></a><em></em><br><em></em>The organization object the membership belongs to.</p> |
| **createdAt**      | <p><em>Date</em></p><p>Date of the time the membership was created.</p>                                                                             |
| **updatedAt**      | <p><em>Date</em></p><p>Date of the last time the membership was updated.</p>                                                                        |

## Methods

### destroy()

`destroy()=>Promise<void>`

Deletes the membership from the organization it belongs to.

### update(params)

`updateMember(params:` UpdateOrganizationMembershipParams`) =>Promise<`OrganizationMembershipResource`>`

Updates the membership. For now only role can be changed.

{% tabs %}
{% tab title="Parameters" %}
| Name     | Description                                                            |
| -------- | ---------------------------------------------------------------------- |
| **role** | <p><em>MembershipRole</em><br><em></em>The role of the new member.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise<_OrganizationInvitationResource_>_

This method returns a `Promise` which resolves to the  `OrganizationMembershipResource` for the updated membership.
{% endtab %}
{% endtabs %}
