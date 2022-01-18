---
description: Navigate immediately to the sign-up URL
---

# \<RedirectToSignUp />

## Overview

Rendering a `<RedirectToSignUp/>` component will navigate to the sign up URL which has been configured in your application instance. You can find the configuration in the [Clerk Dashboard](https://dashboard.clerk.dev). Go to your application, select your instance, then go to **Paths**.&#x20;

The behavior will be just like a server-side (3xx) redirect, and will override the current location in the history stack.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

The example below shows a page which is normally accessible by authenticated users. All other visitors (unauthenticated requests) will get redirected to the sign up page.

The [\<SignedIn/>](../signed-in.md) and [\<SignedOut/>](../signed-out.md) components work together as a conditional, allowing the `<RedirectToSignUp/>` component to be rendered only for unauthenticated requests.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { 
  ClerkProvider,
  SignedIn,
  SignedOut,
  RedirectToSignUp
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
          to the sign up page.
        */}
        <RedirectToSignUp />
      </SignedOut>
    </div>
  );
}
```
{% endtab %}
{% endtabs %}

## Props

| Name                | Description                                                                                                                                                                                                                                                                             |
| ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **redirectUrl?**    | <p><em>string</em></p><p>Full URL or path to navigate to after successful sign in or sign up. Use this instead of setting <code>afterSignInUrl</code> and <code>afterSignUpUrl</code> to the same value.<br><br>To return to the same URL, set to <code>window.location.href</code></p> |
| **afterSignInUrl?** | <p><em>string</em></p><p>The full URL or path to navigate to after a successful sign in.<br><br>Defaults to the Sign-in URL on the Paths page of your Dashboard.</p>                                                                                                                    |
| **afterSignUpUrl?** | <p><em>string</em></p><p>The full URL or path to navigate to after a successful sign up.<br><br>Defaults to the Sign-up URL on the Paths page of your Dashboard.</p>                                                                                                                    |
