<script>
  import { currentUser } from './store';
  import { pb } from './pb';

  let menuOpen = false;

  function toggleMenu() {
    menuOpen = !menuOpen;
  }

  function logout() {
    pb.authStore.clear();
    menuOpen = false;
    window.location.hash = '#/login';
  }
</script>

<header class="header">
  <div class="header-content">
    <a href="#/library" class="logo">Unlost</a>

    {#if $currentUser}
      <button class="hamburger" on:click={toggleMenu} aria-label="Menu">
        <span></span>
        <span></span>
        <span></span>
      </button>
    {/if}
  </div>
</header>

{#if $currentUser && menuOpen}
  <div
    class="menu-overlay"
    role="button"
    tabindex="0"
    on:click={toggleMenu}
    on:keydown={(e) => e.key === 'Escape' && toggleMenu()}
  ></div>
  <nav class="menu" class:menu-open={menuOpen}>
    <div class="menu-header">
      <span class="menu-user">{$currentUser.username}</span>
      <button class="menu-close" on:click={toggleMenu}>×</button>
    </div>
    <ul class="menu-list">
      <li><a href="#/library" on:click={toggleMenu}>Library</a></li>
      <li><button on:click={logout}>Logout</button></li>
    </ul>
  </nav>
{/if}
