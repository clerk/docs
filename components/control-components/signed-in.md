---
description: Conditionally render content only when a user is signed in.
---

# &lt;SignedIn&gt;

## Overview

The `<SignedIn/>` component offers authentication checks as a cross-cutting concern.

Any children components wrapped by a `<SignedIn/>` component will be rendered only if there's a [User](../../reference/clerkjs/user.md) with an active [Session](../../reference/clerkjs/session.md) signed in your application. In that sense, the `<SignedIn/>` component provides a safe context where the current `User` object will be available.

This is a component that controls the rendering flow. It acts as a guard; any content that you place inside a `<SignedIn/>` component will be protected from unauthenticated users.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

A common scenario for using the `<SignedIn/>` component, is having an application with content that anybody can access and content that can only be accessed by authenticated users. 

For example, you might have a page that displays general-knowledge content, but provides additional information to members. The page should be publicly accessible, but part of the content should be visible only to signed in users.

You can use the `<SignedIn/>` to guard the members-only section of your page.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { SignedIn, ClerkProvider } from "@clerk/clerk-react";

function Page() {
  return (
    <ClerkProvider frontendApi="clerk.[your-domain].com">
      <section>
        <div>
          This content is always visible.
        </div>
        <SignedIn>
          <div>
            This content is visible only to 
            signed in users.
          </div>
        </SignedIn>
      </section>
    </ClerkProvider>
  );
}
```
{% endtab %}
{% endtabs %}

### Conditional routing

Another common scenario which better resembles a real world use-case, might be to restrict some of your application's pages to signed in users.

For example, a Saas platform website might have a page with information about the company and this page should be publicly accessible. At the same time, there might be a page for signed in users only, where users can edit their preferences.

Let's see how the `<SignedIn/>` component might help with the above scenario. Notice how we also use the [&lt;ClerkProvider/&gt;](../../reference/clerk-react/clerkprovider.md), [&lt;SignedOut/&gt;](signed-out.md) and [&lt;RedirectToSignIn/&gt;](redirect-to-sign-in.md) components to complete the functionality. The example below uses the popular [React Router](https://reactrouter.com/) routing library, but this is just an implementation detail. You can use any routing mechanism to achieve the same result.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { 
  ClerkProvider,
  SignedIn, 
  SignedOut, 
  RedirectToSignIn 
} from "@clerk/clerk-react";

function App() {
  return (
    <Router>
      <ClerkProvider frontendApi="clerk.[your-domain].com">
        <Route path="/public">
          <div>This page is publicly accessible.</div>
        </Route>
        <Route path="/private">
          <SignedIn>
            <div>
              This content is accessible only to signed
              in users.
            </div>
          </SignedIn>
          <SignedOut>
            {/* 
              Route matches, but no user is signed in. 
              Redirect to the sign in page.
            */}
            <RedirectToSignIn />
          </SignedOut>
        </Route>
      </ClerkProvider>
    </Router>
  );
}
```
{% endtab %}
{% endtabs %}

## Props

This component accepts no props, apart from child components that will conditionally render, only when a user is signed in.

<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>children</b>
      </td>
      <td style="text-align:left">
        <p><em>JSX.Element</em>
        </p>
        <p>Pass any component or components and they will be rendered only if a user
          is signed in.</p>
      </td>
    </tr>
  </tbody>
</table>

