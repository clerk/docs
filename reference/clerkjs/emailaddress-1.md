---
description: The ExternalAccount object describes a User's external account.
---

# ExternalAccount

## Overview

The `ExternalAccount` object is a model around an identification obtained by an external provider (e.g. an OAuth provider such as Google).

External account must be verified, so that we can make sure they can be assigned to their rightful owners. The `ExternalAccount` object holds all necessary state around the verification process.&#x20;

## Attributes



| Name                 | Description                                                                                                                                                                                   |
| -------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **id**               | <p><em>string</em></p><p>A unique identifier for this external account.</p>                                                                                                                   |
| **identificationId** | <p><em>string</em></p><p>The identification with which this external account is associated.</p>                                                                                               |
| **provider**         | <p><em>string</em></p><p>The provider name e.g. <code>oauth_google</code></p>                                                                                                                 |
| **providerUserId**   | <p><em>string</em></p><p>The unique id of the user in the provider.</p>                                                                                                                       |
| **emailAddress**     | <p><em>string</em></p><p>The provided email address of the user.</p>                                                                                                                          |
| **approvedScopes**   | <p><em>string[]</em></p><p>The scopes that the user has granted access to.</p>                                                                                                                |
| **firstName**        | <p><em>string</em></p><p>The provided first name of the user.</p>                                                                                                                             |
| **lastName**         | <p><em>string</em></p><p>The provided last name of the user.</p>                                                                                                                              |
| **avatarUrl**        | <p><em>string</em></p><p>The provided avatar URL of the user. </p>                                                                                                                            |
| **username**         | <p><em>string | null</em></p><p>The provided username of the user.</p>                                                                                                                        |
| **publicMetadata**   | <p><em>object</em></p><p>Extra metadata that are unique for this provider.</p>                                                                                                                |
| **label**            | <p><em>string | null</em></p><p>A descriptive label to differentiate multiple external accounts of the same user for the same provider.</p>                                                   |
| **verification**     | <p><em></em><a href="emailaddress-1.md#verificationresource"><em>VerificationResource</em></a><em></em></p><p>An object holding information on the verification of this external account.</p> |

## Methods

### destroy()

`destroy() => Promise<void>`

Delete this external account.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise\<void>_

This method returns a `Promise` which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

## Interfaces

### VerificationResource

| Property                            | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ----------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **status**                          | <p><em>string | null</em></p><p>The verification status. Possible values are:</p><ul><li><strong>unverified</strong>: The verification process has not been completed.</li><li><strong>verified</strong>: The verification process has completed successfully</li><li><strong>failed</strong>: The verification process has been completed, but failed.</li><li><strong>expired</strong>: The verification is invalid because it wasn't completed in the allowed time.</li></ul> |
| **strategy**                        | <p><em>string | null</em></p><p>The verification strategy. For example <code>oauth_google</code></p>                                                                                                                                                                                                                                                                                                                                                                             |
| **attempts**                        | <p><em>number | null</em></p><p>The number of attempts to complete the verification so far. Usually, a verification allows for maximum 3 attempts to be completed.</p>                                                                                                                                                                                                                                                                                                           |
| **expireAt**                        | <p><em>Date | null</em></p><p>The timestamp when the verification will expire and cease to be valid.</p>                                                                                                                                                                                                                                                                                                                                                                         |
| **error**                           | <p><em>ClerkAPIError | null</em></p><p>Any error that occurred during the verification process from the <a href="../frontend-api-reference/">Clerk API</a>.</p>                                                                                                                                                                                                                                                                                                                  |
| **externalVerificationRedirectURL** | <p><em>URL | null</em></p><p>The URL that the user will be redirected to, after the verification is completed.</p>                                                                                                                                                                                                                                                                                                                                                               |
