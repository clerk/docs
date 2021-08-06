---
description: >-
  The SignUp component is the fastest way to add a sign up flow to your
  application.
---

# SignUp

By default, Clerk hosts the **SignUp** component on the **accounts.\*** subdomain of your root domain. In production, this is configured with a CNAME in your DNS records.

The **SignUp** component can also be mounted within your application, or presented as a modal over your application.

## Mounting

To mount the **SignUp** component, first import it from our NPM package:

```jsx
import { SignUp } from "@clerk/clerk-react";
```

Then, place it anywhere in your JSX:

```jsx
<SignUp />
```

### Path-based routing

The mounted component uses hash-based routing by default. As the user signs up, the hash portion of the URL will update to reflect the current step.

With additional configuration, the mounted component can use path-based routing instead of hash-based routing.

First, ensure your **ClerkProvider** component has [its **navigate** prop](installation.md#4-the-navigate-prop) configured.

Then, add the **path** and **routing** props to your **SignUp** component. Set **path** to the path where the component renders:

```jsx
<SignUp routing="path" path="/sign-up" />
```

**Important:** When using path-based routing, the component must render on **path** and all of its subpaths:

* For Next.js, use an [optional catch-all route](https://nextjs.org/docs/routing/dynamic-routes#optional-catch-all-routes) like **pages/sign-up/\[\[...index\]\].js**
* For React Router, use a [wildcard path](https://reactrouter.com/web/api/Route/path-string-string) like **sign-up/\***

## Presenting as a modal

The **SignUp** component can also be presented as a modal. This is typically used on pages that show content whether or not the user is signed in.

```jsx
import { useClerk } from "@clerk/clerk-react";

const SignUpButton = () => {
  const { openSignUp } = useClerk();
  return <button onClick={() => openSignUp({})}>Sign up</button>;
};
```

