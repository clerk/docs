---
description: How to setup social login with Bitbucket
---

# Bitbucket

## Overview

Adding social login with Bitbucket to your app with Clerk is straightforward -  you will need to populate the **Client ID**, **Client Secret** and **Authorized redirect URI** in your instance settings.

To make the development flow as smooth as possible, Clerk uses preconfigured shared OAuth credentials and redirect URIs for development instances - no other configuration is needed.&#x20;

For production instances, you will need to generate your own Consumer Key and Consumer Secret using your Bitbucket account.

{% hint style="info" %}
The purpose of this guide is to help you create a Bitbucket account and a Bitbucket OAuth Consumer - if you're looking for step-by-step instructions using Clerk to add social login (OAuth) to your application, follow the [Social login (OAuth)](../../popular-guides/social-login-oauth.md) guide.
{% endhint %}

## Before you start

* You need to create a Clerk Application in your [Clerk Dashboard](https://dashboard.clerk.dev). For more information, check out our [Setup your application](../../popular-guides/setup-your-application.md) guide.
* You need to have a Bitbucket account. To create one, [click here](https://bitbucket.org/account/signup).

## Configuring Bitbucket social login

Firstly, you will need to have a Bitbucket workspace.

You can navigate to the [list of all your workspaces](https://bitbucket.org/account/workspaces/) to select an existing workspace or create a new one:

![Bitbucket workspace listing](../../.gitbook/assets/bitbucket\_workspaces.png)

From there, click on the desired existing workspace or create a new one, then navigate to:

**Settings** > **Oauth Consumers**

![Bitbucket Oauth Consumer list](<../../.gitbook/assets/bitbucket\_oauth\_consumers (1).png>)

You will need to create a consumer if you don't have one already.

The first important setting for the connection to work is the **Callback URL**. In this field the **Authorized redirect URI** from your Clerk Bitbucket settings should be pasted:

![Oauth consumer creation](../../.gitbook/assets/bitbucket\_oauth\_consumer\_1.png)

The rest of the form fields (Name, Description, URL, Privacy Policy URL, EULA URL) can be set to match your application settings.

The second important setting is the selection of scopes that Clerk should request from Bitbucket upon connecting.

![Bitbucket Oauth Consumer scopes](<../../.gitbook/assets/image (7).png>)

Under **Account**, the **Email** & **Read** scopes should be enabled so that Clerk can use your basic Bitbucket account info when creating your Clerk user.

After saving the Oauth Consumer, Bitbucket will generate a Consumer Key & Consumer Secret for you. These can be displayed by clicking on the Consumer entry in your Oauth Consumer list.

You will need to copy:

* The Bitbucket Client **Key** to the **Client ID** input on Clerk
* The Bitbucket Client **Secret** to the **Client Secret** input on Clerk

![Client ID & Client Secret inputs for Bitbucket connection](<../../.gitbook/assets/image (8).png>)

That's it, now you are all set to sign up & sign in to Clerk with Bitbucket!
