import { App } from './vues/app.js';
import { AppState } from './models/app.js';

const appState = new AppState();
await appState.init(); // load the todos

new App(appState);
