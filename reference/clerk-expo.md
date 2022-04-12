---
description: Integrate Clerk into your React Native application
---

# Clerk Expo

## Overview

Clerk Expo is a wrapper around [Clerk React](clerk-react/). It is the recommended way to integrate Clerk into your React Native application.

{% hint style="info" %}
Clerk Expo works for IOS and Android. For web based applications use [Clerk React](clerk-react/).
{% endhint %}

Clerk Expo provides [hooks](clerk-react/) and [Clerk Control Components](../components/control-components/) which act as helpers for implementing a seamless authentication experience, give you access to the [Clerk object](clerkjs/clerk.md) and a set of useful helper methods for signing in and signing up.

{% hint style="info" %}
Looking for a quickstart? We created a [demo app](https://github.com/clerkinc/clerk-expo-starter) to show you how to add Clerk to your project.
{% endhint %}

## Setting up Clerk Expo

{% hint style="warning" %}
You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev) before you can set up Clerk React. For more information, check out our [Setup your application](broken-reference) guide.
{% endhint %}

There's an ES module for Clerk Expo, available under the [@clerk/clerk-expo](https://www.npmjs.com/package/@clerk/clerk-expo) npm package. Use `npm` or `yarn` to install the Clerk React module.

```bash
# Install the package with npm.
npm install @clerk/clerk-expo

# Alternative install method with yarn.
yarn add @clerk/clerk-expo
```

## Clerk provider

The [`ClerkProvider`](clerk-react/clerkprovider.md) allows you to render [Clerk Control Components](../components/control-components/) and access the available Clerk React hooks in any nested component. You'll have to wrap your application once with a `<ClerkProvider/>`.

Render a `<ClerkProvider/>` component at the root of your React app so that it is available everywhere you need it.

In order to use`<ClerkProvider/>` first you need to locate the entry point file of your React Native app. In Expo, this is usually your `src/App.js`.

Replace the `frontendApi` prop with the [Frontend API](frontend-api-reference/) host found on the [API Keys page](https://dashboard.clerk.dev/last-active?path=api-keys).

```jsx
import React from "react";
import { Text } from "react-native";
import { SafeAreaProvider } from "react-native-safe-area-context";
import { ClerkProvider } from "@clerk/clerk-expo";

export default function App() {
  return (
    <ClerkProvider frontendApi="clerk.[your-domain].com">
      <SafeAreaProvider>
        <Text>Hello world!</Text>
      </SafeAreaProvider>
    </ClerkProvider>
  );
}
```

## Token Cache

Clerk Expo stores the client JWT token in memory by default. It is highly recommended to use a secure store according to your React Native framework.

For example, [Expo](https://expo.dev) provides a way to encrypt and securely store key–value pairs locally on the device via [expo-secure-store](https://docs.expo.dev/versions/latest/sdk/securestore/). You can use it as your client JWT storage by setting the `tokenCache` prop in your `<ClerkProvider/>` as shown below.

```jsx
import React from "react";
import { Text } from "react-native";
import { SafeAreaProvider } from "react-native-safe-area-context";
import { ClerkProvider } from "@clerk/clerk-expo";

import * as SecureStore from "expo-secure-store";

const tokenCache = {
  getToken(key: string) { 
    return SecureStore.getItemAsync(key);
  }, 
  saveToken(key: string, value: string) {
    return SecureStore.setItemAsync(key, value);
  }
};

export default function App() {
  return (
    <ClerkProvider frontendApi="clerk.[your-domain].com" tokenCache={tokenCache}>
      <SafeAreaProvider>
        <Text>Hello world!</Text>
      </SafeAreaProvider>
    </ClerkProvider>
  );
}
```
