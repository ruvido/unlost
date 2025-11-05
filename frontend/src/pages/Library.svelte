<script>
  import { onMount } from 'svelte';
  import { pb } from '../lib/pb';
  import Header from '../lib/Header.svelte';

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
        <div class="photo-card">
          <img
            src="/library/{photo.path}"
            alt={photo.path}
            loading="lazy"
          />
        </div>
      {/each}
    </div>
  {/if}
</main>
