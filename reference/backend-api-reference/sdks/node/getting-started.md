# Getting started with Node

## Before you get started

### **Set CLERK\_API\_KEY**

The Node SDK will pick automatically the `CLERK_API_KEY` vale from your environment ****variables. If your application is using .**env** files, create a file named **.env.local** in your application root if it doesn't exist already and add the above variable.

{% hint style="warning" %}
Make sure you update ****this variable with the API key found in your dashboard under **Settings** → **API Keys**.
{% endhint %}

{% code title=".env.local" %}
```jsx
CLERK_API_KEY=test_asdf1234
```
{% endcode %}

For detailed usage to the [official documentation of the Node SDK](https://www.npmjs.com/package/@clerk/clerk-sdk-node).

**Multi session applications**

If Clerk is running in multi-session mode, it's important to ensure your frontend sends the Session ID that is making the request. 

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

function handler(req WithSessionProp<NextApiRequest>, res: NextApiResponse) {
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

This strategy mandates that a session be available.  If not, it returns a 401 \(no body\) and your handler is never called.

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

function handler(req RequireSessionProp<NextApiRequest>, res: NextApiResponse) {
    // do something with session.userId
}

export requireSession(handler)
```
{% endtab %}
{% endtabs %}

## Express middleware

```javascript
import { ClerkExpressMiddleware } from '@clerk/clerk-sdk-node';

app.use(ClerkExpressMiddleware());
```

## Manual authentication

### Authenticate a particular session

Highly recommended for authenticating actions.  

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

