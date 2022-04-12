# Express server-side changes

{% hint style="info" %}
Before starting this guide, please complete the [Client-side changes](client-side-changes-all-frameworks.md)
{% endhint %}

### Middleware

Express middleware been updated to mirror the [new `useAuth` hook](client-side-changes-all-frameworks.md#useauth-introduction) on the client-side.

* `ClerkExpressWithSession` is deprecated and replaced with `ClerkExpressWithAuth`
* `ClerkExpressRequireSession` is deprecated and replaced with `ClerkExpressRequireAuth`
* Instead of decorating the Request object with a `session` object, it is now decorated with an `auth` object that mirrors `useAuth()` on the client-side.
  * `const { userId, sessionId, getToken } = req.auth;`

Example usage

```javascript
import { ClerkExpressRequireAuth } from '@clerk/clerk-sdk-node';

app.use(ClerkExpressRequireAuth());
app.get('/', (req, res) => {
  const { userId, sessionId, getToken } = req.auth;
  // Your handler...
})
```
