import { Sessions } from '/models/sessions.js';
import { Observable } from '/reactivity.js'

export class PageState extends Observable {
    /**  @type {Sessions} */
    #sessions;

    get sessions() { return this.#sessions }

    constructor() {
        super();
    }

    async setup() {
        this.#sessions = new Sessions();
        await this.#sessions.list();
    }
}
