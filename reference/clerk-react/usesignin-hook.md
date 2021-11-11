---
description: Access the SignIn object inside your components.
---

# useSignIn

## Overview

The `useSignIn` hook gives you access to the [SignIn](../clerkjs/signin/) object inside your components.&#x20;

You can use the methods of the `SignIn` object to create your own custom sign in flow, as an alternative to using Clerk's pre-built [\<SignIn/>](../../components/sign-in.md) component.

The `SignIn` object will also contain the state of the sign in attempt that is currently in progress, giving you the chance to examine all the details and act accordingly.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

Getting access to the [SignIn](../clerkjs/signin/) object from inside one of your components is easy. Note that the `useSignIn` hook can only be used inside a [\<ClerkProvider/>](clerkprovider.md) context.

The following example accesses the `SignIn` object to check the current sign in attempt's status.

```jsx
import { useSignIn } from "@clerk/clerk-react";

// Your component must be rendered as a 
// descendant of <ClerkProvider />.
function SignInStep() {
  const signIn = useSignIn();
  
  return (
    <div>
      The current sign in attempt status 
      is {signIn.status}.
    </div>
  );
}
```

A more involved example follows below. In this example, we show an approach to create your own custom form for signing in your users.

{% hint style="info" %}
We recommend using the [\<SignIn/>](../../components/sign-in.md) component instead of building your own custom sign in form. It gives you a ready-made form and handles each step of the sign in flow.
{% endhint %}

```jsx
import { useSignIn } from "@clerk/clerk-react";

// Your component must be rendered as a 
// descendant of <ClerkProvider />.
function SignInForm() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  
  const signIn = useSignIn();
  
  async function submit(e) {
    e.preventDefault();
    
    // Check the sign in response to 
    // decide what to do next.
    await signIn.create({
      identifier: email,
      password,
    });
  }
  
  return (
    <form onSubmit={submit}>
      <div>
        <label htmlFor="email">Email</label>
        <input 
          type="email" 
          value={email} 
          onChange={e => setEmail(e.target.value)} 
        />
      </div>
      <div>
        <label htmlFor="password">Password</label>
        <input
          type="password"
          value={password}
          onChange={e => setPassword(e.target.value)}
        />
      </div>
      <div>
        <button>Sign in</button>
      </div>
    </form>
  );
}
```
