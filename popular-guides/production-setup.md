---
description: Learn how to deploy your app to production.
---

# Deploy to production

## Overview

Before you share your app to the world, you'll need to create a production instance.  While development instances are convenient, they are not as secure as Clerk's production environment.

{% hint style="danger" %}
It is strongly recommended that you **do not** go live with a development instance,  we are soon going to put up warning messages around this.
{% endhint %}

{% hint style="info" %}
Deploying on a Clerk provided domain is coming soon.
{% endhint %}

## Before you start

* You should already have your app working with a development instance.
* You will need to be able to add DNS records on your domain.
* You will need Social Login (OAuth) credentials for any providers that you would like to use in production.

## Creating your production instance

![](../.gitbook/assets/screenshot\_20211020\_105434.png)

The modal is relatively straight forward, but there's a few things you should know.  Your development and production instances are separate entities, and don't share anything - to match your development instance you should clone it's settings, which we make easy.  If you ever need to change a setting, you should do it in development first to make sure everything works, before making the change to your production instance. &#x20;

{% hint style="warning" %}
URL and redirects settings may not copy over, you will need to set these values again.
{% endhint %}

{% hint style="warning" %}
Integrations you have setup will not copy over.  You will need to re-enable these.
{% endhint %}

## API Keys and Environment variables

A common mistake when deploying to production is forgetting to change your API keys to your production instances keys.  The best way to set this up, is to make use of **Environment variables.  **All modern hosting providers, such as AWS, GCP, Vercel, Heroku, Render, etc.. make it easy to add these values.  Locally, you should use an `env` file, this way these values are being set dynamically depending on your environment.  Here's a list of Clerk variables you'll need to change:

* Frontend API
  * &#x20;Formatted `clerk.123.abc.lcl.dev` in development, and `clerk.example.com` in production.  and is passed to the **\<ClerkProvider /> **during initialization
* API Key
  * Formatted `test_xxxxx` in development,  and `live_xxxxx` in production.  These values are used to access Clerk's Backend API
* OAuth credentials
  * In development, Clerk provides you with a set of shared OAuth credentials.  These are not secure in production, and you will need to provide your own.



## DNS Records

Clerk uses DNS records to provide session management, and emails verified from your domain.  You will need to go to **Settings > DNS** to see the records that you need to set.  Note: It can take up to 24hrs for DNS Records to fully propagate, so be patient. &#x20;

![](../.gitbook/assets/screenshot\_20211020\_111528.png)

{% hint style="danger" %}
Some DNS providers will automatically append the domain for the 'To' field, if this is the case, you will need to omit your domain from the 'To' field.  So, add 'clerk' instead of 'clerk.example.com'.
{% endhint %}

{% hint style="warning" %}
**Do not proxy your  records**

Clerk uses a DNS check to validate this CNAME record.

If this subdomain is reverse proxied behind a service that points to generic hostnames, such as Cloudflare, the DNS check will fail. Please set the DNS record for this subdomain to a "DNS only" mode on your host to prevent proxying.
{% endhint %}

## Deploy

That's all! Once you've completed all the above steps, you're ready to go to the home page, and press **Deploy**!  If you run into any trouble feel free to reach out to any of our [support channels](https://clerk.dev/support) and we can help you out.

![](<../.gitbook/assets/Screenshot 2021-10-01 at 12.49.54 PM.png>)
