# Next.js server-side changes

{% hint style="info" %}
Before starting this guide, please complete the [Client-side changes](client-side-changes-all-frameworks.md)
{% endhint %}

### API routes

Helpers for API routes have been updated to mirror the [new `useAuth` hook](client-side-changes-all-frameworks.md#useauth-introduction) on the client-side.

* `withSession` is deprecated and replaced with `withAuth`
* `requireSession` is deprecated with replaced with `requireAuth`
* Instead of decorating the Request object with a `session` object, it is now decorated with an `auth` object that mirrors `useAuth()` on the client-side.
  * `const { userId, sessionId, getToken } = req.auth;`

Example usage

```javascript
import { withAuth } from "@clerk/nextjs/api";

export default withAuth(async (req, res) => {
  const { userId, sessionId, getToken } = req.auth;
  const hasuraToken = await getToken({template: "hasura"});
  
  // Your handler
});
```

### Edge middleware

Edge middleware has also been updated to mirror the [new `useAuth` hook](client-side-changes-all-frameworks.md#useauth-introduction) on the client-side. The import path has also been changed to avoid confusion.

```javascript
import { withEdgeMiddlewareAuth } from "@clerk/nextjs/edge-middleware";

export const middleware = withEdgeMiddlewareAuth((req, ev) => {
  const { userId, sessionId, getToken } = req.auth;
  
  // Your middleware
});
```
