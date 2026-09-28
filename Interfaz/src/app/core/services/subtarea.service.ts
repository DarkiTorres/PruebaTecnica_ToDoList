import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { SubTarea } from '../../models/subtarea';
import { ApiService } from './api.service';

import { CrearSubTareaRequest } from '../../models/crear_sub_tarea_request';

@Injectable({
  providedIn: 'root'
})
export class SubTareaService extends ApiService {
    obtenerPorId(id: number): Observable<SubTarea> {
        return this.http.get<SubTarea>(
            `${this.apiUrl}/subtareas/${id}`
        )
    }

    obtenerPorTarea(tareaId: number): Observable<SubTarea[]> {
        return this.http.get<SubTarea[]>(
            `${this.apiUrl}/tareas/${tareaId}/subtareas`
        );
    }

    crear(request: CrearSubTareaRequest): Observable<SubTarea> {
        return this.http.post<SubTarea>(
            `${this.apiUrl}/tareas/${request.tareaId}/subtareas`,
            request
        );
    }

    actualizar(id: number, subTarea: SubTarea): Observable<void> {
        return this.http.put<void>(
            `${this.apiUrl}/subtareas/${id}`,
            subTarea
        );
    }

    eliminar(id: number, modificadoPor: number): Observable<void> {
        return this.http.patch<void>(
            `${this.apiUrl}/subtareas/${id}/eliminar`,
            { 
                modificadoPor 
            }
        );
    }
    completar(id: number, estaTerminada: boolean): Observable<void> {
        return this.http.patch<void>(
            `${this.apiUrl}/subtareas/${id}/completar`,
            {
                estaTerminada
            }
        );
    }
}