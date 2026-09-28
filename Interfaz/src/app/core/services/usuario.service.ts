import { Injectable } from "@angular/core";
import { Observable } from "rxjs";

import { Usuario } from "../../models/usuario";
import { ApiService } from "./api.service";

@Injectable({
    providedIn: 'root'
})
export class UsuarioService extends ApiService{
    obtenerTodos(): Observable<Usuario[]>{
        return this.http.get<Usuario[]>(
            `${this.apiUrl}/usuarios`
        );
    }

    obtenerPorId(id: number): Observable<Usuario> {
        return this.http.get<Usuario>(
            `${this.apiUrl}/usuarios/${id}`
        )
    }
}