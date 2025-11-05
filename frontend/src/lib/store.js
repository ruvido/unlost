import { writable } from 'svelte/store';
import { pb } from './pb';

export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange((token, model) => {
  currentUser.set(model);
});
