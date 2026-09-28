export interface ActualizarTareaRequest {
  titulo: string;
  descripcion: string | null;
  prioridadId: number;
  fechaEntrega: string | null;
  modificadoPor: number;
  asignadoA: number;
  subTareas: ActualizarTareaRequest[];
}