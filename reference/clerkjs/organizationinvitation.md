# OrganizationInvitation

## Overview

The `OrganizationInvitation` object is the model around an organization invitation.

## Attributes

| Name               | Description                                                                                  |
| ------------------ | -------------------------------------------------------------------------------------------- |
| **id**             | <p><em>string</em></p><p>A unique identifier for this organization membership.</p>           |
| **emailAddress**   | <p><em>string</em></p><p>The email address the invitation has been sent.</p>                 |
| **organizationId** | <p><em>string</em></p><p>The organization id of the organization this invitation is for.</p> |
| **role**           | <p><em>MembershipRole</em><br><em></em>The role of the member for this invitation.</p>       |
| **status**         | <p><em>OrganizationInvitationStatus</em><br><em></em>The status of the invitation.</p>       |
| **createdAt**      | <p><em>Date</em></p><p>Date of the time the membership was created.</p>                      |
| **updatedAt**      | <p><em>Date</em></p><p>Date of the last time the membership was updated.</p>                 |

## Methods

### revoke()

`revoke()=>Promise<`OrganizationInvitationResource`>`

Revokes the invitation for the email it corresponds to.
