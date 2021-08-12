---
description: Access the User object inside your components.
---

# useUser

## Overview

The `useUser` hook accesses the active [User](../clerkjs/user.md) object. It can be used to display information about the user's profile, like their name or email address. The hook provides a shortcut for retrieving the [Clerk.session.user](../clerkjs/clerk.md#attributes) property.

The `User` object returned from the hook will hold all state for the currently signed in user. As such, the `useUser` hook must be called from a component that is a descendant of the [&lt;SignedIn/&gt;](../../components/control-components/signed-in.md) component. Otherwise, the hook will throw an error.

There are a couple of [alternative methods](useuser-hook.md#alternatives) to retrieve the currently signed in user outside of a `<SignedIn/>` component.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

A basic example that showcases the `useUser` hook in action is a greeting component that greets the signed in user by their first name. Note that your component must be a descendant of the [&lt;SignedIn/&gt;](../../components/control-components/signed-in.md) component, which in turn needs to be wrapped inside the [&lt;ClerkProvider/&gt;](clerkprovider.md).

```jsx
import { SignedIn, useUser } from "@clerk/clerk-react";

function App() {
  return (
    <SignedIn>
      <Greeting />
    </SignedIn>
  );
}

function Greeting() {
  const user = useUser();
  
  const greeting = user.firstName || "there";
  return (
    <div>Hello {greeting}!</div>
  );
}
```

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

