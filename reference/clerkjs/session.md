---
description: The client-side user session.
---

# Session

## Overview

The `Session` object is an abstraction over an HTTP session. It models the period of information exchange between a user and the server.

The `Session` object includes methods for recording session activity and ending the session client-side. For security reasons, sessions can also expire server-side. 

As soon as a [User](user.md) signs in, we create a Session for the current [Client](client.md). Clients can have more than one sessions at any point in time, but only one of those sessions will be **active**.

In certain scenarios, a session might be replaced by another one. This is often the case with [mutli-session applications](../../popular-guides/popular-guides-multi-session-applications.md).

All sessions that are **expired**, **removed**, **replaced**, **ended** or **abandoned** are not considered valid. 

{% hint style="info" %}
For more details regarding the different session states, see our documentation on [Session management](../../main-concepts/session-management.md).
{% endhint %}

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
      <td style="text-align:left"><b>id</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>A unique identifier for the session.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>user</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="user.md"><em>User</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>The user associated with the session.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>status</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="session.md#sessionstatus"><em>SessionStatus</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>The current state of the session.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>lastActiveAt</b>
      </td>
      <td style="text-align:left">
        <p><em>Date</em>
        </p>
        <p>The time the session was last active on the <a href="client.md">Client</a>.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>abandonAt</b>
      </td>
      <td style="text-align:left">
        <p><em>Date</em>
        </p>
        <p>The time when the session was abandoned by the user.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>expireAt</b>
      </td>
      <td style="text-align:left">
        <p><em>Date</em>
        </p>
        <p>The time the session expires and will seize to be active.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>updatedAt</b>
      </td>
      <td style="text-align:left">
        <p><em>Date</em>
        </p>
        <p>The last time the session recorded activity of any kind.</p>
      </td>
    </tr>
  </tbody>
</table>

## Methods

### end\(\)

`end() => Promise<SessionResource>`

Marks the session as ended. The session will no longer be active for this [Client](client.md) and its status will become **ended**.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SessionResource_](session.md)_&gt;_

This method returns a `Promise` which resolves with a `Session` object.
{% endtab %}
{% endtabs %}

### remove\(\)

`remove() => Promise<SessionResource>`

Marks the session as removed. The session will no longer be active for this [Client](client.md) and its status will become **removed**.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SessionResource_](session.md)_&gt;_

This method returns a `Promise` which resolves with a `Session` object.
{% endtab %}
{% endtabs %}

### touch\(\)

`touch() => Promise<SessionResource>`

Touches the session, signifying some kind of user activity. Use this method to record any updates to user activity.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SessionResource_](session.md)_&gt;_

This method returns a `Promise` which resolves with a `Session` object.
{% endtab %}
{% endtabs %}

## Types

### SessionStatus

`abandoned | active | ended | expired | removed | replaced | revoked`

| Value | Description |
| :--- | :--- |
| **abandoned** | The session was abandoned client-side. |
| **active** | The session is valid and all activity is allowed. |
| **ended** | The user signed out of the session, but the `Session` remains in the [Client](client.md) object. |
| **expired** | The period of allowed activity for this session has passed. |
| **removed** | The user signed out of the session and the `Session` was removed from the [Client](client.md) object. |
| **replaced** | The session has been replaced by another one, but the `Session` remains in the [Client](client.md) object. |
| **revoked** | The application ended the session and the `Session` was removed from the [Client](client.md) object. |

