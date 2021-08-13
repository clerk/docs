---
description: Learn to install and initialize Clerk in a new Create React App.
---

# Get started with React

## 1. Create a Create React App app

### i. Run create command

Start by creating a new Create React App app.

```bash
npx create-react-app yourapp
# or
yarn create react-app yourapp
```

For more information about these commands, please reference the [Create React App documentation](https://reactjs.org/docs/create-a-new-react-app.html#create-react-app).

### ii. Install clerk-react

Install Clerk's NPM package for React applications.

```bash
# From your application's root directory
cd yourapp

npm install @clerk/clerk-react
# or
yarn add @clerk/clerk-react
```

### iii. Add environment variables

Create a file named **.env.local** in your application root. Any variables inside this file with the **REACT\_APP** prefix will be accessible in your frontend via `process.env.REACT_APP_X`.

{% hint style="warning" %}
Make sure you update ****these variables with the Client Host found in your dashboard.
{% endhint %}

```text
REACT_APP_CLERK_FRONTEND_API=clerk.abc123.lcl.dev
```

### iii. Start the dev server

```bash
npm start
# or
yarn start
```

## 2. Configure src/App.js

Clerk requires every page to be wrapped in the `<ClerkProvider>` context. In Create React App, we add this in **src/App.js**.

In the code below, we've also added a sign in requirement, and show a greeting when a user is signed in.

{% code title="src/App.js" %}
```jsx
import { 
  ClerkProvider, 
  RedirectToSignIn, 
  SignedIn, 
  SignedOut, 
  UserButton, 
  useUser 
} from '@clerk/clerk-react';

// Retrieve Clerk settings from the environment
const clerkFrontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;

function App() {
  return (
    <ClerkProvider frontendApi={clerkFrontendApi}>
      <SignedIn>
        <UserButton />
        <Greeting />
      </SignedIn>
      <SignedOut>
        <RedirectToSignIn />
      </SignedOut>
    </ClerkProvider>
  );
}

function Greeting() {
  const { firstName } = useUser();
  return <div>Hello, {firstName}!</div>;
}

export default App;
```
{% endcode %}

## 3. Add a router

A common next step is to add a router. Here we use [React Router](https://reactrouter.com/) to add public and protected routes to our application.

Before using this code, you must install React Router:

```bash
npm install react-router-dom
# or
yarn add react-router-dom
```

This is just one example of how authentication can be handled with Clerk. Feel free to adjust the strategy to better serve your use case!

{% code title="src/App.js" %}
```jsx
import { 
  ClerkProvider, 
  RedirectToSignIn, 
  SignedIn, 
  SignedOut, 
  UserButton, 
  useUser 
} from '@clerk/clerk-react';
import { 
  BrowserRouter as Router, 
  Link, 
  Route, 
  Switch 
} from 'react-router-dom';

// Retrieve Clerk settings from the environment
const clerkFrontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;

function App() {
  return (
    <Router>
      <ClerkProvider frontendApi={clerkFrontendApi}>
        <Switch>
          {/* Public routes, accesible whether or not a user is signed in */}
          <Route path="/public">
            <div>
              Reached the public route. <Link to="/">Return home.</Link>
            </div>
          </Route>

          {/* Private routes, accesible only if a user is signed in */}
          <PrivateRoute path="/private">
            <div>
              Reached the private route. <Link to="/">Return home.</Link>
            </div>
          </PrivateRoute>

          {/* Catch-all route will render if no other route renders */}
          <Route>
            <SignedIn>
              <UserButton />
              <Greeting />
              <div>You are signed in. You can access both routes.</div>
              <Navigation />
            </SignedIn>
            <SignedOut>
              <div>You are signed out. You can access the public route.</div>
              <Navigation />
            </SignedOut>
          </Route>
        </Switch>
      </ClerkProvider>
    </Router>
  );
}

function PrivateRoute(props) {
  // If the route matches but the user is not signed in, redirect to /sign-in
  return (
    <>
      <SignedIn>
        <Route {...props} />
      </SignedIn>
      <SignedOut>
        <RedirectToSignIn />
      </SignedOut>
    </>
  );
}

function Navigation() {
  return (
    <ul>
      <li>
        <Link to="/public">Public route</Link>
      </li>
      <li>
        <Link to="/private">Private route</Link>
      </li>
    </ul>
  );
}

function Greeting() {
  const { firstName } = useUser();
  return <div>Hello, {firstName}!</div>;
}

export default App;
```
{% endcode %}

## 4. Use your backend

### Configure backend requests

Use our [guide for making backend requests]() to ensure your backend can determine the signed in user.

## 5. Mount components \(optional\)

By default, Clerk hosts the **SignIn**, **SignUp**, and **UserProfile** components on the **accounts.\*** subdomain of your root domain.

If you prefer, you can also "mount" these components directly in your application.

### i. Update src/App.js

The code below uses the same logic as above, but with three additional changes:

1. We add public routes for **SignIn** and **SignUp**, and a private route for **UserProfile**.
2. We create a `<ClerkProviderWithNavigate>` component that uses React Router's **useHistory** hook to pass a **navigate** prop to `<ClerkProvider>`. This allows Clerk to navigate inside and between components without conflicting with React Router.
3. We replace our custom `<RedirectToSignIn>` component with React Router's `<Redirect>` component, since now the navigation is internal.

{% code title="src/App.js" %}
```jsx
import {
  ClerkProvider,
  RedirectToSignIn,
  SignedIn,
  SignedOut,
  SignIn,
  SignUp,
  UserButton,
  UserProfile,
  useUser,
} from '@clerk/clerk-react';
import { 
  BrowserRouter as Router, 
  Link, 
  Route, 
  Switch, 
  useHistory 
} from 'react-router-dom';

// Retrieve Clerk settings from the environment
const clerkFrontendApi = process.env.REACT_APP_CLERK_FRONTEND_API;

function App() {
  return (
    <Router>
      <ClerkProviderWithNavigate>
        <Switch>
          {/* Public routes, accesible whether or not a user is signed in */}
          <Route path="/public">
            <div>
              Reached the public route. <Link to="/">Return home.</Link>
            </div>
          </Route>
          <Route path="/sign-in/(.*)?">
            <SignIn routing="path" path="/sign-in" />
          </Route>
          <Route path="/sign-up/(.*)?">
            <SignUp routing="path" path="/sign-up" />
          </Route>

          {/* Private routes, accesible only if a user is signed in */}
          <PrivateRoute path="/private">
            <div>
              Reached the private route. <Link to="/">Return home.</Link>
            </div>
          </PrivateRoute>
          <PrivateRoute path="/user/(.*)?">
            <UserProfile routing="path" path="/user" />
          </PrivateRoute>

          {/* Catch-all route will render if no other route renders */}
          <Route>
            <SignedIn>
              <UserButton />
              <Greeting />
              <div>You are signed in. You can access both routes.</div>
              <Navigation />
            </SignedIn>
            <SignedOut>
              <div>You are signed out. You can access the public route.</div>
              <Navigation />
            </SignedOut>
          </Route>
        </Switch>
      </ClerkProviderWithNavigate>
    </Router>
  );
}

function Navigation() {
  return (
    <ul>
      <li>
        <Link to="/public">Public route</Link>
      </li>
      <li>
        <Link to="/private">Private route</Link>
      </li>
    </ul>
  );
}

function Greeting() {
  const { firstName } = useUser();
  return <div>Hello, {firstName}!</div>;
}

function PrivateRoute(props) {
  // If the route matches but the user is not signed in, redirect to /sign-in
  return (
    <>
      <SignedIn>
        <Route {...props} />
      </SignedIn>
      <SignedOut>
        <RedirectToSignIn />
      </SignedOut>
    </>
  );
}

function ClerkProviderWithNavigate({ children }) {
  const { push } = useHistory();
  return (
    <ClerkProvider
      frontendApi={clerkFrontendApi}
      navigate={(to) => push(to)}
    >
      {children}
    </ClerkProvider>
  );
}

export default App;
```
{% endcode %}

### ii. Update instance settings

The final step is to update your instance settings so that users will be defaulted to your mounted components instead of the "accounts" subdomain.

To change the settings, open your instance in the [Dashboard](https://dashboard.clerk.dev) and navigate to **Settings** → **URLs and Redirects**.

