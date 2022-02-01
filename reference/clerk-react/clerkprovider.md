---
description: >-
  The ClerkProvider component wraps your React application to provide active
  session and user context to Clerk's hooks and other components.
---

# \<ClerkProvider />

## Props

| Name                | Description                                                                                                                                                                                                                                          |
| ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **frontendApi**     | <p><em>string</em></p><p>The frontend API host for your instance. This can be found in your clerk Dashboard.</p>                                                                                                                                     |
| **navigate?**       | <p><em>(to: string) => Promise&#x3C;any> | void</em></p><p>A function which takes the destination path as an argument and performs a "push" navigation.</p>                                                                                          |
| **clerkJsVariant?** | <p><em>string</em></p><p>If your web application uses only <a href="../../components/control-components/">Control components</a> you can set this value to <code>headless</code> and load a minimal ClerkJS bundle for optimal page performance.</p> |

## Usage

The ClerkProvider component must be added to your React entrypoint.  For more information, please reference the Installation guide:

{% content-ref url="installation.md" %}
[installation.md](installation.md)
{% endcontent-ref %}

