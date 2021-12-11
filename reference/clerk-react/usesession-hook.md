---
description: Access the current Session object inside your components.
---

# useSession

## Overview

The `useSession` hook accesses the current user [Session](../clerkjs/session.md) object. It can be used to retrieve information about the session status. The `useSession` hook is a shortcut for retrieving the [Clerk.session](../clerkjs/clerk.md#attributes) property.

As soon as a [User](../clerkjs/user.md) signs in, we create a [Session](../clerkjs/session.md) on the [Client](../clerkjs/client.md) object. Only one session can be active on a single client, and that's exactly the session that is returned by the `useSession` hook.

The `Session` object returned from the hook will hold all state for the currently active session. As such, the `useSession` hook must be called from a component that is a descendant of the [\<SignedIn/>](../../components/signed-in.md) component. Otherwise, the hook will throw an error.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

The following example accesses the [Session](../clerkjs/session.md) object in order to display how long a user has been active in this client session. Note that you can also get to the [User](../clerkjs/user.md) object through the `useSession` hook.&#x20;

```jsx
import { SignedIn, useSession } from "@clerk/clerk-react";

function App() {
  return (
    <SignedIn>
      <Analytics />
    </SignedIn>
  );
}

function Analytics() {
  const session = useSession();
  const { user } = session;
  
  return (
    <div>
      User {user.firstName} has been active 
      since {session.lastActiveAt}.
    </div>
  );
}
```
