---
description: >-
  This guide will help you authenticate requests made to your Next.js API
  routes.
---

# Get started with Next.js API

## 0. Pre-requirements

This guide assumes you have properly configured your frontend to make requests to your backend. If you haven't done so already, please follow the [guide for making backend requests](https://docs.clerk.dev/frontend/react/making-backend-requests).

## 1. Set CLERK\_API\_KEY

Create a file named **.env.local** in your application root if it doesn't exist already. Any variables inside this file will be accessible in your backend with **process.env.VARIABLE\_NAME**.

{% hint style="warning" %}
Make sure you update ****this variable with the API key found in your dashboard under **Settings** → **API Keys**.
{% endhint %}

{% code title=".env.local" %}
```jsx
CLERK_API_KEY=test_asdf1234
```
{% endcode %}

## 2. Install clerk-sdk-node

Install Clerk's NPM package for node backends \(under the hood, Next.js API routes are serverless Node.js functions\).

```bash
npm install @clerk/clerk-sdk-node
# or
yarn add @clerk/clerk-sdk-node
```

## 3. Create an API route

In Next.js, API routes are created by adding a file to the **pages/api** folder.  The filename is used as the API path. New Next.js applications come with a **pages/api/hello.js** file that we can edit.

{% code title="pages/api/hello.js" %}
```jsx
import { requireSession } from "@clerk/clerk-sdk-node";

export default requireSession((req, res) => {
  res.statusCode = 200;
  res.json({ id: req.session.userId });
});
```
{% endcode %}

That's it! Now, if a user is signed in, requests to **/api/hello** will return that user's ID.  This ID will never change for this user, so you can use it as a stable user ID in your own codebase.

If a user is signed out, the **requireSession** helper will automatically return with HTTP status 401 and no response body. If you'd prefer to process a request when there is no user, try the **withSession** helper instead:

{% code title="pages/api/hello.js" %}
```jsx
import { withSession } from "@clerk/clerk-sdk-node";

export default withSession((req, res) => {
  res.statusCode = 200;
  if (req.session){
    res.json({ id: req.session.userId });
  } else {
    res.json({ id: null });
  } 
});
```
{% endcode %}

