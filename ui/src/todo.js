import { API_BASE_URL } from './main.js';
import { Observable, PubSub } from './types.js';

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
            if (!response.ok) {
                throw new Error(`Response status: ${response.status}`);
            }

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
            if (!response.ok) {
                throw new Error(`Response status: ${response.status}`);
            }

            const result = await response.json();
            if (result.status == "success") {
                // Signal to Todos that delete happened so it can remove the todo from the list
                this.fireEvent(EVT_TODO_DELETE, null, this.id);

            } else console.log("Error during api call for complete");

        } catch (error) {
            console.error(error.message);
        }

    }
}

/**  
 * const static de Todos
 * */
const EVT_TODOS_CHANGE = "todos:changed";

export class Todos extends Observable {
    static get EVT_TODOS_CHANGE() { return EVT_TODOS_CHANGE };

    /**  
     * @type {Todo[]}
     * */
    #todos;

    get todos() { return this.#todos }

    /**  
     * @param {Todo[]} newTodos
     * */
    set todos(newTodos) {
        this.#todos = newTodos;
        this.fireEvent(EVT_TODOS_CHANGE, null, this.todos);
    };

    /**  
     * @param {PubSub} pcs
     * */
    constructor() {
        super(new PubSub());
        this.#todos = [];
    }

    async getAll() {
        const endpoint = "/todo";

        try {
            const response = await fetch(API_BASE_URL + endpoint);
            if (!response.ok) {
                throw new Error(`Response status: ${response.status}`);
            }

            const result = await response.json();
            this.todos = result.data.map(t => {
                const todo = new Todo(t.id, t.title, t.description, t.completed);
                todo.addPropertyChangeListener(Todo.EVT_TODO_DELETE, e => this.delete(e.newValue));
                return todo;
            });

        } catch (error) {
            console.error(error.message);
        }
    }

    /**  
     * @param {number} id -- Id of the todo to delete
     * */
    delete(id) {
        this.todos = this.todos.filter(t => t.id != id);
    }
}
