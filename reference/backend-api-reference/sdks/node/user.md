# User

## getUser

{% tabs %}
{% tab title="Arguments" %}
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>userId</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The ID of the User to retrieve.</p>
      </td>
    </tr>
  </tbody>
</table>
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
<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>userId</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The ID of the User to retrieve.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>attributes</b>
      </td>
      <td style="text-align:left"><em>See below.</em>
      </td>
    </tr>
  </tbody>
</table>

Attributes that can be updated.

<table>
  <thead>
    <tr>
      <th style="text-align:left">Name</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>firstName</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s first name.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>lastName</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s first name.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>password</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The user&apos;s password.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>primaryEmailAddressID</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The ID of the email address to set as primary.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>primaryPhoneNumberID</b>
      </td>
      <td style="text-align:left">
        <p><em>string</em>
        </p>
        <p>The ID of the phone number to set as primary.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>privateMetadata</b>
      </td>
      <td style="text-align:left">
        <p><em>{[string]: any}</em>
        </p>
        <p>Secure metadata that can only be accessed with your API key or from your
          Clerk Dashboard. The data is kept hidden from your frontend.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>publicMetadata</b>
      </td>
      <td style="text-align:left">
        <p><em>{[string]: any}</em>
        </p>
        <p>Metadata that is accessible on your frontend with the <a href="../../../clerk-react/useuser-hook.md">useUser() hook</a> and
          on window.Clerk.user.</p>
      </td>
    </tr>
  </tbody>
</table>
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

