import { SubTarea } from "./subtarea";

export interface Tarea {
    id: number;
    asignadoAId?: number;
    asignadoA?: string;
    titulo: string;
    descripcion?: string;
    prioridadId: number;
    fechaEntrega?: string;
    estaTerminada: boolean;
    estaEliminada: boolean;
    creadoEl: string;
    creadoPor: number;
    modificadoEl?: string;
    modificadoPor?: number;
    subTareas?: SubTarea[];
}