---
description: The Clerk object is the core of the ClerkJS SDK.
---

# Clerk

## Overview

The `Clerk` object is a singleton which can act as the entry point for gaining access to other Clerk resources, like the  active [Client](client.md), [Session](session.md) and [User](user.md) objects. It also includes helper methods for mounting [Clerk Components](../../main-concepts/clerk-components.md) to your pages.

The `Clerk` object is always available via `window.Clerk`.

## Attributes

| Name        | Description                                                                                                                                                                                                                                                                                                                                                              |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **client**  | <p><em></em><a href="client.md"><em>Client</em></a><em></em></p><p>The <code>Client</code> object for the current window.</p>                                                                                                                                                                                                                                            |
| environment | Environment                                                                                                                                                                                                                                                                                                                                                              |
| **session** | <p><em></em><a href="session.md"><em>Session</em></a><em> | null | undefined</em></p><p>The currently active <code>Session</code>, which is guaranteed to be one of the sessions in <code>Client.sessions</code>. If there is no active session, this field will be <strong>null</strong>. If the session is loading, this field will be <strong>undefined</strong>.</p> |
| **user**    | <p><em></em><a href="user.md"><em>User</em></a><em> | null | undefined</em></p><p>A shortcut to <code>Session.user</code> which holds the currently active <code>User</code> object. If the session is <strong>null</strong> or <strong>undefined</strong>, the <code>user</code> field will match.</p>                                                                  |
| **version** | <p><em>string</em></p><p>The <a href="./">ClerkJS</a> SDK version.</p>                                                                                                                                                                                                                                                                                                   |

## Methods

### addListener(listener)

`addListener(listener: (resources: Resources) => void) => UnsubscribeCallback`

Registers a listener that triggers a callback whenever an important change in the `Client` object occurs. This method can be used to hook into different steps of the [sign in](../../main-concepts/sign-in-flow.md) and [sign up](../../main-concepts/sign-up-flow.md) flow and execute custom logic.&#x20;

Some things to note for specific changes in the `Client` object include:

* When there is an active session, `user === session.user`**.**
* When there is no active session, `user`** **and `session`** **will both be **null**.
* When a session is loading, `user`** **and `session`** **will be **undefined**.

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

### closeSignIn()

`closeSignIn() => void`

Closes the [\<SignIn/>](../../components/sign-in.md) modal.

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

Closes the [\<SignUp/>](../../components/sign-up.md) modal.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
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
| Name        | Description                                                                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **options** | <p><em></em><a href="clerk.md#componentoptions"><em>ComponentOptions</em></a><em></em></p><p>Configuration and options for initializing the <code>Clerk</code> object and <a href="../../main-concepts/clerk-components.md">Clerk Components</a>.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

### mountSignIn(node, nodeProps?)

`mountSignIn(node: HTMLDivElement, nodeProps?: SignInProps) => void`

Renders a [\<SignIn/>](../../components/sign-in.md) component inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
| Name           | Description                                                                                                                                                                                            |
| -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **node**       | <p><em>HTMLDivElement</em></p><p>An HTML <code>&#x3C;div/></code> element which will render the <a href="../../components/sign-in.md">&#x3C;SignIn/></a> component.</p>                                |
| **nodeProps?** | <p><em></em><a href="clerk.md#signinprops"><em>SignInProps</em></a><em></em></p><p>Additional props that will be passed to the <a href="../../components/sign-in.md">&#x3C;SignIn/></a> component.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### mountSignUp(node, nodeProps?)

`mountSignUp(node: HTMLDivElement, nodeProps?: SignUpProps) => void`

Renders a [\<SignUp/>](../../components/sign-up.md) component inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
| Name           | Description                                                                                                                                                                                            |
| -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **node**       | <p><em>HTMLDivElement</em></p><p>An HTML <code>&#x3C;div/></code> element which will render the <a href="../../components/sign-up.md">&#x3C;SignUp/></a> component.</p>                                |
| **nodeProps?** | <p><em></em><a href="clerk.md#signupprops"><em>SignUpProps</em></a><em></em></p><p>Additional props that will be passed to the <a href="../../components/sign-up.md">&#x3C;SignUp/></a> component.</p> |
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
| Name           | Description                                                                                                                                                                                                                |
| -------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **node**       | <p><em>HTMLElement</em></p><p>The element that the UserProfile should be mounted in.</p>                                                                                                                                   |
| **nodeProps?** | <p><em></em><a href="clerk.md#userprofileprops"><em>UserProfileProps</em></a><em></em></p><p>Additional props that will be passed to the <a href="../../components/user-profile.md">&#x3C;UserProfile/></a> component.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### navigate(to)

`navigate(to: string) => Promise<unknown>`

Helper method which will use the custom push navigation function of your application to navigate to the provided URL or relative path. See the relevant [section on routing](../../main-concepts/routing.md) for more information on navigation.

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

Opens the [\<SignIn/>](../../components/sign-in.md)component as a modal.

{% tabs %}
{% tab title="Parameters" %}
| Name      | Description                                                                                                                                                                                                    |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **props** | <p><em></em><a href="clerk.md#signinprops"><em>SignInProps</em></a><em></em></p><p>Configuration properties that will be passed to the <a href="../../components/sign-in.md">&#x3C;SignIn/></a> component.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### openSignUp(props)

`openSignUp(props: SignUpProps) => void`

Opens the [\<SignUp/>](../../components/sign-up.md) component as a modal.

{% tabs %}
{% tab title="Parameters" %}
| Name      | Description                                                                                                                                                                                                    |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **props** | <p><em></em><a href="clerk.md#signupprops"><em>SignUpProps</em></a><em></em></p><p>Configuration properties that will be passed to the <a href="../../components/sign-up.md">&#x3C;SignUp/></a> component.</p> |
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
| **session**     | <p><em></em><a href="session.md"><em>SessionResource </em></a><em>| string | null</em></p><p>A <code>Session</code> object or <code>Session</code> ID string to be set as the current session, or <strong>null</strong> to simply remove the active session, without setting a new one.</p>           |
| **beforeEmit?** | <p><em>(session: </em><a href="session.md"><em>SessionResource</em></a><em> | null) => Promise&#x3C;any> | void</em></p><p>Callback that will trigger when the current session is set to <strong>undefined</strong>, before finally being set to the passed session. Usually used for navigation.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which doesn't resolve to any value. The `Promise` will resolve after the passed `session` is set.
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

### signOut(callback?)

`signOut(callback?: SignOutCallback) => Promise<void>`

Signs out the active user from all sessions in a [multi-session application](../../popular-guides/popular-guides-multi-session-applications.md), or simply the current session in a single-session context. The current client will be deleted.

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

Signs out the active user from the current session. For a [multi-session application](../../popular-guides/popular-guides-multi-session-applications.md) this means that the rest of the session for the current client will remain, but the user will not be authenticated.

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

Unmounts the [\<SignIn/>](../../components/sign-in.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
| Name     | Description                                                                                                                                          |
| -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| **node** | <p><em>HTMLDivElement</em></p><p>The element that the <a href="../../components/sign-in.md">&#x3C;SignIn/></a> component will be unmounted from.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### unmountSignUp(node)

`unmountSignUp(node: HTMLDivElement) => void`

Unmounts the [\<SignUp/>](../../components/sign-up.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
| Name     | Description                                                                                                                                          |
| -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| **node** | <p><em>HTMLDivElement</em></p><p>The element that the <a href="../../components/sign-up.md">&#x3C;SignUp/></a> component will be unmounted from.</p> |
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

Unmounts the [\<UserProfile/>](../../components/user-profile.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
| Name     | Description                                                                                                                                                    |
| -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **node** | <p><em>HTMLDivElement</em></p><p>The element that the <a href="../../components/user-profile.md">&#x3C;UserProfile/></a> component will be unmounted from.</p> |
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

## Interfaces

### ComponentOptions

| Property                  | Description                                                                                                                                                                                                                                                                                                                            |
| ------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **selectInitialSession?** | <p><em>(client: </em><a href="client.md"><em>ClientResource</em></a><em>) => </em><a href="session.md"><em>SessionResource</em></a><em> | undefined</em></p><p>This function can be used to set the initial session in <a href="../../popular-guides/popular-guides-multi-session-applications.md">multi-session applications</a>.</p> |
| **navigate?**             | <p><em>(to: string) => Promise&#x3C;unknown> | unknown</em></p><p>Provide an implementation for the <a href="clerk.md#navigate-to">Clerk.navigate</a> method.</p>                                                                                                                                                                      |

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

| Property                | Description                                                                                                                                                                                                                                        |
| ----------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **afterSignOutAll?**    | <p><em>string</em></p><p>Full URL or relative path to navigate after sign out is complete and there are no active sessions on this client.</p>                                                                                                     |
| **afterSignOutOne?**    | <p><em>string</em></p><p>Full URL or relative path to navigate after sign out is complete.</p>                                                                                                                                                     |
| **afterSwitchSession?** | <p><em>string</em></p><p>Full URL or relative path to navigate after a successful account change. This property is used only for <a href="../../popular-guides/popular-guides-multi-session-applications.md">multi-session applications</a>.</p>   |
| **signInURL?**          | <p><em>string</em></p><p>Full URL or relative path to navigate on the "Add another account" action. This property is used only for <a href="../../popular-guides/popular-guides-multi-session-applications.md">multi-session applications</a>.</p> |
| **userProfileURL?**     | <p><em>string</em></p><p>Full URL or relative path of the user account management interface.</p>                                                                                                                                                   |

### UserProfileProps

| Property            | Description                                                                                                                                                                                                                                                     |
| ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **routing?**        | <p><em>string</em></p><p>The routing strategy for your pages. Supported values are:</p><ul><li><strong>path</strong>: Path based routing.</li><li><strong>hash</strong>: Hash based routing.</li><li><strong>virtual</strong>: Virtual based routing.</li></ul> |
| **path?**           | <p><em>string</em></p><p>The root URL where the component is mounted on.</p>                                                                                                                                                                                    |
| **hideNavigation?** | <p><em>boolean</em></p><p>Hides the default navigation bar. Can be used when a custom navigation bar is built.</p>                                                                                                                                              |
| **only?**           | <p><em>string</em></p><p>Renders only a specific page of the UserProfile component. Supported values are:</p><ul><li><strong>account</strong>: User account page.</li><li><strong>security</strong>: User security page.</li></ul>                              |

### HandleOAuthCallbackParams

| Name                 | Description                                                                                                                                                                                        |
| -------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **afterSignInUrl?**  | <p><em>string</em></p><p>Full URL or path to navigate after successful sign in.</p>                                                                                                                |
| **afterSignUpUrl?**  | <p><em>string</em></p><p>Full URL or path to navigate after successful sign up.</p>                                                                                                                |
| **redirectUrl?**     | <p><em>string</em></p><p>Full URL or path to navigate after successful sign in  or sign up. The same as setting <code>afterSignInUrl</code> and <code>afterSignUpUrl</code> to the same value.</p> |
| **secondFactorUrl?** | <p><em>string</em></p><p>Full URL or path to navigate during sign in, if 2FA is enabled.</p>                                                                                                       |

## Types

### SignOutCallback

`() => void | Promise<any>`

| Value             | Description                                 |
| ----------------- | ------------------------------------------- |
| **() => void**    | A `Function` which doesn't return anyhing.  |
| **Promise\<any>** | A `Promise` which can resolve to any value. |
