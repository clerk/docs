# Client-side changes (all frameworks)

### Upgrade npm

**Next.js**

```
npm install @clerk/nextjs@latest
```

**Gatsby, Redwood, React**

```
npm install @clerk/clerk-react@latest
```

### useAuth() introduction

`useAuth()` is a new, SSR-compatible hook that is recommended for all **authentication** tasks.

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

* `useSession()` to access `getToken()` or the `sessionId`
* `useUser()` to access `getToken()` for integrations
* `useClerk()` to access `signOut()`
* `<SignedIn>` and `<SignedOut>` in a way that requires extra components

### Hook and component API changes

Version 3 changes several APIs for improved consistency and usability.

#### useUser() and useSession()

Previously, `useUser` and `useSession` could only be called from inside the `<SignedIn/>` component. This restriction has been relaxed, but now developers must manually handle the possibility the `user` and `session` objects may not be loaded when the hook is called.

In a future version, we plan to support React `<Suspense/>` to make this even easier.

| Old API                         | New API                                                                                                                                                                                                                                                                                                            |
| ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `const user = useUser();`       | `const { isLoaded, isSignedIn, user } = useUser();`                                                                                                                                                                                                                                                                |
| `const session = useSession();` | <p><code>const { isLoaded, isSignedIn, session } = useSession();</code><br><code></code><br><code></code>If you only call <code>useSession</code> for the ID or <code>getToken</code>, please switch to <a href="client-side-changes-all-frameworks.md#useauth-introduction"><code>useAuth</code></a> instead.</p> |

#### \<RedirectToSignIn /> and \<RedirectToSignUp/>

The behavior enabled by the `returnBack` prop has now been promoted to default behavior. `returnBack` is deprecated in favor of using an empty

| Old API                          | New API                |
| -------------------------------- | ---------------------- |
| `<RedirectToSignIn returnBack/>` | `<RedirectToSignIn />` |
| `<RedirectToSignIn returnBack/>` | `<RedirectToSignUp />` |

#### \<SignIn /> and \<SignUp />

Props which take URLs (absolute or relative) have been standardized to use the `Url` (camelCase) suffix

| Old props     | New props        |
| ------------- | ---------------- |
| `afterSignIn` | `afterSignInUrl` |
| `afterSignUp` | `afterSignUpUrl` |
| `signUpURL`   | `signUpUrl`      |
| `signInURL`   | `signInUrl`      |

#### \<UserButton />

Props which take URLs (absolute or relative) have been standardized to use the `Url` (camelCase) suffix

| Old API              | New API                             |
| -------------------- | ----------------------------------- |
| `afterSwitchSession` | `afterSwitchSessionUrl`             |
| `userProfileURL`     | `userProfileUrl`                    |
| `signInURL`          | `signInUrl`                         |
| `afterSignOutAll`    | `afterSignOutUrl`                   |
| `afterSignOutOne`    | `afterMultiSessionSingleSignOutUrl` |

#### useSignUp() & useSignIn()

`useSignUp()` and `useSignIn()` are no longer required to be contained in `<SignedOut/>` or `<ClerkLoaded/>` . Instead, they have been updated to include a loading state:

| Old API                       | New API                                     |
| ----------------------------- | ------------------------------------------- |
| `const signUp = useSignUp();` | `const { isLoaded, signUp } = useSignUp();` |
| `const signIn = useSignIn();` | `const { isLoaded, signIn } = useSignIn();` |

### Resource changes

**Method signature changes**

In order to make our APIs easier to use in codebases not using Typescript and make them more consistent across frameworks, we removed support for snake\_cased method parameters. We also updated all method signatures to always accept a param object instead of plain positional params.

We expect that most users will only need to apply these changes to the `user.update()` , `user.createEmailAddress()` and `user.createPhoneNumber()` methods as shown below.

Users who have built custom auth flows using the SignIn and SignUp resources can consult the [SignUp and SignIn](client-side-changes-all-frameworks.md#signup-and-signin) section below.

#### User

`User.getToken("token-name")` has been deprecated and is now accessible through `getToken` from the [`useAuth`](client-side-changes-all-frameworks.md#useauth-introduction) hook or `Session.getToken()`&#x20;

The new `getToken()` method accepts an optional `{ template: string; }` object, used to specify the JWT template you want to use.&#x20;

Please note that If you want to fetch a token for an integration, you **must** prefix the integration name with  `integration_`, eg: `getToken({ template: 'integration_firebase' })`

Moreover, the following methods now accept a params objects instead of positional params: `createEmailAddress()`, `createPhoneNumber()`, `setProfileImage()`

Respecting the camelCase changes we explained at the start of this section, all method params should now use camelCasing. An example for the `create` method is shown below.



| Old API                                           | New API                                                                                                                                                                                                |
| ------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `User.getToken("token-name")`                     | <p><code>const { getToken } = useAuth()</code><br><code>getToken({template: "integration_token-name"})</code><br><br>or<br><br><code>Session.getToken({template: "integration_token-name"})</code></p> |
| `User.createEmailAddress(value)`                  | `User.createEmailAddress({ email: value })`                                                                                                                                                            |
| `User.createPhoneNumber(value)`                   | `User.createPhoneNumber({ phoneNumber: value })`                                                                                                                                                       |
| `User.setProfileImage(file)`                      | `User.setProfileImage({ file })`                                                                                                                                                                       |
| `User.update({ first_name: '', last_name: '' })`  | `User.update({ firstName: '', lastName: '' })`                                                                                                                                                         |

#### SignUp and SignIn

The `createMagicLinkFlow` and `authenticateWithRedirect` calls have had parameters renamed to use the word "redirect" instead of "callback."

| Old API                                                          | New API                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `createMagicLinkFlow({ callbackUrl })`                           | `createMagicLinkFlow({ redirectUrl })`                           |
| `authenticateWithRedirect({ callbackUrl, callbackUrlComplete })` | `authenticateWithRedirect({ redirectUrl, redirectUrlComplete })` |
