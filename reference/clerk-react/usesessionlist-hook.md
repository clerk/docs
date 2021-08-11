---
description: Access the full list of Client Sessions in your components.
---

# useSessionList

## Overview

The `useSessionList` hook provides access to all the sessions on the [Client](../clerkjs/client.md) object. It returns an array of [Session](../clerkjs/session.md) objects that have been registered on the client device.

Out of all the `Session` objects on the client, only one will be **active**. That's the current user's session. The rest of the sessions in the list can be in any of the rest [session states](../clerkjs/session.md#sessionstatus).

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

The example below is very basic, but illustrates how to get a hold of all the sessions on a given [Client](../clerkjs/client.md).

```jsx
import { useSessionList } from "@clerk/clerk-react";

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

