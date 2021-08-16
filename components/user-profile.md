---
description: A full-featured account management component
---

# &lt;UserProfile /&gt;

## Overview

The `<UserProfile/>` component is used to render a beautiful, full-featured account management UI that allows users to manage their profile and security settings. 

![](../.gitbook/assets/user-profile.png)

The `<UserProfile/>` comes with built-in support for:

* Name and profile image management
* Email address and phone number management
* Social sign-in account management
* Multi-factor authentication management
* Trusted device management

Control the look and feel of the `<UserProfile/>` component and match it to your using the [Theme Settings](../popular-guides/setup-your-application.md#theme), [theming props](user-profile.md#customization) or [plain CSS](user-profile.md#customization).

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../reference/clerk-react/installation.md) or [ClerkJS](../reference/clerkjs/installation.md) before running the snippets below.
{% endhint %}

### Mounting in your app

Once you set up the desired functionality and look and feel for the `<UserProfile/>` component, all that's left is to render it inside your page. The default rendering is simple but powerful enough to cover most use-cases. The  theme configuration \(look and feel\) that you've set up in your [Clerk Dashboard](https://dashboard.clerk.dev) will work out of the box.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import {
  ClerkProvider,
  SignedIn,
  SignedOut,
  UserProfile,
  RedirectToSignIn,
} from "@clerk/clerk-react";
import { useHistory, BrowserRouter } from "react-router-dom";

function AppWithRoutes() {
  // Get the navigate/push method from
  // the router your app is using. For this
  // example we will use 'react-router-dom'.
  const { push } = useHistory();

  return (
    // Pass the push method to ClerkProvider
    <ClerkProvider
      frontendApi={"[your-frontend-api]"}
      navigate={(to) => push(to)}
    >
    {/* If a user is signed in, they will see
        the user profile, otherwise thet will get
        redirected to the sign in page */}
      <SignedIn>
        <UserProfile />
      </SignedIn>
      <SignedOut>
        <RedirectToSignIn />
      </SignedOut>
    </ClerkProvider>
  );
}

function App() {
  return (
    <BrowserRouter>
      <AppWithRoutes />
    </BrowserRouter>
  );
}

export default App;


```
{% endtab %}

{% tab title="ClerkJS" %}
```markup
<html>
<body>
    <div id="user-profile"></div>
    
    <script>
        const el = document.getElementById("user-profile");
        // Mount the pre-built Clerk UserProfile component
        // in an HTMLElement on your page. 
        window.Clerk.mountUserProfile(el);
    </script>
</body>
</html>
```
{% endtab %}
{% endtabs %}

### Using path-based routing

The mounted `<UserProfile/>` component uses hash-based routing by default: as the user goes through the different pages, the hash portion of the URL will update to reflect the current page \(eg: `example.com/#/account/first-name`\).

With additional configuration, the mounted component can use path-based routing instead \(eg: `example.com/account/first-name`\):

1. If using Clerk React, ensure your **ClerkProvider** component has [its **navigate** prop](../reference/clerk-react/installation.md#4-the-navigate-prop) configured.
2. Add the **path** and **routing** props to your **UserProfile** component. Set **path** to the path where the component renders

{% hint style="info" %}
When using path-based routing, the `<UserProfile/>` component must render on `path` and all of it's subpaths:

* For NextJS, use an [optional catch-all route](https://nextjs.org/docs/routing/dynamic-routes#optional-catch-all-routes) like `pages/user/[[...index]].js`
* For React Router, use a [wildcard path](https://reactrouter.com/web/api/Route/path-string-string) like `/user(/?.*)` or simply `/user`
{% endhint %}

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import {
  ClerkProvider, SignedIn, SignedOut, 
  UserProfile, RedirectToSignIn,
} from "@clerk/clerk-react";
import {
  useHistory, Switch, Route,
  BrowserRouter, Link,
} from "react-router-dom";

function AppWithRoutes() {
  // Get the navigate/push method from
  // the router your app is using. For this
  // example we will use 'react-router-dom'.
  const { push } = useHistory();

  return (
    // Pass the push method to ClerkProvider
    <ClerkProvider
      frontendApi={"[your-frontend-api]"}
      navigate={(to) => push(to)}
    >
      <Switch>
    {/* Define a /user route. 
        If a user is signed in, they will see
        the user profile, otherwise thet will get
        redirected to the sign in page */}
        <Route path="/user">
          <SignedIn>
            <UserProfile routing="path" path="/user" />
          </SignedIn>
          <SignedOut>
            <RedirectToSignIn />
          </SignedOut>
        </Route>
        {/* Home route that links to user profile */}
        <Route>
          <Link to="/user">
            <h1>Go to user profile</h1>
          </Link>
        </Route>
      </Switch>
    </ClerkProvider>
  );
}

function App() {
  return (
    <BrowserRouter>
      <AppWithRoutes />
    </BrowserRouter>
  );
}

export default App;

```
{% endtab %}

{% tab title="Clerk Next" %}
```jsx
// In _app.jsx:
// Your usual NextJS root component, wrapped by ClerkProvider
import '../styles/globals.css';
import { ClerkProvider } from '@clerk/clerk-react';
import { useRouter } from 'next/router';

function MyApp({ Component, pageProps }) {
  // Get the navigate/push method from
  // the NextJS router
  const { push } = useRouter();

  return (
    // Pass the push method to ClerkProvider
    <ClerkProvider 
        frontendApi={"[your-frontend-api]"} 
        navigate={to => push(to)}
    >
      <Component {...pageProps} />
    </ClerkProvider>
  );
}

export default MyApp;


// In pages/user/[[..index]].tsx
// This is your catch all route that renders the UserProfile component
import { UserProfile } from '@clerk/clerk-react';

export default function UserProfilePage() {
  // Finally, mount the UserProfile component under "/user" 🎉
  // Don't forget to set the "routing" and "path" props
  return <UserProfile routing='path' path='/user' />;
}

```
{% endtab %}

{% tab title="ClerkJS" %}
```markup
<html>
<body>
    <div id="user-profile"></div>
    
    <script>
        const el = document.getElementById("user-profile");
        // Mount the pre-built Clerk UserProfile component
        // in an HTMLElement on your page. 
        window.Clerk.mountUserProfile(el, {
            routing: 'path',
            path: '/user',
        });
    </script>
</body>
</html>
```
{% endtab %}
{% endtabs %}

For more information, see [Routing](../main-concepts/routing.md).

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
  </tbody>
</table>

## Customization

The `<UserProfile/>` component can be highly customized through the Instance settings in the [Clerk Dashboard](https://dashboard.clerk.dev/). This document will be updated soon with all necessary details.

