---
description: >-
  The UserButton component implements the de facto standard for allowing a user
  to manage their account, switch accounts, sign out.
---

# UserButton

The **UserButton** component is commonly placed in the upper right of the page for a signed in user. In it's inactive state, it's just an image of the user \(or a person icon\). In its active state, it allows the user to manage their profile, sign out, or sign in to another account.

![Example UserButton](../../.gitbook/assets/image%20%283%29.png)

To mount the **UserButton** in your application, simply place it anywhere in your JSX:

```jsx
import { UserButton } from "@clerk/clerk-react";

const Header = () => {
  return (
    <header>
      <h1>My application</h1>
      <UserButton />
    </header>
  );
};
```



