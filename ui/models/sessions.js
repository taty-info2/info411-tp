import { Session } from './session.js';
import { Observable } from '/reactivity.js';

/**  
 * const static de sessions
 * */
const EVT_SESSIONS_CHANGE = "sessions:changed";

export class Sessions extends Observable {
    static get EVT_SESSIONS_CHANGE() { return EVT_SESSIONS_CHANGE };

    /**  
     * @type {Session[]}
     * */
    #sessions;

    get sessions() { return this.#sessions }

    /**  
     * @param {Session[]} newSessions
     * */
    set sessions(newSessions) {
        this.#sessions = newSessions;
        this.fireEvent(EVT_SESSIONS_CHANGE, null, this.sessions);
    };

    constructor() {
        super();
        this.#sessions = [];
    }

    async list() {
        const sessions = [
            new Session(1, "Ski", "/assets/ski.png", "18-02-2026", "3 rue Pignot 93200 Paris", "15:00", "16:30", 30, 15),
            new Session(2, "Natation", "/assets/natation.png", "14-03-2026", "14 rue Laplace 93200 Paris", "12:00", "14:00", 22, 22)
        ];
        this.#sessions = sessions;
    }
}
