---
description: Support for server-side rendering with Clerk is now in open beta.
---

# SSR Beta

Thank you for your interest in server-side rendering with Clerk! Our team is incredibly excited to open the SSR beta to the public and eager to hear your feedback. We’d appreciate if you can [connect with us in Discord](https://discord.com/invite/b5rXHjAg7A) to share your feedback, or otherwise contact us by email at support@clerk.dev.

### SSR objectives

We approached SSR with two high-level goals:

1. Expose the user’s authentication state to the framework’s data-loading mechanism (for example, `getServerSideProps` for Next.js, `loader` for Remix)
2. Expose the user’s authentication state to Clerk’s React hooks, so user data is available when React renders the initial HTML on the server

### New to Clerk?

This documentation is intended for developers who have a baseline understanding of Clerk’s functionality. If you’re new to Clerk and need SSR, don’t worry! We recommend trying our [Get Started with Next.js](https://docs.clerk.dev/get-started/nextjs) guide first.

### Demo repo

See SSR in action with our [demo repo](https://github.com/clerkinc/clerk-nextjs-ssr-demo) - the only setup step is configure `.env.local`  with your settings from your own Clerk instance.

### NPM package

SSR support is available by installing `@clerk/nextjs@next` with either `npm` or `yarn`.

### Introduction of the “auth” context

To support SSR, Clerk is adding an “auth” context to both React components and framework-specific data loading functions. This context contains the minimal information needed for data-loading:

* **userId:** The ID of the active user, or `null` when signed out. In data-loaders, this is often the only piece of information needed to securely retrieve the data associated with a request.
* **sessionId:** The ID of the active session, or `null` when signed out. This is primarily used in audit logs to enable device-level granularity instead of user-level.
* **getToken(): R**etrieves the signed JWT associated with the request, or throws when signed out. Used for relaying the JWT to other services which must reverify the signature. _Please note: custom JWT templates like `getToken({template: "hasura"})` are not available yet._

Critically, **Clerk uses stateless JWTs to authenticate SSR in under 1 millisecond** without a database query or other network request.

#### Data-loaders: `context.auth`

```jsx
// Next.js
import { withServerSideAuth } from "@clerk/nextjs/ssr";

export const getServerSideProps = withServerSideAuth(
  async (context) => {
    const { userId, sessionId, getToken } = context.auth;
    // Load any data your application needs and pass to props
    return { props: {} };
  }
);
```

#### React components: `useAuth()`

For `useAuth()` to work during SSR, the state from the data-loader must be passed to `<ClerkProvider>`. This is accomplished simply in the application’s entrypoint:

```jsx
// Next.js
import type { AppProps } from "next/app";
import { ClerkProvider } from "@clerk/nextjs";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ClerkProvider {...pageProps}>
      <Component {...pageProps} />
    </ClerkProvider>
  );
}

export default MyApp;
```

Now, `useAuth()` can be used throughout the application during SSR. Since we updated `<ClerkProvider>`, the `isLoaded` state will be `false` during SSR. It’s only set to `true` during client-side rendering.

```jsx
import { useAuth } from '@clerk/nextjs';

const Page = () => {
  const { isLoaded, userId, sessionId, getToken } = useAuth();
  
  // Handle these cases in case the user signs out while on the page.
  if (!isLoaded || !userId) {
    return null;
  }

  return <div>Hello, {userId}</div>;
};
```

Note: if you’re currently using `useSession()` only for `getToken`, we recommend switching to `useAuth()`.

### Retrieving the User resource during SSR

If desired, Clerk’s complete User objects can also be retrieved during SSR by passing a flag to the data-loader. Since these objects require network requests to retrieve, they are not available by default.

In practice, we expect developers will do this when they want to render user profile data from the server, like the user’s first name.

#### Data-loader: `{loadUser: true}`

```jsx
// Next.js
import { withServerSideAuth } from '@clerk/nextjs/ssr';

export const getServerSideProps = withServerSideAuth(
  context => {
    const { user } = context;
    // Load any data your application needs and pass to props
    return { props: {} };
  },
  { loadUser: true },
);
```

#### React components `useUser()`

For `useUser()` to work during SSR, the state from the data-loader must be passed to `<ClerkProvider>`. This is accomplished simply in the application’s entrypoint.

_If you’ve already done this for `useAuth()` above, no further changes are needed._

```jsx
// Next.js
import type { AppProps } from 'next/app';
import { ClerkProvider } from '@clerk/nextjs';

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ClerkProvider {...pageProps}>
      <Component {...pageProps} />
    </ClerkProvider>
  );
}

export default MyApp;
```

Now, `useUser()` can be used throughout the application during SSR. Since we updated `<ClerkProvider>`, the `isLoaded` state will be `false` during SSR. It’s only set to `true` during client-side rendering.

If you’ve used Clerk in the past, you’ll notice the the addition of the `isLoaded` and `isSignedIn` states. If the `loadUser` or `loadSession` flags are passed to the data-loader, the objects will be preloaded during SSR.

```jsx
import { useUser } from '@clerk/nextjs';

const Page = () => {
  const { isLoaded, isSignedIn, user } = useUser();

  if (isLoaded || !isSignedIn) {
    return null;
  }

  return <>Hello, {user.firstName}</>;
};
```

### Retrieving the Session resource during SSR

The Session resource works identically to the User resource. There is a corresponding `{loadSession: true}` flag as well as a `useSession()` with the same signature.

### Typescript 4.6

For better type support, we highly recommend updating to [Typescript 4.6](https://devblogs.microsoft.com/typescript/announcing-typescript-4-6-beta/), which is also in public beta.

For typed applications, our new hooks benefit significantly from Typescript's new [Control Flow Analysis for Dependent Parameters](https://devblogs.microsoft.com/typescript/announcing-typescript-4-6-beta/#control-flow-analysis-for-dependent-parameters).

### Upgrading from @2 to @next

If you are upgrading an existing application, you will need to make the following changes

* `useUser()` and `useSession()` have changed signatures (see above)
* `<RedirectToSignIn/>` and `<RedirectToSignUp/>` now return users to the same page they came from by default, instead of to the URL in your dashboard.
* If you have not upgraded to “Auth v2", you will need to. Auth v1 is not compatible with SSR. If your application was created after October 21, 2021, you are already using Auth v2. [Upgrade information is available here.](../main-concepts/auth-v2.md)

You can find a detailed migration guide below:

#### useUser()

| Old API                   | New API                                             |
| ------------------------- | --------------------------------------------------- |
| `const user = useUser();` | `const { isLoaded, isSignedIn, user } = useUser();` |

#### useSession()

| Old API                         | New API                                                                                                                                                                                                                         |
| ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `const session = useSession();` | <p><code>const { isLoaded, isSignedIn, session } = useSession();</code></p><p><br>or use the newly introduced <code>useAuth()</code> hook:<br><br><code>const { isLoaded, sessionId, userId, getToken } = useAuth();</code></p> |

#### \<RedirectToSignIn /> and \<RedirectToSignUp/>

| Old API                          | New API                                                                                              |
| -------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `<RedirectToSignIn returnBack/>` | <p><code>&#x3C;RedirectToSignIn /></code><br><code></code>(redirects to current page by default)</p> |
| `<RedirectToSignIn returnBack/>` | <p><code>&#x3C;RedirectToSignUp /></code><br><code></code>(redirects to current page by default)</p> |

#### \<SignIn />

| Old API       | New API          |
| ------------- | ---------------- |
| `afterSignIn` | `afterSignInUrl` |
| `signUpURL`   | `signUpUrl`      |

#### \<SignUp />

| Old API       | New API          |
| ------------- | ---------------- |
| `afterSignUp` | `afterSignUpUrl` |
| `signInURL`   | `signInUrl`      |

#### \<UserButton />

| Old API              | New API                 |
| -------------------- | ----------------------- |
| `afterSwitchSession` | `afterSwitchSessionUrl` |
| `userProfileURL`     | `userProfileUrl`        |
| `signInURL`          | `signInUrl`             |
| `afterSignOutAll`    | `afterSignOutAllUrl`    |
| `afterSignOutOne`    | `afterSignOutOneUrl`    |

#### SignIn & SignUp

| Old API                                                          | New API                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `createMagicLinkFlow({ callbackUrl })`                           | `createMagicLinkFlow({ redirectUrl })`                           |
| `authenticateWithRedirect({ callbackUrl, callbackUrlComplete })` | `authenticateWithRedirect({ redirectUrl, redirectUrlComplete })` |

#### Email

| Old API                                | New API                                |
| -------------------------------------- | -------------------------------------- |
| `createMagicLinkFlow({ callbackUrl })` | `createMagicLinkFlow({ redirectUrl })` |

#### User

| Old API                          | New API                                           |
| -------------------------------- | ------------------------------------------------- |
| `User.getToken("template-name")` | `Session.getToken({ template: "template-name" })` |

#### Session

| Old API            | New API |
| ------------------ | ------- |
| `Session.revoke()` | -       |

#### window.Clerk

| Old API                        | New API                                                                                   |
| ------------------------------ | ----------------------------------------------------------------------------------------- |
| `Clerk.redirectToSignIn(true)` | <p><code>Clerk.redirectToSignIn()</code></p><p>(redirects to current page by default)</p> |
| `Clerk.redirectToSignUp(true)` | <p><code>Clerk.redirectToSignUp()</code></p><p>(redirects to current page by default)</p> |
