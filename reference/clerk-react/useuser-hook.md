---
description: Access the User object inside your components.
---

# useUser

## Overview

The `useUser` hook accesses the active [User](../clerkjs/user.md) object. It can be used to update the user or display information about the user's profile, like their name or email address. The hook provides a shortcut for retrieving the [Clerk.session.user](../clerkjs/clerk.md#attributes) property.

In its simplest form, the `useUser` hook is called without arguments:`useUser()`. The object returned from the hook will hold all state for the currently **signed in** user. As such, the `useUser` hook must be called from a component that is a descendant of the [&lt;SignedIn/&gt;](../../components/control-components/signed-in.md) component - otherwise, the hook will throw an error. For more details, check the [Simple form](useuser-hook.md#simple-form) section.

For more advanced use cases where you want fine-grained control over the state of the `User` object, you can use the second form of the hook: `useUser({ withAssertions: true });`. This form does not have the constraints mentioned above, so it can also be used from components that are not descendants of a `<SignedIn/>`. For more details, check the [Advanced form w/ state helpers](useuser-hook.md#advanced-form-w-state-helpers) section.

If you don't wish to use React hooks, we also offer a couple of [alternative methods](useuser-hook.md#alternatives) to retrieve the currently signed in user outside of a `<SignedIn/>` component.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

### Simple form

A basic example that showcases the `useUser` hook in action is a greeting component that greets the signed in user by their first name. Note that your component must be a descendant of the [&lt;SignedIn/&gt;](../../components/control-components/signed-in.md) component, which in turn needs to be wrapped inside the [&lt;ClerkProvider/&gt;](clerkprovider.md).

{% tabs %}
{% tab title="React JSX" %}
```jsx
import {
  ClerkProvider,
  useUser,
  SignedIn,
  SignedOut,
  SignIn,
} from "@clerk/clerk-react";

const frontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;

function App() {
  return (
    <ClerkProvider frontendApi={frontendApi}>
      <SignedIn>
        {/* Components using useUser() must be rendered
         under the SignedIn component */}
        <Greeting />
      </SignedIn>
      
      <SignedOut>
        <SignIn />
      </SignedOut>
    </ClerkProvider>
  );
}

function Greeting() {
  // Use the useUser hook to get the Clerk.user object
  // This hook causes a re-render on user changes
  const user = useUser();
  return <div>Hello, {user.firstName}!</div>;
}

export default App;
```
{% endtab %}
{% endtabs %}

### Advanced form w/ state helpers

A slightly more advanced example follows. It showcases the `useUser` hook used by the `<Greeting/>`component that greets the signed in user by their first name. Note that your component does not need to be a descendant of the [&lt;SignedIn/&gt;](../../components/control-components/signed-in.md) component.

When `useUser` is called with the `withAssertions` option set to true, in addition to the user object it also returns the following helpers: 

* `isSignedIn` : A helper method that accepts the user and returns true if the user is signed in.
* `isSignedOut` : A helper method that accepts the user and returns true if the user is signed out.
* `isLoading` : A helper method that accepts the user and returns true if Clerk is still loading.

```jsx
import {
  ClerkProvider,
  useUser,
  SignedOut,
  SignIn,
} from "@clerk/clerk-react";

const frontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;

function App() {
  return (
    <ClerkProvider frontendApi={frontendApi}>
      <Greeting />

      <SignedOut>
        <SignIn />
      </SignedOut>
    </ClerkProvider>
  );
}

function Greeting() {
  // Use the useUser hook to get the Clerk.user object
  // This hook causes a re-render on user changes
  const { user, isSignedOut, isLoading } = useUser({
    withAssertions: true,
  });

  if (isLoading(user)) {
    // Inside this block, the type of `user` is `undefined`
    return <div>Clerk is loading...</div>;
  }

  if (isSignedOut(user)) {
    // Inside this block, the type of `user` is `null`
    return <div>You are signed out, `user` is null</div>;
  }

  // Here, the type of `user` is `UserResource`
  // You can also use the isSignedIn helper
  return <div>Hello, {user.firstName}!</div>;
}

export default App;

```

{% hint style="info" %}
The above helpers are **methods** that need to be passed the user object. This API is designed in such a way to help Typescript correctly infer the relationship between the different states and the user object when using object restructuring.
{% endhint %}

## Alternatives

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

There are times where using a hook might be inconvenient. For such cases, there are alternative ways to get access to the [User](../clerkjs/user.md) object.

Clerk provides the [&lt;WithUser/&gt;](useuser-hook.md#withuser-1) component and the [withUser](useuser-hook.md#withuser) higher order component directive that will allow your wrapped components to get access to the currently signed in user.

### withUser

The `withUser` function offers a [Higher Order Component \(HOC\)](https://reactjs.org/docs/higher-order-components.html) that you can wrap your own components with, and they will have access to the active [User](../clerkjs/user.md) object.

The `User` object will be accessible in the wrapped component under the `user` prop.

```jsx
import { withUser } from "@clerk/clerk-react";

// Your component must be a descendant
// of the <ClerkProvider /> component.
class Greeting extends React.Component {  
  render() {
    return (
      <div>
        {this.props.user.firstName
          ? `Hello ${this.props.user.firstName}!`
          : "Hello there!"}
      </div>
    );
  }
}

// Wrap your component with the withUser HOC.
export const GreetingWithUser = withUser(Greeting);
```

### &lt;WithUser /&gt; <a id="withuser-component"></a>

If you really want to stretch JSX capabilities and you cannot use the [withUser](useuser-hook.md#withuser) higher order component, we provide a `<WithUser/>` component that accepts a [Function as a child](https://reactjs.org/docs/jsx-in-depth.html#functions-as-children). Inside this function, the active [User](../clerkjs/user.md) object will be accessible.

```jsx
import { WithUser } from "@clerk/clerk-react";

class Greeting extends React.Component {
  render() {
    return (
      <WithUser>
        {(user) => (
          <div>
            {user.firstName 
              ? `Hello, ${user.firstName}!` 
              : "Hello there!"}
          </div>
        )}
      </WithUser>
    );
  }
}
```

### Typescript support

If you're using Typescript, Clerk provides a `WithUserProp` type to make the compiler aware of the `user` prop for your components.

You can use the `WithUserProp` type in combination with both the [withUser ](useuser-hook.md#withuser)higher order component and the [&lt;WithUser/&gt;](useuser-hook.md#withuser-1) component. 

The example below uses the `withUser` higher order component.

```jsx
import { withUser, WithUserProp } from "@clerk/clerk-react";

class Greeting extends React.Component<WithUserProp<{}>> {
  render() {
    return (
      <div>
        {this.props.user.firstName
          ? `Hello, ${this.props.user.firstName}!`
          : "Hello there!"}
      </div>
    );
  }
}

export const GreetingWithUser = withUser(Greeting);
```

