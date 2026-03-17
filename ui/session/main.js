import { PageState } from "./page.js";
import { Page } from "./vues/page.js";

// Model
const pageState = new PageState();
await pageState.setup();

// Vue
const pageVue = new Page(pageState);
