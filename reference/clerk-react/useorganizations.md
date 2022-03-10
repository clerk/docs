---
description: Access methods for the Organization and related objects.
---

# useOrganizations

## Overview

The `useOrganizations` hook gives you access to [Organization](../clerkjs/organization.md) and related object methods. Calling the hook will return an object structure with the available methods.

For now the available methods are: `createOrganization`, `getOrganizationMemberships` and `getOrganization.`

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) `or` [`Clerk Next.js`](../../get-started/nextjs.md#installing-clerk) `` before running the snippets below.
{% endhint %}

The examples below illustrates simple usage of the methods available. Note that your component must be a descendant of the [\<SignedIn/>](../../components/signed-in.md) component, which in turn needs to be wrapped inside the [\<ClerkProvider/>](clerkprovider.md).

A more guided approach can be found at the [Organizations Popular Guide](../../popular-guides/organizations.md) section.

{% code title="Memberships" %}
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

export default function OrganizationMemberships() {
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
{% endcode %}
