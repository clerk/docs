---
description: How to setup social login with LinkedIn
---

# LinkedIn

## Overview

Adding social login with LinkedIn to your app with Clerk is simple -  you only need to set the **Client ID**, **Client Secret** and **Authorized redirect URI** in your instance settings.

To make the development flow as smooth as possible, Clerk uses preconfigured shared OAuth credentials and redirect URIs for development instances - no other configuration is needed.&#x20;

For production instances, you will need to generate your own Client ID and Client secret using your LinkedIn account.

{% hint style="info" %}
The purpose of this guide is to help you create a LinkedIn account and a LinkedIn OAuth app - if you're looking for step-by-step instructions using Clerk to add social login (OAuth) to your application, follow the [Social login (OAuth)](../../popular-guides/social-login-oauth.md) guide.
{% endhint %}

## Before you start

* You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). For more information, check out our [Setup your application](../../popular-guides/setup-your-application.md) guide.
* You need to have a LinkedIn account. To create one, [click here](https://developer.linkedin.com).

## Configuring LinkedIn social login

First, you need to create a new OAuth LinkedIn app.

![Creating an OAuth LinkedIn app](../../.gitbook/assets/oauth\_linkedin\_create\_app.png)

You need to set a name, associate a LinkedIn page with it and finally upload a logo for your new application. Go to the [Clerk Dashboard](https://dashboard.clerk.dev), select your application **** and instance and go to **Authentication -> Social Login**. Click the **Manage connection** button under the LinkedIn provider, copy the **Authorized redirect URI** and paste the value into the **Redirect URL**, as shown below.

![Obtaining the Application ID and Client secret](../../.gitbook/assets/oauth\_linkedin\_credentials.png)

The last thing you have to do is to enable the _Sign in with LinkedIn_ product for your OAuth application.&#x20;

![Enable Sign in with LinkedIn product](../../.gitbook/assets/oauth\_linkedin\_enable\_sign\_in.png)

Copy the **Application ID** and **Secret** from the previous screen**.** Go back to the Clerk Dashboard and paste them into the respective fields.

Don't forget to click **Apply** in the Clerk dashboard. Social login with LinkedIn is now configured 🔥
