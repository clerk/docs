---
description: How to setup social login with Notion
---

# Notion

## Overview

Adding social login with Notion to your app with Clerk is simple - you only need to set the **Client ID**, **Client Secret** and **Authorized redirect URI** in your instance settings.

To make the development flow as smooth as possible, Clerk uses preconfigured shared OAuth credentials and redirect URIs for development instances - no other configuration is needed.

For production instances, you will need to generate your own Client ID and Client secret using your LinkedIn account.

{% hint style="info" %}
The purpose of this guide is to help you create a Notion developer account and a Notion OAuth app - if you're looking for step-by-step instructions using Clerk to add social login (OAuth) to your application, follow the [Social login (OAuth)](broken-reference) guide.
{% endhint %}

## Before you start

* You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). For more information, check out our [Setup your application](broken-reference) guide.
* You need to have a Notion developer account. To create one, [click here](https://developers.notion.com).

## Configuring Notion social login

First, you need to create a new OAuth Notion app.

![Creating an OAuth Notion app](../../.gitbook/assets/oauth-notion-create-app.png)

You need to set a name, a logo and associate a Notion workspace with it. Make sure you are also requesting the user email address along with the other information and click Submit.

![](../../.gitbook/assets/oauth-notion-public-integration.png)

Go to the [Social Login page](https://dashboard.clerk.dev/last-active?path=authentication/social) and enable Notion. In the modal that opened, ensure **Use custom credentials** is enabled and copy **Authorized redirect URI**. Paste the value into the **Redirect URIs**, as shown below, after changing the integration type to **Public**. Fill also any other information required from Notion and click Submit.

![Obtaining the Client ID and Client secret](../../.gitbook/assets/oauth-notion-credentials.png)

Copy the **Client ID** and **Client secret** as shown in the above image\*\*.\*\* Go back to the Clerk Dashboard and paste them into the respective fields.

Don't forget to click **Apply** in the Clerk dashboard. Social login with Notion is now configured 🔥
