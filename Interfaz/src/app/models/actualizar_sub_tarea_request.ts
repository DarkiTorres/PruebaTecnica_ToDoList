export interface ActualizarSubTareaRequest {
  id: number;
  titulo: string;
  estaTerminada: boolean;
  estaEliminada: boolean;
  creadoPor: number;
  modificadoPor: number;
}