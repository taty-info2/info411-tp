import { Observable } from '/reactivity.js';

export class Session extends Observable {
    /**  @type {number} */
    id;

    /**  @type {string} */
    activiteNom;

    /**  @type {string} */
    activiteImgUrl;

    /**  @type {string} */
    date;

    /**  @type {string} */
    adresse;

    /**  @type {string} */
    heure_debut;

    /**  @type {string} */
    heure_fin;

    /**  @type {number} */
    nb_places;

    /**  @type {number} */
    nb_inscrits;

    /**  @type {boolean} */
    est_complete;

    /**  
     * @param {number} id
     * @param {string} activiteNom
     * @param {string} activiteImgUrl
     * @param {string} date
     * @param {string} adresse
     * @param {string} heure_debut
     * @param {string} heure_fin
     * @param {number} nb_places
     * @param {number} nb_inscrits
     * */
    constructor(id, activiteNom, activiteImgUrl, date, adresse, heure_debut, heure_fin, nb_places, nb_inscrits) {
        super();
        this.id = id;
        this.activiteNom = activiteNom;
        this.activiteImgUrl = activiteImgUrl;
        this.date = date;
        this.adresse = adresse;
        this.heure_debut = heure_debut;
        this.heure_fin = heure_fin;
        this.nb_places = nb_places;
        this.nb_inscrits = nb_inscrits;
        this.est_complete = nb_places == nb_inscrits;
    };

    /**  
     * Inscrit l'utilisateur à la session
     * */
    async inscrire() { }

    /**  
     * Inscrit l'utilisateur en file d'attente
     * */
    async inscrire_attente() { }

    /**  
     * Desinscrit l'utilisateur de la session
     * */
    async desinscrit() { }
}
