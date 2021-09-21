---
description: Complete a custom OAuth flow
---

# &lt;AuthenticateWithRedirectCallback /&gt;

## Overview

The `<AuthenticateWithRedirectCallback/>` is used to complete a custom OAuth flow started by calling either [`SignIn.authenticateWithRedirect(params)`](../../reference/clerkjs/signin.md#signinwithoauth) or [`SignUp.authenticateWithRedirect(params)`](../../reference/clerkjs/signup.md#signinwithoauth). 

Internally, it calls the [`Clerk.handleRedirectCallback()`](../../reference/clerkjs/clerk.md#handleredirectcallback-params) method

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

The most common scenario for using the `<AuthenticateWithRedirectCallback/>` component, is to complete a custom OAuth sign in or sign up flow in React and NextJS apps. Simply render the component under the route you passed as `callbackUrl` to the `authenticateWithRedirect` methods.

For a more detailed walkthrough, you can check the [Social Login \(OAuth\) guide](../../popular-guides/social-login-oauth.md).

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import React from "react";
import {
  ClerkProvider,
  AuthenticateWithRedirectCallback,
  UserButton,
  SignedOut,
  useSignIn,
} from "@clerk/clerk-react";

const frontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;

function App() {
  return (
    //  react-router-dom requires your app to be wrapped with a Router
    <BrowserRouter>
      <ClerkProvider frontendApi={frontendApi}>
        <Switch>
          {/* Define a / route that displays the OAuth button */}
          <Route path="/">
            <SignedOut>
              <SignInOAuthButtons />
            </SignedOut>
          </Route>
         
           {/* Define a /sss-callback route that handle the OAuth redirect flow */}
          <Route path="/sso-callback">
            {/* Simply render the component */}
            <AuthenticateWithRedirectCallback />
          </Route>
        </Switch>
      </ClerkProvider>
    </BrowserRouter>
  );
}

function SignInOAuthButtons() {
  const { authenticateWithRedirect } = useSignIn();
  const signInWithGoogle = () =>
    authenticateWithRedirect({
      strategy: "oauth_google",
      callbackUrl: "/sso-callback",
      callbackUrlComplete: "/",
    });
  return <button onClick={signInWithGoogle}>Sign in with Google</button>;
}
```
{% endtab %}
{% endtabs %}



## Props

<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>afterSignInUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or path to navigate after successful sign in.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSignUpUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or path to navigate after successful sign up.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>redirectUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or path to navigate after successful sign in or sign up. The
          same as setting <code>afterSignInUrl</code> and <code>afterSignUpUrl</code> to
          the same value.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>secondFactorUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or path to navigate during sign in, if 2FA is enabled.</p>
      </td>
    </tr>
  </tbody>
</table>

