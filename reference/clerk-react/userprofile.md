---
description: >-
  The UserProfile component is the fastest way to add user profile management to
  your application.
---

# UserProfile

By default, Clerk hosts the **UserProfile** component on your "accounts" subdomain, which you configure in production with a CNAME in your DNS records.

The **UserProfile** component can also be mounted within your application.

## Mounting

To mount the **UserProfile** component, first import it from our NPM package:

```jsx
import { UserProfile } from "@clerk/clerk-react";
```

Then, place it anywhere in your JSX:

```jsx
<UserProfile />
```

### Path-based routing

The mounted component uses hash-based routing by default. As the user navigates between different settings pages, the hash portion of the URL will update to reflect the current page.

With additional configuration, the mounted component can use path-based routing instead of hash-based routing.

First, ensure your **ClerkProvider** component has [its **navigate** prop](installation.md#4-the-navigate-prop) configured.

Then, add the **path** and **routing** props to your **UserProfile** component. Set **path** to the path where the component renders:

```jsx
<UserProfile routing="path" path="/user" />
```

**Important:** When using path-based routing, the component must render on **path** and all of its subpaths.

* For Next.js, use an [optional catch-all route](https://nextjs.org/docs/routing/dynamic-routes#optional-catch-all-routes) like **pages/user/\[\[...index\]\].js**
* For React Router, use a [wildcard path](https://reactrouter.com/web/api/Route/path-string-string) like **user/\***

