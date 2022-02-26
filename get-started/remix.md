---
description: Learn to install and initialize Clerk in a Remix application.
---

# Get started with Remix

## Overview

{% hint style="info" %}
Remix support is still in beta and the APIs presented below are subject to change. Please report any feedback, issues, or ideas [to support](https://clerk.dev/support) - we'd love to hear from you!&#x20;
{% endhint %}

Clerk is the easiest way to add authentication and user management to your [Remix](https://remix.run) application. This guide will walk you through the necessary steps to install and use Clerk in a new Next.js application.

After following this guide, you should have a working Remix app complete with:&#x20;

* Fully fledged sign in and sign up flows.
* Social Login with the vendors of your choice (Google, Twitter, etc)
* Secure password or passwordless authentication.
* A prebuilt user profile page.

## Before you start

You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). For more information, check out our [Setup your application](../popular-guides/setup-your-application.md) guide.

## Creating a new Remix app

If you haven't done so already, start by creating a new Remix application

```
npx create-remix@latest
```

For more information, you can reference the [Remix documentation](https://remix.run/docs/en/v1/tutorials/blog).

## Install @clerk/remix

Once you have a Remix app ready, you need to install the Clerk Remix SDK. This will give you access to our prebuilt Clerk Components and React hooks.

```
npm install @clerk/remix@next
```

Now, we need to retrieve your Backend API Key from the [Clerk Dashboard](https://dashboard.clerk.dev). Select your **Application**, **** and find the value on the Development instance Home page.

![Home page with Frontend API key highlighted](<../.gitbook/assets/home - frontend api key highlighted.png>)

Remix does not include a mechanism for setting environment variables, so developers must configure one themselves.&#x20;

For the purposes of this guide, we suggest using the `dotenv` npm package as described in Remix's [Guide to Environment Variables](https://remix.run/docs/en/v1/guides/envvars).

After you've installed `dotenv` and made the necessary change to your package.json `dev` script, create a `.env` file with your key:

```bash
# Add environment variable to .env.local file
# Replace [your-backend-api-key] with the actual Backend API key
echo "CLERK_API_KEY=[your-backend-api-key]" > .env
```

Clerk is now installed, but still needs to be initialized.  We can go ahead and start the development server:

```
npm run dev
```

## Initialize Clerk in app/root.tsx

In Remix, `app/root.tsx` wraps your entire application in both server and browser contexts. Clerk requires modifications to this file so we can propagate the authentication state through your Remix routes.

### Set \`rootAuthLoader\`

First, let's define a loader in root.tsx.  Users who are not signed in will be redirected to your sign-in page:

```javascript
import { rootAuthLoader } from "@clerk/remix/ssr.server";
import { redirect } from "remix";

export const loader: LoaderFunction = (args) =>
  rootAuthLoader(args, ({ params, request, context, auth }) => {
    if(!auth.userId){
      return redirect("https://accounts.foo.bar.lcl.dev/sign-in");
    }
    // auth.userId will be set if there is a signed in user
    return { yourData: "here" };
  });
```

{% hint style="info" %}
Currently, your entire application must require authentication. This limitation will be resolved before February 11.
{% endhint %}

### Connect Clerk to React

Clerk provides a `ConnectClerk` helper to provide the authentication state to your React tree. This helper works with SSR out-of-the-box and follows the "higher-order component" paradigm.

First, update App so it is not exported as the default module:

```javascript
// Remove "export default" from in front of `function App() {`
function App() {
  ...
}
```

Then, wrap `App` with `ConnectClerk` and pass in your `frontendApi` variable from your instance dashboard:

```javascript
import { ConnectClerk } from "@clerk/remix";

export default ConnectClerk(App, {
  // Replace with your own value
  frontendApi: "clerk.foo.bar.lcl.dev",
});
```

{% hint style="info" %}
Currently, the frontendApi value must be hardcoded. We are exploring alternatives and will release a new version soon.
{% endhint %}

That's all! The rest of your application can now leverage authentication as needed.

## Using authentication state

Now that your `root.tsx` is configured, you can use the user's authentication information anywhere in your application.

### In loaders and actions

Use the `getAuth` helper to retrieve the user ID and session ID from a loader or action. Most commonly, developers only require the user ID, while the session ID my be helpful auditing use cases.

```javascript
import { getAuth } from "@clerk/remix/ssr.server";

const { userId, sessionId } = await getAuth(request);
```

### In React

In React, the same information is available via the `useAuth` hook:

```javascript
import { useAuth } from "@clerk/remix";

const { userId, sessionId } = useAuth();
```

A complete reference to our React helpers is [available here](broken-reference).

### Demo repository

[Clerk Remix Demo Repository](https://github.com/nikosdouvlis/clerk-remix-demo)
