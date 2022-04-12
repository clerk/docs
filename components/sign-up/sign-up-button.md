---
description: Show a button that links to the sign-up URL or opens the sign-up modal
---

# \<SignUpButton />

## Overview

`<SignUpButton/>` is a button that links to the sign-up page or displays the sign-up modal. By default, it is a `<button>` tag that says **Sign up**, but it is completely customizable by passing children.

## Usage

The component can be imported from our React or Next.js SDK:

```javascript
// React
import { SignUpButton } from "@clerk/clerk-react";

// Next.js
import { SignUpButton } from "@clerk/nextjs";
```

### Simple button

```jsx
<SignUpButton />
```

### Custom button, open a modal

```xml
<SignUpButton mode="modal">
  <button className="btn">
    Sign up
  </button>
</SignUpButton>
```

## Props



| Name                | Description                                                                                                                                                                                                                                                                             |
| ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **mode**            | <p><em>"redirect" (default) | "modal"</em> </p><p>If mode is set to "redirect", the button will redirect to the sign-up page. If mode is set to "modal", the button will open a modal instead<em>.</em></p><p></p><p>Defaults to "redirect".</p>                                        |
| **redirectUrl?**    | <p><em>string</em></p><p>Full URL or path to navigate to after successful sign in or sign up. Use this instead of setting <code>afterSignInUrl</code> and <code>afterSignUpUrl</code> to the same value.<br><br>To return to the same URL, set to <code>window.location.href</code></p> |
| **afterSignInUrl?** | <p><em>string</em></p><p>The full URL or path to navigate to after a successful sign in.<br><br>Defaults to the Sign-in URL on the Paths page of your Dashboard.</p>                                                                                                                    |
| **afterSignUpUrl?** | <p><em>string</em></p><p>The full URL or path to navigate to after a successful sign up.<br><br>Defaults to the Sign-up URL on the Paths page of your Dashboard.</p>                                                                                                                    |
