---
description: Learn to install and initialize Clerk in a new Web3 application.
---

# Get started with Web3

### Overview

Clerk is the easiest way to add Web3 authentication to your application. This guide will walk you through the necessary steps to install and use Web3 authentication in a new Next.js application.

After following this guide, you should have a working Next.js app complete with:&#x20;

* Web3 authentication
* Session management
* Self-serve multifactor authentication
* Self-serve profile enrichment

### Before you start

You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). Configure the application to use "Sign in with Metamask"

![Choose the "Sign in with Metamask" authentication strategy](<../.gitbook/assets/Screen Shot 2022-01-21 at 2.57.32 PM.png>)

### Creating a new Next.js app

Start by creating a new Next.js application - this is usually done using the CLI:

{% tabs %}
{% tab title="npm" %}
```bash
npx create-next-app
```
{% endtab %}

{% tab title="yarn" %}
```
yarn create next-app
```
{% endtab %}
{% endtabs %}

If you wish to use TypeScript, just add `--typescript` to the commands above. Clerk is written in TypeScript, so it works out of the box without any extra configuration. For more information, you can reference the [Next.js documentation](https://nextjs.org/docs/api-reference/create-next-app).

### Installing Clerk

Once you have a Next.js app ready, you need to install the [Clerk React SDK](../reference/clerk-react/). This will give you access to our [prebuilt Clerk Components](broken-reference) and React hooks.

{% tabs %}
{% tab title="npm" %}
```bash
# Navigate to your application's root directory
cd yourapp

# Install the clerk/nextjs package 
npm install @clerk/nextjs
```
{% endtab %}

{% tab title="yarn" %}
```bash
# Navigate to your application's root directory
cd yourapp

# Install the clerk/nextjs package 
yarn add @clerk/nextjs
```
{% endtab %}
{% endtabs %}

Now, we need to set the `CLERK_FRONTEND_API` environment variable. Go to the [API Keys page](https://dashboard.clerk.dev/last-active?path=api-keys) and copy the **Frontend API Key** field.

Then, create a file named `.env.local` in your application root. Any variables inside this file with the `NEXT_PUBLIC_` prefix will be accessible in your Next.js code via `process.env.NEXT_PUBLIC_VAR_NAME`. Create a `NEXT_PUBLIC_CLERK_FRONTEND_API` variable and set it to the `Frontend API` you copied earlier:

```bash
# Add environment variable to .env.local file
# Replace [your-frontend-api] with the actual Frontend API key
echo "NEXT_PUBLIC_CLERK_FRONTEND_API=[your-frontend-api]" > .env.local
```

Clerk is now successfully installed   🎉 &#x20;

To run your app, start the development server and navigate to [https://localhost:3000](http://localhost:3000).

{% tabs %}
{% tab title="npm" %}
```bash
npm run dev
```
{% endtab %}

{% tab title="yarn" %}
```bash
yarn dev
```
{% endtab %}
{% endtabs %}

For more details, consult the [Clerk React installation](../reference/clerk-react/installation.md) page.

### Adding \<ClerkProvider />

Clerk requires your application to be wrapped in the [`<ClerkProvider/>`](../reference/clerk-react/clerkprovider.md) component. In Next.js, we add this in `pages/_app.js`.&#x20;

{% code title="pages/_app.js" %}
```jsx
import { ClerkProvider } from '@clerk/nextjs';

function MyApp({ Component, pageProps }) {

  return (
    // Wrap your entire app with ClerkProvider
    <ClerkProvider>
      <Component {...pageProps} />
    </ClerkProvider>
  );
}

export default MyApp;
```
{% endcode %}

`<ClerkProvider/>` from the `@clerk/nextjs` package is already configured to using the same routing logic with Next.js.  This makes sure that navigating between subpages and when navigating to callback URLs, routing remains consistent.

Your app is now configured! 🎉  &#x20;

Next, let's see how you can use Clerk to require authentication before navigating to a protected page.

### Requiring Web3 authentication

The easiest way to require authentication before showing a protected page is to use our Control Components:       &#x20;

* [`<SignedIn/>`](../components/signed-in.md): Renders its children only when a user is signed in.
* [`<SignedOut/>`](../components/signed-out.md): Renders its children only when there's no active user.

As a child of `<SignedOut/>`, we will include our [`<SignInWithMetamaskButton/>`](../components/sign-in-with-metamask-button.md) so the user can trigger the Metamask authentication flow.

The following example shows you how to compose our flexible [Control Components](../components/control-components/) to build authentication flows that match your needs. Please note that you don't need to use any additional APIs, everything shown below is just Javascript.

{% code title="pages/_app.js" %}
```jsx
import '../styles/globals.css'
import { ClerkProvider, SignedIn, SignedOut, SignInWithMetamaskButton } from '@clerk/nextjs';
import { useRouter } from 'next/router';

//  List pages you want to be publicly accessible, or leave empty if
//  every page requires authentication. Use this naming strategy:
//   "/"              for pages/index.js
//   "/foo"           for pages/foo/index.js
//   "/foo/bar"       for pages/foo/bar.js
//   "/foo/[...bar]"  for pages/foo/[...bar].js
const publicPages = [];

function MyApp({ Component, pageProps }) {
  // Get the pathname
  const { pathname } = useRouter();

  // Check if the current route matches a public page
  const isPublicPage = publicPages.includes(pathname);

  // If the current route is listed as public, render it directly
  // Otherwise, use Clerk to require authentication
  return (
    <ClerkProvider>
      {isPublicPage ? (
        <Component {...pageProps} />
      ) : (
        <>
          <SignedIn>
            <Component {...pageProps} />
          </SignedIn>
          <SignedOut>
            <SignInWithMetamaskButton />
          </SignedOut>
        </>
      )}
    </ClerkProvider>
  );
}

export default MyApp;
```
{% endcode %}

Visit [https://localhost:3000](https://localhost:3000) to see your page. The home `"/"` page is not listed in the `publicPages` array, so you'll immediately see the Sign in with Metamask button - go ahead and try it!

### Hello, world!

After signing in with Metamask, you'll be presented with the Next.js default start screen.  Let's update `pages/index.js` to retrieve add a `UserButton` and retrieve the current users Metamask address:

{% code title="pages/index.js" %}
```jsx
import { UserButton, useUser } from "@clerk/nextjs";

export default function Home() {
  const { primaryWeb3Wallet } = useUser();
  return (
    <>
      <header
        style={{
          display: "flex",
          justifyContent: "space-between",
          padding: "1rem",
        }}
      >
        <span>My first DApp</span>
        <UserButton />
      </header>
      <main style={{ padding: "1rem" }}>
        <p>Address: {primaryWeb3Wallet.web3Wallet}</p>
      </main>
    </>
  );
}
```
{% endcode %}

### Accessing the Web3 address from the backend

#### Set CLERK\_API\_KEY

Go to the Clerk Dashboard, select your Application, copy the Backend API Key field from the Development instance Home page and add it as `CLERK_API_KEY` in the **.env.local** in your application root.

![Home page with Backend API key highlighted](<../.gitbook/assets/home - backend api key highlighted.png>)

After the addition, your **.env.local** file will have two lines, and look something like this (except with your own unique values):

{% code title=".env.local" %}
```jsx
NEXT_PUBLIC_CLERK_FRONTEND_API=clerk.foo12.bar34.lcl.dev
CLERK_API_KEY=test_asdf1234
```
{% endcode %}

#### Add authentication to a serverless API route

A new Next.js app comes with a sample API route in **/pages/api/hello.js**. Let's modify it to retrieve the user's Web3 address:

```javascript
import { requireSession } from "@clerk/nextjs/api";

export default requireSession(async (req, res) => {
  try {
    // SDK helpers for this request are coming soon!
    const user = await fetch(
      `https://api.clerk.dev/v1/users/${req.session.userId}`,
      {
        headers: { Authorization: `Bearer ${process.env.CLERK_API_KEY}` },
      }
    ).then((r) => r.json());

    // Retrieve the address
    const address = user.web3_wallets[0].web3_wallet;
    res.status(200).json({ address: address });
  } catch (e) {
    console.log(e);
    res.status(500).json({
      error: "Something went wrong retrieving the web3wallet from Clerk",
    });
  }
});
```

### Next steps

You now have Web3 authentication on a new Next.js application, and you know how to retrieve the users Web3 address from both the frontend and the backend.

* When you're ready, learn how to [deploy your app to production](../popular-guides/production-setup.md).
* Get support or at least say hi in our [Discord channel](https://discord.com/invite/b5rXHjAg7A) 👋. Web3 is a new feature and we'd love to hear your feedback
