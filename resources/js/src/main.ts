import { mount } from 'svelte';
import '@fontsource-variable/rubik/wght.css';
import "./app.css"
import App from './App.svelte';
mount(App, { target: document.querySelector('#app')! });
