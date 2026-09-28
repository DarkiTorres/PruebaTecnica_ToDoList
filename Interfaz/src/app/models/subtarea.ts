export interface SubTarea {
    id: number;
    tareaId: number;
    titulo: string;
    estaTerminada: boolean;
    estaEliminada: boolean;
    creadoEl: string;
    creadoPor: number;
    modificadoEl?: string;
    modificadoPor?: number;
}