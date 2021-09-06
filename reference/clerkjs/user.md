---
description: The User object contains the user's account information.
---

# User

## Overview

The `User` object holds all the information for a user of your application and provides a set of methods to manage their account.

Users have a unique authentication identifier which might be their email address, phone number or a username.

Users can be contacted at their primary email address or primary phone number. They can have more than one registered email addresses, but only one of them will be their primary email address. Same thing with phone numbers; they can have more than one, but only one will be their primary.

Finally, `User` objects hold profile data like their name, a profile image and a set of metadata that can be used internally to store arbitrary information. The metadata are split into public and private. Both types are set from the [Backend API](../backend-api-reference/), but public metadata can be accessed from the [Frontend API](../frontend-api-reference/) and [Backend API](../backend-api-reference/).

The ClerkJS SDK provides some helper [methods](user.md#methods) on the `User` object to help retrieve and update user information and authentication status.

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
        <p>A unique identifier for the user.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>firstName</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The user&apos;s first name.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>lastName</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The user&apos;s last name.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>fullName</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The user&apos;s full name, a combination of their first and last name.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>username</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The user&apos;s username.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>profileImageUrl</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The URL for the user&apos;s profile image.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>primaryEmailAddress</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="emailaddress.md"><em>EmailAddress</em></a><em> | null</em>
        </p>
        <p>Information about the user&apos;s primary email address.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>primaryEmailAddressId</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The unique identifier for the <a href="emailaddress.md">EmailAddress</a> that
          the user has set as primary.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>emailAddresses</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="emailaddress.md"><em>EmailAddress</em></a><em>[]</em>
        </p>
        <p>Any array of all the <a href="emailaddress.md">EmailAddress</a> objects
          associated with the user. Includes the primary.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>primaryPhoneNumber</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="phonenumber.md"><em>PhoneNumber </em></a><em>| null</em>
        </p>
        <p>Information about the user&apos;s primary phone number.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>primaryPhoneNumberId</b>
      </td>
      <td style="text-align:left">
        <p><em>string | null</em>
        </p>
        <p>The unique identifier for the <a href="phonenumber.md">PhoneNumber</a> that
          the user has set as primary.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>phoneNumbers</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="phonenumber.md"><em>PhoneNumber</em></a><em>[]</em>
        </p>
        <p>Any array of all the <a href="phonenumber.md">PhoneNumber</a> objects associated
          with the user. Includes the primary.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>externalAccounts</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="user.md#externalaccount"><em>ExternalAccount</em></a><em>[]</em>
        </p>
        <p>Any array of all the <a href="user.md#externalaccount">ExternalAccount</a> objects
          associated with the user with OAuth.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>passwordEnabled</b>
      </td>
      <td style="text-align:left">
        <p><em>boolean</em>
        </p>
        <p>A boolean indicating whether the user has a password on their account.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>publicMetadata</b>
      </td>
      <td style="text-align:left">
        <p><em>{[string]: any} | null</em>
        </p>
        <p>Metadata that can be read from the <a href="../frontend-api-reference/">Frontend API</a> and
          <a
          href="../backend-api-reference/">Backend API</a>and can be set only from the Backend API .</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>privateMetadata</b>
      </td>
      <td style="text-align:left">
        <p><em>{[string]: any} | null</em>
        </p>
        <p>Metadata that can be read and set only from the <a href="../backend-api-reference/">Backend API</a>.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>unsafeMetadata</b>
      </td>
      <td style="text-align:left">
        <p><em>{[string]: any} | null</em>
        </p>
        <p>Metadata that can be read and set from the frontend. <em><br /></em>One
          common use case for this attribute is to use it to implement custom fields
          that will be attached to the User object.
          <br />Please note that there is also an <code>unsafeMetadata</code> attribute
          in the <a href="signup.md">SignUp</a> object. The value of that field will
          be automatically copied to the user&apos;s unsafe metadata once the sign
          up is complete.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>createdAt</b>
      </td>
      <td style="text-align:left">
        <p><em>Date</em>
        </p>
        <p>Date when the user was first created.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>updatedAt</b>
      </td>
      <td style="text-align:left">
        <p><em>Date</em>
        </p>
        <p>Date of the last time the user was updated.</p>
      </td>
    </tr>
  </tbody>
</table>

## Methods

### createEmailAddress\(email\)

`createEmailAddress(email: string) => Promise<EmailAddressResource>`

Adds an email address for the user. A new [EmailAddress](emailaddress.md) will be created and associated with the user.

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
      <td style="text-align:left"><b>email</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The email address to associate with the user.</p>
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

### createPhoneNumber\(phoneNumber\)

`createPhoneNumber(phoneNumber: string) => Promise<PhoneNumberResource>`

Adds a phone number for the user. A new [PhoneNumber](phonenumber.md) will be created and associated with the user.

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
      <td style="text-align:left"><b>phoneNumber</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The phone number to associate with the user.</p>
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

### delete\(\)

`delete() => Promise<void>`

Deletes the user.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;void&gt;_

This method returns a `Promise` which doesn't resolve to any value.
{% endtab %}
{% endtabs %}

### getSessions\(\)

`getSessions() => Promise<SessionWithActivities[]>`

Retrieves all **active** sessions for this user. This method uses a cache so a network request will only be triggered only once.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_SessionWithActivities_](sessionwithactivities.md)_\[\]&gt;_

This method returns a `Promise` which resolves to an array of `SessionWithActivities` objects.
{% endtab %}
{% endtabs %}

### getToken\(service, options?\)

`getToken(service: JWTService, options?: UserGetTokenOptions) => Promise<string>`

Retrieves the user's token for the given integration service. This method uses a cache so a network request will only be made if the token in memory has expired. The TTL for each token is one minute.

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
      <td style="text-align:left"><b>service</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="user.md#jwtservice"><em>JWTService</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>The name of the service that you&apos;ve integrated with.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>options?</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="user.md#usergettokenoptions"><em>UserGetTokenOptions</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>Optionally pass the leeway for expiring the cache. Default leeway is 10
          seconds and cannot exceed the token TTL, which is 1 minute.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;string&gt;_  
Returns a `Promise` that resolves to a `string`. The string is the user's token for the provided integration service.
{% endtab %}
{% endtabs %}

### setProfileImage\(file\)

`setProfileImage(file: Blob | File) => Promise<ImageResource>`

Add the user's profile image or replace it if one already exists. This method will upload an image and associate it with the user.

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
      <td style="text-align:left"><b>file</b>
      </td>
      <td style="text-align:left">
        <p><em>Blob | File</em>
        </p>
        <p>A file or file-like object (raw data). Should be an image.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_ImageResource_](user.md#imageresource)_&gt;_

This method
{% endtab %}
{% endtabs %}

### twoFactorEnabled\(\)

`twoFactorEnabled() => boolean`

Checks if the user has enabled 2-step verification \(2FA\) for their account.

{% tabs %}
{% tab title="Parameters" %}
This method accepts no parameters.
{% endtab %}

{% tab title="Returns" %}
_boolean_

This method returns `true` when the user has enabled 2-factor authentication, `false` otherwise.
{% endtab %}
{% endtabs %}

### update\(params\)

`update(params: UpdateUserParams) => Promise<UserResource>`

Updates the user's attributes. Use this method to save information you collected about the user. 

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
        <p>&lt;em&gt;&lt;/em&gt;<a href="user.md#updateuserparams"><em>UpdateUserParams</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>User profile related attributes.</p>
      </td>
    </tr>
  </tbody>
</table>
{% endtab %}

{% tab title="Returns" %}
_Promise&lt;_[_UserResource_](user.md)_&gt;_

This method returns a `Promise` which resolves to a `User` object.
{% endtab %}
{% endtabs %}

## Interfaces

### ExternalAccount

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
        <p>A unique identifier for this external account.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>provider</b>
      </td>
      <td style="text-align:left">
        <p>&lt;em&gt;&lt;/em&gt;<a href="user.md#oauthprovider"><em>OAuthProvider</em></a>&lt;em&gt;&lt;/em&gt;</p>
        <p>The name of the OAuth provider.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>externalId</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s unique identifier at the OAuth provider.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>emailAddress</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s email address registered with the OAuth provider.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>approvedScopes</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>A list of <a href="https://oauth.net/2/scope/">OAuth scopes</a> as returned
          by the OAuth provider.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>firstName</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s first name as registered with the OAuth provider.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>lastName</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s first name as registered with the OAuth provider.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>picture</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>URL for the user&apos;s profile picture (avatar) that&apos;s registered
          with the OAuth provider.</p>
      </td>
    </tr>
  </tbody>
</table>

### ImageResource

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
        <p>A unique identifier for this image.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>name</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>A name for this image. The image title.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>publicUrl</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The URL which can be used to access this image.</p>
      </td>
    </tr>
  </tbody>
</table>

### UpdateUserParams

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>username</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The username for the user.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>firstName</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s first name.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>lastName</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s last name.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>primaryEmailAddressId</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Use this attribute to reference an <a href="emailaddress.md">EmailAddress</a> object
          by ID and associate it with the user.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>primaryPhoneNumberId</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>Use this attribute to reference a <a href="phonenumber.md">PhoneNumber</a> object
          by ID and associate it with the user.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>password</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s password.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>unsafeMetadata</b>
      </td>
      <td style="text-align:left">
        <p><em>{[string]: any} | null</em>
        </p>
        <p>Metadata that can be read and set from the frontend. <em><br /></em>One
          common use case for this attribute is to use it to implement custom fields
          that will be attached to the User object.</p>
      </td>
    </tr>
  </tbody>
</table>

### UserGetTokenOptions

<table>
  <thead>
    <tr>
      <th style="text-align:left">Property</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>leewayInSeconds?</b>
      </td>
      <td style="text-align:left">
        <p><em>number</em>
        </p>
        <p>The number of seconds that we allow the token to be cached.</p>
      </td>
    </tr>
  </tbody>
</table>

## Types

### JWTService

`clerk | firebase | hasura`

| Value | Description |
| :--- | :--- |
| **clerk** | Get a short-lived Clerk JWT. |
| **firebase** | Integration with [Firebase](../../integrations/firebase.md). |
| **hasura** | Integration with [Hasura](../../integrations/hasura.md). |

### OAuthProvider

`facebook | github | google | hubspot | tiktok`

| Value | Description |
| :--- | :--- |
| **facebook** | Facebook OAuth provider. |
| **github** | Github OAuth provider. |
| **google** | Google OAuth provider. |
| **hubspot** | Hubspot OAuth provider. |
| **tiktok** | TikTok OAuth provider. |

