---
description: Navigate to the sign in URL.
---

# &lt;RedirectToSignIn /&gt;

## Overview

Rendering a `<RedirectToSignIn/>` component will navigate to the sign in URL which has been configured in your application instance. You can find the configuration in the [Clerk Dashboard](https://dashboard.clerk.dev). Go to your application, select your instance, then go to **Settings** &gt; **URL & redirects**. 

The behavior will be just like a server-side \(3xx\) redirect, and will override the current location in the history stack.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

The example below shows a page which is normally accessible by authenticated users. All other visitors \(unauthenticated requests\) will get redirected to the sign in page.

The [&lt;SignedIn/&gt;](signed-in.md) and [&lt;SignedOut/&gt;](signed-out.md) components work together as a conditional, allowing the `<RedirectToSignIn/>` component to be rendered only for unauthenticated requests.

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

This component accepts no props.

