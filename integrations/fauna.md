# Fauna

### Overview

Fauna combines the flexibility of NoSQL systems with the relational querying capabilities of SQL databases into a modern, developer-friendly, transactional database with GraphQL support.

Although Fauna offers built-in identity, authentication, and password management functionality, it requires that you manage the user data yourself and does not provide features like social SSO, passwordless authentication, and multi-session management. Clerk provides these features and more without the hassle of managing your own user and identity service.

The Clerk integration with Fauna enables you to authenticate queries to your Fauna database using a JSON Web Token (JWT) created with our [JWT Templates](../popular-guides/jwt-templates.md) feature.

This guide will walk you through the necessary steps to integrate Clerk as the external authentication provider for Fauna.

### Getting started

The first step is to create a new Clerk application from your Clerk Dashboard if you haven’t done so already. You can choose whichever authentication strategy and social login providers you prefer. For more information, check out our [Set up your application](../popular-guides/setup-your-application.md) guide.

After your Clerk application has been created, scroll down to the **Connect your application** section of the dashboard and copy your **Frontend API key**.

![](../.gitbook/assets/01\_frontend-api-key.png)

The next step is to head over to the Fauna database that you want Clerk-authenticated users to access. Click on the **Security** link on the left-hand navigation and then the **Providers** tab.

![](../.gitbook/assets/02\_fauna-security.png)

Click the **New Access Provider** button.

Enter a name in the **Name** field that will identify this access provider. The recommendation is to use “Clerk”.

Paste the **Frontend API key** you copied from the Clerk Dashboard into the **Issuer** field and then prefix it with the `https://` protocol so it follows `https://<YOUR_FRONTEND_API>`

**Note**: Make sure you do not add a trailing slash `/` to the end of the Issuer URL or it will fail.

Assuming you didn’t use a custom signing key, set the **JWKS endpoint** field to the JSON Web Key Set (JWKS) URL Clerk automatically created with your Frontend API at `https://<YOUR_FRONTEND_API>/.well-known/jwks.json`

The last part is to select a role that will be assigned to authenticated users. The [Role](https://docs.fauna.com/fauna/current/security/roles) defines the access privileges and domain-specific security rules. You can specify multiple roles if necessary.

The completed provider form should resemble the following:

![](../.gitbook/assets/03\_provider-form.png)

Before clicking the Save button, copy the Audience URL then go ahead and save.

### JWT template

Navigate back to the Clerk Dashboard and click on the **JWT Templates** configuration section.

Click on the button to create a new **Blank** template.

![](../.gitbook/assets/04\_jwt-template.png)

Give your JWT template a short, but meaningful name (e.g. `fauna`).

Add in the `aud` claim set to the Audience URL you copied from Fauna. The URL format should be `https://db.fauna.com/db/<YOUR_FAUNA_DB_ID>`

![](../.gitbook/assets/05\_token-claims.png)

You can include additional claims if you’d like, but `aud` is the only required one. [Shortcodes](../popular-guides/jwt-templates.md#shortcodes) are available to make adding dynamic user values easy.

Click the “Apply changes” button and your Fauna JWT template will be saved.

### Client connection

To validate the token authentication with Fauna, you will need to have a client application that has a Clerk JavaScript library installed.

Clerk has an official [clerk-fauna-starter repo](https://github.com/clerkinc/clerk-fauna-starter) built with Next.js that you can use. If you already have an application with the Clerk library installed, you can [skip ahead](fauna.md#authenticating-fauna-queries).

Clone the starter repo from GitHub to your local machine, change into the project directory, and install the package dependencies with the following commands:

```shell
git clone https://github.com/clerkinc/clerk-fauna-starter
cd clerk-fauna-starter
npm install
```

A sample environment variable config file is available. Copy this file to `.env.local` with:

```shell
cp .env.local.sample .env.local
```

Set `CLERK_API_KEY` to your Clerk Backend API key and `NEXT_PUBLIC_CLERK_FRONTEND_API` to your Frontend API key. Both are available on the Clerk Dashboard or in the **API Keys** section.

After setting those environment variables, spin up the Next.js development server:

```bash
npm run dev
```

Browse to `http://localhost:3000>` and you should see the following screen:

![](../.gitbook/assets/06\_starter-screen.png)

If you haven’t already signed up for an account in the application do so now. Once you have signed in, you will be presented with the following:

![](<../.gitbook/assets/07\_app-screen (1).png>)

Click the first button to query Fauna. If the JWT authentication has been successful, you will see a user ID in the response output:

```json
{
  "id": "user_258ezWEGNPzIjNvpMpqRYbo9Eab"
}
```

If all went well, hooray! You can now make authenticated queries to your Fauna database.

If the current user ID is not returned, check the browser console for errors and reach out for [support](https://clerk.dev/support) if needed.

### Authenticating Fauna queries

The way to authenticate requests to Fauna with Clerk is to use a JWT. After adding the necessary claims in a JWT template, you can generate the token by calling the `getToken` method from the `useSession` hook provided by Clerk:

```jsx
import { useSession } from '@clerk/nextjs';
import faunadb from 'faunadb';
const q = faunadb.query;

const Example = () => {
  const { getToken } = useSession();
  const [userId, setUserId] = React.useState();
  const makeQuery = async () => {
    try {
      // TODO: Update with your JWT template name
      const secret = await getToken({ template: 'fauna' });
      const client = new faunadb.Client({ secret, keepAlive: false });
      const id = await client.query(q.CurrentIdentity());
      setUserId(id);
    } catch (e) {
      // Handle error
      setUserId(null);
    }
  };

  return (
    <>
      <button onClick={makeQuery}>Make authenticated query</button>
      <p>User ID: {userId}</p>
    </>
  );
};
```

**Note**: The `getToken({ template: <your-template-name> })` call is asynchronous and returns a Promise that needs to be resolved before accessing the token value. This token is short-lived for better security and should be called before every request to your Fauna database. The caching and refreshing of the token are handled automatically by Clerk.

### Next Steps

* Check out our official [Clerk and Fauna with Next.js starter](https://github.com/clerkinc/clerk-fauna-starter) repo
* Replace the [Fauna JavaScript driver](https://docs.fauna.com/fauna/current/drivers/javascript) with a GraphQL client like [Apollo](https://github.com/apollographql/apollo-client) or [graphql-request](https://github.com/prisma-labs/graphql-request)
* Try adding some [custom claims](../popular-guides/jwt-templates.md#template-basics) to the JWT token
* Get support or at least say hi in our [Discord channel](https://discord.com/invite/b5rXHjAg7A) 👋
