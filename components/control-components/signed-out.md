---
description: Conditionally render content only when there's no signed in user.
---

# &lt;SignedOut /&gt;

## Overview

The `<SignedOut/>` component offers authentication checks as a cross-cutting concern.

Any child nodes wrapped by a `<SignedOut/>` component will be rendered only if there's no [User](../../reference/clerkjs/user.md) signed in your application. 

This is a component that controls the rendering flow. It acts as a guard; any content that you place inside a `<SignedOut/>` component will not be accessible to authenticated users.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](../../reference/clerk-react/installation.md) before running the snippets below.
{% endhint %}

A common scenario for using the `<SignedOut/>` component is having an application with a page that should be accessible only by signed in users. With the `<SignedOut/>` component, you can show some special content to your visitors.

{% tabs %}
{% tab title="Clerk React" %}
```jsx
import { SignedOut, ClerkProvider } from "@clerk/clerk-react";

function Page() {
  return (
    <ClerkProvider frontendApi="clerk.[your-domain].com">
      <section>
        <div>
          This content is always visible by both
          signed in users and guests.
        </div>
        <SignedOut>
          <div>
            This content is not visible to 
            signed in users.
          </div>
        </SignedOut>
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

Let's see how the `<SignedOut/>` component might help with the above scenario. There's a restricted route and the `<SignedOut/>` component acts as a conditional inside it, in order to redirect unauthenticated requests to the sign in page.

Notice how we also use the [&lt;ClerkProvider/&gt;](../../reference/clerk-react/clerkprovider.md), [&lt;SignedIn/&gt;](signed-in.md) and [&lt;RedirectToSignIn/&gt;](redirect-to-sign-in.md) components to complete the functionality. The example below uses the popular [React Router](https://reactrouter.com/) routing library, but this is just an implementation detail. You can use any routing mechanism to achieve the same result.

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

This component accepts no props, apart from child components that will conditionally render, only when there's no signed in user.

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
        <p>Pass any component or components and they will be rendered only if no
          user is signed in.</p>
      </td>
    </tr>
  </tbody>
</table>



