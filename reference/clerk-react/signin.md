---
description: >-
  The SignIn component is the fastest way to add a sign in flow to your
  application.
---

# SignIn

By default, Clerk hosts the **SignIn** component on the **accounts.\*** subdomain of your root domain. In production, this is configured with a CNAME in your DNS records.

The **SignIn** component can also be mounted within your application, or presented as a modal over your application.

## Mounting

To mount the **SignIn** component, first import it from our NPM package:

```jsx
import { SignIn } from "@clerk/clerk-react";
```

Then, place it anywhere in your JSX:

```jsx
<SignIn />
```

### Path-based routing

The mounted component uses hash-based routing by default. As the user signs in, the hash portion of the URL will update to reflect the current step.

With additional configuration, the mounted component can use path-based routing instead of hash-based routing.

First, ensure your **ClerkProvider** component has [its **navigate** prop](installation.md#4-the-navigate-prop) configured.

Then, add the **path** and **routing** props to your **SignUp** component. Set **path** to the path where the component renders:

```jsx
<SignIn routing="path" path="/sign-in" />
```

**Important:** When using path-based routing, the component must render on **path** and all of its subpaths:

* For Next.js, use an [optional catch-all route](https://nextjs.org/docs/routing/dynamic-routes#optional-catch-all-routes) like **pages/sign-in/\[\[...index\]\].js**
* For React Router, use a [wildcard path](https://reactrouter.com/web/api/Route/path-string-string) like **sign-in/\***

## Presenting as a modal

The **SignIn** component can also be presented as a modal. This is typically used on pages that show content whether or not the user is signed in.

```jsx
import { useClerk } from "@clerk/clerk-react";

const SignInButton = () => {
  const { openSignIn } = useClerk();
  return <button onClick={() => openSignIn({})}>Sign in</button>;
};
```

