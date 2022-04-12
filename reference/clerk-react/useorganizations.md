---
description: Access methods for the Organization and related objects.
---

# useOrganizations

## Overview

The `useOrganizations` hook gives you access to [Organization](../clerkjs/organization.md) and related object methods. Calling the hook will return an object structure with the available methods.

For now the available methods are: `createOrganization`, `getOrganizationMemberships` and `getOrganization.`

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) `or` [`Clerk Next.js`](broken-reference) `` before running the snippets below.
{% endhint %}

The examples below illustrates simple usage of the methods available. Note that your component must be a descendant of the [\<SignedIn/>](../../components/signed-in.md) component, which in turn needs to be wrapped inside the [\<ClerkProvider/>](clerkprovider.md).

A more guided approach can be found at the [Organizations Popular Guide](broken-reference) section.

{% tabs %}
{% tab title="createOrganization" %}
```jsx
import { useState } from "react";
import { SignedIn, useOrganizations } from "@clerk/clerk-react";

function App() {
  return (
    <SignedIn>
      <NewOrganization />
    </SignedIn>
  );
}

// Form to create a new organization. The current user
// will become the organization administrator.
function NewOrganization() {
  const [name, setName] = useState("");
  const router = useRouter();

  const { createOrganization } = useOrganizations();

  async function submit(e) {
    e.preventDefault();
    try {
      // Create a new organization.
      await createOrganization({ name });
      setName("");
      // Do anything after the action is complete
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

{% tab title="getMemberships" %}
```jsx
// Retrieve OrganizationMemberships of the current user.
import { SignedIn, useOrganizations } from "@clerk/clerk-react";

function App() {
  return (
    <SignedIn>
      <OrganizationMemberships />
    </SignedIn>
  );
}

function OrganizationMemberships() {
  const [organizationMemberships, setOrganizationMemberships] = useState([]);

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
            <p
              key={organization.id}
            >
              {organization.name}
            </p>
          ))}
      </ul>
    </div>
  );
}
```
{% endtab %}

{% tab title="getOrganization" %}
```jsx
import { useState } from "react";
import { SignedIn, useOrganizations } from "@clerk/clerk-react";

function App() {
  return (
    <SignedIn>
      <GetOrganization orgId={'test_org_id'} />
    </SignedIn>
  );
}

function GetOrganization({ orgId }){
  const [organization, setOrganization] = useState(null);

  const { getOrganization } = useOrganizations();

  useEffect(() => {
    async function fetchOrganization() {
      try {
        const org = await getOrganization(orgId);
        setOrganization(orgs);
      } catch (err) {
        console.error(err);
      }
    }

    fetchOrganization();
  }, [orgId]);
  
  if(!organization){
    return null;
  }
  
  return (
    <div>
      <h2>Organization {organization.name} with id: {organization.id}</h2>
    </div>
  );
}
```
{% endtab %}
{% endtabs %}

#### View all organizations the current user belongs to

