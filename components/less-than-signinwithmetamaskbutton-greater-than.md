---
description: Sign in with Metamask Web3 wallet in the browser
---

# \<SignInWithMetamaskButton/>

## Overview

The `<SignInWithMetamaskButton/>` component is used to complete a one-click, cryptographically-secure login flow using [MetaMask](https://metamask.io), with all data stored on the Clerk [user](../main-concepts/user-object.md).

Internally, it calls the `Clerk.authenticateWithMetamask()` method.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

The most common scenario for using the `<SignInWithMetamaskButton/>` component is to execute a [Metamask](https://metamask.io) sign in or sign up flow in React and NextJS apps.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import React from 'react';
import { render } from 'react-dom';
import {
  ClerkProvider,
  SignedIn,
  SignedOut,
  UserButton,
  SignInWithMetamaskButton,
} from '@clerk/clerk-react';

const frontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;

function App() {
  return (
    <ClerkProvider frontendApi={frontendApi}>
      <SignedIn>
        <UserButton />
      </SignedIn>
      <SignedOut>
        <SignInWithMetamaskButton />
      </SignedOut>
    </ClerkProvider>
  );
}

render(<App />, document.getElementById('root'));
```
{% endtab %}
{% endtabs %}

### Props

| Name             | Description                                                                                     |
| ---------------- | ----------------------------------------------------------------------------------------------- |
| **redirectUrl?** | <p><em>string</em></p><p>Full URL or path to navigate after successful sign in  or sign up.</p> |

### Simple button

```jsx
<SignInWithMetamaskButton />
```

### Custom button

```xml
<SignInWithMetamaskButton>
  <button className="btn">
    Web3 authentication rocks!
  </button>
</SignInWithMetamaskButton>
```
