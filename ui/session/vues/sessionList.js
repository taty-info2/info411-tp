// La liste des sessions
import { Sessions } from '/models/sessions.js';
import { Session } from '/models/session.js';
import { SessionInList } from './sessionInList.js';

export class SessionList {
    /**  @type {Sessions} */
    #sessions;

    /**  @type {HTMLOListElement} */
    #el;

    get el() { return this.#el };

    /**  
     * @param {Sessions} sessions -- Modèle
     * */
    constructor(sessions) {
        this.#sessions = sessions;
        this.#sessions.addPropertyChangeListener(Sessions.EVT_SESSIONS_CHANGE, e => this.#render(e.newValue));

        this.#init();
        this.#render(this.#sessions.sessions);
    }

    #init() {
        this.#el = document.querySelector("#session-list").content.cloneNode(true).firstElementChild;
    }

    /**  
     * @param {Session[]} sessions
     * */
    #render(sessions) {
        this.#el.replaceChildren(
            ...sessions.map(session => new SessionInList(session).el)
        );
    }

}
