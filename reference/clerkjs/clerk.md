---
description: The Clerk object is the core of the ClerkJS SDK.
---

# Clerk

## Overview

The `Clerk` object is a singleton which can act as the entry point for gaining access to other Clerk resources, like the  active [Client](client.md), [Session](session.md) and [User](user.md) objects. It also includes helper methods for mounting [Clerk Components](broken-reference) to your pages.

The `Clerk` object is always available via `window.Clerk`.

## Attributes

| Name        | Description                                                                                                                                                                                                                                                                                                                                                              |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **client**  | <p><em></em><a href="client.md"><em>Client</em></a><em></em></p><p>The <code>Client</code> object for the current window.</p>                                                                                                                                                                                                                                            |
| **session** | <p><em></em><a href="session.md"><em>Session</em></a> <em>| null | undefined</em></p><p>The currently active <code>Session</code>, which is guaranteed to be one of the sessions in <code>Client.sessions</code>. If there is no active session, this field will be <strong>null</strong>. If the session is loading, this field will be <strong>undefined</strong>.</p> |
| **user**    | <p><em></em><a href="user.md"><em>User</em></a> <em>| null | undefined</em></p><p>A shortcut to <code>Session.user</code> which holds the currently active <code>User</code> object. If the session is <strong>null</strong> or <strong>undefined</strong>, the <code>user</code> field will match.</p>                                                                  |
| **version** | <p><em>string</em></p><p>The <a href="./">ClerkJS</a> SDK version.</p>                                                                                                                                                                                                                                                                                                   |

## Methods

### addListener(listener)

`addListener(listener: (resources: Resources) => void) => UnsubscribeCallback`

Registers a listener that triggers a callback whenever a change in the`Client`, `Session`, or `User` object occurs. This method is primary used to build frontend SDKs like [@clerk/clerk-react](https://www.npmjs.com/package/@clerk/clerk-react).

`Resources` is an interface with the following definition. To import our types, please add [@clerk/types](https://www.npmjs.com/package/@clerk/types):

```typescript
interface Resources {
  client: ClientResource;
  session?: ActiveSessionResource | null;
  user?: UserResource | null;
}
```

Please note that the `session`  and `user` object have a special relationship that the type definition alone does not capture:

* When there is an active session, `user === session.user`**.**
* When there is no active session, `user` **** and `session` **** will both be **null**.
* When a session is loading, `user` **** and `session` **** will be **undefined**.

{% tabs %}
{% tab title="Parameters" %}
| Name         | Description                                                                                                                |
| ------------ | -------------------------------------------------------------------------------------------------------------------------- |
| **listener** | <p><em>(resources: Resources) => void</em></p><p>A function to be called when the <code>Client</code> object changes. </p> |
{% endtab %}

{% tab title="Returns" %}
_() => void_

This method returns a function that can be used to clean up the registered listener.&#x20;
{% endtab %}
{% endtabs %}

### authenticateWithMetamask(params)

`authenticateWithMetamask(params: AuthenticateWithMetamaskParams) => void`

Starts an authentication flow that uses the Metamask browser extension to authenticate the user using their public wallet address.

{% tabs %}
{% tab title="Parameters" %}
| Name | Description                                                                                                                                                                                                                                 |
| ---- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|      | <p><em></em><a href="clerk.md#handlemagiclinkverificationparams"><em>AuthenticateWithMetamaskParams</em></a><em></em></p><p>These props allow you to define the URL where the user should be redirected to on successful authentication</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### closeSignIn()

`closeSignIn() => void`

Closes the [\<SignIn/>](../../components/sign-in/sign-in.md) modal.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### closeSignUp()

`closeSignUp() => void`

Closes the [\<SignUp/>](../../components/sign-up/sign-up.md) modal.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### handleMagicLinkVerification(params)

`handleRedirectCallback(params: HandleMagicLinkVerificationParams) => Promise<unknown>`

Completes a magic link verification flow once we've reached the magic link results URL.

When users click on magic links they get redirected to the URL that was provided during magic link verification flow initialization. The URL will contain a couple of important query parameters added by Clerk. These are called `__clerk_status` and `__clerk_created_session`.

The `__clerk_status` query parameter is the outcome of the verification and can take three values: **verified**, **failed** or **expired**.

The `__clerk_created_session` query parameter will hold the ID of the new session, if one was created as a result of the verification. Since the magic link can be opened at any device and not the one that originated the verification, the new session ID might not be available under [Client.sessions](client.md#attributes).

Magic link flows can be completed on the same device that they were initiated or on a completely different browser. For example, a user might start the magic link flow on their desktop browser, but click on the magic link from their mobile phone.

The `handleMagicLinkVerification()` method takes care of finalizing the magic link flow, depending on the verification outcome.&#x20;

Upon successful verification, the method will figure out if the sign in or sign up attempt was completed and redirect the user accordingly. As such, it accepts different parameters for the URL it should redirect when sign in/up is completed and the URL which it should redirect when the sign in/up attempt is still pending. Both parameters are optional, but you can provide them to the method up front. The final redirect will depend on the sign in/up attempt's status.

Additionally, the `handleMagicLinkVerification()` method allows you to execute a callback if the successful verification happened on another device.

In case the magic link verification wasn't successful, the `handleMagicLinkVerification()` method will throw a [`MagicLinkError`](clerk.md#undefined). You can check the error's `code` property to see if the magic link expired, or the verification simply failed.

Take a look at the function parameters description below for more usage details.

{% tabs %}
{% tab title="Parameters" %}
| Name       | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **params** | <p><em></em><a href="clerk.md#handleoauthcallbackparams"><em>HandleMagicLinkVerificationParams</em></a><em></em></p><p>These props allow you to define the URLs where the user should be redirected to on successful verification and:</p><ul><li>completed sign in or sign up attempt.</li><li>pending sign in or sign up attempt.</li></ul><p>If the magic link is successfully verified on another device, there's a callback function parameter that allows custom code execution.</p> |


{% endtab %}

{% tab title="Throws" %}
__[_MagicLinkError_](clerk.md#undefined)__

This method will throw a `MagicLinkError` if the magic link verification failed or the link expired. Check the error's `code` property for details.
{% endtab %}
{% endtabs %}

### handleRedirectCallback(params)

`handleRedirectCallback(params: HandleOAuthCallbackParams) => Promise<void>`

Completes a custom OAuth flow started by calling either [`SignIn.authenticateWithRedirect(params)`](signin.md#signinwithoauth) or [`SignUp.authenticateWithRedirect(params)`](signup.md#signinwithoauth)

{% tabs %}
{% tab title="Parameters" %}
| Name        | Description                                                                                                                                                                                                                 |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **params?** | <p><em></em><a href="clerk.md#handleoauthcallbackparams"><em>HandleOAuthCallbackParams</em></a><em></em></p><p>Additional props that define where the user will be redirected to at the end of a successful OAuth flow.</p> |
{% endtab %}
{% endtabs %}

### isReady()

`isReady() => boolean`

Returns **true** when the [ClerkJS](./) library has fully loaded and the `Clerk` object is ready for use, **false** otherwise.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_boolean_

This method will return **true** when the `Clerk` object is ready, **false** otherwise.
{% endtab %}
{% endtabs %}

### load(options?)

`load(options?: ComponentOptions) => Promise<void>`

Initializes the `Clerk` object and loads all necessary environment configuration and Instance settings from the [Frontend API](../frontend-api-reference/).

It is absolutely necessary to call this method before using the `Clerk` object in your code. Refer to the [ClerkJS installation](installation.md#setting-up-clerkjs) guide for more details.

{% tabs %}
{% tab title="Parameters" %}
| Name        | Description                                                                                                                                                                                                                    |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **options** | <p><em></em><a href="clerk.md#componentoptions"><em>ComponentOptions</em></a><em></em></p><p>Configuration and options for initializing the <code>Clerk</code> object and <a href="broken-reference">Clerk Components</a>.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

### mountSignIn(node, nodeProps?)

`mountSignIn(node: HTMLDivElement, nodeProps?: SignInProps) => void`

Renders a [\<SignIn/>](../../components/sign-in/sign-in.md) component inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
| Name           | Description                                                                                                                                                                                                    |
| -------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **node**       | <p><em>HTMLDivElement</em></p><p>An HTML <code>&#x3C;div/></code> element which will render the <a href="../../components/sign-in/sign-in.md">&#x3C;SignIn/></a> component.</p>                                |
| **nodeProps?** | <p><em></em><a href="clerk.md#signinprops"><em>SignInProps</em></a><em></em></p><p>Additional props that will be passed to the <a href="../../components/sign-in/sign-in.md">&#x3C;SignIn/></a> component.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### mountSignUp(node, nodeProps?)

`mountSignUp(node: HTMLDivElement, nodeProps?: SignUpProps) => void`

Renders a [\<SignUp/>](../../components/sign-up/sign-up.md) component inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
| Name           | Description                                                                                                                                                                                                    |
| -------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **node**       | <p><em>HTMLDivElement</em></p><p>An HTML <code>&#x3C;div/></code> element which will render the <a href="../../components/sign-up/sign-up.md">&#x3C;SignUp/></a> component.</p>                                |
| **nodeProps?** | <p><em></em><a href="clerk.md#signupprops"><em>SignUpProps</em></a><em></em></p><p>Additional props that will be passed to the <a href="../../components/sign-up/sign-up.md">&#x3C;SignUp/></a> component.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### mountUserButton(node, nodeProps?)

`mountUserButton(node: HTMLDivElement, nodeProps?: UserButtonProps) => void`

Renders a [\<UserButton/>](../../components/user-button.md) component for the active user inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
| Name         | Description                                                                                                                                                                                                            |
| ------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **node**     | <p><em>HTMLDivElement</em></p><p>An HTML <code>&#x3C;div/></code> element which will render the <a href="../../components/user-button.md">&#x3C;UserButton/></a> component.</p>                                        |
| **options?** | <p><em></em><a href="clerk.md#userbuttonprops"><em>UserButtonProps</em></a><em></em></p><p>Additional props that will be passed to the <a href="../../components/user-button.md">&#x3C;UserButton/></a> component.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### mountUserProfile(node, nodeProps?)

`mountUserProfile(node: HtmlDivElement, nodeProps?: UserProfileProps) => void`

Renders a [\<UserProfile/>](../../components/user-button.md) component for the active user inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
| Name           | Description                                                                                                                                                                                                                             |
| -------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **node**       | <p><em>HTMLElement</em></p><p>The element that the UserProfile should be mounted in.</p>                                                                                                                                                |
| **nodeProps?** | <p><em></em><a href="clerk.md#userprofileprops"><em>UserProfileProps</em></a><em></em></p><p>Additional props that will be passed to the <a href="../../components/user-profile/user-profile.md">&#x3C;UserProfile/></a> component.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### navigate(to)

`navigate(to: string) => Promise<unknown>`

Helper method which will use the custom push navigation function of your application to navigate to the provided URL or relative path. See the relevant [section on routing](broken-reference) for more information on navigation.

{% tabs %}
{% tab title="Parameters" %}
| Name   | Description                                                            |
| ------ | ---------------------------------------------------------------------- |
| **to** | <p><em>string</em></p><p>Full URL or relative path to navigate to.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<unknown>_

This method returns a `Promise` that can resolve with any value.
{% endtab %}
{% endtabs %}

### openSignIn(props)

`openSignIn(props: SignInProps) => void`

Opens the [\<SignIn/>](../../components/sign-in/sign-in.md)component as a modal.

{% tabs %}
{% tab title="Parameters" %}
| Name      | Description                                                                                                                                                                                                            |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **props** | <p><em></em><a href="clerk.md#signinprops"><em>SignInProps</em></a><em></em></p><p>Configuration properties that will be passed to the <a href="../../components/sign-in/sign-in.md">&#x3C;SignIn/></a> component.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### openSignUp(props)

`openSignUp(props: SignUpProps) => void`

Opens the [\<SignUp/>](../../components/sign-up/sign-up.md) component as a modal.

{% tabs %}
{% tab title="Parameters" %}
| Name      | Description                                                                                                                                                                                                            |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **props** | <p><em></em><a href="clerk.md#signupprops"><em>SignUpProps</em></a><em></em></p><p>Configuration properties that will be passed to the <a href="../../components/sign-up/sign-up.md">&#x3C;SignUp/></a> component.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### redirectToSignIn()

`redirectToSignIn() => Promise<unknown>`

Redirects to the sign in URL, as configured in your Instance settings. This method uses the [Clerk.navigate](clerk.md#navigate-to) method under the hood.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
Promise\<unknown>

This method returns a `Promise` which can resolve to any value.
{% endtab %}
{% endtabs %}

### redirectToSignUp()

`redirectToSignUp() => Promise<unknown>`

Redirects to the sign up URL, as configured in your Instance settings. This method uses the [Clerk.navigate](clerk.md#navigate-to) method under the hood.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
Promise\<unknown>

This method returns a `Promise` which can resolve to any value.
{% endtab %}
{% endtabs %}

### redirectToUserProfile()

`redirectToUserProfile() => Promise<unknown>`

Redirects to the user profile management URL, as configured in your Instance settings. This method uses the [Clerk.navigate](clerk.md#navigate-to) method under the hood.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
Promise\<unknown>

This method returns a `Promise` which can resolve to any value.
{% endtab %}
{% endtabs %}

### setSession(session, beforeEmit?)

`setSession(session: SessionResource | string | null, beforeEmit?: (session: SessionResource | null) => any) => Promise<void>`

Set the current session on this client to the provided session. The provided session can be either a complete [Session](session.md) object or simply its unique identifier.

Passing null as the session will result in the current session to be removed from the client.

If an active session already exists, it will be replaced with the new one. The change happens in three steps:

1. The current `Session` object is set to **undefined**, which causes the control components to stop rendering their children as though Clerk is still initializing.
2. The **beforeEmit** callback is executed. If a `Promise` is returned, Clerk waits for the `Promise` to resolve.
3. The current `Session` is set to the passed **session**. This causes the control components to render their children again.

{% tabs %}
{% tab title="Parameters" %}
| Name            | Description                                                                                                                                                                                                                                                                                           |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **session**     | <p><em></em><a href="session.md"><em>SessionResource</em> </a><em>| string | null</em></p><p>A <code>Session</code> object or <code>Session</code> ID string to be set as the current session, or <strong>null</strong> to simply remove the active session, without setting a new one.</p>           |
| **beforeEmit?** | <p><em>(session:</em> <a href="session.md"><em>SessionResource</em></a> <em>| null) => Promise&#x3C;any> | void</em></p><p>Callback that will trigger when the current session is set to <strong>undefined</strong>, before finally being set to the passed session. Usually used for navigation.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which doesn't resolve to any value. The `Promise` will resolve after the passed `session` is set.
{% endtab %}
{% endtabs %}

### signOut(callback?)

`signOut(callback?: SignOutCallback) => Promise<void>`

Signs out the active user from all sessions in a [multi-session application](broken-reference), or simply the current session in a single-session context. The current client will be deleted.

{% tabs %}
{% tab title="Parameters" %}
| Name          | Description                                                                                                                                                       |
| ------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **callback?** | <p><em></em><a href="clerk.md#signoutcallback"><em>SignOutCallback</em></a><em></em></p><p>A callback function that will be called after successful sign out.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which does not resolve to any value.
{% endtab %}
{% endtabs %}

### signOutOne(callback?)

`signOutOne(callback?: SignOutCallback) => Promise<void>`

Signs out the active user from the current session. For a [multi-session application](broken-reference) this means that the rest of the session for the current client will remain, but the user will not be authenticated.

{% tabs %}
{% tab title="Parameters" %}
| Name          | Description                                                                                                                                                       |
| ------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **callback?** | <p><em></em><a href="clerk.md#signoutcallback"><em>SignOutCallback</em></a><em></em></p><p>A callback function that will be called after successful sign out.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which does not resolve to any value.
{% endtab %}
{% endtabs %}

### unmountSignIn(node)

`unmountSignIn(node: HTMLDivElement) => void`

Unmounts the [\<SignIn/>](../../components/sign-in/sign-in.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
| Name     | Description                                                                                                                                                  |
| -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **node** | <p><em>HTMLDivElement</em></p><p>The element that the <a href="../../components/sign-in/sign-in.md">&#x3C;SignIn/></a> component will be unmounted from.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### unmountSignUp(node)

`unmountSignUp(node: HTMLDivElement) => void`

Unmounts the [\<SignUp/>](../../components/sign-up/sign-up.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
| Name     | Description                                                                                                                                                  |
| -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **node** | <p><em>HTMLDivElement</em></p><p>The element that the <a href="../../components/sign-up/sign-up.md">&#x3C;SignUp/></a> component will be unmounted from.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### unmountUserButton(node)

`unmountUserButton(node: HTMLDivElement) => void`

Unmounts the [\<UserButton/>](../../components/user-button.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
| Name     | Description                                                                                                                                                  |
| -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **node** | <p><em>HTMLDivElement</em></p><p>The element that the <a href="../../components/user-button.md">&#x3C;UserButton/></a> component will be unmounted from.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### unmountUserProfile(node)

`unmountUserProfile(node: HTMLDivElement) => void`

Unmounts the [\<UserProfile/>](../../components/user-profile/user-profile.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
| Name     | Description                                                                                                                                                                 |
| -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **node** | <p><em>HTMLDivElement</em></p><p>The element that the <a href="../../components/user-profile/user-profile.md">&#x3C;UserProfile/></a> component will be unmounted from.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

## Interfaces

### ComponentOptions

| Property                  | Description                                                                                                                                                                                                                                                                           |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **selectInitialSession?** | <p><em>(client:</em> <a href="client.md"><em>ClientResource</em></a><em>) =></em> <a href="session.md"><em>SessionResource</em></a> <em>| undefined</em></p><p>This function can be used to set the initial session in <a href="broken-reference">multi-session applications</a>.</p> |
| **navigate?**             | <p><em>(to: string) => Promise&#x3C;unknown> | unknown</em></p><p>Provide an implementation for the <a href="clerk.md#navigate-to">Clerk.navigate</a> method.</p>                                                                                                                     |

### AuthenticateWithMetamaskParams



| Property         | Description                                                                                    |
| ---------------- | ---------------------------------------------------------------------------------------------- |
| **redirectUrl?** | <p><em>string</em></p><p>Full URL or path to navigate after successful sign in or sign up.</p> |

### HandleMagicLinkVerificationParams

| Name                         | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| ---------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **redirectUrl?**             | <p><em>string</em></p><p>Full URL or path to navigate after successful magic link verification on the same device, but a sign in or sign up attempt that cannot complete yet. For example, let's say a user has enabled <a href="broken-reference">two-factor authentication</a> (2FA). In the sign in flow, the magic link verification may be successfully verified but user needs to complete 2FA before they can log in. Use this parameter to redirect to your 2FA route.</p> |
| **redirectUrlComplete?**     | <p><em>string</em></p><p>Full URL or path to navigate after successful magic link verification on the same device and the sign in or sign up attempt has completed. As an example, user tries to sign in via a magic link and they're successfully logged in. Use this parameter to decide where the user will be redirected to.</p>                                                                                                                                               |
| **onVerifiedOnOtherDevice?** | <p><em>Function</em></p><p>Use this callback function to run any custom code on a successful magic link verification on a device different than the one which originated the magic link verification flow. This function provides an opportunity to handle the case where a magic link is opened from another device.</p>                                                                                                                                                          |

### HandleOAuthCallbackParams

| Name                 | Description                                                                                                                                                                                        |
| -------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **afterSignInUrl?**  | <p><em>string</em></p><p>Full URL or path to navigate after successful sign in.</p>                                                                                                                |
| **afterSignUpUrl?**  | <p><em>string</em></p><p>Full URL or path to navigate after successful sign up.</p>                                                                                                                |
| **redirectUrl?**     | <p><em>string</em></p><p>Full URL or path to navigate after successful sign in  or sign up. The same as setting <code>afterSignInUrl</code> and <code>afterSignUpUrl</code> to the same value.</p> |
| **secondFactorUrl?** | <p><em>string</em></p><p>Full URL or path to navigate during sign in, if 2FA is enabled.</p>                                                                                                       |

### SignInProps

| Property         | Description                                                                                                                                                                                                                                                     |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **routing?**     | <p><em>string</em></p><p>The routing strategy for your pages. Supported values are:</p><ul><li><strong>path</strong>: Path based routing.</li><li><strong>hash</strong>: Hash based routing.</li><li><strong>virtual</strong>: Virtual based routing.</li></ul> |
| **path?**        | <p><em>string</em></p><p>The root URL where the component is mounted on.</p>                                                                                                                                                                                    |
| **afterSignIn?** | <p><em>string</em></p><p>The full URL or path to navigate after a successful sign in.</p>                                                                                                                                                                       |
| **signUpURL?**   | <p><em>string</em></p><p>Full URL or path to the sign up page. Use this property to provide the target of the "Sign Up" link that's rendered.</p>                                                                                                               |

### SignUpProps

| Property         | Description                                                                                                                                                                                                                                                     |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **routing?**     | <p><em>string</em></p><p>The routing strategy for your pages. Supported values are:</p><ul><li><strong>path</strong>: Path based routing.</li><li><strong>hash</strong>: Hash based routing.</li><li><strong>virtual</strong>: Virtual based routing.</li></ul> |
| **path?**        | <p><em>string</em></p><p>The root URL where the component is mounted on.</p>                                                                                                                                                                                    |
| **afterSignUp?** | <p><em>string</em></p><p>The full URL or path to navigate after a successful sign up.</p>                                                                                                                                                                       |
| **signInURL?**   | <p><em>string</em></p><p>Full URL or path to the sign in page. Use this property to provide the target of the "Sign In" link that's rendered.</p>                                                                                                               |

### UserButtonProps

| Property                | Description                                                                                                                                                                                       |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **afterSignOutAll?**    | <p><em>string</em></p><p>Full URL or relative path to navigate after sign out is complete and there are no active sessions on this client.</p>                                                    |
| **afterSignOutOne?**    | <p><em>string</em></p><p>Full URL or relative path to navigate after sign out is complete.</p>                                                                                                    |
| **afterSwitchSession?** | <p><em>string</em></p><p>Full URL or relative path to navigate after a successful account change. This property is used only for <a href="broken-reference">multi-session applications</a>.</p>   |
| **signInURL?**          | <p><em>string</em></p><p>Full URL or relative path to navigate on the "Add another account" action. This property is used only for <a href="broken-reference">multi-session applications</a>.</p> |
| **userProfileURL?**     | <p><em>string</em></p><p>Full URL or relative path of the user account management interface.</p>                                                                                                  |

### UserProfileProps

| Property            | Description                                                                                                                                                                                                                                                     |
| ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **routing?**        | <p><em>string</em></p><p>The routing strategy for your pages. Supported values are:</p><ul><li><strong>path</strong>: Path based routing.</li><li><strong>hash</strong>: Hash based routing.</li><li><strong>virtual</strong>: Virtual based routing.</li></ul> |
| **path?**           | <p><em>string</em></p><p>The root URL where the component is mounted on.</p>                                                                                                                                                                                    |
| **hideNavigation?** | <p><em>boolean</em></p><p>Hides the default navigation bar. Can be used when a custom navigation bar is built.</p>                                                                                                                                              |
| **only?**           | <p><em>string</em></p><p>Renders only a specific page of the UserProfile component. Supported values are:</p><ul><li><strong>account</strong>: User account page.</li><li><strong>security</strong>: User security page.</li></ul>                              |

## Types

### MagicLinkError

Custom error for magic links. Raised when the magic link verification doesn't succeed, either because the link has expired or a general failure. The error's `code` property will indicate the outcome, its values being:

* `MagicLinkErrorCode.Expired`
* `MagicLinkErrorCode.Failed`

### SignOutCallback

`() => void | Promise<any>`

| Value             | Description                                 |
| ----------------- | ------------------------------------------- |
| **() => void**    | A `Function` which doesn't return anyhing.  |
| **Promise\<any>** | A `Promise` which can resolve to any value. |
