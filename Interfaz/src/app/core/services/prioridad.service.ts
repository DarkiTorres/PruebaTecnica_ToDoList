import { Injectable } from "@angular/core";
import { Observable } from "rxjs";

import { Prioridad } from "../../models/prioridad";
import { ApiService } from "./api.service";

@Injectable({
    providedIn: 'root'
})
export class PrioridadService extends ApiService {
    obtenerTodas(): Observable<Prioridad[]> {
        return this.http.get<Prioridad[]>(
            `${this.apiUrl}/prioridades`
        );
    }
}