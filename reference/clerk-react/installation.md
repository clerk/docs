---
description: Learn how to make Clerk React components available in your project.
---

# Installation

## Overview

Clerk React is a wrapper around [ClerkJS](../clerkjs/). It is the recommended way to integrate Clerk into your React application.

Clerk React provides [React.js](https://reactjs.org) implementations of [Clerk Components](../../main-concepts/clerk-components.md); highly customizable, pre-built components that you can use to build beautiful user management applications. You can find display components for building [sign in](../../components/sign-in.md), [sign up](../../components/sign-up.md), [account switching](../../components/user-button.md) and [user profile management](../../components/user-profile.md) pages as well as flow [control components](../../components/control-components/) which act as helpers for implementing a seamless authentication experience.

Clerk React comes loaded with custom [hooks](./). These hooks give you access to the [Clerk object](../clerkjs/clerk.md), and a set of useful helper methods for signing in and signing up.

{% hint style="info" %}
This doc assumes that you already have a basic working knowledge of [React](https://reactjs.org) and that you have already set up a React project. If you're new to React, you can check our Getting Started guides on [React](../../get-started/create-react-app.md) or [Next.js](../../get-started/nextjs.md).
{% endhint %}

## Setting up Clerk React

{% hint style="warning" %}
You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev) before you can set up Clerk React. For more information, check out our [Setup your application](../../popular-guides/setup-your-application.md) guide.
{% endhint %}

There's an ES module for Clerk React, available under the [@clerk/clerk-react](https://www.npmjs.com/package/@clerk/clerk-react) npm package. Use `npm` or `yarn` to install the Clerk React module.

```bash
# Install the package with npm.
npm install @clerk/clerk-react

# Alternative install method with yarn.
yarn add @clerk/clerk-react
```

## Clerk provider

The [`ClerkProvider`](clerkprovider.md) allows you to render [Clerk Components](../../main-concepts/clerk-components.md) and access the available Clerk React hooks in any nested component. You have to wrap your application once with a `<ClerkProvider/>`.

Render a `<ClerkProvider/>` component at the root of your React app so that it is available everywhere you need it.

In order to use`<ClerkProvider/>,` first you need to locate the entry point file of your React app. Usually this is your `src/index.js` (Create React App) or `pages/_app` (Next.js) file. In general, you're looking for the file where the `ReactDOM.render` function gets called.

Replace the `frontendApi` prop with the [Frontend API](../frontend-api-reference/) host found in your [Clerk Dashboard](https://dashboard.clerk.dev).

{% tabs %}
{% tab title="React" %}
```jsx
import ReactDOM from "react-dom";
import { ClerkProvider } from "@clerk/clerk-react";

// App is the top-level component. Wrap your whole
// DOM tree with a ClerkProvider. Pass the Frontend 
// API host as a prop.
function App() {
  return (
    <ClerkProvider frontendApi="clerk.[your-domain].com">
      {/* Your application tree goes here. */}
    </ClerkProvider>
  );
}

// Assuming the React app will render inside an 
// HTMLElement with id "app".
const appEl = document.getElementById("app");
ReactDOM.render(<App />, appEl);
```
{% endtab %}

{% tab title="Next.js" %}
```jsx
import { ClerkProvider } from '@clerk/clerk-react';

// App is the top-level component. Wrap your whole
// DOM tree with a ClerkProvider. Pass the Frontend
// API host as a prop.
function App({ Component, pageProps }) {
  return (
    <ClerkProvider frontendApi="clerk.[your-domain].com">
      <Component {...pageProps} />
    </ClerkProvider>
  );
}
```
{% endtab %}
{% endtabs %}

## Routing

A common scenario when using React is to navigate through your application without a full page reload. Our display components use this prop when navigating between subpages and when navigating to callback URLs.

The `<ClerkProvider/>` [navigate prop](clerkprovider.md#props) allows you to provide a function which takes the destination URL as an argument and performs a "push" navigation. You should not implement the push yourself, but instead wrap the `navigate` function provided by your router.

You can find examples of setting up path-based routing for [React Router](https://reactrouter.com) and Next.js.

{% tabs %}
{% tab title="React" %}
```jsx
import { BrowserRouter as Router, useNavigate } from "react-router-dom";
import { ClerkProvider } from "@clerk/clerk-react";

// App is the top-level component. Place the Router 
// above ClerkProvider and provide the React Router 
// navigate function to ClerkProvider.
function App() {
  const navigate = useNavigate();

  return (
    <Router>
      <ClerkProvider
        frontendApi="clerk.[your-domain].com"
        navigate={(to) => navigate(to)}
      >
        {/* Your application tree goes here. */}
      </ClerkProvider>
    </Router>
  );
}
```
{% endtab %}

{% tab title="Next.js" %}
```jsx
import { useRouter } from "next/router";
import { ClerkProvider } from "@clerk/clerk-react";

function App({ Component, pageProps }) {
  const { push } = useRouter();
  
  return (
    <ClerkProvider 
      frontendApi="clerk.[your-domain].com" 
      navigate={(to) => push(to)}
    >
      <Component {...pageProps} />
    </ClerkProvider>
  );
}
```
{% endtab %}
{% endtabs %}
