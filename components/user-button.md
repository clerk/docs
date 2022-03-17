---
description: >-
  A familiar component that allows users to manage their account, switch
  accounts, or sign out
---

# \<UserButton />

## Overview

Originally popularized by Google, users have come to expect that little photo of themselves in the top-right of the page – it’s the access point to manage their account, switch accounts, or sign out. The `<UserButton/>` component is used to render this familiar user button UI. It renders a clickable user avatar - when clicked, the full UI opens as a popup.&#x20;



![The default \<UserButton/> component](../.gitbook/assets/user-button.png)

Clerk is the only provider with multi-session support, allowing users to sign into multiple accounts at once and switch between them. For multisession apps, the `<UserButton/>` automatically supports instant account switching, without the need of a full page reload. For more information, you can check out the [Multi-session applications](../popular-guides/popular-guides-multi-session-applications.md) guide.

Control the look and feel of the `<UserButton/>` component and match it to your using the [Theme Settings](../popular-guides/setup-your-application.md#theme), [theming props](user-profile/user-profile.md#customization) or [plain CSS](user-profile/user-profile.md#customization).

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../reference/clerk-react/installation.md) or [ClerkJS](../reference/clerkjs/installation.md) before running the snippets below.
{% endhint %}

### Mounting in your app

Once you set up the desired functionality and look and feel for the `<UserButton/>` component, all that's left is to render it inside your page. The default rendering is simple but powerful enough to cover most use-cases. The theme configuration (look and feel) that you've set up on the [Theme page](https://dashboard.clerk.dev/last-active?path=customization/theme) will work out of the box.

{% tabs %}
{% tab title="Clerk React" %}
{% code title="src/App.jsx" %}
```jsx
import {
  ClerkProvider,
  SignIn,
  SignedOut,
  UserButton,
} from "@clerk/clerk-react";

const frontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;

function App() {
  return (
    <ClerkProvider frontendApi={frontendApi}>
      <Header />
      {/* If the user is signed out, show the SignIn component */}
      {/* After signing in, the user button will be visible */}
      <SignedOut>
        <SignIn />
      </SignedOut>
    </ClerkProvider>
  );
}

function Header() {
  return (
    <header>
      <h1>My application</h1>
      {/* Mount the UserButton component */}
      <UserButton />
    </header>
  );
}

export default App;
```
{% endcode %}
{% endtab %}

{% tab title="ClerkJS" %}
```markup
<html>
<body>
    <div id="user-button"></div>
    
    <script>
        const el = document.getElementById("user-button");
        // Mount the pre-built Clerk UserButton component
        // in an HTMLElement on your page. 
        window.Clerk.mountUserButton(el);
    </script>
</body>
</html>
```
{% endtab %}
{% endtabs %}

## Props

| Name                       | Description                                                                                                                                                                        |
| -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **showName?**              | <p><em>string</em></p><p>Controls if the user name is displayed next to the user image button.</p>                                                                                 |
| **signInUrl?**             | <p><em>string</em></p><p>The full URL or path to navigate to when the "Add another account" button is clicked.</p>                                                                 |
| **userProfileUrl?**        | <p><em>string</em><br><em></em>The full URL or path leading to the user management interface.</p>                                                                                  |
| **afterSignOutAllUrl?**    | <p><em>string</em></p><p>The full URL or path to navigate to after a signing out from all accounts (multi-session apps) or the currently active account (single-session apps).</p> |
| **afterSignOutOneUrl?**    | <p><em>string</em></p><p>The full URL or path to navigate to after a signing out from currently active account (multisession apps) .</p>                                           |
| **afterSwitchSessionUrl?** | <p><em>string</em></p><p>Full URL or path to navigate to after a successful account change (multi-session apps).</p>                                                               |

## Customization

The `<UserButton/>` component can be highly customized through the Instance settings on the [Theme page](https://dashboard.clerk.dev/last-active?path=customization/theme). This document will be updated soon with all necessary details.
