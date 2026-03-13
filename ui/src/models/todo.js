import { API_BASE_URL } from './app.js';
import { Observable, PubSub } from '../types.js';

/**  
 * const static de Todo
 * */
const EVT_TODO_COMPLETE = "todo:completed",
    EVT_TODO_DELETE = "todo:deleted";

export class Todo extends Observable {
    static get EVT_TODO_COMPLETE() { return EVT_TODO_COMPLETE };
    static get EVT_TODO_DELETE() { return EVT_TODO_DELETE };

    /**  @type {number} */
    id;

    /**  @type {string} */
    title;

    /**  @type {string} */
    description;

    /**  @type {boolean} */
    #completed;

    get completed() { return this.#completed };

    /**  
     * @param {Boolean} newCompleted
     * */
    set completed(newCompleted) {
        const oldValue = this.#completed;
        this.#completed = newCompleted;
        this.fireEvent(EVT_TODO_COMPLETE, oldValue, newCompleted);
    };

    /**  
     * @param {PubSub} pcs
     * @param {number} id
     * @param {string} title
     * @param {string} description
     * @param {boolean} completed
     * @param {()=>void} onDelete
     * */
    constructor(id, title, description, completed) {
        super(new PubSub());
        this.id = id;
        this.title = title;
        this.description = description;
        this.#completed = completed;
    };

    async complete() {
        const endpoint = `/todo/${this.id}`;

        try {
            const response = await fetch(API_BASE_URL + endpoint, {
                "method": "PATCH",
                "body": JSON.stringify({ "completed": !this.completed }),
                "headers": {
                    "Content-Type": "application/json"
                }
            });
            if (!response.ok) throw new Error(`Response status: ${response.status}`);

            const result = await response.json();
            if (result.status == "success") this.completed = !this.completed;
            else console.log("Error during api call for complete");

        } catch (err) {
            console.error(err.message);
        }
    }

    // Signal to Todos that deletion should happen
    async delete() {
        const endpoint = `/todo/${this.id}`;

        try {
            const response = await fetch(API_BASE_URL + endpoint, {
                "method": "DELETE"
            });
            if (!response.ok) throw new Error(`Response status: ${response.status}`);

            const result = await response.json();
            if (result.status == "success") {
                // Signal to Todos that delete happened so it can remove the todo from the list
                this.fireEvent(EVT_TODO_DELETE, null, this.id);

            } else console.log("Error during api call for complete");

        } catch (err) {
            console.error(err.message);
        }

    }
}

