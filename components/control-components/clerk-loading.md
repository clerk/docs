---
description: Renders while Clerk is loading.
---

# &lt;ClerkLoading&gt;

## Overview

The `<ClerkLoading>` renders its children while Clerk is loading, and is helpful for showing a custom loading state.

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
        Loading
      </ClerkLoading>
      <ClerkLoaded>
        <Page />
      </ClerkLoaded>
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

<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>children</b>
      </td>
      <td style="text-align:left">
        <p><em>JSX.Element</em>
        </p>
        <p>Pass any component or components to render while Clerk is loading.</p>
      </td>
    </tr>
  </tbody>
</table>



