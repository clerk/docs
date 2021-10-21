---
description: Access the Clerk object inside your components.
---

# useClerk

## Overview

The `useClerk` hook accesses the [Clerk](../clerkjs/clerk.md) object. It can be used to retrieve any object in the [ClerkJS](../clerkjs/) SDK. Moreover, it allows access to all of the [Clerk object's methods](../clerkjs/clerk.md#methods), giving you the freedom to build alternatives to any [Clerk Component](../../main-concepts/clerk-components.md).

{% hint style="info" %}
If you want to access a [ClerkJS](../clerkjs/) object directly, like [User](../clerkjs/user.md) or [Session](../clerkjs/session.md), Clerk provides more specific hooks like [useUser](useuser-hook.md) or [useSession](usesession-hook.md). There's probably a more fine-grained [Clerk React](./) hook that you can turn to before resorting to the `useClerk` hook.
{% endhint %}

There are a couple of [alternative methods](useclerk-hook.md#alternatives) to retrieve the `Clerk` object if you cannot use hooks.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

An example of the `useClerk` hook in action is shown below. We get access to the [Clerk](../clerkjs/clerk.md) object in order to render a button that opens the [sign in form](../../components/sign-in.md) as a modal. Note that if your component is not a descendant of the [\<ClerkProvider/>](clerkprovider.md) component, the `useClerk` hook will throw an error.

```jsx
import { useClerk } from "@clerk/clerk-react";

// Your component must be a descendant
// of the <ClerkProvider /> component.
function SignInButton() {
  const clerk = useClerk();
  
  return (
    <button onClick={() => clerk.openSignIn({})}>
      Sign in
    </button>
  );
}
```

## Alternatives

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

There are times where using a hook might be inconvenient. For such cases, there are alternative ways to get access to the [Clerk](../clerkjs/clerk.md) object.

Clerk provides the [\<WithClerk/>](useclerk-hook.md#withclerk-component) component and the [withClerk](useclerk-hook.md#withclerk) higher order component directive that will allow your wrapped components to get access to the `Clerk` object.

### withClerk

The `withClerk` function offers a [Higher Order Component (HOC)](https://reactjs.org/docs/higher-order-components.html) that you can wrap your own components with, and they will have access to the [Clerk](../clerkjs/clerk.md) object.

The `Clerk` object will be accessible in the wrapped component under the `clerk` prop.

```jsx
import { withClerk } from "@clerk/clerk-react";

class SignInButton extends React.Component {
  render() {
    return (
      <button onClick={() => this.props.clerk.openSignIn({})}>
        Sign in
      </button>
    );
  }
}

// Wrap your component with the useClerk HOC.
export const SignInButtonWithClerk = withClerk(SignInButton);
```

### \<WithClerk /> <a href="withclerk-component" id="withclerk-component"></a>

f you really want to stretch JSX capabilities and you cannot use the [withClerk](useclerk-hook.md#withclerk) higher order component, we provide a `<WithClerk/>` component that accepts a [Function as a child](https://reactjs.org/docs/jsx-in-depth.html#functions-as-children). Inside this function, the [Clerk](../clerkjs/clerk.md) object will be accessible.

```jsx
import { WithClerk } from "@clerk/clerk-react";

class SignInButton extends React.Component {
  render() {
    return (
      <WithClerk>
        {(clerk) => (
          <button onClick={() => clerk.openSignIn({})}>
            Sign in
          </button>
        )}
      </WithClerk>
    );
  }
}
```

### Typescript support

If you're using Typescript, Clerk provides a `WithClerkProp` type to make the compiler aware of the `clerk` prop for your components.

You can use the `WithClerkProp` type in combination with both the [withClerk](useclerk-hook.md#withclerk) higher order component and the [\<WithClerk/>](useclerk-hook.md#withclerk-component) component.&#x20;

The example below uses the `withClerk` higher order component.

```jsx
import { withClerk, WithClerkProp } from "@clerk/clerk-react";

class SignInButton extends React.Component<WithClerkProp<{}>> {
  render() {
    return (
      <button onClick={() => this.props.clerk.openSignUp({})}>
        Sign in
      </button>
    );
  }
}

export const SignInButtonWithClerk = withClerk(SignInButton);
```
