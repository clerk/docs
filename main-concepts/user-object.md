---
description: The User object models a user's account information.
---

# The User object

## Overview

The `User` object contains all account information that describes a user of your app in Clerk. Users can authenticate and manage their accounts, updating their personal and contact info, or setting up security features for their account.

{% hint style="info" %}
For more information on users, see the[ User object reference](https://docs.clerk.dev/reference/clerkjs/user#attributes).
{% endhint %}

## Create users

Users can be created through the [Sign up flow](sign-up-flow.md). They have to provide a unique authentication identifier which can be an email address, phone number or username.

If the sign up attempt succeeds, a new session and user is created for your application instance.&#x20;

After creation, the user's authentication identifier can be used at any time to complete the [Sign in flow](sign-in-flow.md) and get an active [Session](../reference/clerkjs/session.md).

Whenever you have an active session, that means you can access the currently signed in `User` object, through the `Clerk` object.

{% hint style="info" %}
The current user can be accessed at any time through the [Clerk object](../reference/clerkjs/clerk.md), or the [useUser](../reference/clerk-react/useuser-hook.md) React hook.
{% endhint %}

Users can also be created through the [Backend API](../reference/clerkjs/signin/backend-api-reference/users.md#create-a-user). Finally, you can [import your existing users](../reference/import-users.md).

## Identify users

You can uniquely identify your users in two ways; either through their Clerk unique identifier (their `id`), or their unique authentication identification, as described above.&#x20;

It is strongly recommended that you verify your user's identification (email address, phone number) upon sign up. Clerk currently asks for verification by default.

There are currently many different ways to verify your user's unique authentication identifier and they vary depending on what that identifier is.

* Email addresses can be verified via magic links, [one-time codes](../popular-guides/passwordless-authentication.md) or through a trusted [social (OAuth) provider](../popular-guides/social-login-oauth.md).
* Phone numbers can be verified via [one-time codes](../popular-guides/passwordless-authentication.md) that are sent with SMS.
* Usernames must be guaranteed to be unique and can be verified only through a trusted [social (OAuth) provider](../popular-guides/social-login-oauth.md).

## Contact users

Users can be contacted at their primary email address or primary phone number.&#x20;

They can have more than one registered email addresses, but only one of them will be their primary email address. Same thing with phone numbers; they can have more than one, but only one will be their primary.

Clerk takes care of automatically assigning a verified email address or phone number as the user's primary contact method, but users are allowed to pick one through their account settings as well.

## Custom user attributes

`User` objects hold a set of metadata that can be used internally to store arbitrary information.&#x20;

The metadata are split into public and private. Both types are set from the [Backend API](../reference/clerkjs/signin/backend-api-reference/), but public metadata can be accessed from the [Frontend API](../reference/frontend-api-reference/) and [Backend API](../reference/clerkjs/signin/backend-api-reference/). These types of metadata can be used to access non-sensitive information about a user, such as their profile picture or their name. They can be used to build UIs where a user might not be signed in, but you still want to show some of their profile info.

There's also a third kind of metadata which are called "unsafe", since they can be set and accessed by both the [Frontend API](../reference/frontend-api-reference/) and the [Backend API](../reference/clerkjs/signin/backend-api-reference/). They provide a quick method to add custom attributes to a user. These attributes will be stored on the `User` object and will be available for access at all times.&#x20;

The "unsafe" custom attributes can be set upon sign up, when [creating](../reference/clerkjs/signup.md#create-params) or [updating](../reference/clerkjs/signup.md#update-params) a [`SignUp`](../reference/clerkjs/signup.md) object. After a successful sign up, these attributes will be copied to the `User` object. From that point on they can be accessed as a direct [attribute](../reference/clerkjs/user.md#attributes) of the [`User`](../reference/clerkjs/user.md) object.

Here's a comparison of the different metadata attributes on the User object and how each can be used.

| Metadata    | Frontend API | Backend API |
| ----------- | ------------ | ----------- |
| **private** | -            | Read/Write  |
| **public**  | Read         | Read/Write  |
| **unsafe**  | Read/Write   | Read/Write  |

Choosing which metadata field you should use for setting custom attributes on your users becomes a matter of access and visibility.&#x20;

Do you need to set the custom attributes from the front-end (using our ClerkJS library or the Frontend API)? You should choose the **unsafeMetadata** property.

Will the custom attributes contain sensitive information that should not be displayed on the front-end? Use the **privateMetadata** property.

Do you need to set some metadata from your back-end and have them displayed as read-only on the front-end? Use the **publicMetadata** property.

{% hint style="info" %}
Looking for inspiration on how to set custom user attributes? Our [Workflow starter repo](https://github.com/clerkinc/clerk-workflow-next) demonstrates storing additional user info like their company, country and whether they accept the terms and conditions.
{% endhint %}

## Update user profiles

`User` objects hold profile data like their name, a profile photo, email addresses, phone numbers, connected OAuth accounts and a set of metadata that can be used internally to store arbitrary information.&#x20;

Users are free to [update their profile info](../reference/clerkjs/user.md#update-params) and [change their profile picture](../reference/clerkjs/user.md#setprofileimage-file). They can add as many [email addresses](../reference/clerkjs/user.md#createemailaddress-email) or [phone numbers](../reference/clerkjs/user.md#createphonenumber-phonenumber) as they want. They can also manage their connected OAuth accounts.

Clerk provides ready-made [components](../components/user-profile.md) and [hosted pages](clerk-hosted-pages.md) that enable updating a user's profile out of the box.&#x20;

## User account security

Regarding account security, Clerk provides features that strengthen a user's account with little effort.

Users can enable [multi-factor authentication](../popular-guides/multi-factor-authentication.md) for their accounts. In this case, every time they successfully authenticate, they must complete an additional challenge to verify their identity. Currently, only SMS based multi-factor authentication is supported, and it can be enabled by [creating a phone number](../reference/clerkjs/user.md#createphonenumber-phonenumber) for the `User` object and then configuring it for multi-factor authentication.&#x20;

Users can also change their password. Frequently updating a password leads to safer authentication. Even if a password is compromised, the new one can still be used to sign in. You can change a user's password by simply [updating](../reference/clerkjs/user.md#update-params) the `User` object.

Finally, users can get a [list of their active sessions](../reference/clerkjs/user.md#getsessions) and information about the device they signed in from.

Clerk provides ready-made [components](../components/user-profile.md) and [hosted pages](clerk-hosted-pages.md) that enable updating a user's profile out of the box.&#x20;

## Deleting users

Clerk allows user accounts to be deleted. Currently, this can be done only through the [Backend API](../reference/clerkjs/signin/backend-api-reference/users.md#delete-a-user).
