import { Todos } from '../models/todos.js';
import { TodoInList } from './todo.js';

// Template
const TODOS_TEMPLATE = document.querySelector("#todo-list");

export class TodoList {
    /**  @type {Todos} */
    #todos;

    /**  @type {HTMLOListElement} */
    #el;

    get el() { return this.#el };

    /**  
     * @param {Todos} todos -- Le modèle
     * */
    constructor(todos) {
        this.#todos = todos;
        this.#init();
        this.#todos.addPropertyChangeListener(Todos.EVT_TODOS_CHANGE, _ => this.#render());
        this.#render();
    }

    #init() {
        const fragment = TODOS_TEMPLATE.content.cloneNode(true);
        this.#el = fragment.firstElementChild;
    }

    #render() {
        this.#el.replaceChildren(
            ...this.#todos.todos.map(todo => {
                const li = document.createElement("li");
                li.appendChild(new TodoInList(todo).el);
                return li;
            })
        );
    }

}
