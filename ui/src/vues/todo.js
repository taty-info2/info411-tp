import { Todo } from '../models/todo.js';

// template
const TODO_IN_LIST_TEMPLATE = document.querySelector("#todo-in-list");

export class TodoInList {
    /**  @type {Todo} */
    #todo;

    /**  @type {HTMLDivElement} */
    #el;

    get el() { return this.#el };

    /**  @type {HTMLHeadingElement} */
    #title;

    /**  @type {HTMLParagraphElement} */
    #description;

    /**  @type {HTMLDivElement} */
    #completeCb;

    /**  @type {HTMLButtonElement} */
    #delBtn;


    /**  
     * @param {Todo} todo -- Modèle
     * */
    constructor(todo) {
        this.#todo = todo;
        this.#init();
        this.#todo.addPropertyChangeListener(Todo.EVT_TODO_COMPLETE, e => this.#completeCb.checked = e.newValue)
        this.#render();

    }

    #init() {
        const fragment = TODO_IN_LIST_TEMPLATE.content.cloneNode(true);
        this.#el = fragment.firstElementChild;
        this.#title = this.#el.querySelector("[data-title]");
        this.#description = this.#el.querySelector("[data-description]");

        this.#completeCb = this.#el.querySelector("[data-completed]");
        this.#completeCb.addEventListener("click", e => {
            e.preventDefault();
            this.#todo.complete();
        });

        this.#delBtn = this.#el.querySelector("[data-delete]");
        this.#delBtn.addEventListener("click", _ => this.#todo.delete());
    }

    #render() {
        this.#title.innerText = this.#todo.title;
        this.#description.innerText = this.#todo.description;
        this.#completeCb.checked = this.#todo.completed;
    }
}
