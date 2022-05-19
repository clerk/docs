# User

## getUser

{% tabs %}
{% tab title="Arguments" %}
| Name       | Description                                                  |
| ---------- | ------------------------------------------------------------ |
| **userId** | <p><em>string</em></p><p>The ID of the User to retrieve.</p> |
{% endtab %}

{% tab title="Return" %}
Returns a promise that resolves to a User object.
{% endtab %}

{% tab title="Example" %}
```javascript
import { setClerkApiKey, users } from '@clerk/clerk-sdk-node';

// This is unnecessary if CLERK_API_KEY is set as
// an environment variable
setClerkApiKey("api-key")

userId = "user-id";
const user = await users.getUser(userId);
```
{% endtab %}
{% endtabs %}

## updateUser

{% tabs %}
{% tab title="Arguments" %}
| Name           | Description                                                  |
| -------------- | ------------------------------------------------------------ |
| **userId**     | <p><em>string</em></p><p>The ID of the User to retrieve.</p> |
| **attributes** | _See below._                                                 |

Attributes that can be updated.

| Name                      | Description                                                                                                                                                                 |
| ------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **firstName**             | <p><em>string</em></p><p>The user's first name.</p>                                                                                                                         |
| **lastName**              | <p><em>string</em></p><p>The user's first name.</p>                                                                                                                         |
| **password**              | <p><em>string</em></p><p>The user's password.</p>                                                                                                                           |
| **primaryEmailAddressID** | <p><em>string</em></p><p>The ID of the email address to set as primary.</p>                                                                                                 |
| **primaryPhoneNumberID**  | <p><em>string</em></p><p>The ID of the phone number to set as primary.</p>                                                                                                  |
| **privateMetadata**       | <p><em>{[string]: any}</em></p><p>Secure metadata that can only be accessed with your API key or from your Clerk Dashboard. The data is kept hidden from your frontend.</p> |
| **publicMetadata**        | <p><em>{[string]: any}</em></p><p>Metadata that is accessible on your frontend with the <a href="broken-reference">useUser() hook</a> and on window.Clerk.user.</p>         |
{% endtab %}

{% tab title="Return" %}
Returns a promise that resolves to a User object.
{% endtab %}

{% tab title="Example" %}
```javascript
import { setClerkApiKey, users } from '@clerk/clerk-sdk-node';

// This is unnecessary if CLERK_API_KEY is set as
// an environment variable
setClerkApiKey("api-key")

userId = "user-id";
try {
  await users.updateUser(userId, {
    publicMetadata: {foo: "bar"}
  });
} catch (error) {
  // handle error
}
```
{% endtab %}
{% endtabs %}
