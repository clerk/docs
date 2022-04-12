---
description: Renders while Clerk is loading.
---

# \<ClerkLoading>

## Overview

The `<ClerkLoading>` renders its children while Clerk is loading, and is helpful for showing a custom loading state.\
\
This component only renders during client-side rendering (CSR only).

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { useEffect } from "react";
import { ClerkLoaded, ClerkLoading, ClerkProvider } from "@clerk/clerk-react";

function App() {
  return (
    <ClerkProvider frontendApi="clerk.[your-domain].com">
      <ClerkLoading>
        <div>Clerk is loading...</div>
      </ClerkLoading>
      <ClerkLoaded>
        <Page/>
      </ClerkLoaded>
      <div>This div is always visible</div>
    </ClerkProvider>
  );
}

function Page() {
  return (
    <div>The content</div>
  );
}
```
{% endtab %}
{% endtabs %}

## Props

This component accepts no props, apart from any child components that will render while Clerk is loading.

| Name         | Description                                                                                          |
| ------------ | ---------------------------------------------------------------------------------------------------- |
| **children** | <p><em>JSX.Element</em></p><p>Pass any component or components to render while Clerk is loading.</p> |

