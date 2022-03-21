---
description: Learn how to install Clerk in a RedwoodJS application
---

# Get started with RedwoodJS

### Overview

RedwoodJS is an opinionated, full-stack web application framework focused on providing an integrated developer experience from build to deployment. It follows standard-based conventions with an eye towards the modern architecture of Jamstack and serverless apps.

Although Redwood contains a built-in authentication system (dbAuth), it requires that you manage the user database yourself and does not provide features like social SSO, passwordless authentication, and multi-session management. Clerk provides these features and more without the hassle of managing your own user auth service.

The Clerk integration with Redwood enables you to focus on building the frontend and GraphQL API aspects of your application without needing to add database fields to the Users table and maintain a secret key to encrypt the session cookie. It’s nice to not have to be concerned with these things.

This guide will walk you through the necessary steps to integrate Clerk as the external authentication provider for RedwoodJS.

### Getting started

The first step is to create a new Clerk application from your [Clerk Dashboard](https://dashboard.clerk.dev) if you haven’t done so already. You can choose whichever authentication strategy and social login providers you prefer. For more information, check out our [Set up your application](../popular-guides/setup-your-application.md) guide.

Now, we need to retrieve your API keys from the [API Keys page](https://dashboard.clerk.dev/last-active?path=api-keys).

In your Redwood app directory, create a `.env` file (if one does not currently exist) and set the following environment variables to the respective values from your Clerk dashboard:

```bash
CLERK_API_KEY=<YOUR_BACKEND_API_KEY>
CLERK_FRONTEND_API_URL=<YOUR_FRONTEND_API_KEY>
```

In order for the Frontend API to be accessible to the Web side in production, you need to add `CLERK_FRONTEND_API_URL` to the `includeEnvironmentVariables` array in the `redwood.toml` file:

```jsx
// redwood.toml

[web]
  includeEnvironmentVariables = ['CLERK_FRONTEND_API_URL']
```

### Set up auth

The next step is to run a [Redwood CLI command](https://redwoodjs.com/docs/authentication.html#clerk) to install the required packages and generate some boilerplate code:

```bash
yarn rw setup auth clerk --force
```

**Note**: The `--force` flag is needed to overwrite the existing dbAuth logic.

You should see terminal output similar to the following:

```bash
✔ Generating auth lib...
    ✔ Successfully wrote file `./api/src/lib/auth.js`
  ✔ Adding auth config to web...
  ✔ Adding auth config to GraphQL API...
  ✔ Adding required web packages...
  ✔ Adding required api packages...
  ✔ Installing packages...
  ✔ One more thing...

  You will need to add two environment variables with your Clerk URL and API key.
  Check out web/src/App.{js,tsx} for the variables you need to add.
  See also: <https://redwoodjs.com/docs/authentication#clerk>
```

If you already followed the instructions to add your environment variables, you should be all set. If you didn’t, add them now.

In your code editor of choice, open up `web/src/App.tsx` (or `App.js` if you’re not using TypeScript).

Wrap the Redwood `<AuthProvider type="clerk">` component with `<ClerkAuthProvider>`

```jsx
const App = () => (
  <FatalErrorBoundary page={FatalErrorPage}>
    <RedwoodProvider titleTemplate="%PageTitle | %AppTitle">
      <ClerkAuthProvider>
        <AuthProvider type="clerk">
          <RedwoodApolloProvider>
            <Routes />
          </RedwoodApolloProvider>
        </AuthProvider>
      </ClerkAuthProvider>
    </RedwoodProvider>
  </FatalErrorBoundary>
)
```

### Add Clerk components and hooks

Now that Clerk auth is set up, restart the dev server with `yarn rw dev`. If you had the dev server running, it must be restarted to read the newly added environment variables.

You can now access Clerk functions through the Redwood `useAuth()` hook or you can use the Clerk components directly.

For example, you can write:

```jsx
import { useAuth } from '@redwoodjs/auth'

const HomePage = () => {
  const { currentUser, isAuthenticated, logIn, logOut } = useAuth()

  return (
    <>
      {isAuthenticated ? (
        <button onClick={logOut}>Log out</button>
      ) : (
        <button onClick={logIn}>Log in</button>
      )}
      {isAuthenticated && <h1>Hello {currentUser.firstName}</h1>}
    </>
  )
}

export default HomePage
```

The `isAuthenticated` property checks if there is an active user session. Clicking the “Log in” button opens a modal window allowing you to sign in with the authentication methods chosen when you set up the project in the Clerk dashboard.

Since Clerk React is installed, you can use the Clerk components instead:

```jsx
import { SignInButton, UserButton } from '@clerk/clerk-react'
import { useAuth } from '@redwoodjs/auth'

const HomePage = () => {
  const { currentUser, isAuthenticated } = useAuth()

  return (
    <>
      {isAuthenticated ? (
        <UserButton afterSignOutAllUrl={window.location.href} />
      ) : (
        <SignInButton mode="modal" />
      )}
      {isAuthenticated && <h1>Hello {currentUser.firstName}</h1>}
    </>
  )
} 

export default HomePage
```

Clerk makes it super easy to add in these authentication components. There are further options for customization available as well.

### Next steps

* Get started with our official [clerk-redwood-starter](https://github.com/clerkinc/clerk-redwood-starter) repo
* Check out our [RedwoodJS Blog Tutorial with Clerk](https://clerk.dev/tutorials/redwoodjs-blog-tutorial-with-clerk)
* Learn more about the [Redwood Authentication API](https://redwoodjs.com/docs/authentication#api)
* Update the Sign In and Sign Up pages with [mounted Clerk components](../components/sign-in/sign-in.md)
* Get support or at least say hi in our [Discord channel](https://discord.com/invite/b5rXHjAg7A) 👋
