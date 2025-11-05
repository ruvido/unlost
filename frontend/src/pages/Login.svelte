<script>
  import { pb } from '../lib/pb';

  let email = '';
  let password = '';
  let error = '';
  let loading = false;

  async function handleLogin() {
    error = '';
    loading = true;
    try {
      await pb.collection('users').authWithPassword(email, password);
      window.location.hash = '#/library';
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  }
</script>

<main>
  <h1 style="text-align: center; margin-bottom: 2rem;">Login</h1>

  <form on:submit|preventDefault={handleLogin}>
  <label>
    Email
    <input
      type="email"
      name="email"
      autocomplete="email"
      bind:value={email}
      required
    />
  </label>

  <label>
    Password
    <input
      type="password"
      name="password"
      autocomplete="current-password"
      bind:value={password}
      required
    />
  </label>

  {#if error}
    <div class="error">{error}</div>
  {/if}

  <button type="submit" disabled={loading}>
    {loading ? 'Loading...' : 'Login'}
  </button>

  <p style="margin-top: 1rem;">
    No account? <a href="#/signup">Sign up</a>
  </p>
</form>
</main>
