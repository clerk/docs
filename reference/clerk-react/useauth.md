---
description: Access the auth state inside your React components.
---

# useAuth

## Overview

The `useAuth` hook is a convenient way to access the current auth state. This hook provides the minimal information needed for data-loading and helper methods to manage the current active session.

This hook is compatible with SSR and is recommended for all authentication tasks.&#x20;

```javascript
const {
  userId,
  sessionId,
  getToken,
  isLoaded,
  isSignedIn,
  signOut
} = useAuth();
```

This new hook should be used instead of:&#x20;

* `useSession()` , to access `getToken()` or the `sessionId`
* `useUser()` , to access `getToken()` for integrations
* `useClerk()` , to access `signOut()`

## Usage

In the following basic example, `useAuth` is used to access the auth state and generate a new JWT for a custom Supabase integration in a NextJS application. Note that your component must be a descendant of [\<ClerkProvider/>](clerkprovider.md).

{% hint style="info" %}
#### Typescript

For better type support, we highly recommend updating to **at least** [**Typescript 4.6**](https://devblogs.microsoft.com/typescript/announcing-typescript-4-6/)**.** Applications using Typescript will benefit significantly from Typescript's new [Control Flow Analysis for Dependent Parameters](https://devblogs.microsoft.com/typescript/announcing-typescript-4-6/#control-flow-analysis-for-dependent-parameters) when using our hooks.
{% endhint %}

{% tabs %}
{% tab title="NextJS" %}
```jsx
// pages/supabase.tsx
import { useAuth } from '@clerk/nextjs';
import { supabase } from './supabaseClient';

function SupabasePage() {
  const { getToken, isLoaded, isSignedIn } = useAuth();
  
  if (!isLoaded || !isSignedIn) {
    // You can handle the loading or signed state separately
    return null;
  }
  
  const fetchDataFromSupabase = () => {
    const token = await getToken({ template: 'supabase' });
    supabase.auth.setAuth(token);
    const { data } = await supabase.from('your-table').select();
    return data;
  }
  
  return <div>...</div>;
}

export default App;
```
{% endtab %}
{% endtabs %}

## Alternative ways to access the auth object

There are times when you may need to access the auth object during SSR, API handlers and Edge middleware. The auth object is also available in Next.js, Remix, and Gatsby API routes, even during server-side rendering.

The following snippets are quick examples for NextJS applications. For more details, look at the starter guide for your framework of choice.

#### API Routes: `req.auth`

{% code title="pages/api/hello.ts" %}
```jsx
import { withAuth } from "@clerk/nextjs/api";

export default withAuth(async ( req, res ) => {
    const { userId, sessionId, getToken } = req.auth;
    const supabaseToken = await getToken({template: "supabase"})
    // Load any data your application needs for the API route
    return { props: {} };
});
```
{% endcode %}

#### Server-side rendering: `req.auth`

{% code title="pages/my-page.tsx" %}
```jsx
import { withServerSideAuth } from "@clerk/nextjs/ssr";

export default MyPage(){
   return ...;
}

export const getServerSideProps = withServerSideAuth(
  async ({ req }) => {
    const { userId, sessionId, getToken } = req.auth;
    const hasuraToken = await getToken({template: "hasura"})    
    // Load any data your application needs and pass to props
    return { props: {} };
  }
);
```
{% endcode %}

#### Edge Middleware: `req.auth`

{% code title="pages/_middleware.js" %}
```jsx
import { withEdgeMiddlewareAuth } from "@clerk/nextjs/edge-middleware";

export default withEdgeMiddlewareAuth(
  async ( req ) => {
    const { userId, sessionId, getToken } = req.auth;
    const supabaseToken = await getToken({template: "supabase"})
    
    // Run your middleware

    // Complete response
    return NextResponse.next();
  }
);
```
{% endcode %}
