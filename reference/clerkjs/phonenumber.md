---
description: The PhoneNumber object describes a User's phone number.
---

# PhoneNumber

## Overview

The `PhoneNumber` object describes a phone number. Phone numbers can be used as a proof of identification for users, or simply as a means of contacting users.

Phone numbers must be verified, so that we can make sure they can be assigned to their rightful owners. The `PhoneNumber` object holds all the necessary state around the verification process.

The verification process always starts with the [PhoneNumber.prepareVerification\(\)](phonenumber.md#prepareverification) method, which will send a one-time verification code via an SMS message. The second and final step involves an attempt to complete the verification by calling the [PhoneNumber.attemptVerification\(\)](phonenumber.md#attemptverification-code) method, passing the one-time code as a parameter.

Finally, phone numbers are used as part of [multi-factor authentication](../../popular-guides/multi-factor-authentication.md). Users receive an SMS message with a one-time code that they need to provide as an extra verification step.

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
        <p>A unique identifier for this phone number.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>phoneNumber</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The value of this phone number, in <a href="https://en.wikipedia.org/wiki/E.164">E.164</a> format.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>reservedForSecondFactor</b>
      </td>
      <td style="text-align:left">
        <p>boolean</p>
        <p>Value will be <code>true</code> if this phone number is reserved for <a href="../../popular-guides/multi-factor-authentication.md">multi-factor authentication</a> (2FA), <code>false</code> otherwise.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>verification</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="phonenumber.md#verificationresource"><em>VerificationResource</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object holding information on the verification of this phone number.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>linkedTo</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="phonenumber.md#identificationlinkresource"><em>IndentificationLinkResource</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object containing information about any other identification that might
          be linked to this phone number.</p>
      </td>
    </tr>
  </tbody>
</table>

## Methods

### attemptVerification\(code\)

`attemptVerification(code: string) => Promise<PhoneNumberResource>`

Attempts to verify this phone number, passing the one-time code that was sent as an SMS message. The code will be sent when calling the [PhoneNumber.prepareVerification\(\)](phonenumber.md#prepareverification) method.

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
      <td style="text-align:left"><b>code</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The one-time code to be checked. Code needs to match the value that was
          sent via SMS.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_PhoneNumberResource_](phonenumber.md)_&gt;_

This method returns a `Promise` which resolves with a `PhoneNumber` object.
{% endtab %}
{% endtabs %}

### destroy\(\)

`destroy() => Promise<void>`

Delete this phone number.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;void&gt;_

This method returns a `Promise` which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

### prepareVerification\(\)

`prepareVerification() => Promise<PhoneNumberResource>`

Kick off the verification process for this phone number. An SMS message with a one-time code will be sent to the  phone number value.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_PhoneNumberResource_](phonenumber.md)_&gt;_

This method returns a `Promise` which resolves with a `PhoneNumber` object.
{% endtab %}
{% endtabs %}

### setReservedForSecondFactor\(\)

`setReservedForSecondFactor() => Promise<PhoneNumberResource>`

Marks this phone number as reserved for [mutli-factor authentication](../../popular-guides/multi-factor-authentication.md) \(2FA\).

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_PhoneNumberResource_](phonenumber.md)_&gt;_

This method returns a `Promise` which resolves with a `PhoneNumber` object.
{% endtab %}
{% endtabs %}

### toString\(\)

`toString() => string | null`

Returns the phone number value in [E.164 format](https://en.wikipedia.org/wiki/E.164). The value is taken from the [PhoneNumber.phoneNumber](phonenumber.md#attributes) attribute.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_string \| null_

This method returns the string value of this phone number.
{% endtab %}
{% endtabs %}

## Interfaces

### IdentificationLinkResource

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
        <p>A unique identifier for the identification.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>type</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The identification type.</p>
      </td>
    </tr>
  </tbody>
</table>

### VerificationResource

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>status</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The verification status. Possible values are:</p>
        <ul>
          <li><b>unverified</b>: The verification process has not been completed.</li>
          <li><b>verified</b>: The verification process has completed successfully.</li>
          <li><b>transferable</b>: The verification can be transferred as part of an
            external account verification process.</li>
          <li><b>failed</b>: The verification process has been completed, but failed.</li>
          <li><b>expired</b>: The verification is invalid because it wasn&apos;t completed
            in the allowed time.</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>strategy</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The verification strategy. Possible <code>strategy</code> values are:</p>
        <ul>
          <li><b>email_code</b>: User will receive a one-time authentication code via
            email. At least one email address should be on file for the user. The <code>email_address_id</code> parameter
            can also be specified to select one of the user&apos;s known email addresses.</li>
          <li><b>phone_code</b>: User will receive a one-time authentication code in
            their phone, via SMS. At least one phone number should be on file for the
            user. The <code>phone_number_id</code> parameter can also be specified to
            select which of the user&apos;s known phone numbers the SMS will go to.</li>
          <li><b>password</b>: The user needs to provide their password in order to
            be authenticated.</li>
          <li><b>oauth_google</b>: The user will be authenticated with their Google
            account (Google SSO).</li>
          <li><b>oauth_facebook</b>: The user will be authenticated with their Facebook
            account (Facebook SSO).</li>
          <li><b>oauth_hubspot</b>: The user will be authenticated with their Hubspot
            account (Hubspot SSO).</li>
          <li><b>oauth_github</b>: The user will be authenticated with their Github
            account (Github SSO).</li>
          <li><b>oauth_tiktok</b>: The user will be authenticated with their TikTok
            account (TikTok SSO).</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>attempts</b>
      </td>
      <td style="text-align:left">
        <p><em>number | null</em>
        </p>
        <p>The number of attempts to complete the verification so far. Usually, a
          verification allows for maximum 3 attempts to be completed.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>expireAt</b>
      </td>
      <td style="text-align:left">
        <p><em>Date | null</em>
        </p>
        <p>The timestamp when the verification will expire and cease to be valid.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>error</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href><em>ClerkAPIError</em></a><em> | null</em>
        </p>
        <p>Any error that occurred during the verification process from the <a href="../frontend-api-reference/">Clerk API</a>.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>externalVerificationRedirectURL</b>
      </td>
      <td style="text-align:left">
        <p><em>URL | null</em>
        </p>
        <p>If this is a verification that is based on an external account (usually <b>oauth_*</b>),
          this is the URL that the user will be redirected to after the verification
          is completed.</p>
      </td>
    </tr>
  </tbody>
</table>



