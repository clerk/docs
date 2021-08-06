---
description: The Clerk object is the core of the ClerkJS SDK.
---

# Clerk

## Overview

The `Clerk` object is a singleton which can act as the entry point for gaining access to other Clerk resources, like the  active [Client](client.md), [Session](session.md) and [User](user.md) objects. It also includes helper methods for mounting [Clerk Components](../../main-concepts/clerk-components.md) to your pages.

The `Clerk` object is always available via `window.Clerk`.

## Attributes

<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>client</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="client.md"><em>Client</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>The <code>Client</code> object for the current window.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">environment</td>
      <td style="text-align:left">Environment</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>session</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="session.md"><em>Session</em></a><em> | null | undefined</em>
        </p>
        <p>The currently active <code>Session</code>, which is guaranteed to be one
          of the sessions in <code>Client.sessions</code>. If there is no active session,
          this field will be <b>null</b>. If the session is loading, this field will
          be <b>undefined</b>.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>user</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="user.md"><em>User</em></a><em> | null | undefined</em>
        </p>
        <p>A shortcut to <code>Session.user</code> which holds the currently active <code>User</code> object.
          If the session is <b>null</b> or <b>undefined</b>, the <code>user</code> field
          will match.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>version</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The <a href="./">ClerkJS</a> SDK version.</p>
      </td>
    </tr>
  </tbody>
</table>

## Methods

### addListener\(listener\)

`addListener(listener: (resources: Resources) => void) => UnsubscribeCallback`

Registers a listener that triggers a callback whenever an important change in the `Client` object occurs. This method can be used to hook into different steps of the [sign in](../../main-concepts/sign-in-flow.md) and [sign up](../../main-concepts/sign-up-flow.md) flow and execute custom logic. 

Some things to note for specific changes in the `Client` object include:

* When there is an active session, `user === session.user`**.**
* When there is no active session, `user` ****and `session` ****will both be **null**.
* When a session is loading, `user` ****and `session` ****will be **undefined**.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>listener</b>
      </td>
      <td style="text-align:left">
        <p><em>(resources: Resources) =&gt; void</em>
        </p>
        <p>A function to be called when the <code>Client</code> object changes.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_\(\) =&gt; void_

This method returns a function that can be used to clean up the registered listener. 
{% endtab %}
{% endtabs %}

### closeSignIn\(\)

`closeSignIn() => void`

Closes the [&lt;SignIn/&gt;](../../components/sign-in.md) modal.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### closeSignUp\(\)

`closeSignUp() => void`

Closes the [&lt;SignUp/&gt;](../../components/sign-up.md) modal.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### isReady\(\)

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

### load\(options?\)

`load(options?: ComponentOptions) => Promise<void>`

Initializes the `Clerk` object and loads all necessary environment configuration and Instance settings from the [Frontend API](../frontend-api-reference/).

It is absolutely necessary to call this method before using the `Clerk` object in your code. Refer to the [ClerkJS installation](installation.md#setting-up-clerkjs) guide for more details.

### mountSignIn\(node, nodeProps?\)

`mountSignIn(node: HTMLDivElement, nodeProps?: SignInProps) => void`

Renders a [&lt;SignIn/&gt;](../../components/sign-in.md) component inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>node</b>
      </td>
      <td style="text-align:left">
        <p><em>HTMLDivElement</em>
        </p>
        <p>An HTML <code>&lt;div/&gt;</code> element which will render the <a href="../../components/sign-in.md">&lt;SignIn/&gt;</a> component.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>nodeProps?</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="clerk.md#signinprops"><em>SignInProps</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>Additional props that will be passed to the <a href="../../components/sign-in.md">&lt;SignIn/&gt;</a> component.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### mountSignUp\(node, nodeProps?\)

`mountSignUp(node: HTMLDivElement, nodeProps?: SignUpProps) => void`

Renders a [&lt;SignUp/&gt;](../../components/sign-up.md) component inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>node</b>
      </td>
      <td style="text-align:left">
        <p><em>HTMLDivElement</em>
        </p>
        <p>An HTML <code>&lt;div/&gt;</code> element which will render the <a href="../../components/sign-up.md">&lt;SignUp/&gt;</a> component.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>nodeProps?</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="clerk.md#signupprops"><em>SignUpProps</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>Additional props that will be passed to the <a href="../../components/sign-up.md">&lt;SignUp/&gt;</a> component.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### mountUserButton\(node, nodeProps?\)

`mountUserButton(node: HTMLDivElement, nodeProps?: UserButtonProps) => void`

Renders a [&lt;UserButton/&gt;](../../components/user-button.md) component for the active user inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>node</b>
      </td>
      <td style="text-align:left">
        <p><em>HTMLDivElement</em>
        </p>
        <p>An HTML <code>&lt;div/&gt;</code> element which will render the <a href="../../components/user-button.md">&lt;UserButton/&gt;</a> component.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>options?</b>
      </td>
      <td style="text-align:left">
        <p><em>UserButtonProps</em>
        </p>
        <p>Additional props that will be passed to the <a href="../../components/user-button.md">&lt;UserButton/&gt;</a> component.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### mountUserProfile\(node, nodeProps?\)

`mountUserProfile(node: HtmlDivElement, nodeProps?: UserProfileProps) => void`

Renders a [&lt;UserProfile/&gt;](../../components/user-button.md) component for the active user inside the provided HTML element, allowing to pass any props that configure the component.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>node</b>
      </td>
      <td style="text-align:left">
        <p><em>HTMLElement</em>
        </p>
        <p>The element that the UserProfile should be mounted in.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>nodeProps?</b>
      </td>
      <td style="text-align:left">
        <p><em>UserProfileProps</em>
        </p>
        <p>Additional props that will be passed to the <a href="../../components/user-profile.md">&lt;UserProfile/&gt;</a> component.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### navigate\(to\)

`navigate(to: string) => Promise<unknown>`

Helper method which will use the custom push navigation function of your application to navigate to the provided URL or relative path. See the relevant [section on routing](../../main-concepts/routing.md) for more information on navigation.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>to</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or relative path to navigate to.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;unknown&gt;_

This method returns a `Promise` that can resolve with any value.
{% endtab %}
{% endtabs %}

### openSignIn\(props\)

`openSignIn(props: SignInProps) => void`

Opens the [&lt;SignIn/&gt;](../../components/sign-in.md)component as a modal.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>props</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="clerk.md#signinprops"><em>SignInProps</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>Configuration properties that will be passed to the <a href="../../components/sign-in.md">&lt;SignIn/&gt;</a> component.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### openSignUp\(props\)

`openSignUp(props: SignUpProps) => void`

Opens the [&lt;SignUp/&gt;](../../components/sign-up.md) component as a modal.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>props</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="clerk.md#signupprops"><em>SignUpProps</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>Configuration properties that will be passed to the <a href="../../components/sign-up.md">&lt;SignUp/&gt;</a> component.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### redirectToSignIn\(\)

`redirectToSignIn() => Promise<unknown>`

Redirects to the sign in URL, as configured in your Instance settings. This method uses the [Clerk.navigate](clerk.md#navigate-to) method under the hood.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
Promise&lt;unknown&gt;

This method returns a `Promise` which can resolve to any value.
{% endtab %}
{% endtabs %}

### redirectToSignUp\(\)

`redirectToSignUp() => Promise<unknown>`

Redirects to the sign up URL, as configured in your Instance settings. This method uses the [Clerk.navigate](clerk.md#navigate-to) method under the hood.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
Promise&lt;unknown&gt;

This method returns a `Promise` which can resolve to any value.
{% endtab %}
{% endtabs %}

### redirectToUserProfile\(\)

`redirectToUserProfile() => Promise<unknown>`

Redirects to the user profile management URL, as configured in your Instance settings. This method uses the [Clerk.navigate](clerk.md#navigate-to) method under the hood.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
Promise&lt;unknown&gt;

This method returns a `Promise` which can resolve to any value.
{% endtab %}
{% endtabs %}

### setSession\(session, beforeEmit?\)

`setSession(session: SessionResource | string | null, beforeEmit?: (session: SessionResource | null) => any) => Promise<void>`

Set the current session on this client to the provided session. The provided session can be either a complete [Session](session.md) object or simply its unique identifier.

Passing null as the session will result in the current session to be removed from the client.

If an active session already exists, it will be replaced with the new one. The change happens in three steps:

1. The current `Session` object is set to **undefined**, which causes the control components to stop rendering their children as though Clerk is still initializing.
2. The **beforeEmit** callback is executed. If a `Promise` is returned, Clerk waits for the `Promise` to resolve.
3. The current `Session` is set to the passed **session**. This causes the control components to render their children again.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>session</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="session.md"><em>SessionResource </em></a><em>| string | null</em>
        </p>
        <p>A <code>Session</code> object or <code>Session</code> ID string to be set
          as the current session, or <b>null</b> to simply remove the active session,
          without setting a new one.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>beforeEmit?</b>
      </td>
      <td style="text-align:left">
        <p><em>(session: </em><a href="session.md"><em>SessionResource</em></a><em> | null) =&gt; Promise&lt;any&gt; | void</em>
        </p>
        <p>Callback that will trigger when the current session is set to <b>undefined</b>,
          before finally being set to the passed session. Usually used for navigation.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;void&gt;_

This method returns a `Promise` which doesn't resolve to any value. The `Promise` will resolve after the passed `session` is set.
{% endtab %}
{% endtabs %}

### signOut\(callback?\)

`signOut(callback?: SignOutCallback) => Promise<void>`

Signs out the active user from all sessions in a [multi-session application](../../popular-guides/popular-guides-multi-session-applications.md), or simply the current session in a single-session context. The current client will be deleted.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>callback?</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="clerk.md#signoutcallback"><em>SignOutCallback</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>A callback function that will be called after successful sign out.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;void&gt;_

This method returns a `Promise` which does not resolve to any value.
{% endtab %}
{% endtabs %}

### signOutOne\(callback?\)

`signOutOne(callback?: SignOutCallback) => Promise<void>`

Signs out the active user from the current session. For a [multi-session application](../../popular-guides/popular-guides-multi-session-applications.md) this means that the rest of the session for the current client will remain, but the user will not be authenticated.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>callback?</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="clerk.md#signoutcallback"><em>SignOutCallback</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>A callback function that will be called after successful sign out.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;void&gt;_

This method returns a `Promise` which does not resolve to any value.
{% endtab %}
{% endtabs %}

### unmountSignIn\(node\)

`unmountSignIn(node: HTMLDivElement) => void`

Unmounts the [&lt;SignIn/&gt;](../../components/sign-in.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>node</b>
      </td>
      <td style="text-align:left">
        <p><em>HTMLDivElement</em>
        </p>
        <p>The element that the <a href="../../components/sign-in.md">&lt;SignIn/&gt;</a> component
          will be unmounted from.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### unmountSignUp\(node\)

`unmountSignUp(node: HTMLDivElement) => void`

Unmounts the [&lt;SignUp/&gt;](../../components/sign-up.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>node</b>
      </td>
      <td style="text-align:left">
        <p><em>HTMLDivElement</em>
        </p>
        <p>The element that the <a href="../../components/sign-up.md">&lt;SignUp/&gt;</a> component
          will be unmounted from.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### unmountUserButton\(node\)

`unmountUserButton(node: HTMLDivElement) => void`

Unmounts the [&lt;UserButton/&gt;](../../components/user-button.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>node</b>
      </td>
      <td style="text-align:left">
        <p><em>HTMLDivElement</em>
        </p>
        <p>The element that the <a href="../../components/user-button.md">&lt;UserButton/&gt;</a> component
          will be unmounted from.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

### unmountUserProfile\(node\)

`unmountUserProfile(node: HTMLDivElement) => void`

Unmounts the [&lt;UserProfile/&gt;](../../components/user-profile.md) component from the specified HTML element.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>node</b>
      </td>
      <td style="text-align:left">
        <p><em>HTMLDivElement</em>
        </p>
        <p>The element that the <a href="../../components/user-profile.md">&lt;UserProfile/&gt;</a> component
          will be unmounted from.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
This method has no return value.
{% endtab %}
{% endtabs %}

## Interfaces

### SignInProps

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>routing?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The routing strategy for your pages. Supported values are:</p>
        <ul>
          <li><b>path</b>: Path based routing.</li>
          <li><b>hash</b>: Hash based routing.</li>
          <li><b>virtual</b>: Virtual based routing.</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>path?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The root URL where the component is mounted on.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSignIn?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The full URL or path to navigate after a successful sign in.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>signUpURL?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or path to the sign up page. Use this property to provide the
          target of the &quot;Sign Up&quot; link that&apos;s rendered.</p>
      </td>
    </tr>
  </tbody>
</table>

### SignUpProps

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>routing?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The routing strategy for your pages. Supported values are:</p>
        <ul>
          <li><b>path</b>: Path based routing.</li>
          <li><b>hash</b>: Hash based routing.</li>
          <li><b>virtual</b>: Virtual based routing.</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>path?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The root URL where the component is mounted on.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSignUp?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The full URL or path to navigate after a successful sign up.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>signInURL?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or path to the sign in page. Use this property to provide the
          target of the &quot;Sign In&quot; link that&apos;s rendered.</p>
      </td>
    </tr>
  </tbody>
</table>

### UserButtonProps

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>afterSignOutAll?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or relative path to navigate after sign out is complete and there
          are no active sessions on this client.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSignOutOne?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or relative path to navigate after sign out is complete.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSwitchSession?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or relative path to navigate after a successful account change.
          This property is used only for <a href="../../popular-guides/popular-guides-multi-session-applications.md">multi-session applications</a>.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>signInURL?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or relative path to navigate on the &quot;Add another account&quot;
          action. This property is used only for <a href="../../popular-guides/popular-guides-multi-session-applications.md">multi-session applications</a>.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>userProfileURL?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or relative path of the user account management interface.</p>
      </td>
    </tr>
  </tbody>
</table>

### UserProfileProps

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>routing?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The routing strategy for your pages. Supported values are:</p>
        <ul>
          <li><b>path</b>: Path based routing.</li>
          <li><b>hash</b>: Hash based routing.</li>
          <li><b>virtual</b>: Virtual based routing.</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>path?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The root URL where the component is mounted on.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSignOutAll?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or relative path to navigate after successful sign out, with
          no other active accounts on the client.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>afterSignOutOne?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Full URL or relative path to navigate after successful sign out.</p>
      </td>
    </tr>
  </tbody>
</table>

## Types

### SignOutCallback

`() => void | Promise<any>`

| Value | Description |
| :--- | :--- |
| **\(\) =&gt; void** | A `Function` which doesn't return anyhing. |
| **Promise&lt;any&gt;** | A `Promise` which can resolve to any value. |

