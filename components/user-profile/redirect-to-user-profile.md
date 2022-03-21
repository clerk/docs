---
description: Navigate to the user profile URL.
---

# \<RedirectToUserProfile />

## Overview

Rendering a `<RedirectToUserProfile/>` component will navigate to the user profile URL which has been configured in your application instance. You can find the configuration on the [Paths page](https://dashboard.clerk.dev/last-active?path=paths).

This component will use the custom `navigate` function from the [`<ClerkProvider/>` component](../../reference/clerk-react/clerkprovider.md) if one is given - otherwise it will trigger a full page reload with the new URL location.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

The example below shows a page which is normally accessible by authenticated users. All other visitors (unauthenticated users) will get redirected to the sign in page.

The [\<SignedIn/>](../signed-in.md) and [\<SignedOut/>](../signed-out.md) components work together as a conditional, allowing the `<RedirectToUserProfile />` component to be rendered only for authenticated users.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import {
  ClerkProvider,
  SignedOut,
  SignedIn,
  RedirectToSignIn,
  RedirectToUserProfile,
} from "@clerk/clerk-react";

const frontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;

function App() {
  return (
    <ClerkProvider frontendApi={frontendApi}>
      {/* Signed in users will get redirected to
       the user profile page */}
      <SignedIn>
        <RedirectToUserProfile />
      </SignedIn>

      <SignedOut>
        <RedirectToSignIn />
      </SignedOut>
    </ClerkProvider>
  );
}

export default App;
```
{% endtab %}
{% endtabs %}

## Props

This component accepts no props.
