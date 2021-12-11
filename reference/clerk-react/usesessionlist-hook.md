---
description: Access the full list of Client Sessions in your components.
---

# useSessionList

## Overview

The `useSessionList` hook provides access to all the sessions on the [Client](../clerkjs/client.md) object. It returns an array of [Session](../clerkjs/session.md) objects that have been registered on the client device.

Out of all the `Session` objects on the client, at least one will will be **active**. In a multi-session application there can be more than one active sessions, but in a single-session application there's only one active session. The rest of the sessions in the list can be in any of the rest [session states](../clerkjs/session.md#sessionstatus).

Whatever the case, the [Client.session](../clerkjs/client.md#attributes) attribute will hold the current user's active session.&#x20;

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

The example below is very basic, but illustrates how to get a hold of all the sessions on a given [Client](../clerkjs/client.md). Please note that the `useSessionList` hook will throw an error unless your component is a descendant of the [\<ClerkLoaded/>](../../components/control-components/clerk-loaded.md) component. The [\<SignedIn/>](../../components/signed-in.md) and [\<SignedOut/>](../../components/signed-out.md) components are already wrapped inside a [\<ClerkLoaded/>](../../components/control-components/clerk-loaded.md) component, so any of these will be enough.

```jsx
import { SignedIn, useSessionList } from "@clerk/clerk-react";

function App() {
  return (
    <SignedIn>
      <Activity />
    </SignedIn>
  );
}

function Activity() {
  const sessions = useSessionList();
  
  return (
    <div>
      Welcome back. You've been here 
      {sessions.length} times before.
    </div>
  );
}
```
