---
description: The client-side user session.
---

# Session

## Overview

The `Session` object is an abstraction over an HTTP session. It models the period of information exchange between a user and the server.

The `Session` object includes methods for recording session activity and ending the session client-side. For security reasons, sessions can also expire server-side.&#x20;

As soon as a [User](user.md) signs in, we create a Session for the current [Client](client.md). Clients can have more than one sessions at any point in time, but only one of those sessions will be **active**.

In certain scenarios, a session might be replaced by another one. This is often the case with [mutli-session applications](../../popular-guides/popular-guides-multi-session-applications.md).

All sessions that are **expired**, **removed**, **replaced**, **ended** or **abandoned** are not considered valid.&#x20;

{% hint style="info" %}
For more details regarding the different session states, see our documentation on [Session management](broken-reference).
{% endhint %}

## Attributes

| Property           | Description                                                                                                                    |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| **id**             | <p><em>string</em></p><p>A unique identifier for the session.</p>                                                              |
| **user**           | <p><em></em><a href="user.md"><em>User</em></a><em></em></p><p>The user associated with the session.</p>                       |
| **publicUserData** | <p><em>PublicUserData</em></p><p>Public information about the user that this session belongs to.</p>                           |
| **status**         | <p><em></em><a href="session.md#sessionstatus"><em>SessionStatus</em></a><em></em></p><p>The current state of the session.</p> |
| **lastActiveAt**   | <p><em>Date</em></p><p>The time the session was last active on the <a href="client.md">Client</a>.</p>                         |
| **abandonAt**      | <p><em>Date</em></p><p>The time when the session was abandoned by the user.</p>                                                |
| **expireAt**       | <p><em>Date</em></p><p>The time the session expires and will seize to be active.</p>                                           |
| **updatedAt**      | <p><em>Date</em></p><p>The last time the session recorded activity of any kind.</p>                                            |

## Methods

### end()

`end() => Promise<SessionResource>`

Marks the session as ended. The session will no longer be active for this [Client](client.md) and its status will become **ended**.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise<_[_SessionResource_](session.md)_>_

This method returns a `Promise` which resolves with a `Session` object.
{% endtab %}
{% endtabs %}

### remove()

`remove() => Promise<SessionResource>`

Marks the session as removed. The session will no longer be active for this [Client](client.md) and its status will become **removed**.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise<_[_SessionResource_](session.md)_>_

This method returns a `Promise` which resolves with a `Session` object.
{% endtab %}
{% endtabs %}

### touch()

`touch() => Promise<SessionResource>`

Touches the session, signifying some kind of user activity. Use this method to record any updates to user activity.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise<_[_SessionResource_](session.md)_>_

This method returns a `Promise` which resolves with a `Session` object.
{% endtab %}
{% endtabs %}

### getToken(options?)

`getToken(options?: GetSessionTokenOptions) => Promise<string>`

Retrieves the user's session token for the given template or the default clerk token. This method uses a cache so a network request will only be made if the token in memory has expired. The TTL for clerk token is one minute.

{% tabs %}
{% tab title="Parameters" %}
|              |                                                                                                                                                                                            |
| ------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **options?** | <p><a href="session.md#getsessiontokenoptions">GetSessionTokenOptions</a><br>Optionally pass the JWT template name and the leeway for expiring the cache. Default leeway is 10 seconds</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise\<string>_\
Returns a `Promise` that resolves to a `string`. The string is the user's session token.
{% endtab %}
{% endtabs %}

## Interfaces

### PublicUserData

| Property            | Description                                                                                                                                                                                                                 |
| ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **firstName**       | <p><em>string | null</em></p><p>The user's first name. This attribute will only be populated if name is enabled in <a href="../../popular-guides/setup-your-application.md#personal-information">instance settings</a>.</p> |
| **lastName**        | <p><em>string | null</em></p><p>The user's last name. This attribute will only be populated if name is enabled in <a href="../../popular-guides/setup-your-application.md#personal-information">instance settings</a>.</p>  |
| **profileImageUrl** | <p><em>string</em></p><p>The URL of the user's profile image.</p>                                                                                                                                                           |
| **identifier**      | <p><em>string</em></p><p>The user's identifier (email address, phone number, username, etc) that was used for authentication when this session was created.</p>                                                             |

### GetSessionTokenOptions

| Property             | Description                                                                                                                                                                                                                                          |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **leewayInSeconds?** | <p><em>number</em></p><p>The number of seconds that we allow the token to be cached.</p>                                                                                                                                                             |
| **template?**        | <p><em>string</em></p><p>The name of the JWT template to generate a new token from.</p>                                                                                                                                                              |
| **throwOnError?**    | <p><em>boolean</em></p><p>Whether to throw an error or return an empty string, if an error occurs.</p>                                                                                                                                               |
| **skipCache?**       | <p><em>boolean</em><br>Whether to skip the cache lookup and force a call to the server instead, even within the TTL. Useful if the token claims are time-sensitive or depend on data that can be updated (e.g. user fields).<br>(default: false)</p> |

## Types

### SessionStatus

`abandoned | active | ended | expired | removed | replaced | revoked`

| Value         | Description                                                                                                |
| ------------- | ---------------------------------------------------------------------------------------------------------- |
| **abandoned** | The session was abandoned client-side.                                                                     |
| **active**    | The session is valid and all activity is allowed.                                                          |
| **ended**     | The user signed out of the session, but the `Session` remains in the [Client](client.md) object.           |
| **expired**   | The period of allowed activity for this session has passed.                                                |
| **removed**   | The user signed out of the session and the `Session` was removed from the [Client](client.md) object.      |
| **replaced**  | The session has been replaced by another one, but the `Session` remains in the [Client](client.md) object. |
| **revoked**   | The application ended the session and the `Session` was removed from the [Client](client.md) object.       |
