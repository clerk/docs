---
description: Full-featured UI for signing users in your application.
---

# \<SignIn />

## Overview

The `<SignIn/>` component renders a UI for signing in users. Most of the times, the `<SignIn/>` component is all you need for completing sign ins. It supports any authentication scheme, from [Email/password authentication](../../popular-guides/email-and-password.md), and [Passwordless](../../popular-guides/passwordless-authentication.md), to [Social Login (OAuth)](../../popular-guides/social-login-oauth.md) and [Multi-factor verification](../../popular-guides/multi-factor-authentication.md).

The contents and functionality of the `<SignIn/>` component are controlled for the most part by the instance settings you specify in your [Clerk Dashboard](https://dashboard.clerk.dev). Your instance settings also allow for customization of the look and feel of the `<SignIn/>` component.

You can further customize your `<SignIn/>` component by passing additional [properties](sign-in.md#props) at the time of rendering.

Here's an example of what the component looks like once it's rendered.

![](<../../.gitbook/assets/sign-in (1).png>)

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) or [ClerkJS](../../reference/clerkjs/installation.md) before running the snippets below.
{% endhint %}

Once you set up the desired functionality and look and feel for the `<SignIn/>` component, all that's left is to render it inside your page. The default rendering is simple but powerful enough to cover most use-cases. The authentication and display (look and feel) configuration that you've set up in your [Clerk Dashboard](https://dashboard.clerk.dev) will work out of the box.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { SignIn } from "@clerk/react";

// Render the sign in component in your
// custom sign in page.
function SignInPage() {
  return (
    <SignIn />
  );
}
```
{% endtab %}

{% tab title="ClerkJS" %}
```javascript
// Mount the sign in component inside the HTML element
// with id "sign-in".
window.Clerk.mountSignIn(
  document.getElementById("sign-in")
);

// Open the sign in component as a modal.
window.Clerk.openSignIn();
```
{% endtab %}
{% endtabs %}

### Routing & Paths

The mounted `<SignIn/>` component uses hash-based routing by default. As the sign in flow progresses, the hash portion of the URL will update to reflect the current step. An example of such a URL would be `example.com/sign-in#/factor-one`.

You can instead use path-based routing with some additional configuration. With path-based routing, the above URL would become`example.com/sign-in/factor-one`.

There are two props that control the routing type and the path. These are called `routing` and `path`. You can read more about them in the [Props section](sign-in.md#props) of this document.

Below is an example that uses path-based routing to mount the `<SignIn/>` component under the "/sign-in" path. The [Clerk React](../../reference/clerk-react/) snippet uses the popular [React Router](https://reactrouter.com) library, but it can be easily adapted to use the  routing library of your choice. We've also added an example for [Next.js](https://nextjs.org), which showcases integration with Next.js routing.&#x20;

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { 
  BrowserRouter as Router,
  Route,
  Switch,
  useHistory 
} from 'react-router-dom';
import { ClerkProvider, SignIn } from '@clerk/clerk-react';

function App() {
  // Get the navigate/push method from
  // your router library. Here we'll use
  // react-router-dom.
  const { push } = useHistory();

  return (
    <Router>
      // Pass the push method to ClerkProvider
      <ClerkProvider 
        frontendApi="clerk.[your-domain].com" 
        navigate={(to) => push(to)}
      >
        <Switch>
          // Declare a /sign-in route
          <Route path="/sign-in">
            // Mount the SignIn component under "/sign-in".
            // The routing is set to path-based.
            <SignIn routing="path" path="/sign-up" />
          </Route>
        </Switch>
      </ClerkProvider>
    </Router>
  );
}

export default App;
```
{% endtab %}

{% tab title="Clerk Next" %}
```jsx
// In _app.jsx:
// Your usual NextJS root component, wrapped by ClerkProvider
import { ClerkProvider } from '@clerk/clerk-react';
import { useRouter } from 'next/router';

function MyApp({ Component, pageProps }) {
  // Get the navigate/push method from
  // the NextJS router
  const { push } = useRouter();

  return (
    // Pass the push method to ClerkProvider
    <ClerkProvider 
      frontendApi="clerk.[your-domain].com"
      navigate={(to) => push(to)}
    >
      <Component {...pageProps} />
    </ClerkProvider>
  );
}

export default MyApp;


// In pages/sign-in/[[..index]].jsx:
// This is your catch all route that renders the SignIn 
// component
import { SignIn } from '@clerk/clerk-react';

export default function SignInPage() {
  // Mount the SignIn component under "/sign-up". 
  // The routing is set to path-based.
  return <SignIn routing="path" path="/sign-in" />;
}
```
{% endtab %}

{% tab title="ClerkJS" %}
```javascript
<html>
<body>
    <div id="sign-in"></div>
    
    <script>
        const el = document.getElementById("sign-in");
        // Mount the pre-built Clerk SignIn component
        // inside an HTMLElement on your page. 
        window.Clerk.mountSignIn(el, {
            routing: "path",
            path: "/sign-up",
        });
    </script>
</body>
</html>
```
{% endtab %}
{% endtabs %}

### Override URLs

By default, the `<SignIn/>` component will use the [Clerk Hosted Pages](broken-reference) URL for sign ups. You can override this at runtime, by passing the `signUpURL` property to the component.

Similarly, you can override the redirect URL after successful sign ins by providing the `afterSignIn` property to the component.

Both URL property values can be either relative paths (like **/home**) or full URLs (like **https://my-domain.com/home**).

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { SignIn } from "@clerk/react";

// Root path points to your homepage and sign up URL
// is the full URL to your sign up page. These can be
// either relative paths or full URLs.
import { rootPath, signUpURL } from "src/routes";

// Render the sign in component in your custom sign in 
// page overriding the pre-configured URLs.
function SignInPage() {
  return (
    <SignIn 
      afterSignIn={rootPath}
      signUpURL={signUpURL}
    />
  );
}
```
{% endtab %}

{% tab title="ClerkJS" %}
```javascript
// Mount the sign in component inside the
// HTML element with id "sign-in".
window.Clerk.mountSignIn(
  document.getElementById("sign-in"),
  {
    afterSignIn: "/home",
    signUpURL: "/sign-up",
  },
);

// Open the sign in component as a modal.
window.Clerk.openSignIn({
  afterSignIn: "/home",
  signUpURL: "/sign-up",  
});
```
{% endtab %}
{% endtabs %}

## Props

| Name                | Description                                                                                                                                                                                                                                                     |
| ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **routing?**        | <p><em>string</em></p><p>The routing strategy for your pages. Supported values are:</p><ul><li><strong>path</strong>: Path based routing.</li><li><strong>hash</strong>: Hash based routing.</li><li><strong>virtual</strong>: Virtual based routing.</li></ul> |
| **path?**           | <p><em>string</em></p><p>The root URL where the component is mounted on. This prop is ignored in <strong>hash</strong> based routing.</p>                                                                                                                       |
| **redirectUrl?**    | <p><em>string</em></p><p>Full URL or path to navigate after successful sign in or sign up.<br><br>The same as setting <code>afterSignInUrl</code> and <code>afterSignUpUrl</code> to the same value.</p>                                                        |
| **afterSignInUrl?** | <p><em>string</em></p><p>The full URL or path to navigate after a successful sign in.</p>                                                                                                                                                                       |
| **afterSignUpUrl?** | <p><em>string</em></p><p>The full URL or path to navigate after a successful sign up.</p>                                                                                                                                                                       |
| **signUpUrl?**      | <p><em>string</em></p><p>Full URL or path to the sign up page. Use this property to provide the target of the "Sign Up" link that's rendered.</p>                                                                                                               |

## Customization

The `<SignIn/>` component can be highly customized through the Instance settings in the [Clerk Dashboard](https://dashboard.clerk.dev). This document will be updated soon with all necessary details.
