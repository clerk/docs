# Supabase

### Overview

Supabase is the open source alternative to Firebase. It offers cloud-based access to a full PostgreSQL database as well as realtime subscriptions and storage solutions. Although Supabase offers a built-in authentication and user management service with Supabase Auth, it requires the user data to be stored in your own database. With Clerk, you can get the benefits of the backend services they provide without the hassle of managing your own user database.

The Clerk integration with Supabase enables you to authenticate requests to your Supabase database using a JSON Web Token (JWT) created with our [JWT Templates](../popular-guides/jwt-templates.md) feature.

This guide will walk you through the necessary steps to integrate Clerk as the external authentication provider for Supabase.

### Get started

The first step is to create a new Clerk application from your Clerk Dashboard if you haven’t done so already. You can choose whichever authentication strategy and social login providers you prefer. For more information, check out our [Set up your application](../popular-guides/setup-your-application.md) guide.

After your Clerk application has been created, use the lefthand menu to navigate to the **JWT Templates** page.

Click on the button to create a new template based on Supabase.

![](../.gitbook/assets/01\_supabase-template.png)

This will pre-populate the default claims required by Supabase. You can include additional claims or modify them as necessary. [Shortcodes](../popular-guides/jwt-templates.md#shortcodes) are also available to make adding dynamic user values easy.

ℹ️ Note the name of the JWT template (which you can change) because this will be needed later.

![](../.gitbook/assets/02\_jwt-claims.png)

Supabase requires that JWTs be signed with the H256 signing algorithm and use their signing key. You can locate the JWT secret key in your Supabase project under **Settings** > **API** in the Config section.

![](../.gitbook/assets/03\_jwt-secret.png)

Reveal the JWT secret to copy it and then paste it in the Signing key field in the Clerk JWT template.

![](../.gitbook/assets/04\_signing-key.png)

After the key is added, you can click the Apply Changes button to save your template.

### Configure your client

To configure your client, you need to set some local environment variables. Assuming a React application, set the following:

```bash
REACT_APP_CLERK_FRONTEND_API=your-frontend-api
REACT_APP_SUPABASE_URL=your-supabase-url
REACT_APP_SUPABASE_KEY=your-supabase-anon-key
```

{% hint style="info" %}
**Note**: If you’re using Next.js, replace the `REACT_APP` prefix with `NEXT_PUBLIC`
{% endhint %}

Your Clerk Frontend API can be found on the Clerk Dashboard home in the **Connect your application** section.

![](../.gitbook/assets/05\_clerk-frontend-api.png)

To get the ones needed for Supabase, navigate to the same **Settings** > **API** page as before and locate the anon public key and URL.



![](../.gitbook/assets/06\_supabase-keys.png)

{% hint style="info" %}
**Note**: Supabase recommends enabling [Row Level Security](https://supabase.com/docs/guides/auth/row-level-security) (RLS) for your database tables and configuring access policies as needed.
{% endhint %}

After setting those three environment variables, you should be able to start up your application development server.

Install the JavaScript client for Supabase with:

```bash
npm install @supabase/supabase-js
```

You can then initialize the Supabase client by passing it the environment variables:

```jsx
import { createClient } from '@supabase/supabase-js';

const supabaseUrl = process.env.REACT_APP_SUPABASE_URL;
const supabaseKey = process.env.REACT_APP_SUPABASE_KEY;

export const supabase = createClient(supabaseUrl, supabaseKey);
export default supabase;
```

In order to access the custom JWT, you can use the `getToken` function returned by the Clerk `useAuth` hook and pass it the name of your template (hopefully you remembered from earlier).

{% hint style="info" %}
**Note**: The `getToken({ template: <your-template-name> })` call is asynchronous and returns a Promise that needs to be resolved before accessing the token value. This token is short-lived for better security and should be called before every request to your Supabase backend. The caching and refreshing of the token is handled automatically by Clerk.
{% endhint %}

Call `supabase.auth.setAuth(token)` to override the JWT on the current client. The JWT will then be sent to Supabase with all subsequent network requests.

```jsx
import { useAuth } from '@clerk/clerk-react';
import supabase from './lib/supabaseClient';

function App() {
  const { getToken } = useAuth();

  const fetchData = async () => {
    // TODO #1: Replace with your JWT template name
    const token = await getToken({ template: 'supabase' });

    supabase.auth.setAuth(token);

    // TODO #2: Replace with your database table name
    const { data, error } = await supabase.from('your_table').select();

    // TODO #3: Handle the response
  };

  return (
    <div className="app">
      <button onClick={fetchData}>Fetch data</button>
    </div>
  );
}
```

### Access user ID in RLS policies

It is common practice to need access to the user identifier on the database level, especially when working with RLS policies in Postgres. Although Supabase provides a special function `auth.uid()` to extract the user ID from the JWT, this [does not currently work](https://github.com/supabase/supabase/discussions/4954) with Clerk. The workaround is to write a custom SQL function to read the `sub` property from the JWT claims.

In the **SQL Editor** section of the Supabase dashboard, click New Query and enter the following:

```sql
create or replace function requesting_user_id()
returns text 
language sql stable
as $$
  select nullif(current_setting('request.jwt.claims', true)::json->>'sub', '')::text;
$$;
```

This will create a `requesting_user_id()` function that can be used within an RLS policy.

For example, this policy would check that the user making the request is authenticated and matches the `user_id` column of a todos table.

```sql
CREATE POLICY "Authenticated users can update their own todos" 
	ON public.todos FOR UPDATE USING (
		auth.role() = 'authenticated'::text
	) WITH CHECK (
		requesting_user_id() = user_id
	);
```

### Next steps

* Get started with our official [clerk-supabase-starter](https://github.com/clerkinc/clerk-supabase-starter) repo
* Check out our [Next.js + Supabase + Clerk tutorial](https://clerk.dev/blog/nextjs-supabase-todos-with-multifactor-authentication)
* Try adding some [custom claims](../popular-guides/jwt-templates.md#template-basics) to the JWT template in `app_metadata` or `user_metadata`
* Get support or at least say hi in our [Discord channel](https://discord.com/invite/b5rXHjAg7A) 👋
