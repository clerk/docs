---
description: Learn how to manage organizations and their members
---

# Organizations

{% hint style="info" %}
Organizations is a premium feature. Please get in touch if you would like us to enable it for your account. You can contact us at [support@clerk.dev](mailto:support@clerk.dev).
{% endhint %}

## Overview

Organizations are shared accounts, useful for project and team leaders. Organization members can usually collaborate across shared resources. There are members with elevated privileges who can manage member access to the organization's data and resources.

Each member of an organization needs to have a user account in your application. All organization members have access to most of the organization resources, but some members can take advantage of administrative features.

### Roles

This distinction in member permissions is possible with organization member roles. Roles determine a user's level of access to the organization. There are currently two roles; administrators and members.

| Role              | Description                                                                                                                                                                                                                                                                              |
| ----------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **admin**         | The "admin" role offers full access to organization resources. Members with the admin role have administrator privileges and can fully manage organizations and organization memberships.                                                                                                |
| **basic\_member** | The "basic\_member" role is the standard role for a user that is part of the organization. Access to organization resources is limited. Basic members cannot manage organizations and organization memberships, but can view information about the organization and other members in it. |

### Available actions

You can use Clerk's organizations feature to provide team grouping and sharing functionality for users of your applications.

Your users can create organizations. Organization owners effectively get the "admin" role.&#x20;

Administrators can then invite other users to join the organization. An invitation email is sent out, and organization invitations support adding existing users of your application, or new ones. They can register once they accept the invitation.

Administrators can also revoke an invitation for a user that hasn't joined yet, as well as remove a user who's already a member from the organization or change their role. When removing organization members or updating their role, there needs to be at least one administrator for the organization at all times.

Finally, all members of an organization, regardless of their role can view information about other members in the same organization.

## Before you start

* You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). For more information, check out our [Setup your application](setup-your-application.md) guide.
* You need to install [Clerk React](../reference/clerk-react/) or [ClerkJS](../reference/clerkjs/) to your application.
* The organizations feature needs to be enabled for your account. Contact us at [support@clerk.dev](mailto:support@clerk.dev) to turn the feature on.

## Custom flow

Let's follow a simple scenario for setting up an organization and managing its members.&#x20;

In the following guide, we'll create a new organization and invite a couple of members. We'll see how to revoke one of the two invitations, and update the other member's role. Finally, we'll remove the other member from the organization.

This tutorial assumes that a [signed in user](../main-concepts/sign-in-flow.md) already exists. All snippets below require a session and user to be present.

#### Create a new organization

{% tabs %}
{% tab title="Clerk Next.js" %}
```jsx
// pages/organizations/new.js
import { useState } from "react";
import { useRouter } from "next/router";
import { useOrganizations } from "@clerk/nextjs";

// Form to create a new organization. The current user
// will become the organization administrator.
export default function NewOrganization() {
  const [name, setName] = useState("");
  const router = useRouter();

  const { createOrganization } = useOrganizations();

  async function submit(e) {
    e.preventDefault();
    try {
      // Create a new organization.
      await createOrganization({ name });
      setName("");
      router.push("/organizations");
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div>
      <h2>Create an organization</h2>
      <form onSubmit={submit}>
        <div>
          <label>Name</label>
          <br />
          <input
            name="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <button>Create organization</button>
      </form>
    </div>
  );
}
```
{% endtab %}

{% tab title="Clerk React" %}
```jsx
import { useState } from "react";
import { useOrganizations } from "@clerk/react";

// Form to create a new organization. The current user
// will become the organization administrator.
export default function NewOrganization() {
  const [name, setName] = useState("");

  const { createOrganization } = useOrganizations();

  async function submit(e) {
    e.preventDefault();
    try {
      // Create a new organization.
      await createOrganization({ name });
      setName("");
      // Assuming there's an organizations index 
      // page, navigate there.
      window.location = "/organizations";
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div>
      <h2>Create an organization</h2>
      <form onSubmit={submit}>
        <div>
          <label>Name</label>
          <br />
          <input
            name="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <button>Create organization</button>
      </form>
    </div>
  );
}
```
{% endtab %}

{% tab title="ClerkJS" %}
```html
<!-- Form to create an organization -->
<form id="new_organization">
  <div>
    <label>Name</label>
    <br />
    <input name="name" />
  </div>
  <button>Create organization</button>
</form>

<script>
  const form = document.getElementById("new_organization");
  form.addEventListener('submit', function(e) {
    e.preventDefault();
    const inputEl = form.getElementsByTagName("input")[0];
    if (!inputEl) {
      return;
    }
    try {
      await window.Clerk.createOrganization({ name: inputEl.value });
    } catch (err) {
      console.error(err);
    }
  });
</script>

```
{% endtab %}
{% endtabs %}

#### View all organizations the current user belongs to

{% tabs %}
{% tab title="Clerk Next.js" %}
```jsx
import { useState, useEffect } from "react";
import Link from "next/link";
import { useOrganizations } from "@clerk/nextjs";
import { OrganizationMembershipResource } from "@clerk/types";

// Lists all organization the user is a member of.
// Each entry is a link to a page to manage organization
// members.
export default function Organizations() {
  const [organizationMemberships, setOrganizationMemberships] = useState<
    OrganizationMembershipResource[]
  >([]);

  const { getOrganizationMemberships } = useOrganizations();

  useEffect(() => {
    async function fetchOrganizationMemberships() {
      try {
        const orgs = await getOrganizationMemberships();
        setOrganizationMemberships(orgs);
      } catch (err) {
        console.error(err);
      }
    }

    fetchOrganizationMemberships();
  }, []);

  return (
    <div>
      <h2>Your organizations</h2>
      <ul>
        {organizationMemberships.map(({ organization }) => (
            <Link
              key={organization.id}
              href={`/organizations/${organization.id}`}
            >
              {organization.name}
            </Link>
          ))}
      </ul>
    </div>
  );
}


```
{% endtab %}

{% tab title="Clerk React" %}
```jsx
import { useState, useEffect } from "react";
import { useOrganizations } from "@clerk/react";

// Lists all organization the user is a member of.
// Each entry is a link to a page to manage organization
// members.
export default function Organizations() {
  const [organizationMemberships, setOrganizationMemberships] = useState<
    OrganizationMembershipResource[]
  >([]);

  const { getOrganizationMemberships } = useOrganizations();

  useEffect(() => {
    async function fetchOrganizationMemberships() {
      try {
        const orgs = await getOrganizationMemberships();
        setOrganizationMemberships(orgs);
      } catch (err) {
        console.error(err);
      }
    }

    fetchOrganizationMemberships();
  }, []);

  return (
    <div>
      <h2>Your organizations</h2>
      <ul>
        {organizationMemberships.map(({ organization }) => (
          <a key={organization.id} href={`/organizations/${organization.id}`}>
            {organization.name}
          </a>
        ))}
      </ul>
    </div>
  );
}

```
{% endtab %}

{% tab title="ClerkJS" %}
```html
<ul id="organizations_list"></ul>

<script>
  const list = document.getElementById("organizations_list");
  try {
    const organizationMemberships = await window.Clerk.getOrganizationMemberships();
    organizationMemberships.map(({ organization }) => {
      const li = document.createElement("li");
      li.textContent = `${organization.name} - ${organization.role}`;
      list.appendChild(li);
    });
  } catch (err) {
    console.error(err);
  }
</script>

```
{% endtab %}
{% endtabs %}

#### Manage organization members

{% tabs %}
{% tab title="Clerk Next.js" %}
```jsx
// pages/organizations/[id].js
import { useState, useEffect } from "react";
import { useRouter } from "next/router";
import { useOrganizations } from "@clerk/nextjs";

// View and manage organization members, along with any
// pending invitations.
// Invite new members.
export default function Organization() {
  const [organizationMemberships, setOrganizationMemberships] = useState([]);

  const { query } = useRouter();
  const organizationId = query.id;

  const { getOrganizationMemberships } = useOrganizations();

  useEffect(() => {
    async function fetchOrganizationMemberships() {
      try {
        const orgMemberships = await getOrganizationMemberships();
        setOrganizationMemberships(orgMemberships);
      } catch (err) {
        console.log(err);
      }
    }

    fetchOrganizationMemberships();
  }, [organizationId, getOrganizationMemberships]);

  const currentOrganizationMembership = organizationMemberships.find(membership => 
    membership.organization.id === organizationId);

  if (!currentOrganizationMembership) {
    return null;
  }
  
  const isAdmin = currentOrganizationMembership.role === "admin";

  return (
    <div>
      <h2>{currentOrganizationMembership.organization.name}</h2>

      <Memberships organization={currentOrganizationMembership.organization} isAdmin={isAdmin} />
      {isAdmin && <Invitations organization={currentOrganizationMembership.organization} />}
    </div>
  );
}

// List of organization memberships. Administrators can
// change member roles or remove members from the organization.
function Memberships({ organization, isAdmin }) {
  const [memberships, setMemberships] = useState([]);

  const getMemberships = organization.getMemberships;

  const fetchMemberships = useCallback(async () => {
    try {
      const membs = await getMemberships();
      setMemberships(membs);
    } catch (err) {
      console.error(err);
    }
  }, [getMemberships]);

  useEffect(() => {
    fetchMemberships();
  }, [fetchMemberships]);

  async function remove(userId) {
    try {
      await organization.removeMember(userId);
      fetchMemberships();
    } catch (err) {
      console.error(err);
    }
  }

  async function switchRole(membership) {
    const role = membership.role === "admin" ? "basic_member" : "admin";
    try {
      await membership.update({ role });
      fetchMemberships();
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div>
      <b>Members</b>
      <ul>
        {memberships.map((membership) => (
          <li key={membership.id}>
            {membership.publicUserData.identifier} - {membership.role}
            {isAdmin && (
              <>
                <br />
                <button
                  onClick={(e) => {
                    switchRole(membership);
                  }}
                >
                  Change role
                </button>
                &nbsp;or&nbsp;
                <button
                  onClick={(e) => {
                    e.preventDefault();
                    remove(membership.publicUserData.userId);
                  }}
                >
                  Remove
                </button>
              </>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
}

// List of organization pending invitations. 
// You can invite new organization members and 
// revoke already sent invitations.
function Invitations({ organization }) {
  const [invitations, setInvitations] = useState([]);
  const [emailAddress, setEmailAddress] = useState("");

  const getPendingInvitations = organization.getPendingInvitations;

  const fetchInvitations = useCallback(async () => {
    try {
      const invites = await getPendingInvitations();
      setInvitations(invites);
    } catch (err) {
      console.error(err);
    }
  }, [getPendingInvitations]);

  useEffect(() => {
    fetchInvitations();
  }, [fetchInvitations]);

  async function invite(e) {
    e.preventDefault();
    try {
      await organization.inviteMember({ 
        emailAddress, 
        role: "basic_member", 
      });
      setEmailAddress("");
      fetchInvitations();
    } catch (err) {
      console.error(err);
    }
  }

  async function revoke(invitation) {
    try {
      await invitation.revoke();
      fetchInvitations();
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div>
      <b>Pending invitations</b>
      <ul>
        {invitations.map((invitation) => (
          <li key={invitation.id}>
            {invitation.emailAddress}
            &nbsp;
            <button
              onClick={(e) => {
                e.preventDefault();
                revoke(invitation);
              }}
            >
              Revoke
            </button>
          </li>
        ))}
      </ul>

      <b>Invite new member</b>
      <form onSubmit={invite}>
        <div>
          <label>Email address</label>
          <br />
          <input
            type="email"
            name="email_address"
            value={emailAddress}
            onChange={(e) => setEmailAddress(e.target.value)}
          />
        </div>
        <button>Invite</button>
      </form>
    </div>
  );
}
```
{% endtab %}

{% tab title="Clerk React" %}
```jsx
import { useState, useEffect } from "react";
import { useOrganizations } from "@clerk/react";

// View and manage organization members, along with any
// pending invitations.
// Invite new members.
export default function Organization({ organizationId }) {
  const [organizationMemberships, setOrganizationMemberships] = useState([]);

  const { getOrganizationMemberships } = useOrganizations();

  useEffect(() => {
    async function fetchOrganizationMemberships() {
      try {
        const orgMemberships = await getOrganizationMemberships();
        setOrganizationMemberships(orgMemberships);
      } catch (err) {
        console.log(err);
      }
    }

    fetchOrganizationMemberships();
  }, [organizationId, getOrganizationMemberships]);

  const currentOrganizationMembership = organizationMemberships.find(membership => 
    membership.organization.id === organizationId);

  if (!currentOrganizationMembership) {
    return null;
  }
  
  const isAdmin = currentOrganizationMembership.role === "admin";

  return (
    <div>
      <h2>{currentOrganizationMembership.organization.name}</h2>

      <Memberships organization={currentOrganizationMembership.organization} isAdmin={isAdmin} />
      {isAdmin && <Invitations organization={currentOrganizationMembership.organization} />}
    </div>
  );
}

// List of organization memberships. Administrators can
// change member roles or remove members from the organization.
function Memberships({ organization, isAdmin }) {
  const [memberships, setMemberships] = useState([]);

  const getMemberships = organization.getMemberships;

  const fetchMemberships = useCallback(async () => {
    try {
      const membs = await getMemberships();
      setMemberships(membs);
    } catch (err) {
      console.error(err);
    }
  }, [getMemberships]);

  useEffect(() => {
    fetchMemberships();
  }, [fetchMemberships]);

  async function remove(userId) {
    try {
      await organization.removeMember(userId);
      fetchMemberships();
    } catch (err) {
      console.error(err);
    }
  }

  async function switchRole(membership) {
    const role = membership.role === "admin" ? "basic_member" : "admin";
    try {
      await membership.update({ role });
      fetchMemberships();
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div>
      <b>Members</b>
      <ul>
        {memberships.map((membership) => (
          <li key={membership.id}>
            {membership.publicUserData.identifier} - {membership.role}
            {isAdmin && (
              <>
                <br />
                <button
                  onClick={(e) => {
                    switchRole(membership);
                  }}
                >
                  Change role
                </button>
                &nbsp;or&nbsp;
                <button
                  onClick={(e) => {
                    e.preventDefault();
                    remove(membership.publicUserData.userId);
                  }}
                >
                  Remove
                </button>
              </>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
}

// List of organization pending invitations. 
// You can invite new organization members and 
// revoke already sent invitations.
function Invitations({ organization }) {
  const [invitations, setInvitations] = useState([]);
  const [emailAddress, setEmailAddress] = useState("");

  const getPendingInvitations = organization.getPendingInvitations;

  const fetchInvitations = useCallback(async () => {
    try {
      const invites = await getPendingInvitations();
      setInvitations(invites);
    } catch (err) {
      console.error(err);
    }
  }, [getPendingInvitations]);

  useEffect(() => {
    fetchInvitations();
  }, [fetchInvitations]);

  async function invite(e) {
    e.preventDefault();
    try {
      await organization.inviteMember({ 
        emailAddress, 
        role: "basic_member", 
      });
      setEmailAddress("");
      fetchInvitations();
    } catch (err) {
      console.error(err);
    }
  }

  async function revoke(invitation) {
    try {
      await invitation.revoke();
      fetchInvitations();
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div>
      <b>Pending invitations</b>
      <ul>
        {invitations.map((invitation) => (
          <li key={invitation.id}>
            {invitation.emailAddress}
            &nbsp;
            <button
              onClick={(e) => {
                e.preventDefault();
                revoke(invitation);
              }}
            >
              Revoke
            </button>
          </li>
        ))}
      </ul>

      <b>Invite new member</b>
      <form onSubmit={invite}>
        <div>
          <label>Email address</label>
          <br />
          <input
            type="email"
            name="email_address"
            value={emailAddress}
            onChange={(e) => setEmailAddress(e.target.value)}
          />
        </div>
        <button>Invite</button>
      </form>
    </div>
  );
}
```
{% endtab %}

{% tab title="ClerkJS" %}
```html
<ul id="memberships_list"></ul>
<ul id="invitations_list"></ul>

<form id="new_invitation">
  <div>
    <label>Email address</div>
    <br />
    <input type="email" name="email_address" />
  </div>
  <button>Invite</button>
</form>

<script>
  async function renderMemberships(organization) {
    const list = document.getElementById("memberships_list");
    try {
      const isAdmin = organization.role === "admin";

      const memberships = await organization.getMembers();
      memberships.map((membership) => {
        const li = document.createElement("li");
        li.textContent = `${membership.identifier} - ${membership.role}`;

        // Add administrative actions; update role and remove member.
        if (isAdmin) {
          const updateBtn = document.createElement("button");
          updateBtn.textContent = "Change role";
          updateBtn.addEventListener("click", async function(e) {
            e.preventDefault();
            const role = membership.role === "admin" ?
              "basic_member" :
              "admin";
            await membership.update({ role });
          });
          li.appendChild(updateBtn);

          const removeBtn = document.createElement("button");
          removeBtn.textContent = "Remove";
          removeBtn.addEventListener("click", async function(e) {
            e.preventDefault();
            await currentOrganization.removeMember(membership.userId);
          });
          li.appendChild(removeBtn);
        }

        // Add the entry to the list
        list.appendChild(li);
      });
    } catch (err) {
      console.error(err);
    }
  }

  async function renderInvitations(organization) {
    const list = document.getElementById("invitations_list");
    try {
      const isAdmin = organization.role === "admin";

      const invitations = await organization.getPendingInvitations();
      invitations.map((invitation) => {
        const li = document.createElement("li");
        li.textContent = `${invitation.emailAddress} - ${invitation.role}`;

        // Add administrative actions; revoke invitation
        if (isAdmin) {
          const revokeBtn = document.createElement("button");
          revokeBtn.textContent = "Revoke";
          revokeBtn.addEventListener("click", async function(e) {
            e.preventDefault();
            await invitation.revoke();
          });
          li.appendChild(revokeBtn);
        }
        // Add the entry to the list
        list.appendChild(li);
      });
    } catch (err) {
      console.error(err);
    }
  }

  async function init() {
    // This is the current organization ID.
    const organizationId = "org_XXXXXXX";
    const currentOrganization = await window.Clerk.getOrganization(organizationId)
    if (!currentOrganization) {
      return;
    }
    const isAdmin = currentOrganization.role === "admin";

    renderMemberships(currentOrganization);
    renderInvitations(currentOrganization);

    if (isAdmin) {
      const form = document.getElementById("new_invitation");
      form.addEventListener("submit", async function(e) {
        e.preventDefault();
        const inputEl = form.getElementsByTagName("input");
        if (!inputEl) {
          return;
        }

        try {
          await currentOrganization.inviteMember({
            emailAddress: inputEl.value,
            role: "basic_member",
          });
        } catch (err) {
          console.error(err);
        }
      });
    }
  }

  init();
</script>

```
{% endtab %}
{% endtabs %}

