# Getting started with Node

## Before you get started

### **Set CLERK\_API\_KEY**

The Node SDK will pick automatically the `CLERK_API_KEY` vale from your environment **** variables. If your application is using .**env** files, create a file named **.env.local** in your application root if it doesn't exist already and add the above variable.

{% hint style="warning" %}
Make sure you update **** this variable with the API key found in your dashboard under **Settings** → **API Keys**.
{% endhint %}

{% code title=".env.local" %}
```jsx
CLERK_API_KEY=test_asdf1234
```
{% endcode %}

For detailed usage to the [official documentation of the Node SDK](https://www.npmjs.com/package/@clerk/clerk-sdk-node).

**Multi session applications**

If Clerk is running in multi-session mode, it's important to ensure your frontend sends the Session ID that is making the request.&#x20;

Our middlewares will look for a query string parameter named **\_clerk\_session\_id.**  If this parameter is not found, the middleware will instead choose the last active session, which may be subject to race conditions and should not be relied on for authenticating actions.

## Next.js middleware

### Optional session

This strategy allows you to detect whether or not there's an active session, and handle each case separately.

{% tabs %}
{% tab title="Javascript" %}
```javascript
import { withSession } from '@clerk/clerk-sdk-node';

function handler(req, res) {
    if (req.session) {
        // do something with session.userId
    } else {
        // Respond with 401 or similar
    }
}

export default withSession(handler);
```
{% endtab %}

{% tab title="Typescript" %}
```typescript
import { withSession, WithSessionProp } from '@clerk/clerk-sdk-node';

function handler(req: WithSessionProp<NextApiRequest>, res: NextApiResponse) {
    if (req.session) {
        // do something with session.userId
    } else {
        // Respond with 401 or similar
    }
}

export withSession(handler);
```
{% endtab %}
{% endtabs %}

### Required session

This strategy mandates that a session be available.  If not, it returns a 401 (no body) and your handler is never called.

{% tabs %}
{% tab title="Javascript" %}
```javascript
import { requireSession } from '@clerk/clerk-sdk-node';

function handler(req, res) {
    // do something with session.userId
}

export default requireSession(handler)
```
{% endtab %}

{% tab title="Typescript" %}
```typescript
import { requireSession, RequireSessionProp } from '@clerk/clerk-sdk-node';

function handler(req: RequireSessionProp<NextApiRequest>, res: NextApiResponse) {
    // do something with session.userId
}

export requireSession(handler)
```
{% endtab %}
{% endtabs %}

## Express middleware

For usage with <a href="https://github.com/expressjs/express" target="_blank">Express</a>, this package also exports `ClerkExpressWithSession` (lax) & `ClerkExpressRequireSession` (strict)
middlewares that can be used in the standard manner:

```ts
import { ClerkWithSession } from '@clerk/clerk-sdk-node';

// Initialize express app the usual way

app.use(ClerkWithSession());
```

The `ClerkWithSession` middleware will set the Clerk session on the request object as `req.session` and then call the next middleware.

You can then implement your own logic for handling a logged-in or logged-out user in your express endpoints or custom
middleware, depending on whether your users are trying to access a public or protected resource.

If you want to use the express middleware of your custom `Clerk` instance, you can use:

```ts
app.use(clerk.expressWithSession());
```

Where `clerk` is your own instance.

If you prefer that the middleware renders a 401 (Unauthenticated) itself, you can use the following variant instead:

```ts
import { ClerkExpressRequireSession } from '@clerk/clerk-sdk-node';

app.use(ClerkExpressRequireSession());
```

### onError option

The Express middleware supports an `options` object as an optional argument.
The only key currently supported is `onError` for providing your own error handler.

The `onError` function, if provided, should take an `Error` argument (`onError(error)`).

Depending on the return value, it can affect the behavior of the middleware as follows:

- If an `Error` is returned, the middleware will call `next(err)` with that error. If the `err` has a `statusCode` it will indicate to Express what HTTP code the response should have.
- If anything other than an `Error` is returned (or nothing is returned at all), then the middleware will call `next()` without arguments

The default implementations unless overridden are:

```ts
// defaultOnError swallows the error
defaultOnError(error: Error) {
  console.error(error.message);
}

// strictOnError returns the error so that Express will halt the request chain
strictOnError(error: Error) {
  console.error(error.message);
  return error;
}
```

`defaultOnError` is used in the lax middleware variant and `strictOnError` in the strict variant.

### Express Error Handlers

Not to be confused with the `onError` option mentioned above, Express comes with a default error handler for errors encountered in the middleware chain.

Developers can also implement their own custom error handlers as detailed <a href="https://expressjs.com/en/guide/error-handling.html" target="_blank">here</a>.

An example error handler can be found in the [Express examples folder](https://github.com/clerkinc/clerk-sdk-node/tree/main/examples/express):

```js
// Note: this is just a sample errorHandler that pipes clerk server errors through to your API responses
// You will want to apply different handling in your own app to avoid exposing too much info to the client
function errorHandler(err, req, res, next) {
  const statusCode = err.statusCode || 500;
  const body = err.data || { error: err.message };

  res.status(statusCode).json(body);
}
```

If you are using the strict middleware variant, the `err` pass to your error handler will contain enough context for you to respond as you deem fit.

## Manual authentication

### Authenticate a particular session

Highly recommended for authenticating actions. &#x20;

```javascript
import { sessions } from '@clerk/clerk-sdk-node';
import Cookies from 'cookies';

// Retrieve the particular session ID from a
// query string parameter
const sessionId = req.query._clerk_session_id;

// Note: Clerk stores the clientToken in a cookie 
// named "__session" for Firebase compatibility
const cookies = new Cookies(req, res);
const clientToken = cookies.get('__session');  

const session = await sessions.verifySession(sessionId, clientToken);
```

### Authenticate the last active session

Using the last active session is appropriate when determining the user after a navigation.

```javascript
import { clients, sessions } from '@clerk/clerk-sdk-node';

// Note: Clerk stores the clientToken in a cookie 
// named "__session" for Firebase compatibility
const cookies = new Cookies(req, res);
const clientToken = cookies.get('__session');  

const client = await clients.verifyClient(sessionToken);
const sessionId = client.lastActiveSessionId;

const session = await sessions.verifySession(sessionId, clientToken);
```
