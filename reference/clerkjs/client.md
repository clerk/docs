---
description: >-
  The Client object tracks sessions, as well as the state of any sign in and
  sign up attempts.
---

# Client

## Overview

The `Client` object keeps track of the authenticated sessions in the current device. The device can be a browser, a native application or any other medium that is usually the requesting part in a request/response architecture.

The `Client` object also holds information about any sign in or sign up attempts that might be in progress, tracking the sign in or sign up progress.

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
      <td style="text-align:left"><b>signIn</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md"><em>SignInResource</em></a><em> | null</em>
        </p>
        <p>The current sign in attempt, or null if there is none.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>signUp</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signup.md"><em>SignUpResource</em></a><em> | null</em>
        </p>
        <p>The current sign up attempt, or null if there is none.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>sessions</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="session.md"><em>Session</em></a><em>[]</em>
        </p>
        <p>A list of sessions that have been created on this client.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>activeSessions</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="session.md"><em>Session</em></a><em>[]</em>
        </p>
        <p>A list of active sessions on this client.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>lastActiveSessionId</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Unique identifier of the last active <a href="session.md">Session</a> on
          this client.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>updatedAt</b>
      </td>
      <td style="text-align:left">
        <p><em>Date</em>
        </p>
        <p>Timestamp of last update for the client.</p>
      </td>
    </tr>
  </tbody>
</table>

## Methods

### create\(\)

`create() => Promise<ClientResource>`

Creates a new client for the current instance along with its cookie.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_ClientResource_](client.md)_&gt;_

This method returns a `Promise` which resolves with a `Client` object.
{% endtab %}
{% endtabs %}

### destroy\(\)

`destroy() => Promise<void>`

Deletes the client. All sessions will be reset.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;void&gt;_

This method returns a Promise which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

### isNew\(\) 

`isNew() => boolean`

Returns **true** if this client hasn't been saved \(created\) yet in the [Frontend API](../frontend-api-reference/), **false** otherwise.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
This method returns **true** for new clients, **false** otherwise.
{% endtab %}
{% endtabs %}

