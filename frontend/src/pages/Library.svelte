<script>
  import { onMount } from 'svelte';
  import { pb } from '../lib/pb';
  import Header from '../lib/Header.svelte';
  import GLightbox from 'glightbox';
  import 'glightbox/dist/css/glightbox.min.css';

  let photos = [];
  let loading = true;
  let error = '';

  onMount(async () => {
    try {
      const records = await pb.collection('media').getList(1, 50, {
        sort: '-taken_at',
        filter: `owner = "${pb.authStore.model.id}"`,
      });
      photos = records.items;

      // Initialize lightbox after photos loaded
      setTimeout(() => {
        GLightbox({ touchNavigation: true, keyboardNavigation: true });
      }, 100);
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  });
</script>

<Header />

<main>
  {#if loading}
    <div class="loading">Loading photos...</div>
  {:else if error}
    <div class="error">{error}</div>
  {:else if photos.length === 0}
    <div class="loading">No photos yet</div>
  {:else}
    <div class="photo-grid">
      {#each photos as photo}
        <a href="/thumbs/view/{photo.path}" class="glightbox photo-card">
          <img
            src="/thumbs/small/{photo.path}"
            alt={photo.path}
            loading="lazy"
          />
        </a>
      {/each}
    </div>
  {/if}
</main>
