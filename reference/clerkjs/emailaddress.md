---
description: The EmailAddress object describes a User's email address.
---

# EmailAddress

## Overview

The `EmailAddress` object is a model around an email address. Email addresses are used to provide identification for users. 

Email addresses must be verified, so that we can make sure they can be assigned to their rightful owners. The `EmailAddress` object holds all necessary state around the verification process. 

The verification process always starts with the [EmailAddress.prepareVerification\(\)](emailaddress.md#prepareverification) method, which will send a one-time verification code via an email message. The second and final step involves an attempt to complete the verification by calling the [EmailAddress.attemptVerification\(\)](emailaddress.md#attemptverification-code) method, passing the one-time code as a parameter.

Finally, email addresses can be linked to other identifications.

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
        <p>A unique identifier for this email address.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>emailAddress</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The value of this email address, the actual email box address.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>verification</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="emailaddress.md#verificationresource"><em>VerificationResource</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object holding information on the verification of this email address.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>linkedTo</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="emailaddress.md#identificationlinkresource"><em>IndentificationLinkResource</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object containing information about any identifications that might
          be linked to this email address.</p>
      </td>
    </tr>
  </tbody>
</table>

## Methods

### attemptVerification\(code\)

`attemptVerification(code: string) => Promise<EmailAddressResource>`

Attempts to verify this email address, passing the one-time code that was sent as an email message. The code will be sent when calling the [EmailAddress.prepareVerification\(\)](emailaddress.md#prepareverification) method.

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
          sent via email.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_EmailAddressResource_](emailaddress.md)_&gt;_

This method returns a `Promise` which resolves with an `EmailAddress` object.
{% endtab %}
{% endtabs %}

### destroy\(\)

`destroy() => Promise<void>`

Delete this email address.

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

`prepareVerification() => Promise<EmailAddressResource>`

Kick off the verification process for this email address. An email message with a one-time code will be sent to the email address box.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_EmailAddressResource_](emailaddress.md)_&gt;_

This method returns a `Promise` which resolves with an `EmailAddress` object.
{% endtab %}
{% endtabs %}

### toString\(\)

`toString() => string | null`

Returns the value for this email address. Can also be accessed via the [EmailAddress.emailAddress](emailaddress.md#attributes) attribute.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_string \| null_

This method returns the email address attribute.
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
        <p><em>ClerkAPIError | null</em>
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

