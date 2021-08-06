---
description: >-
  The ClerkProvider component wraps your React application to provide active
  session and user context to Clerk's hooks and other components.
---

# &lt;ClerkProvider /&gt;

## Props

<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>frontendApi</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The frontend API host for your instance. This can be found in your clerk
          Dashboard.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>navigate?</b>
      </td>
      <td style="text-align:left">
        <p><em>(to: string) =&gt; Promise&lt;any&gt; | void</em>
        </p>
        <p>A function which takes the destination path as an argument and performs
          a &quot;push&quot; navigation.</p>
      </td>
    </tr>
  </tbody>
</table>

## Usage

The ClerkProvider component must be added to your React entrypoint.  For more information, please reference the Installation guide:

{% page-ref page="installation.md" %}



