// Un session dans une liste de sessions
import { Session } from '/models/session.js';

export class SessionInList {
    /**  @type {Session} */
    #session;

    /**  @type {HTMLOListElement} */
    #el;

    get el() { return this.#el };

    /**  
     * @param {Session} session
     * */
    constructor(session) {
        this.#session = session;
        this.#init();
        this.#render();
    }

    #init() {
        this.#el = document.querySelector("#session-in-list").content.cloneNode(true).firstElementChild;
    }

    #render() {
        const places_restantes = this.#session.nb_places - this.#session.nb_inscrits;
        this.#el.querySelector("[data-titre]").innerText =
            `${this.#session.activiteNom} (${places_restantes} pl. restantes)`;
        this.#el.querySelector("[data-date]").innerText = this.#session.date;
        this.#el.querySelector("[data-heure]").innerText = `${this.#session.heure_debut}-${this.#session.heure_fin}`;
        this.#el.querySelector("[data-adresse]").innerText = this.#session.adresse;
        this.#el.querySelector("[data-image]").src = this.#session.activiteImgUrl;
    }
}
