import { Todos, Todo } from './todo.js';

class TodoInList {
    /**  @type {Todo} */
    #todo;

    /**  @type {HTMLDivElement} */
    #el;

    get el() { return this.#el };

    /**  @type {HTMLDivElement} */
    #completeCb;

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
        this.#el = document.createElement("div");

        const completeCb = document.createElement("input");
        completeCb.type = "checkbox";
        this.#completeCb = completeCb;
    }

    #render() {
        const title = document.createElement("h1");
        title.innerText = this.#todo.title;

        const description = document.createElement("p");
        description.innerText = this.#todo.description;

        this.#completeCb.checked = this.#todo.completed;
        this.#completeCb.addEventListener("click", e => {
            e.preventDefault();
            this.#todo.complete();
        });

        const delBtn = document.createElement("button");
        delBtn.innerText = "X";
        delBtn.addEventListener("click", _ => this.#todo.delete());

        this.#el.appendChild(title);
        this.#el.appendChild(description);
        this.#el.appendChild(this.#completeCb);
        this.#el.appendChild(delBtn);
    }
}

export class TodosList {
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
        this.#todos.addPropertyChangeListener(Todos.EVT_TODOS_CHANGE, _ => this.render());
        this.render();
    }

    #init() {
        this.#el = document.createElement("ol");
    }

    render() {
        this.#el.replaceChildren(
            ...this.#todos.todos.map(todo => {
                const li = document.createElement("li");
                li.appendChild(new TodoInList(todo).el);
                return li;
            })
        );
    }

}
