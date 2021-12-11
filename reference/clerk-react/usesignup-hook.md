---
description: Access the SignUp object inside your components.
---

# useSignUp

## Overview

The `useSignUp` hook gives you access to the [SignUp](../clerkjs/signup.md) object inside your components.&#x20;

You can use the methods of the `SignUp` object to create your own custom sign up flow, as an alternative to using Clerk's pre-built [\<SignUp/>](../../components/sign-up/sign-up.md) component.

The `SignUp` object will also contain the state of the sign up attempt that is currently in progress, giving you the chance to examine all the details and act accordingly.

## Usage

{% hint style="warning" %}
Make sure you've followed the installation guide for [Clerk React](installation.md) before running the snippets below.
{% endhint %}

Getting access to the [SignUp](../clerkjs/signup.md) object from inside one of your components is easy. Note that the `useSignUp` hook can only be used inside a [\<ClerkProvider/>](clerkprovider.md) context.

The following example accesses the `SignUp` object to check the current sign up attempt's status.

```jsx
import { useSignUp } from "@clerk/clerk-react";

// Your component must be rendered as a 
// descendant of <ClerkProvider />.
function SignUpStep() {
  const signUp = useSignUp();
  
  return (
    <div>
      The current sign up attempt status 
      is {signUp.status}.
    </div>
  );
}
```

A more involved example follows below. In this example, we show an approach to create your own custom form for registering users.

{% hint style="info" %}
We recommend using the [\<SignUp/>](../../components/sign-up/sign-up.md) component instead of building your own custom registration form. It gives you a ready-made form and handles each step of the sign up flow.
{% endhint %}

```jsx
import { useSignUp } from "@clerk/clerk-react";

// Your component must be rendered as a 
// descendant of <ClerkProvider />.
function SignUpForm() {
  const [emailAddress, setEmailAddress] = useState('');
  const [password, setPassword] = useState('');
  
  const signUp = useSignUp();
  
  async function submit(e) {
    e.preventDefault();
    
    // Check the sign up response to 
    // decide what to do next.
    await signUp.create({
      emailAddress,
      password,
    });
  }
  
  return (
    <form onSubmit={submit}>
      <div>
        <label htmlFor="email">Email</label>
        <input 
          type="email" 
          value={emailAddress} 
          onChange={e => setEmailAddress(e.target.value)} 
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
        <button>Sign up</button>
      </div>
    </form>
  );
}
```
