import { PageState } from '../page.js';
import { Query } from './query.js';
import { SessionList } from './sessionList.js';

export class Page {
    /**  @type {PageState} -- Modele */
    #pageState;

    /**  @type {HTMLBodyElement} */
    #el;

    /**  
     * @param {PageState} pageModel
     * */
    constructor(pageModel) {
        this.#pageState = pageModel;
        this.#init();
    }

    #init() {
        this.#el = document.querySelector("body");
        this.#el.appendChild(new Query().el);
        this.#el.appendChild(new SessionList(this.#pageState.sessions).el);
    }
}
