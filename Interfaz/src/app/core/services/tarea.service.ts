import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Tarea } from '../../models/tarea';
import { ApiService } from './api.service';
import { ActualizarTareaRequest } from '../../models/actualizar_tarea_request';

@Injectable({
  providedIn: 'root'
})
export class TareaService extends ApiService {
    
    obtenerTodas(): Observable<Tarea[]> {
        return this.http.get<Tarea[]>(
            `${this.apiUrl}/tareas`
        );
    }
    
    obtenerPorId(id: number): Observable<Tarea> {
        return this.http.get<Tarea>(
            `${this.apiUrl}/tareas/${id}`
        );
    }
    
    obtenerPorUsuario(usuarioId: number): Observable<Tarea[]> {
        return this.http.get<Tarea[]>(
            `${this.apiUrl}/usuarios/${usuarioId}/tareas`
        );
    }

    crear(tarea: CrearTareaRequest): Observable<Tarea> {
        return this.http.post<Tarea>(
            `${this.apiUrl}/tareas`,
            tarea
        );
    }

    actualizar(id: number, tarea: ActualizarTareaRequest): Observable<Tarea> {
        return this.http.put<Tarea>(
            `${this.apiUrl}/tareas/${id}`,
            tarea
        );
    }

    eliminar(id: number, modificadoPor: number): Observable<void> {
        return this.http.patch<void>(
            `${this.apiUrl}/tareas/${id}/eliminar`,
            {
                modificadoPor
            }
        );
    }

    eliminarFisicamente(id: number): Observable<void> {
        return this.http.delete<void>(
            `${this.apiUrl}/tareas/${id}`
        );
    }


    completar(id: number, estaTerminada: boolean): Observable<void> {
        return this.http.patch<void>(
            `${this.apiUrl}/tareas/${id}/completar`,
            {
                estaTerminada
            }
        );
    }
}