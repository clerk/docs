---
description: How to setup social login with GitHub
---

# GitHub

## Overview

Adding social login with GitHub to your app with Clerk is simple -  you only need to set the **Client ID**, **Client Secret** and **Authorized redirect URI** in your instance settings.

To make the development flow as smooth as possible, Clerk uses preconfigured shared OAuth credentials and redirect URIs for development instances - no other configuration is needed.&#x20;

For production instances, you will need to generate your own Client ID and Client secret using your GitHub account.

{% hint style="info" %}
The purpose of this guide is to help you create a GitHub account and a GitHub OAuth app - if you're looking for step-by-step instructions using Clerk to add social login (OAuth) to your application, follow the [Social login (OAuth)](../../popular-guides/social-login-oauth.md) guide.
{% endhint %}

## Before you start

* You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). For more information, check out our [Setup your application](../../popular-guides/setup-your-application.md) guide.
* You need to have a GitHub account. To create one, [click here](https://github.com/signup) .

## Configuring GitHub social login

First, you need to register a new OAuth GitHub app. Follow the official GitHub instructions on [how to create an OAuth app](https://docs.github.com/en/developers/apps/building-oauth-apps/creating-an-oauth-app).

![Registering an OAuth GitHub app](../../.gitbook/assets/screely-1628426717069.png)

Go to the [Social Login page](https://dashboard.clerk.dev/last-active?path=authentication/social) and enable GitHub. In the modal that opened, ensure **Use custom credentials** is enabled and copy **Authorized redirect URI**. Go back to the GitHub panel, paste the value into the **Authorization callback URL** field and complete the registration.

Once registration is complete, you'll get redirected to project's admin panel. Click the **Generate a new client secret** button to get your new client secret. Then, copy the **Client ID** and **Client secret.** Go back to the Clerk Dashboard and paste them into the respective fields.

![Obtaining the Client ID and Client secret](../../.gitbook/assets/screely-1628427343412.png)

Don't forget to click **Apply** in the Clerk dashboard. Social login with GitHub is now configured 🔥&#x20;

## Next Steps

Learn how to add social login with GitHub to your Clerk application by following the [Social login (OAuth)](../../popular-guides/social-login-oauth.md) guide.
