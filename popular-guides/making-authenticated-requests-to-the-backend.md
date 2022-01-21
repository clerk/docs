# Making authenticated requests to the backend

A request is considered “authenticated” when the backend can securely identify which user and which device is making the request. The reasons for making authenticated requests to the backend include:

* Associating the user with the action being performed
* Ensuring the user has permission to make the request
* Keeping an audit log of which device the user is performing actions from

In order to authenticate the user on the backend using the [Clerk SDK](../reference/backend-api-reference/sdks/), the short-lived session token needs to be passed to the server.

### Frontend

To make authenticated requests from the frontend, the approach differs based on whether your client and server are on the same origin. The origin includes the protocol, hostname, and port (optional):

```html
<protocol>//<hostname>[:<port>]
```

### Same origin

If your client and server are on the same origin (e.g. making an API call to `foo.com/api` from JavaScript running on `foo.com`), the session token is automatically passed to the backend in a cookie. This means all requests to same origin endpoints are authenticated by default.

#### Using Fetch

You can use the native browser [Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch\_API) as you normally would and the request will be authenticated.

```jsx
fetch('/api/foo').then(res => res.json());
```

#### useSWR hook

If you are using React or Next.js and would like to make use of the [`useSWR`](https://swr.vercel.app) hook, it’s as simple as supplying the API route with whichever fetcher function you are using.

```jsx
import useSWR from 'swr';

const fetcher = (...args) => fetch(...args).then(res => res.json());
const { data, error } useSWR('/api/foo', fetcher);
```

#### react-query

If you are using [React Query](https://react-query.tanstack.com) in your application, you need to use a query function that throws errors. Since the native Fetch API does not, you can add it in. Also make sure that you have your application wrapped in `<QueryClientProvider />` with a `QueryClient` set.

```jsx
import { useQuery } from 'react-query';

const { data, error } = useQuery('foo', async () => {
  const res = await fetch('/api/foo');
  if (!res.ok) {
    throw new Error('Network response error')
  }
  return res.json()
});
```

### Cross-origin

If your client and server are on different origins (e.g. `api.foo.com` and `foo.com`), the session token needs to be passed in a network request header. There are a few different ways this can be done on the front-end.

#### Using Fetch with React

In order to pass the session token using the browser Fetch API, it should be put inside a Bearer token in the Authorization header. To retrieve the session token, use the `getToken()` method from the client package (e.g. `@clerk/clerk-react`, `@clerk/nextjs`). Be mindful that `getToken` is an async function that returns a Promise which needs to be resolved.

```javascript
import { useSession } from '@clerk/nextjs';

export default function useFetch() {
  const { getToken } = useSession();
  const authenticatedFetch = async (...args) => {
    return fetch(...args, {
      headers: { Authorization: `Bearer ${await getToken()}` }
    }).then(res => res.json());
  };
  return authenticatedFetch;
}
```

#### useSWR hook

If you are using React or Next.js and would like to make use of the `useSWR` hook, you can create a custom hook with `useSession` from Clerk that supplies the token with the requests.

```jsx
import useSWR from 'swr';
import { useSession } from '@clerk/nextjs';

export default function useClerkSWR(url) {
  const { getToken } = useSession();
  const fetcher = async (...args) => {
    return fetch(...args, {
      headers: { Authorization: `Bearer ${await getToken()}` }
    }).then(res => res.json());
  };
  return useSWR(url, fetcher);
}
```

#### react-query

If you are using [React Query](https://react-query.tanstack.com), it will follow a similar pattern composing the `useSession` hook.

```jsx
import { useQuery } from 'react-query';
import { useSession } from '@clerk/nextjs';

export default function useClerkQuery(url) {
  const { getToken } = useSession();
  return useQuery(url, async () => {
    const res = await fetch(url, {
      headers: { Authorization: `Bearer ${await getToken()}` }
    });
    if (!res.ok) {
      throw new Error('Network response error')
    }
    return res.json()
  });
}
```

#### Using Fetch with ClerkJS

If you are not using React or Next.js, you can access the `getToken` method from the `session` property of the `Clerk` object. This assume you have already followed the instructions on [setting up ClerkJS](../reference/clerkjs/installation.md) and provided it with your Frontend API URL.

```jsx
(async () => {
  fetch('/api/foo', { 
    headers: { 
      Authorization: `Bearer ${await Clerk.session.getToken()}` 
    } 
   }).then(res => res.json());
})();
```

### Backend

If you are using Next.js or Express with Node.js, Clerk provides middleware that sets the session property for easy access and can also require a session be available.

#### Next.js Serverless API Route

In Next.js, the `withSession` and `requireSession` helper functions can be used in your API routes to access the authenticated user.

```jsx
import { withSession } from '@clerk/nextjs/api';

export default withSession((req, res) => {
  if (req.session){
    res.status(200).json({ id: req.session.userId });
  } else {
    res.status(401).json({ id: null });
  } 
});
```

#### Gatsby Functions

The Gatsby Function uses the `requireSession` helper function from the Node SDK directly.

```jsx
import { requireSession } from '@clerk/clerk-sdk-node';

// `requireSession` automatically throws an
// error when no user session is found.

export default requireSession((req, res) => {
  res.status(200).json({ id: req.session.userId })
})
```

#### Go Middleware

The Clerk Go SDK provides a simple middleware that adds the active session to the request’s context. You can see an example code [implementation in Go here](../reference/backend-api-reference/sdks/golang/verifying-a-session.md).

#### Rack Middleware for Ruby on Rails

The Clerk Ruby SDK comes with Rack middleware to lazily load the Clerk session and user. If added as a gem to Rails application, the `Clerk::Authenticatable` concern can be added to your controller. Learn more about how to [integrate Clerk with Rack and Rails](../reference/backend-api-reference/sdks/ruby/rack-rails-integration.md).

#### Manual Authentication

If there is not middleware available for your preferred language or framework, you can extract the session token manually. For same origin requests, the session token is included in the `__session` cookie and you can use an open source library to parse the cookie on the back-end. For cross-origin requests, the Bearer token inside the `Authorization` header contains the session token. You can read more about [validating the session token](validating-session-tokens.md) for additional information.
