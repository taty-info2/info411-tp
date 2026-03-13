import { AppState } from '../models/app.js';
import { TodoList } from './todos.js';
import { TodoForm } from './form-todo.js';

export class App {
    /**  @type {AppState} */
    #app;

    /**  @type {HTMLBodyElement} */
    #el;

    /**  
     * @param {AppState} app
     * */
    constructor(app) {
        this.#app = app;
        this.#init();
    }

    #init() {
        this.#el = document.querySelector("body");

        const addTodoBtn = document.querySelector("#formToggleBtn");
        addTodoBtn.addEventListener("click", _ => this.#app.isDialFormTodoOpen = true);
        this.#el.appendChild(addTodoBtn);

        const todoList = new TodoList(this.#app.todos);
        this.#el.appendChild(todoList.el);

        const todoForm = new TodoForm(this.#app);
        this.#el.appendChild(todoForm.el);
    }
}
