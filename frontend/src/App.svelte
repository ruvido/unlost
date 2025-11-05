<script>
  import { onMount } from 'svelte';
  import { currentUser } from './lib/store';
  import { pb } from './lib/pb';
  import Login from './pages/Login.svelte';
  import Signup from './pages/Signup.svelte';
  import Library from './pages/Library.svelte';

  let currentRoute = 'login';

  function updateRoute() {
    const hash = window.location.hash.slice(1) || '/login';
    const path = hash.split('/')[1] || 'login';
    currentRoute = path;

    // Redirect to login if not authenticated and trying to access library
    if (path === 'library' && !pb.authStore.isValid) {
      window.location.hash = '#/login';
      currentRoute = 'login';
    }
  }

  onMount(() => {
    updateRoute();
    window.addEventListener('hashchange', updateRoute);
    return () => window.removeEventListener('hashchange', updateRoute);
  });
</script>

{#if currentRoute === 'login'}
  <Login />
{:else if currentRoute === 'signup'}
  <Signup />
{:else if currentRoute === 'library'}
  <Library />
{/if}
