interface CrearTareaRequest {
    titulo: string;
    descripcion: string | null;
    prioridadId: number;
    fechaEntrega: string | null;
    creadoPor: number;
    asignadoA: number;
}