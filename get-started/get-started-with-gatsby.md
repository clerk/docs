---
description: Learn to install and initialize Clerk in a new Gatsby application.
---

# Get started with Gatsby

## Overview

Clerk is the easiest way to add authentication and user management to your [Gatsby](https://www.gatsbyjs.com) application. This guide will walk you through the necessary steps to install and use Clerk in a new [Gatsby](https://www.gatsbyjs.com) application.&#x20;

After following this guide, you should have a working [Gatsby](https://www.gatsbyjs.com) app complete with:&#x20;

* Fully fledged sign in and sign up flows.
* Google Social Login.
* Secure email/password authentication.
* A prebuilt user profile page.

## Before you start

You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). For more information, check out our [Setup your application](../popular-guides/setup-your-application.md) guide.

## Installing the plugin

Once you have a Gatsby app ready, you need to install the [Clerk React SDK](../reference/clerk-react/) and [gatsby-plugin-clerk](https://www.gatsbyjs.com/plugins/gatsby-plugin-clerk). This will give you access to our [prebuilt Clerk Components](broken-reference) and [React hooks](../reference/clerk-react/).

{% tabs %}
{% tab title="npm" %}
```bash
# Navigate to your application's root directory
cd your_app

# Install the necessary Clerk packages
npm install gatsby-plugin-clerk @clerk/clerk-react
```
{% endtab %}

{% tab title="yarn" %}
```bash
# Navigate to your application's root directory
cd your_app

# Install the necessary Clerk packages
yarn add gatsby-plugin-clerk @clerk/clerk-react
```
{% endtab %}
{% endtabs %}

As a next step, you'll need the `frontendApi` key of your Clerk application. To find it, go to the [API Keys page](https://dashboard.clerk.dev/last-active?path=api-keys) and copy the **Frontend API Key** field.

Now, let's configure the plugin on `gatsby-config.js.`

```bash
// gatsby-config.js

module.exports = {
  plugins: [
    {
      resolve: 'gatsby-plugin-clerk',
      options: {
        frontendApi: <YOUR_FRONTEND_API_KEY>
      }
    }
  ]
}
```

Clerk is now successfully installed   🎉 &#x20;

From here onwards, everything should work just the same. You can start using components like `SignedIn` and `SignedOut` from the root of your app.

{% code title="src/pages/index.js" %}
```jsx
import React from 'react'
import {
  SignIn,
  SignedIn,
  SignedOut,
  UserButton
  } from '@clerk/clerk-react'

export default function IndexPage() {
  return (
    <>
      <SignedIn>
        <UserButton />
      </SignedIn>
      <SignedOut>
        <SignIn />
      </SignedOut>
    </>
  )
}

```
{% endcode %}

And that's it, in just a few steps, we added easy and secure authentication with beautiful and complete user management to your Gatsby app.

## Clerk + Gatsby starter repository

To make it even easier for you, we went ahead and created a [Clerk + Gatsby starter repository](https://github.com/clerkinc/clerk-gatsby-starter). It has Clerk integrated with [Gatsby's default starter](https://github.com/gatsbyjs/gatsby-starter-default).

Inside `src/api` you can also find the new Gatsby Functions in action. We added a couple of examples there, so you know how to use Gatsby's serverless functions with Clerk's backend API.

Fork it, clone it and build it!
