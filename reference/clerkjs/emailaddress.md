---
description: The EmailAddress object describes a User's email address.
---

# EmailAddress

## Overview

The `EmailAddress` object is a model around an email address. Email addresses are used to provide identification for users.&#x20;

Email addresses must be verified, so that we can make sure they can be assigned to their rightful owners. The `EmailAddress` object holds all necessary state around the verification process.&#x20;

The verification process always starts with the [EmailAddress.prepareVerification()](emailaddress.md#prepareverification) method, which will send a one-time verification code via an email message. The second and final step involves an attempt to complete the verification by calling the [EmailAddress.attemptVerification()](emailaddress.md#attemptverification-code) method, passing the one-time code as a parameter.

Finally, email addresses can be linked to other identifications.

## Attributes

| Name             | Description                                                                                                                                                                                                                         |
| ---------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **id**           | <p><em>string</em></p><p>A unique identifier for this email address.</p>                                                                                                                                                            |
| **emailAddress** | <p><em>string</em></p><p>The value of this email address, the actual email box address.</p>                                                                                                                                         |
| **verification** | <p><em></em><a href="emailaddress.md#verificationresource"><em>VerificationResource</em></a><em></em></p><p>An object holding information on the verification of this email address.</p>                                            |
| **linkedTo**     | <p><em></em><a href="emailaddress.md#identificationlinkresource"><em>IndentificationLinkResource</em></a><em></em></p><p>An object containing information about any identifications that might be linked to this email address.</p> |

## Methods

### attemptVerification(code)

`attemptVerification(code: string) => Promise<EmailAddressResource>`

Attempts to verify this email address, passing the one-time code that was sent as an email message. The code will be sent when calling the [EmailAddress.prepareVerification()](emailaddress.md#prepareverification) method.

{% tabs %}
{% tab title="Parameters" %}
| Name     | Description                                                                                                          |
| -------- | -------------------------------------------------------------------------------------------------------------------- |
| **code** | <p><em>string</em></p><p>The one-time code to be checked. Code needs to match the value that was sent via email.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise<_[_EmailAddressResource_](emailaddress.md)_>_

This method returns a `Promise` which resolves with an `EmailAddress` object.
{% endtab %}
{% endtabs %}

### createMagicLinkFlow()

`createMagicLinkFlow() => CreateMagicLinkFlowParams<StartMagicLinkFlowParams, EmailAddressResource>`

Sets up an email verification with magic link flow. Calling `createMagicLinkFlow()` will return two functions.&#x20;

The first function is async and starts the magic link flow, preparing a magic link verification. It sends the magic link email and starts polling for verification results. The signature is `startMagicLinkFlow({ callbackUrl: string }) => Promise<EmailAddressResource>`.

The second function can be used to stop polling at any time, allowing for full control of the flow and cleanup. The signature is `cancelMagicLinkFlow() => void`.

{% tabs %}
{% tab title="Returns" %}
__[_CreateMagicLinkFlowReturn_](emailaddress.md#createmagiclinkflowreturn-less-than-startmagiclinkparams-emailaddressresource-greater-than)_<_[_StartMagicLinkFlowParams_](emailaddress.md#startmagiclinkflowparams)_, _[_EmailAddressResource_](emailaddress.md#attributes)_>_

This method returns two functions. One to start the magic link flow and the other to cancel waiting for the results.
{% endtab %}
{% endtabs %}

### destroy()

`destroy() => Promise<void>`

Delete this email address.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

### prepareVerification()

`prepareVerification() => Promise<EmailAddressResource>`

Kick off the verification process for this email address. An email message with a one-time code will be sent to the email address box.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise<_[_EmailAddressResource_](emailaddress.md)_>_

This method returns a `Promise` which resolves with an `EmailAddress` object.
{% endtab %}
{% endtabs %}

### toString()

`toString() => string | null`

Returns the value for this email address. Can also be accessed via the [EmailAddress.emailAddress](emailaddress.md#attributes) attribute.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_string | null_

This method returns the email address attribute.
{% endtab %}
{% endtabs %}

## Interfaces

### IdentificationLinkResource

| Property | Description                                                              |
| -------- | ------------------------------------------------------------------------ |
| **id**   | <p><em>string</em></p><p>A unique identifier for the identification.</p> |
| **type** | <p><em>string</em></p><p>The identification type.</p>                    |

### StartMagicLinkFlowParams

| Name            | Description                                                                                                                            |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| **callbackUrl** | <p><em>string</em></p><p>The magic link target URL. Users will be redirected here once they click the magic link from their email.</p> |

### VerificationResource

| Property                            | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ----------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **status**                          | <p><em>string | null</em></p><p>The verification status. Possible values are:</p><ul><li><strong>unverified</strong>: The verification process has not been completed.</li><li><strong>verified</strong>: The verification process has completed successfully.</li><li><strong>transferable</strong>: The verification can be transferred as part of an external account verification process.</li><li><strong>failed</strong>: The verification process has been completed, but failed.</li><li><strong>expired</strong>: The verification is invalid because it wasn't completed in the allowed time.</li></ul>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| **strategy**                        | <p><em>string | null</em></p><p>The verification strategy. Possible <code>strategy</code> values are:</p><ul><li><strong>email_code</strong>: User will receive a one-time authentication code via email. At least one email address should be on file for the user. The <code>email_address_id</code> parameter can also be specified to select one of the user's known email addresses.</li><li><strong>phone_code</strong>: User will receive a one-time authentication code in their phone, via SMS. At least one phone number should be on file for the user. The <code>phone_number_id</code> parameter can also be specified to select which of the user's known phone numbers the SMS will go to.</li><li><strong>password</strong>: The user needs to provide their password in order to be authenticated.</li><li><strong>oauth_google</strong>: The user will be authenticated with their Google account (Google SSO).</li><li><strong>oauth_facebook</strong>: The user will be authenticated with their Facebook account (Facebook SSO).</li><li><strong>oauth_hubspot</strong>: The user will be authenticated with their Hubspot account (Hubspot SSO).</li><li><strong>oauth_github</strong>: The user will be authenticated with their Github account (Github SSO).</li><li><strong>oauth_tiktok</strong>: The user will be authenticated with their TikTok account (TikTok SSO).</li></ul> |
| **attempts**                        | <p><em>number | null</em></p><p>The number of attempts to complete the verification so far. Usually, a verification allows for maximum 3 attempts to be completed.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| **expireAt**                        | <p><em>Date | null</em></p><p>The timestamp when the verification will expire and cease to be valid.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| **error**                           | <p><em>ClerkAPIError | null</em></p><p>Any error that occurred during the verification process from the <a href="../frontend-api-reference/">Clerk API</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| **externalVerificationRedirectURL** | <p><em>URL | null</em></p><p>If this is a verification that is based on an external account (usually <strong>oauth_*</strong>), this is the URL that the user will be redirected to after the verification is completed.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |

## Types

### CreateMagicLinkFlowReturn\<StartMagicLinkParams, EmailAddressResource>

`{ startMagicLinkFlow: (params: StartMagicLinkFlowParams) => Promise<EmailAddressResource>, cancelMagicLinkFlow: () => void }`

| **startMagicLinkFlow**  | Function that starts the magic link flow. It prepares a magic link verification and polls for the verification result. Accepts [StartMagicLinkFlowParams](emailaddress.md#undefined). Returns a `Promise` which resolves to an [EmailAddressResource](emailaddress.md#attributes). |
| ----------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **cancelMagicLinkFlow** | Function to cleanup the magic link flow. Stops waiting for verification results.                                                                                                                                                                                                   |
