export class Query {
    /**  @type {HTMLOListElement} */
    #el;

    get el() { return this.#el };

    constructor() {
        this.#init();
    }

    #init() {
        this.#el = document.querySelector("#queries").content.cloneNode(true).firstElementChild;

        // TODO:  Réintroduire le modèle query_session.js, submit le form depuis app qui a un reference vers Query
        this.#el.addEventListener("change", e => {
            e.preventDefault();
            const params = new URLSearchParams(new FormData(this.#el));
            history.pushState(null, "", "?" + params)
        });

    }
}
