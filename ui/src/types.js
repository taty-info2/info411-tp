/**  @type {PropertyChangeEvent => void} PropertyChangeListener */


/**  Un event émit par un modèle */
export class PropertyChangeEvent {
    /**  @type {string} */
    name;

    /**  @type {*} */
    oldValue;

    /**  @type {*} */
    newValue;

    /**  
     * @param {string} name
     * @param {any} oldValue
     * @param {any} newValue
     * */
    constructor(name, oldValue, newValue) {
        this.name = name;
        this.oldValue = oldValue;
        this.newValue = newValue;
    }
}

/**  
 * Permet aux modèles d'émettre des event
 * Permet aux vues d'enregistrer un callback à executer à l'emission d'un event
 * */
export class PubSub {

    /**
     * Associe un nom d'Event à une liste de functions
     * @type Map<string, PropertyChangeListener[]>
     * */
    #listeners;

    constructor() {
        this.#listeners = new Map();
    }

    /**  
     * Appelle tous les listeners associés à un event
     * @param {PropertyChangeEvent} e
     * */
    notify(e) {
        this.#listeners.get(e.name)?.forEach(l => l(e));
    };

    /**  
     * Souscrit un listener à un nom d'event
     * @param {string} eventName
     * @param {PropertyChangeListener} pcl
     * */
    subscribe(eventName, pcl) {
        if (this.#listeners.has(eventName)) this.#listeners.get(eventName).push(pcl);
        else this.#listeners.set(eventName, [pcl]);
    };
}

/**
 * Base class pour les modèles
 * @abstract
 */
export class Observable {
    /**  @type PubSub */
    #pcs;

    constructor() {
        this.#pcs = new PubSub();
    }

    /**
     * Enregistre un callback appelé lors de l’émission d’un event
     * @param {string} eventName
     * @param {PropertyChangeListener} pcl
     */
    addPropertyChangeListener(eventName, pcl) { this.#pcs.subscribe(eventName, pcl) };

    /**  
     * Émet un event
     * @param {String} eventName
     * @param {any} oldValue
     * @param {any} newValue
     * */
    fireEvent(eventName, oldValue, newValue) {
        this.#pcs.notify(new PropertyChangeEvent(eventName, oldValue, newValue))
    };
}

