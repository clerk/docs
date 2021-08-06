---
description: >-
  The SignedIn and SignedOut components are used to control rendering depending
  on whether or not a visitor is signed in.
---

# SignedIn and SignedOut

## Usage

The **SignedIn** component renders if the visitor is signed in, while the **SignedOut** component renders if the visitor is signed out.

```jsx
import { SignedIn, SignedOut } from "@clerk/clerk-react";

const SignInStatus = () => (
  <div>
    <SignedIn>The user is signed in.</SignedIn>
    <SignedOut>The user is signed out.</SignedOut>
  </div>
);
```

More precisely, these components reference the **session** property from the [**useClerk** hook](../clerkjs/clerk.md) to determine whether their children should render.

* When **session** is a [Session object](../clerkjs/session.md), the **SignedIn** component will render.
* When **session** is **null**, the **SignedOut** component will render.
* When **session** is **undefined**, neither component will render. This indicates that Clerk is still loading or that **session** is in the process of changing.

