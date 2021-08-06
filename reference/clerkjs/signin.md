---
description: >-
  The SignIn object contains the state of the current sign-in that the user has
  initiated.
---

# SignIn

## Overview

The `SignIn` object holds all the state of the current sign in and provides helper methods to navigate and complete the sign in process.

There are two important steps in the sign in flow. 

1. Users _must_ complete a first factor verification. This can be something like providing a password, or a one-time code \(OTP\), or providing proof of their identity through an external social account \(SSO/OAuth\). 
2. After that, users _might need to_ go through a second verification process. This is the second factor \(2FA\).

The `SignIn` object's properties can be split into logical groups, with each group providing information on different aspects of the sign in flow. These groups can be:

* Information about the current sign in status in general and which authentication identifiers, authentication methods and verifications are supported.
* Information about the user and the provided authentication identifier value \(email address, phone number or username\).
* Information about each verification, either the first factor \(logging in\) or the second factor \(2FA\).

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
      <td style="text-align:left"><b>status</b>
      </td>
      <td style="text-align:left">
        <p><em>string </em>
        </p>
        <p>The current status of the sign-in. It can take the following values:</p>
        <ul>
          <li><b>needs_identifier</b>: The authentication identifier hasn&apos;t been
            provided.</li>
          <li><b>needs_first_factor</b>: First factor verification for the provided
            identifier needs to be prepared and verified.</li>
          <li><b>needs_second_factor</b>: Second factor verification (2FA) for the provided
            identifier needs to be prepared and verified.</li>
          <li><b>complete</b>: The sign-in is complete and the user is authenticated.</li>
          <li><b>abandoned</b>: The sign-in has been inactive for a long period of time,
            thus it&apos;s considered as abandoned and need to start over.</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>supportedIdentifiers</b>
      </td>
      <td style="text-align:left">
        <p><em>string[]</em>
        </p>
        <p>Array of all the authentication identifiers that are supported for this
          sign in.</p>
        <p>Examples of this could be <b>email_address</b>, <b>phone_number</b> or
          <br
          /><b>username</b>.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>identifier</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The authentication identifier value for the current sign-in.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>supportedExternalAccounts</b>
      </td>
      <td style="text-align:left">
        <p><em>string[]</em>
        </p>
        <p>Array of all the external accounts that can be used in this sign in, e.g. <b>oauth_google</b>, <b>oauth_facebook</b>,
          etc.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>supportedFirstFactors</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#signinfactor"><em>SignInFactor</em></a><em>[]</em>
        </p>
        <p>Array of the first factors that are supported in the current sign-in.</p>
        <p>Each factor contains information about the verification strategy that
          can</p>
        <p>be used, e.g. <b>email_code</b> for email addresses, <b>phone_code</b> for
          phone numbers, etc., as well as the identifier that the factor refers to.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>supportedSecondFactors</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#signinfactor"><em>SignInFactor</em></a><em>[] | null</em>
        </p>
        <p>Array of the second factors that are supported in the current sign-in.
          Each factor contains information about the verification strategy that can</p>
        <p>be used, e.g. <b>email_code</b> for email addresses, <b>phone_code</b> for
          phone numbers, etc., as well as the identifier that the factor refers to.</p>
        <p>Please note that this property is populated only when the first factor
          is verified.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>firstFactorVerification</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#verificationresource"><em>VerificationResource</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>The state of the verification process for the selected first factor.</p>
        <p>Please note that this property contains an empty verification object initially,
          since there is no first factor selected.
          <br />You need to call the <a href="signin.md#preparefirstfactor">SignIn.prepareFirstFactor</a> method
          in order to start the verification process.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>secondFactorVerification</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#verificationresource"><em>VerificationResource</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>The state of the verification process for the selected second factor.</p>
        <p>Similar to <b>firstFactorVerification</b>, this property contains an empty
          verification object initially, since there is no second factor selected.
          <br
          />You need to call the <a href="signin.md#preparesecondfactor">SignIn.prepareSecondFactor</a> method
          in order to start the verification process.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>userData</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#userdata"><em>UserData</em></a><em> | null</em>
        </p>
        <p>An object containing information about the user of the current sign in.</p>
        <p>This property is populated only once an identifier is given to the <code>SignIn</code> object.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>createdSessionId</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The identifier of the session that was created upon completion of the
          current sign-in.</p>
        <p>The value of this property is <b>null</b> if the sign-in status is not <b>complete</b>.</p>
      </td>
    </tr>
  </tbody>
</table>

## Methods

### create\(params\) <a id="create"></a>

`create(params: SignInParams) => Promise<SignInResource>`

Use this method to kick-off the sign in flow. It creates a **SignIn** object and stores the sign in lifecycle state. 

Depending on the use-case and the `params` you pass to the `create` method, it can either complete the sign in process in one go, or simply collect part of the necessary data for completing authentication at a later stage.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Name</b>
      </th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>params</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#signinparams"><em>SignInParams</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>Depending on the keys and values passed here, the <code>create</code> method&apos;s
          behavior changes.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SignInResource_](signin.md#attributes)_&gt;_

This method returns a `Promise` which resolves with a `SignIn` object. Check the `status` attribute to see if the `SignIn` has been completed or a hint on what needs to happen next.
{% endtab %}
{% endtabs %}

### prepareFirstFactor\(params\) <a id="preparefirstfactor"></a>

`prepareFirstFactor(params: PrepareFirstFactorParams) => Promise<SignInResource>`

Begins the first factor verification process. This is a required step in order to complete a sign in, as users should be verified at least by one factor of authentication. 

Common scenarios are one-time code \(OTP\) or social account \(SSO\) verification. This is determined by the accepted `strategy` parameter values. Each authentication identifier supports different strategies.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Name</b>
      </th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>params</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#preparefirstfactorparams"><em>PrepareFirstFactorParams</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object that allows you to specify a verification strategy and any strategy-specific
          additional data.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SignInResource_](signin.md#attributes)_&gt;_

This method returns a `Promise` which resolves with a `SignIn` object. Check the `firstFactorVerification` attribute for the status of the first factor verification process.
{% endtab %}
{% endtabs %}

### attemptFirstFactor\(params\) <a id="attemptfirstfactor"></a>

`attemptFirstFactor(params: AttemptFactorParams) => Promise<SignInResource>`

Attempts to complete the first factor verification process. This is a required step in order to complete a sign in, as users should be verified at least by one factor of authentication.

Make sure that a `SignIn` object already exists before you call this method, either by first calling [SignIn.create](signin.md#create) or [SignIn.prepareFirstFactor](signin.md#preparefirstfactor). 

The only strategy that does not require a verification to have already been prepared before attempting to complete it, is the **password** strategy.

Depending on the strategy that was selected when the verification was prepared, the method parameters should be different.

{% tabs %}
{% tab title="Parameters" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Name</b>
      </th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>params</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#attemptfactorparams"><em>AttemptFactorParams</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object that includes the verification strategy and the evidence that&apos;s
          needed to identify the user and complete the verification process.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SignInResource_](signin.md#attributes)_&gt;_

This method returns a `Promise` which resolves with a `SignIn` object. Check the `firstFactorVerification` attribute for the status of the first factor verification process.
{% endtab %}
{% endtabs %}

### prepareSecondFactor\(params\) <a id="preparesecondfactor"></a>

`prepareSecondFactor(params: PrepareSecondFactorParams) => Promise<SignInResource>`

Begins the second factor verification process. This step is optional in order to complete a sign in. 

A common scenario for the second step verification \(2FA\) is to require a one-time code \(OTP\) as proof of identity. This is determined by the accepted `strategy` parameter values. Each authentication identifier supports different strategies.

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
      <td style="text-align:left"><b>params</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#preparesecondfactorparams"><em>PrepareSecondFactorParams</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object that specifies the verification strategy.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SignInResource_](signin.md#attributes)_&gt;_

This method returns a `Promise` which resolves with a `SignIn` object. Check the `secondFactorVerification` attribute for the status of the second factor verification process.
{% endtab %}
{% endtabs %}

### attemptSecondFactor\(params\) <a id="attemptsecondfactor"></a>

`attemptSecondFactor(params: AttemptFactorParams) => Promise<SignInResource>`

Attempts to complete the second factor verification process \(2FA\). This step is optional in order to complete a sign in.

Make sure that a verification has already been prepared before you call this method, by first calling [SignIn.prepareSecondFactor](signin.md#preparesecondfactor). 

Depending on the strategy that was selected when the verification was prepared, the method parameters should be different.

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
      <td style="text-align:left"><b>params</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#attemptfactorparams"><em>AttemptFactorParams</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object that specifies a verification strategy and any strategy-specific
          additional data.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SignInResource_](signin.md#attributes)_&gt;_

This method returns a `Promise` which resolves with a `SignIn` object. Check the `secondFactorVerification` attribute for the status of the second factor verification process.
{% endtab %}
{% endtabs %}

### signInWithOAuth\(strategy, callbacks\) <a id="signinwithoauth"></a>

`signInWithOAuth(strategy: OAuthStrategy, callbacks: OAuthCallbacks) => Promise<void>`

Signs in users via OAuth. This is commonly known as Single Sign On \(SSO\), where an external account is used for verifying the user's identity.

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
      <td style="text-align:left"><b>strategy</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#oauthstrategy"><em>OAuthStrategy</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>The verification strategy. Select one of the supported OAuth providers.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>callbacks</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="signin.md#oauthcallbacks"><em>OAuthCallbacks</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>An object that specifies the URL for the OAuth provider redirect to, as
          well as the URL that the user should be redirected to upon successful sign
          in.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;void&gt;_

This method returns a `Promise` which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

## Interfaces

### AttemptFactorParams

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>strategy</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The <code>strategy</code> value depends on the object&apos;s <code>identifier</code> value.
          Each authentication identifier supports different verification strategies.
          Possible <code>strategy</code> values are:</p>
        <ul>
          <li><b>email_code</b>: User will receive a one-time authentication code via
            email. At least one email address should be on file for the user. The <code>email_address_id</code> parameter
            can also be specified to select one of the user&apos;s known email addresses.</li>
          <li><b>phone_code</b>: User will receive a one-time authentication code in
            their phone, via SMS. At least one phone number should be on file for the
            user. The <code>phone_number_id</code> parameter can also be specified to
            select which of the user&apos;s known phone numbers the SMS will go to.</li>
          <li><b>password</b>: The verification will attempt to be completed with the
            user&apos;s password.</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>code?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The one-time code that was sent to the user as part of this verification
          step. Required when <code>strategy</code> is <b>email_code</b> or <b>phone_code</b>.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>password?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s password string. Required when <code>strategy</code> is <b>password</b>.</p>
      </td>
    </tr>
  </tbody>
</table>

### OAuthCallbacks

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>callbackUrl</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The URL that the OAuth provider should redirect to, on successful authorization
          on their part.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>callbackUrlComplete</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The URL that the user will be redirected to, after successful authorization
          from the OAuth provider and Clerk sign in</p>
      </td>
    </tr>
  </tbody>
</table>

### PrepareFirstFactorParams

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>strategy</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The <code>strategy</code> value depends on the object&apos;s <code>identifier</code> value.
          Each authentication identifier supports different verification strategies.
          Possible <code>strategy</code> values are:</p>
        <ul>
          <li><b>email_code</b>: User will receive a one-time authentication code via
            email. At least one email address should be on file for the user. The <code>email_address_id</code> parameter
            can also be specified to select one of the user&apos;s known email addresses.</li>
          <li><b>phone_code</b>: User will receive a one-time authentication code in
            their phone, via SMS. At least one phone number should be on file for the
            user. The <code>phone_number_id</code> parameter can also be specified to
            select which of the user&apos;s known phone numbers the SMS will go to.</li>
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
      <td style="text-align:left"><b>email_address_id?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Unique identifier for the user&apos;s email address that will receive
          an email message with the one-time authentication code. This parameter
          will work only when the <b>email_code</b> strategy is specified.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>phone_number_id?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Unique identifier for the user&apos;s phone number that will receive an
          SMS message with the one-time authentication code. This parameter will
          work only when the <b>phone_code</b> strategy is specified.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>redirectUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The URL that the OAuth provider should redirect to, on successful authorization
          on their part. This parameter is required only for OAuth strategies (<b>oauth_*</b>).</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>actionCompleteRedirectUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The URL that the user will be redirected to, after successful authorization
          from the OAuth provider and Clerk sign in. This parameter is required only
          for OAuth strategies (<b>oauth_*</b>).</p>
      </td>
    </tr>
  </tbody>
</table>

### PrepareSecondFactorParams

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>strategy</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The strategy used for verification. Each authentication identifier supports
          different verification strategies. Possible <code>strategy</code> values
          are:</p>
        <ul>
          <li><b>phone_code</b>: User will receive a one-time authentication code in
            their phone, via SMS. The SMS message will be sent to the user&apos;s phone
            number that was added when setting up 2FA.</li>
        </ul>
      </td>
    </tr>
  </tbody>
</table>

### SignInFactor

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>strategy</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The strategy used for the authentication factor. The <code>strategy</code> value
          depends on the object&apos;s <code>identifier</code> value. Each authentication
          identifier supports different verification strategies. Possible <code>strategy</code> values
          are:</p>
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
      <td style="text-align:left"><b>email_address_id?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Unique identifier for the user&apos;s email address that will receive
          an email message with the one-time authentication code. This parameter
          will be present only for the <b>email_code</b> strategy.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>phone_number_id?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Unique identifier for the user&apos;s phone number that will receive an
          SMS message with the one-time authentication code. This parameter will
          be present only for the <b>phone_code</b> strategy.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>safe_identifier?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The actual value of the user&apos;s email address (for <b>email_code</b> strategy)
          or phone number (<b>phone_code</b> strategy). This parameter will be present
          only for the <b>email_code</b> and <b>phone_code</b> strategies.</p>
      </td>
    </tr>
  </tbody>
</table>

### SignInParams

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>identifier</b>
      </td>
      <td style="text-align:left"><em>string</em>
        <br />The authentication identifier for the sign in. This can be the value of
        the user&apos;s email address, phone number or username.</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>strategy?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Optional. Select the first factor verification strategy. The <code>strategy</code> value
          depends on the object&apos;s <code>identifier</code> value. Each authentication
          identifier supports different verification strategies. Possible <code>strategy</code> values
          are:</p>
        <ul>
          <li><b>email_code</b>: User will receive a one-time authentication code via
            email. At least one email address should be on file for the user. The <code>email_address_id</code> parameter
            can also be specified to select one of the user&apos;s known email addresses.</li>
          <li><b>phone_code</b>: User will receive a one-time authentication code in
            their phone, via SMS. At least one phone number should be on file for the
            user. The <code>phone_number_id</code> parameter can also be specified to
            select which of the user&apos;s known phone numbers the SMS will go to.</li>
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
      <td style="text-align:left"><b>password?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Supply the user&apos;s password, if <b>password</b> strategy has been specified.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>redirectUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The URL that the OAuth provider should redirect to, on successful authorization
          on their part. This parameter is required only for OAuth strategies (<b>oauth_*</b>).</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>actionCompleteRedirectUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The URL that the user will be redirected to, after successful authorization
          from the OAuth provider and Clerk sign in. This parameter is required only
          for OAuth strategies (<b>oauth_*</b>).</p>
      </td>
    </tr>
  </tbody>
</table>

### UserData

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>firstName?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s first name.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>lastName?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s last name.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>profileImageUrl?</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The URL of the user&apos;s profile image.</p>
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
        <p>&lt;em&gt;&lt;/em&gt;<a href="../errors.md#clerkapierror"><em>ClerkAPIError</em></a><em> | null</em>
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

## Types

### OAuthStrategy

`oauth_facebook | oauth_github | oauth_google | oauth_hubspot | oauth_tiktok`

| Value | Description |
| :--- | :--- |
| **oauth\_facebook** | Specify Facebook as the verification OAuth provider. |
| **oauth\_github** | Specify Github as the verification OAuth provider. |
| **oauth\_google** | Specify Google as the verification OAuth provider. |
| **oauth\_hubspot** | Specify Hubspot as the verification OAuth provider. |
| **oauth\_tiktok** | Specify TikTok as the verification OAuth provider. |

