import { writable } from 'svelte/store';

const accessTokenStore = writable('');

function setAccessToken(token: string) {
  accessTokenStore.set(token);
  localStorage.setItem('access_token', token);
}

function clearAccessToken() {
  localStorage.removeItem('access_token');
}

export { accessTokenStore, setAccessToken, clearAccessToken};