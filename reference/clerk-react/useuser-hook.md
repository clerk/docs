---
description: Access the User object inside your components.
---

# useUser

## Overview

The useUser hook returns the current user state: `{ isLoaded, isSignedIn, user }` . You can use the `user` object to access the currently active [User](../clerkjs/user.md). It can be used to update the user or display information about the user's profile, like their name or email address.

While hooks are the recommended way to use the Clerk APIs, If you don't wish to use React hooks we also offer a couple of [alternative methods](useuser-hook.md#alternatives) to retrieve the currently signed in user.

## Usage

A basic example that showcases the `useUser` hook in action is a greeting component that greets the signed in user by their first name. Note that your component must be a descendant of [\<ClerkProvider/>](clerkprovider.md).

{% hint style="info" %}
#### Typescript

For better type support, we highly recommend updating to **at least** [**Typescript 4.6**](https://devblogs.microsoft.com/typescript/announcing-typescript-4-6/)**.** Applications using Typescript will benefit significantly from Typescript's new [Control Flow Analysis for Destructured Discriminated Unions](https://devblogs.microsoft.com/typescript/announcing-typescript-4-6/#control-flow-analysis-for-destructured-discriminated-unions) when using our hooks.
{% endhint %}

{% tabs %}
{% tab title="React JSX" %}
```jsx
import { ClerkProvider,  useUser, SignIn } from "@clerk/clerk-react";

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
  const { isLoaded, isSignedIn, user } = useUser();
  
  if (!isLoaded || !isSignedIn) {
    // You can handle the loading or signed state separately
    return null;
  }
  
  return <div>Hello, {user.firstName}!</div>;
}

export default App;
```
{% endtab %}
{% endtabs %}

## Alternatives

There are times where using a hook might be inconvenient. For such cases, there are alternative ways to get access to the [User](../clerkjs/user.md) object.

Clerk provides the [\<WithUser/>](useuser-hook.md#withuser-1) component and the [withUser](useuser-hook.md#withuser) higher order component directive that will allow your wrapped components to get access to the currently signed in user.

### withUser

The `withUser` function offers a [Higher Order Component (HOC)](https://reactjs.org/docs/higher-order-components.html) that you can wrap your own components with, and they will have access to the active [User](../clerkjs/user.md) object.

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

### \<WithUser /> <a href="#withuser-component" id="withuser-component"></a>

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

### Usage with Typescript

If you're using Typescript, Clerk provides a `WithUserProp` type to make the compiler aware of the `user` prop for your components.

You can use the `WithUserProp` type in combination with both the [withUser ](useuser-hook.md#withuser)higher order component and the [\<WithUser/>](useuser-hook.md#withuser-1) component.&#x20;

The example below uses the `withUser` higher order component.

```jsx
import { withUser, WithUserProp } from "@clerk/clerk-react";

type GreetingProps = {
  greeting: string;
}

const Greeting = (props: WithUserProp<GreetingProps>) => {
  const { user, greeting } = props;
  return (
    <h1>{ greeting }</h1>
    <div>
      { user.firstName
        ? `Hello, ${user.firstName}!`
        : "Hello there!" }
    </div>
  );
}

export const GreetingWithUser = withUser(Greeting);
```
