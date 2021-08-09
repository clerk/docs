---
description: >-
  A wrapper over the client-side user session, with additional user activity
  information.
---

# SessionWithActivities

## Overview

The `SessionWithActivities` object is a modified [Session](session.md) object. It contains most of the information that the `Session` object stores, adding extra information about the current session's latest activity.

The additional data included in the latest activity are useful for analytics purposes. A [SessionActivity](sessionwithactivities.md#sessionactivity) object will provide information about the user's location, device and browser.

While the `SessionWithActivities` object wraps the most important information around a `Session` object, the two objects have entirely different methods.

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
      <td style="text-align:left"><b>status</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="sessionwithactivities.md#sessionstatus"><em>SessionStatus</em></a>&lt;em&gt;&lt;/em&gt;</p>
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
      <td style="text-align:left"><b>latestActivity</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="sessionwithactivities.md#sessionactivity"><em>SessionActivity</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object that provides additional information about this session, focused
          around user activity data.</p>
      </td>
    </tr>
  </tbody>
</table>

## Methods

### revoke\(\)

`revoke() => Promise<SessionWithActivitiesResource>`

Marks this session as revoked. If this is the active session, the attempt to revoke it will fail. 

Users can revoke only their own sessions.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SessionWithActivitiesResource_](sessionwithactivities.md)_&gt;_

This method returns a `Promise` which resolves with a `SessionWithActivities` object.
{% endtab %}
{% endtabs %}

## Interfaces

### SessionActivity

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
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
        <p>A unique identifier for the session activity record.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>browserName?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The name of the browser from which this session activity occurred.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>browserVersion?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The version of the browser from which this session activity occurred.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>deviceType?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The type of the device which was used in this session activity.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>ipAddress?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The IP address from which this session activity originated.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>city?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The city from which this session activity occurred. Resolved by IP address
          geo-location.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>country?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The country from which this session activity occurred. Resolved by IP
          address geo-location.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>isMobile?</b>
      </td>
      <td style="text-align:left">
        <p><em>boolean</em>
        </p>
        <p>Will be set to <b>true</b> if the session activity came from a mobile device, <b>false</b> otherwise.</p>
      </td>
    </tr>
  </tbody>
</table>

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



