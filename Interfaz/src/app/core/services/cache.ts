import { Injectable } from '@angular/core';
import { Tarea } from '../../models/tarea';
import { Prioridad } from '../../models/prioridad';
import { Usuario } from '../../models/usuario';

@Injectable({
  providedIn: 'root',
})
export class Cache {
  private obtenerClaveTareas(usuarioId: number, rolId: number) : string {
    return `tarea-${usuarioId}-${rolId}`;
  }

  guardarTareas(usuarioId: number, rolId: number, tareas: Tarea[]): void {
    localStorage.setItem(this.obtenerClaveTareas(usuarioId, rolId), JSON.stringify(tareas));
  }

  obtenerTareas(
    usuarioId: number,
    rolId: number
  ): Tarea[] {

    const datos = localStorage.getItem(
      this.obtenerClaveTareas(usuarioId, rolId)
    );

    if (!datos) {
      return [];
    }

    try {
      return JSON.parse(datos);
    } catch {
      return [];
    }
  }

  guardarPrioridades(prioridades: Prioridad[]): void {
    localStorage.setItem(
      'prioridades-cache',
      JSON.stringify(prioridades)
    );
  }

  guardarUsuarios(usuarios: Usuario[]): void {
    localStorage.setItem(
      'usuarios-cache',
      JSON.stringify(usuarios)
    );
  }

  obtenerUsuarios(): Usuario[] {
    const datos = localStorage.getItem(
      'usuarios-cache'
    );

    if (!datos) {
      return [];
    }

    try {
      return JSON.parse(datos);
    } catch {
      return [];
    }
  }

  guardarUsuarioActual(usuario: Usuario): void {
    localStorage.setItem(
      'usuario-actual-cache',
      JSON.stringify(usuario)
    );
  }

  obtenerUsuarioActual(): Usuario | null {

    const datos = localStorage.getItem(
      'usuario-actual-cache'
    );

    if (!datos) {
      return null;
    }

    try {
      return JSON.parse(datos);
    } catch {
      return null;
    }
  }

  obtenerPrioridades(): Prioridad[] {

    const datos = localStorage.getItem(
      'prioridades-cache'
    );

    if (!datos) {
      return [];
    }

    try {
      return JSON.parse(datos);
    } catch {
      return [];
    }
  }
}
