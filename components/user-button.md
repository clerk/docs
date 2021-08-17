---
description: >-
  A familiar component that allows users to manage their account, switch
  accounts, or sign out
---

# &lt;UserButton /&gt;

## Overview

Originally popularized by Google, users have come to expect that little photo of themselves in the top-right of the page – it’s the access point to manage their account, switch accounts, or sign out. The `<UserButton/>` component is used to render this familiar user button UI. It renders a clickable user avatar - when clicked, the full UI opens as a popup. 



![The default &amp;lt;UserButton/&amp;gt; component](../.gitbook/assets/user-button.png)

Clerk is the only provider with multi-session support, allowing users to sign into multiple accounts at once and switch between them. For multisession apps, the `<UserButton/>` automatically supports instant account switching, without the need of a full page reload. For more information, you can check out the [Multi-session applications](../popular-guides/popular-guides-multi-session-applications.md) guide.

Control the look and feel of the `<UserButton/>` component and match it to your using the [Theme Settings](../popular-guides/setup-your-application.md#theme), [theming props](user-profile.md#customization) or [plain CSS](user-profile.md#customization).

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../reference/clerk-react/installation.md) or [ClerkJS](../reference/clerkjs/installation.md) before running the snippets below.
{% endhint %}

### Mounting in your app

Once you set up the desired functionality and look and feel for the `<UserButton/>` component, all that's left is to render it inside your page. The default rendering is simple but powerful enough to cover most use-cases. The  theme configuration \(look and feel\) that you've set up in your [Clerk Dashboard](https://dashboard.clerk.dev) will work out of the box.

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

<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>routing?</b>
      </td>
      <td style="text-align:left">
        <p><em>RoutingStrategy</em>
        </p>
        <p>The routing strategy for your pages. Supported values are:</p>
        <ul>
          <li><b>hash: </b>(default) Hash based routing.</li>
          <li><b>path</b>: Path based routing.</li>
          <li><b>virtual</b>: Virtual based routing.</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>path?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The root URL where the component is mounted on.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>signInUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The full URL or path to navigate to when the &quot;Add another account&quot;
          button is clicked.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>userProfileUrl?</b>
      </td>
      <td style="text-align:left"><em>string<br /></em>The full URL or path leading to the user management
        interface.</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSignOutAllUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The full URL or path to navigate to after a signing out from all accounts
          (multi-session apps) or the currently active account (single-session apps).</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSignOutOneUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The full URL or path to navigate to after a signing out from currently
          active account (multisession apps) .</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSwitchSessionUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or path to navigate to after a successful account change (multi-session
          apps).</p>
      </td>
    </tr>
  </tbody>
</table>

## Customization

The `<UserButton/>` component can be highly customized through the Instance settings in the [Clerk Dashboard](https://dashboard.clerk.dev/). This document will be updated soon with all necessary details.

