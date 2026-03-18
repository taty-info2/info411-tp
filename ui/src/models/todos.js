import { API_BASE_URL } from './app.js';
import { Todo } from './todo.js';
import { Observable, PubSub } from '../types.js';

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

    constructor() {
        super();
        this.#todos = [];
    }

    /**  
     * Create a Todo, add pcl and return it
     * @return {Todo}
     * */
    #createTodo({ id, title, description, completed }) {
        const todo = new Todo(id, title, description, completed);
        todo.addPropertyChangeListener(Todo.EVT_TODO_DELETE, e => this.delete(e.newValue));
        return todo;
    }

    async getAll() {
        const endpoint = "/todo";

        try {
            const response = await fetch(API_BASE_URL + endpoint);
            if (!response.ok) throw new Error(`Response status: ${response.status}`);

            const result = await response.json();
            this.todos = result.data.map(t => this.#createTodo(t));

        } catch (error) {
            console.error(error.message);
        }
    }

    /**  
     * React to EVT_TODO_DELETE
     * @param {number} id -- Id of the todo to delete
     * */
    delete(id) {
        this.todos = this.todos.filter(t => t.id != id);
    }

    /**  
     * @param {string} title
     * @param {string} description
     * */
    async add({ title, description }) {
        const endpoint = "/todo";

        try {
            const response = await fetch(API_BASE_URL + endpoint, {
                "method": "POST",
                "body": JSON.stringify({ "title": title, "description": description }),
                "headers": {
                    "Content-Type": "application/json"
                }
            });

            if (!response.ok) throw new Error(`Response status: ${response.status}`);

            const result = await response.json();
            if (result.status == "success") this.todos = [...this.todos, this.#createTodo(result.data)];
            else console.log("Error during api call for todos.add");

        } catch (err) {
            console.error(err.message);
        }
    }
}
