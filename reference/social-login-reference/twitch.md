---
description: How to setup social login with Twitch
---

# Twitch

## Overview

Adding social login with Twitch to your app with Clerk is simple - you only need to set the **Client ID**, **Client Secret** and **Authorized redirect URI** in your instance settings.

To make the development flow as smooth as possible, Clerk uses preconfigured shared OAuth credentials and redirect URIs for development instances - no other configuration is needed.

For production instances, you will need to generate your own Client ID and Client secret using your Twitch account.

{% hint style="info" %}
The purpose of this guide is to help you create a Twitch account and a Twitch OAuth app - if you're looking for step-by-step instructions using Clerk to add social login (OAuth) to your application, follow the [Social login (OAuth)](broken-reference) guide.
{% endhint %}

## Before you start

* You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). For more information, check out our [Setup your application](broken-reference) guide.
* You need to have a Twitch account. To create one, [click here](https://www.twitch.tv).

## Configuring Twitch social login

First, you need to register a new OAuth Twitch app at the [Twitch Developers Console](https://dev.twitch.tv/console).

![](../../.gitbook/assets/twitch-create-oauth-app-1.png) ![Creating an OAuth Twitch app](../../.gitbook/assets/twitch-create-oauth-app-2.png)

Set a name and the a category for your new application. You also need to add the **OAuth Redirect URLs.** Go to the [Social Login page](https://dashboard.clerk.dev/last-active?path=authentication/social) and enable Twitch. In the modal that opened, ensure **Use custom credentials** is enabled and copy **Authorized redirect URI**. Paste the value into the **OAuth Redirect URLs** input and click create.

Once all the above are complete, copy the **Client ID** and **Client Secret.** Go back to the Clerk Dashboard and paste them into the respective fields.

![](../../.gitbook/assets/twitch-credentials.png)

Don't forget to click **Apply** in the Clerk dashboard. Social login with Twitch is now configured 🔥

## Next Steps

Learn how to add social login with Twitch to your Clerk application by following the [Social login (OAuth)](broken-reference) guide.
