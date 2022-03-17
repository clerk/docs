---
description: How to setup social login with Dropbox
---

# Dropbox

## Overview

Adding social login with Dropbox to your app with Clerk is simple -  you only need to set the **Client ID**, **Client Secret** and **Authorized redirect URI** in your instance settings.

To make the development flow as smooth as possible, Clerk uses preconfigured shared OAuth credentials and redirect URIs for development instances - no other configuration is needed.&#x20;

For production instances, you will need to generate your own Client ID and Client secret using your Dropbox account.

{% hint style="info" %}
The purpose of this guide is to help you create a Dropbox account and a Dropbox OAuth app - if you're looking for step-by-step instructions using Clerk to add social login (OAuth) to your application, follow the [Social login (OAuth)](../../popular-guides/social-login-oauth.md) guide.
{% endhint %}

## Before you start

* You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). For more information, check out our [Setup your application](../../popular-guides/setup-your-application.md) guide.
* You need to have a Dropbox account. To create one, [click here](https://www.twitch.tv).

## Configuring Dropbox social login

First, you need to register a new OAuth Dropbox app at the [Dropbox App Console](https://www.dropbox.com/developers/apps).

![Creating an OAuth Dropbox app](../../.gitbook/assets/oauth\_dropbox\_create\_app.png)

First you need to choose the API type, the App's type of access and to set a name for your new application. You also need to add the **OAuth Redirect URLs.** Go to the [Social Login page](https://dashboard.clerk.dev/last-active?path=authentication/social) and enable Dropbox. In the modal that opened, ensure **Use custom credentials** is enabled and copy **Authorized redirect URI**. Paste the value into the **OAuth Redirect URIs** input and click create.

Once all the above are complete, copy the **Client ID** and **Client Secret.** Go back to the Clerk Dashboard and paste them into the respective fields.

![Obtaining the Client ID and Client secret](../../.gitbook/assets/oauth\_dropbox\_credentials.png)

Don't forget to click **Apply** in the Clerk dashboard. Social login with Dropbox is now configured 🔥&#x20;
