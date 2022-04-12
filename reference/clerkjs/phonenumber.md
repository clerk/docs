---
description: The PhoneNumber object describes a User's phone number.
---

# PhoneNumber

## Overview

The `PhoneNumber` object describes a phone number. Phone numbers can be used as a proof of identification for users, or simply as a means of contacting users.

Phone numbers must be verified, so that we can make sure they can be assigned to their rightful owners. The `PhoneNumber` object holds all the necessary state around the verification process.

The verification process always starts with the [PhoneNumber.prepareVerification()](phonenumber.md#prepareverification) method, which will send a one-time verification code via an SMS message. The second and final step involves an attempt to complete the verification by calling the [PhoneNumber.attemptVerification()](phonenumber.md#attemptverification-code) method, passing the one-time code as a parameter.

Finally, phone numbers are used as part of [multi-factor authentication](broken-reference). Users receive an SMS message with a one-time code that they need to provide as an extra verification step.

## Attributes

| Name                        | Description                                                                                                                                                                                                                                                                              |
| --------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **id**                      | <p><em>string</em></p><p>A unique identifier for this phone number.</p>                                                                                                                                                                                                                  |
| **phoneNumber**             | <p><em>string</em></p><p>The value of this phone number, in <a href="https://en.wikipedia.org/wiki/E.164">E.164</a> format.</p>                                                                                                                                                          |
| **reservedForSecondFactor** | <p><em>boolean</em></p><p>Value will be <code>true</code> if this phone number is reserved for <a href="broken-reference">multi-factor authentication</a> (2FA), <code>false</code> otherwise.</p>                                                                                       |
| **defaultSecondFactor**     | <p><em>boolean</em></p><p>Value will be <code>true</code> if this phone number is the default second factor, <code>false</code> otherwise.</p><p>A user must have exactly one default second factor, if <a href="broken-reference">multi-factor authentication</a> (2FA) is enabled.</p> |
| **verification**            | <p><em></em><a href="phonenumber.md#verificationresource"><em>VerificationResource</em></a><em></em></p><p>An object holding information on the verification of this phone number.</p>                                                                                                   |
| **linkedTo**                | <p><em></em><a href="phonenumber.md#identificationlinkresource"><em>IndentificationLinkResource</em></a><em></em></p><p>An object containing information about any other identification that might be linked to this phone number.</p>                                                   |

## Methods

### prepareVerification()

`prepareVerification() => Promise<PhoneNumberResource>`

Kick off the verification process for this phone number. An SMS message with a one-time code will be sent to the  phone number value.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise<_[_PhoneNumberResource_](phonenumber.md)_>_

This method returns a `Promise` which resolves with a `PhoneNumber` object.
{% endtab %}
{% endtabs %}



### attemptVerification(code)

`attemptVerification(params: AttemptPhoneNumberVerificationParams) => Promise<PhoneNumberResource>`

Attempts to verify this phone number, passing the one-time code that was sent as an SMS message. The code will be sent when calling the [PhoneNumber.prepareVerification()](phonenumber.md#prepareverification) method.

{% tabs %}
{% tab title="Parameters" %}
| Name       | Description                                                                                                   |
| ---------- | ------------------------------------------------------------------------------------------------------------- |
| **params** | An object of type [AttemptPhoneNumberVerificationParams](phonenumber.md#attemptphonenumberverificationparams) |
{% endtab %}

{% tab title="Returns" %}
_Promise<_[_PhoneNumberResource_](phonenumber.md)_>_

This method returns a `Promise` which resolves with a `PhoneNumber` object.
{% endtab %}
{% endtabs %}

### destroy()

`destroy() => Promise<void>`

Delete this phone number.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

### setReservedForSecondFactor(reserved)

`setReservedForSecondFactor(reserved: boolean) => Promise<PhoneNumberResource>`

Marks this phone number as reserved for [mutli-factor authentication](broken-reference) (2FA) or not.

{% tabs %}
{% tab title="Parameters" %}
| Name         | Description                                                                                                                                                             |
| ------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **reserved** | <p><em>boolean</em></p><p>Pass <strong>true</strong> to mark this phone number as reserved for 2FA, or <strong>false</strong> to disable 2FA for this phone number.</p> |
{% endtab %}

{% tab title="Returns" %}
_Promise<_[_PhoneNumberResource_](phonenumber.md)_>_

This method returns a `Promise` which resolves with a `PhoneNumber` object.
{% endtab %}
{% endtabs %}

### makeDefaultSecondFactor()

`makeDefaultSecondFactor() => Promise<PhoneNumberResource>`

Marks this phone number as the default second factor for [mutli-factor authentication](broken-reference) (2FA). A user can have exactly one default second factor.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters
{% endtab %}

{% tab title="Returns" %}
_Promise<_[_PhoneNumberResource_](phonenumber.md)_>_

This method returns a `Promise` which resolves with a `PhoneNumber` object.
{% endtab %}
{% endtabs %}

### toString()

`toString() => string | null`

Returns the phone number value in [E.164 format](https://en.wikipedia.org/wiki/E.164). The value is taken from the [PhoneNumber.phoneNumber](phonenumber.md#attributes) attribute.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_string | null_

This method returns the string value of this phone number.
{% endtab %}
{% endtabs %}

## Interfaces

### AttemptPhoneNumberVerificationParams

| Property | Description                                                                                                                               |
| -------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| **code** | <p><em>string</em></p><p>The one-time code that was sent to the user's phone number when <code>prepareVerification</code> was called.</p> |

### IdentificationLinkResource

| Property | Description                                                              |
| -------- | ------------------------------------------------------------------------ |
| **id**   | <p><em>string</em></p><p>A unique identifier for the identification.</p> |
| **type** | <p><em>string</em></p><p>The identification type.</p>                    |

### VerificationResource

| Property                            | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ----------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **status**                          | <p><em>string | null</em></p><p>The verification status. Possible values are:</p><ul><li><strong>unverified</strong>: The verification process has not been completed.</li><li><strong>verified</strong>: The verification process has completed successfully.</li><li><strong>transferable</strong>: The verification can be transferred as part of an external account verification process.</li><li><strong>failed</strong>: The verification process has been completed, but failed.</li><li><strong>expired</strong>: The verification is invalid because it wasn't completed in the allowed time.</li></ul>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| **strategy**                        | <p><em>string | null</em></p><p>The verification strategy. Possible <code>strategy</code> values are:</p><ul><li><strong>email_code</strong>: User will receive a one-time authentication code via email. At least one email address should be on file for the user. The <code>email_address_id</code> parameter can also be specified to select one of the user's known email addresses.</li><li><strong>phone_code</strong>: User will receive a one-time authentication code in their phone, via SMS. At least one phone number should be on file for the user. The <code>phone_number_id</code> parameter can also be specified to select which of the user's known phone numbers the SMS will go to.</li><li><strong>password</strong>: The user needs to provide their password in order to be authenticated.</li><li><strong>oauth_google</strong>: The user will be authenticated with their Google account (Google SSO).</li><li><strong>oauth_facebook</strong>: The user will be authenticated with their Facebook account (Facebook SSO).</li><li><strong>oauth_hubspot</strong>: The user will be authenticated with their Hubspot account (Hubspot SSO).</li><li><strong>oauth_github</strong>: The user will be authenticated with their Github account (Github SSO).</li><li><strong>oauth_tiktok</strong>: The user will be authenticated with their TikTok account (TikTok SSO).</li></ul> |
| **attempts**                        | <p><em>number | null</em></p><p>The number of attempts to complete the verification so far. Usually, a verification allows for maximum 3 attempts to be completed.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| **expireAt**                        | <p><em>Date | null</em></p><p>The timestamp when the verification will expire and cease to be valid.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| **error**                           | <p><em></em><a href="broken-reference"><em>ClerkAPIError</em></a> <em>| null</em></p><p>Any error that occurred during the verification process from the <a href="../frontend-api-reference/">Clerk API</a>.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| **externalVerificationRedirectURL** | <p><em>URL | null</em></p><p>If this is a verification that is based on an external account (usually <strong>oauth_*</strong>), this is the URL that the user will be redirected to after the verification is completed.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |

