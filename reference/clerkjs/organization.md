---
description: The Organization object describes a user created organization structure.
---

# Organization

## Overview

The `Organization` object is the model around the organization entity.



## Attributes

| Name           | Description                                                                                    |
| -------------- | ---------------------------------------------------------------------------------------------- |
| **id**         | <p><em>string</em></p><p>A unique identifier for this organization.</p>                        |
| **name**       | <p><em>string</em></p><p>The organization name.</p>                                            |
| **instanceId** | <p><em>string</em></p><p>Identifier of the instance which the organization was created at.</p> |
| **createdAt**  | <p><em>Date</em></p><p>Date of the time the organization was created.</p>                      |
| **updatedAt**  | <p><em>Date</em></p><p>Date of the last time the user was updated.</p>                         |

## Methods

### getMembers(params)

`getMembers(params:` GetMembersParams`) =>Promise<`OrganizationMembershipResource`>`

Retrieve the members of the selected organization.

{% tabs %}
{% tab title="Parameters" %}
| Name       | Description                                                                     |
| ---------- | ------------------------------------------------------------------------------- |
| **limit**  | <p><em>number</em></p><p>Limit of the results returned.</p>                     |
| **offset** | <p><em>number</em><br><em></em>The offset of results to start the retrieval</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise<_OrganizationMembershipResource_>_

This method returns a `Promise` which resolves with a `OrganizationMembershipResource` object.
{% endtab %}
{% endtabs %}

###
