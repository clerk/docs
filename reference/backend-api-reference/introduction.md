# Introduction

The Clerk Backend API is built to be used from your backend code. We make our best attempt to make your Clerk resources available to you in a predictable manner. All requests accept form-encoded request bodies, and respond with a JSON-encoded body.

> Base URL: **https://api.clerk.dev/**

### Authentication

The Clerk backend API uses API keys to authenticate these requests.  You can find and create your API keys on the Clerk dashboard via an instance's  Settings>API keys. These API keys should never be shared with anyone, as they allow the holder to access all of your Clerk resources.  Be careful not to expose these keys in your Git repository, frontend code, or anywhere else that is public. &#x20;
