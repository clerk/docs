---
description: The Web3Wallet object describes a User's web3 wallet public address.
---

# Web3Wallet

## Overview

The `Web3Wallet` object describes a Web3 wallet public address. The public address can be used as a proof of identification for users.&#x20;

Web3 public addresses must be verified so that we can make sure they can be assigned to their rightful owners. The verification is completed via Web3 wallet browser extensions such as [Metamask](https://metamask.io). The `Web3Wallet3` object holds all the necessary state around the verification process.&#x20;

The verification process always starts with the signUp.prepareWeb3WalletVerification() or signIn.prepareFirstFactor() method, which will send the public wallet to Clerk Frontend API and will receive a nonce that needs to be signed by the Web3 wallet browser extension.

The second and final step involves an attempt to complete the verification by calling signUp.attemptWeb3WalletVerification() or signIn.attemptFirstFactor() method, passing the generated signature as a parameter.

## Attributes

| Name             | Description                                                                                                                                                                                                    |
| ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **id**           | <p><em>string</em></p><p>A unique identifier for this web3 wallet.</p>                                                                                                                                         |
| **web3Wallet**   | <p><em>string</em></p><p>In <a href="https://docs.metamask.io/guide/common-terms.html#address-public-key">Ethereum</a>, the address is made up of <code>0x</code> + <code>40 hexadecimal characters</code></p> |
| **verification** | <p><em></em><a href="web3wallet.md#verificationresource"><em>VerificationResource</em></a><em></em></p><p>An object holding information on the verification of this phone number.</p>                          |

## Methods

### toString()

`toString() => string`

Returns the `web3Wallet` hexadecimal string.

## Interfaces

### VerificationResource

| Property     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| ------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **status**   | <p><em>string | null</em></p><p>The verification status. Possible values are:</p><ul><li><strong>unverified</strong>: The verification process has not been completed.</li><li><strong>verified</strong>: The verification process has completed successfully.</li><li><strong>failed</strong>: The verification process has been completed, but failed.</li><li><strong>expired</strong>: The verification is invalid because it wasn't completed in the allowed time.</li></ul> |
| **strategy** | <p><em>string | null</em></p><p>The verification strategy. Possible <code>strategy</code> values are:</p><ul><li><strong>web3_metamask_signature</strong>: Currently Clerk supports <a href="https://docs.metamask.io/guide/signing-data.html#signing-data-with-metamask">Metamask <code>personal_sign</code> method</a>. </li></ul>                                                                                                                                              |
| **nonce**    | <p><em>string | null</em></p><p>A unique nonce that will be signed in the browser.</p>                                                                                                                                                                                                                                                                                                                                                                                            |
| **attempts** | <p><em>number | null</em></p><p>The number of attempts to complete the verification so far. Usually, a verification allows for maximum 3 attempts to be completed.</p>                                                                                                                                                                                                                                                                                                            |
| **expireAt** | <p><em>Date | null</em></p><p>The timestamp when the verification will expire and cease to be valid.</p>                                                                                                                                                                                                                                                                                                                                                                          |
| **error**    | <p><em></em><a href="broken-reference"><em>ClerkAPIError</em></a> <em>| null</em></p><p>Any error that occurred during the verification process from the <a href="../frontend-api-reference/">Clerk API</a>.</p>                                                                                                                                                                                                                                                                  |
