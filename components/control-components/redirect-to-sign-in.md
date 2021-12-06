---
description: Navigate to the sign in URL.
---

# \<RedirectToSignIn />

## Overview

Rendering a `<RedirectToSignIn/>` component will navigate to the sign in URL which has been configured in your application instance. You can find the configuration in the [Clerk Dashboard](https://dashboard.clerk.dev). Go to your application, select your instance, then go to **Settings** > **URL & redirects**.&#x20;

The behavior will be just like a server-side (3xx) redirect, and will override the current location in the history stack.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

The example below shows a page which is normally accessible by authenticated users. All other visitors (unauthenticated requests) will get redirected to the sign in page.

The [\<SignedIn/>](signed-in.md) and [\<SignedOut/>](signed-out.md) components work together as a conditional, allowing the `<RedirectToSignIn/>` component to be rendered only for unauthenticated requests.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { 
  ClerkProvider,
  SignedIn,
  SignedOut,
  RedirectToSignIn 
} from "@clerk/clerk-react";

function PrivatePage() {
  return (
    <ClerkProvider frontendApi="clerk.[your-domain].com">
      <SignedIn>
        Content that is displayed to signed in
        users.
      </SignedIn>
      <SignedOut>
        {/* 
          Non-authenticated visitors will be redirected
          to the sign in page.
        */}
        <RedirectToSignIn />
      </SignedOut>
    </div>
  );
}
```
{% endtab %}
{% endtabs %}

## Props

| Name                | Description                                                                                                                                                                                              |
| ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **redirectUrl?**    | <p><em>string</em></p><p>Full URL or path to navigate after successful sign in or sign up.<br><br>The same as setting <code>afterSignInUrl</code> and <code>afterSignUpUrl</code> to the same value.</p> |
| **afterSignInUrl?** | <p><em>string</em></p><p>The full URL or path to navigate after a successful sign in.</p>                                                                                                                |
| **afterSignUpUrl?** | <p><em>string</em></p><p>The full URL or path to navigate after a successful sign up.</p>                                                                                                                |
