// The app state
import { Todos } from './todos.js';
import { Observable, PubSub } from '../types.js';

export const API_BASE_URL = "http://localhost:3001";

/**  
 * const static de App
 * */
const EVT_APP_DIAL_FORM_TODO = "app:dial-form-todo";

export class AppState extends Observable {
    static get EVT_APP_DIAL_FORM_TODO() { return EVT_APP_DIAL_FORM_TODO };

    /**  @type {boolean} -- Boolean representing the state of the dialog used to create a new Todo */
    #isDialFormTodoOpen;

    get isDialFormTodoOpen() { return this.#isDialFormTodoOpen };

    set isDialFormTodoOpen(newState) {
        const oldValue = this.#isDialFormTodoOpen;
        this.#isDialFormTodoOpen = newState;
        this.fireEvent(EVT_APP_DIAL_FORM_TODO, oldValue, this.#isDialFormTodoOpen);
    };

    /**  @type {Todos} */
    #todos;

    get todos() { return this.#todos };

    constructor() {
        super(new PubSub());
    }

    // To call in main
    // TODO: find a better way
    async init() {
        const todos = new Todos();
        await todos.getAll();
        this.#todos = todos;
    }
}
