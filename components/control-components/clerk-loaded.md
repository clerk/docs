---
description: Guarantee that the Clerk object exists.
---

# \<ClerkLoaded>

## Overview

The `<ClerkLoaded>` component guarantees that the [Clerk](../../reference/clerkjs/clerk.md) object has loaded and will be available under `window.Clerk`.

It essentially provides a wrapper, enabling child components to access the [Clerk](../../reference/clerkjs/clerk.md) object without the need to check if it exists.&#x20;

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

Rendering the `<ClerkLoaded/>` component allows immediate access to the [Clerk](../../reference/clerkjs/clerk.md) object without the need to check if it exists.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { useEffect } from "react";
import { ClerkLoaded, ClerkProvider } from "@clerk/clerk-react";

function App() {
  return (
    <ClerkProvider frontendApi="clerk.[your-domain].com">
      <ClerkLoaded>
        <Page />
      </ClerkLoaded>
    </ClerkProvider>
  );
}

function Page() {
  useEffect(() => {
    // No need to check if the 
    // Clerk object exists.
    document.title = "This page uses Clerk " + 
      document.window.Clerk.version;
  }, []);
  
  return (
    <div>The content</div>
  );
}
```
{% endtab %}
{% endtabs %}

## Props

This component accepts no props, apart from any child components that will be able to safely access the [Clerk](../../reference/clerkjs/clerk.md) object.

| Name         | Description                                                                                                                                                      |
| ------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **children** | <p><em>JSX.Element</em></p><p>Pass any component or components and they will have access to the <a href="../../reference/clerkjs/clerk.md">Clerk</a> object.</p> |

