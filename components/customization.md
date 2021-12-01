---
description: >-
  Learn to customize the look and feel of Clerk's components to match your
  application.
---

# Customization

## Overview

When your users are going through the authentication process, you need to make sure they have  a seamless experience. Clerk allows you to customize the look and feel of your components to match the rest of your application using theming options.

By using the theming options, you can change visual elements of your authentication flow like:&#x20;

* Font colors
* Button colors
* Border styles
* and more...

{% hint style="info" %}
The customization options detailed here are a code-side alternative to [instance-level theming](../popular-guides/setup-your-application.md#theme) configured through the Clerk dashboard. Every theming option provided at code-side will be **merged** with the theming defined in your application dashboard.
{% endhint %}

Components can be customized with the **`theme`** prop or by defining Clerk-specific [CSS variables](https://developer.mozilla.org/en-US/docs/Web/CSS/Using\_CSS\_custom\_properties). Both ways are described in this guide.

{% hint style="warning" %}
Customization options have an order of precedence from which overrides are applied. This precedence order is shown below:\
\
1\. CSS Variables _(takes precedence over)_ **→**\
__2_._ theme prop **** (_takes precedence over)_  **→**\
****3. Dashboard settings
{% endhint %}

## Using the `theme` prop

To customize your components using the **`theme`** prop, you need to provide an object to `<ClerkProvider />` with the parameters you wish to change.

{% tabs %}
{% tab title="React JSX" %}
```jsx
import { ClerkProvider } from "@clerk/clerk-react";

const customTheme = {...};

function App(){
    
    return (
        <ClerkProvider theme={customTheme}>
            {/* Your wrapped application. */}
        </ClerkProvider>
    );
}
```
{% endtab %}

{% tab title="React TSX" %}
```jsx
import { ClerkProvider } from "@clerk/clerk-react";
import { ClerkThemeOptions } from "@clerk/types" 

const customTheme: ClerkThemeOptions = {...};

function App(){
    
    return (
        <ClerkProvider theme={customTheme}>
            {/* Your wrapped application. */}
        </ClerkProvider>
    );
}
```
{% endtab %}
{% endtabs %}

Now all components rendered as a child of `<ClerkProvider />` will inherit the theme options you have defined. There is no need to change any individual component, the theme changes are applied automatically.

### Custom theme properties

The theme object should conform to the `ClerkThemeOptions` interface. \
An example of the options you can pass on the theme prop is shown below:

```javascript
const customThemeOptions = {
  general: {
    /* Primary color */
    color: "#f1f1f1",
    
    /* Components background color */
    backgroundColor: "#f2f2f2",
    
    /* Components font color */
    fontColor: "#f3f3f3",
   
     /* Components font family */
    fontFamily: "Inter, sans serif",
    
    /* Components label font weight */
    labelFontWeight: "500",
    
    /* Padding mod for selected Clerk elements */
    padding: "1em",
    
    /* Border radius for selected Clerk elements */
    borderRadius: "20px",
    
    /* Box shadow applied on form Clerk components */
    boxShadow: "0 2px 8px rgba(0, 0, 0, 0.2)",
  },
  buttons: {
    /* Clerk buttons font color */
    fontColor: "#f4f4f4",
    
    /* Clerk buttons font family */
    fontFamily: "Inter, sans serif"
    
    /* Clerk buttons font weight */
    fontWeight: "300",
  },
}
```

## Using CSS variables

To customize components using CSS variables, you need to define the properties you wish to override under the `.cl-component` scope.

```css
.cl-component {
  // your Clerk theming options go here
}
```

### &#x20;Available CSS variables

The available CSS variables you can override can be seen below:

```css
.cl-component {
    /* Primary color */
    --clerk-primary: #335bf1;
    
    /* Components background color */
    --clerk-background-color: #ffffff;
    
    /* Components font color */    
    --clerk-font-color: #151515;

    /* Components font family */
    --clerk-font-family: "Inter", sans-serif;

    /* Components label font weight */
    --clerk-label-font-weight: 600;    

    /* Padding mod for selected Clerk elements */    
    --clerk-padding-mod: 1;

    /* Border radius for selected Clerk elements */
    --clerk-border-radius: 0.5em;
    
    /* Box shadow applied on form Clerk components */
    --clerk-box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
    
    /* Clerk buttons font color */    
    --clerk-button-font-color: #ffffff;

    /* Clerk buttons font family */
    --clerk-button-font-family: "Inter", sans-serif;

    /* Clerk buttons font weight */    
    --clerk-button-font-weight: 600;
  }
```

## Showcase

Below you can find three different examples showcasing the capabilities of Clerk's theme customization.

{% tabs %}
{% tab title="Futuristic" %}
```javascript
// Supplying a theme prop
const futuristicThemeOptions = {
  general: {
    color: "#05f4b7",
    backgroundColor: "#12151f",
    fontColor: "#05f4b7",
    borderRadius: "20px",
    boxShadow: "2px 4px 12px #05f4b7",
  },
  buttons: {
    fontColor: "#fff",
  },
}



// Using CSS variables.
.cl-component {
  --clerk-border-radius: 20px;
  --clerk-primary: #05f4b7;
  --clerk-font-color: #05f4b7;
  --clerk-background-color: #12151f;
  --clerk-button-font-color: #fff;
  --clerk-box-shadow: 2px 4px 12px #05f4b7;
}
```

![](<../.gitbook/assets/clerk-html-starter-1yvjin.stackblitz.io\_ (2).png>)
{% endtab %}

{% tab title="Professional" %}
```javascript
// Supplying a theme prop
const professionalThemeOptions = {
  general: {
    color: "#fb8122",
    backgroundColor: "#e1e2e2",
    fontColor: "#1d2228",
    labelFontWeight: "300",
    borderRadius: "0px",
  },
  buttons: {
    fontColor: "#fff",
  },
}



// Using CSS variables.
.cl-component {
  --clerk-border-radius: 0px;
  --clerk-primary: #fb8122;
  --clerk-font-color: #1d2228;
  --clerk-label-font-weight: 300;
  --clerk-background-color: #e1e2e2;
  --clerk-button-font-color: #fff;
}
```

![](<../.gitbook/assets/clerk-html-starter-1yvjin.stackblitz.io\_ (3).png>)
{% endtab %}

{% tab title="Hearty" %}
```javascript
// Supplying a theme prop
const heartyThemeOptions = {
  general: {
    color: "#ef0d50",
    backgroundColor: "#e1f2f7",
    fontColor: "#ef0d50",
    labelFontWeight: "bold",
    borderRadius: "12px",
    boxShadow: "10px 0px 12px 6px #ef0d50",
  },
  buttons: {
    fontColor: "#fff",
  },
}

// Using CSS variables.
.cl-component {
  --clerk-border-radius: 12px;
  --clerk-primary: #ef0d50;
  --clerk-font-color: #ef0d50;
  --clerk-label-font-weight: bold;
  --clerk-background-color: #e1f2f7;
  --clerk-button-font-color: #fff;
  --clerk-box-shadow: 10px 0px 12px 6px #ef0d50;
}
```

![](<../.gitbook/assets/clerk-html-starter-1yvjin.stackblitz.io\_ (5).png>)
{% endtab %}
{% endtabs %}

## Need more?

If you need more customization capabilities for your use-case, please reach out to any of our [community channels](https://clerk.dev/support).  We are constantly evaluating new customization options and would love to hear your feedback!

If you require complete customization, you can also implement [ClerkJS custom flows](../popular-guides/email-and-password.md#custom-flow), which leverage the same underlying APIs that are used to build our components.

