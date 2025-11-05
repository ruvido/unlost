<script>
  import { pb } from '../lib/pb';

  let email = '';
  let username = '';
  let password = '';
  let passwordConfirm = '';
  let error = '';
  let loading = false;

  async function handleSignup() {
    error = '';
    loading = true;
    try {
      await pb.collection('users').create({
        email,
        username,
        password,
        passwordConfirm
      });
      // Auto-login after signup
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
  <h1 style="text-align: center; margin-bottom: 2rem;">Sign Up</h1>

  <form on:submit|preventDefault={handleSignup}>
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
    Username
    <input
      type="text"
      name="username"
      autocomplete="username"
      bind:value={username}
      required
    />
  </label>

  <label>
    Password
    <input
      type="password"
      name="password"
      autocomplete="new-password"
      bind:value={password}
      required
    />
  </label>

  <label>
    Confirm Password
    <input
      type="password"
      name="password-confirm"
      autocomplete="new-password"
      bind:value={passwordConfirm}
      required
    />
  </label>

  {#if error}
    <div class="error">{error}</div>
  {/if}

  <button type="submit" disabled={loading}>
    {loading ? 'Creating account...' : 'Sign Up'}
  </button>

  <p style="margin-top: 1rem;">
    Have an account? <a href="#/login">Login</a>
  </p>
</form>
</main>
