export interface CrearTareaForm {
    titulo: string;
    descripcion: string | null;
    prioridadId: number;
    asignadoA: number;
    fechaEntrega: string | null;
}