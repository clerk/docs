---
description: Learn how to make ClerkJS available in your project.
---

# Installation

## Overview

ClerkJS is our foundational JavaScript library for building user management and authentication. It  enables you to register, sign in, verify and manage users for your application using highly customizable flows. 

The primary integration path through ClerkJS is with [Clerk Components](../../main-concepts/clerk-components.md), as ClerkJS provides the foundation upon which all Clerk's UI offerings are built. However, integrating with ClerkJS directly gives you full freedom to use Clerk any way you see fit, dropping down to a lower level of primitives and commands for sign ins, sign ups and user profile management.

## Setting up ClerkJS

There are two ways you can include ClerkJS in your project. You can either [import the ClerkJS npm module](installation.md#using-clerkjs-as-a-module) or [load ClerkJS directly in your browser](installation.md#loading-clerkjs-as-a-script).

In either case, after the ClerkJS library is loaded, you need to call the [~~`Clerk.load()`~~](clerk.md#load) method to initialize ClerkJS, passing the desired options.

{% hint style="warning" %}
You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev) before you can set up ClerkJS. For more information, check out our [Setup your application](../../popular-guides/setup-your-application.md) guide.
{% endhint %}

### Using ClerkJS as a module

There's an ES module for ClerkJS, available under the [@clerk/clerk-js](https://www.npmjs.com/package/@clerk/clerk-js) npm package. Use `npm` or `yarn` to install the ClerkJS module.

```bash
# Install the package with npm.
npm install @clerk/clerk-js

# Alternative install method with yarn.
yarn add @clerk/clerk-js
```

Once you've installed the ClerkJS package, you first have to import it in your own code. The [Clerk object](clerk.md) constructor needs the [Frontend API](../frontend-api-reference/) URL as a parameter. You can find the URL in your **Instance Home** page in the [Clerk Dashboard](https://dashboard.clerk.dev).

```javascript
import Clerk from "@clerk/clerk-js";

const clerkFrontendApi = "clerk.[your-domain].com";
const clerk = new Clerk(clerkFrontendApi);
await clerk.load({
  // Set load options here...
});
```

### Loading ClerkJS as a script

You can also load ClerkJS with a `<script/>` tag in your website, straight from your [Frontend API](../frontend-api-reference/) URL. You can find the URL in your **Instance Home** page in the [Clerk Dashboard](https://dashboard.clerk.dev).

For security reasons, the ClerkJS library can only be loaded from the same domain. Make sure your website runs on the same domain as your Frontend API.

Add the following script in your site's `<body>` element.

```javascript
<script> 
  // Get this URL from the Clerk Dashboard.
  const frontendApi = "clerk.[your-domain].com";
  
  // Create a script that will be loaded asynchronously in 
  // your page. 
  const script = document.createElement('script');
  script.setAttribute('data-clerk-frontend-api', frontendApi);
  script.async = true;
  script.src = `https://${frontendApi}/npm/@clerk/clerk-js@1/dist/clerk.browser.js`
  // Add a listener so you can initialize ClerkJS
  // once it's loaded.
  script.addEventListener('load', async function(){
    await window.Clerk.load({
      // Set load options here...
    });
  });
  document.body.appendChild(script);
</script>
```

## TypeScript support

The [ClerkJS npm package](https://www.npmjs.com/package/@clerk/clerk-js) includes TypeScript declarations for many of the objects returned from the [Clerk Frontend API](../frontend-api-reference/). 

The type declarations in `@clerk/clerk-js` for these objects will always track the latest version of the Clerk Frontend API. If you would like to use these types but are using an older version of the Clerk Frontend API, we recommend [u](https://stripe.com/docs/upgrades#how-can-i-upgrade-my-api)pdating to the latest version, or ignoring and overriding the type definitions as necessary. 

Note that we may release new [minor and patch](https://semver.org/) versions of `@clerk/clerk-js` with small but backwards-incompatible fixes to the type declarations. These changes will not affect ClerkJS itself.

