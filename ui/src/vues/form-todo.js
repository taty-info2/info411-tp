import { AppState } from '../models/app.js';

// template
const TODO_FORM_TEMPLATE = document.querySelector("#form-todo");

export class TodoForm {
    /**  @type {AppState} */
    #appState;

    /**  @type {HTMLDialogElement} */
    #el;

    get el() { return this.#el };

    /**  @type {HTMLFormElement} */
    #form;

    /**  
     * @param {AppState} appState
     * */
    constructor(appState) {
        this.#appState = appState;
        this.#init();
        this.#render();
    }

    #init() {
        const fragment = TODO_FORM_TEMPLATE.content.cloneNode(true);
        this.#el = fragment.firstElementChild;

        this.#form = this.#el.querySelector("form");
        this.#form.addEventListener("submit", e => {
            e.preventDefault();

            const formData = new FormData(this.#form);
            const values = Object.fromEntries(formData.entries());

            this.#appState.todos.add(values);
        });

        this.#appState.addPropertyChangeListener(AppState.EVT_APP_DIAL_FORM_TODO, e => e ? this.#el.showModal() : this.#el.close());
    }

    #render() {
        this.el.open = false;
    }
}
