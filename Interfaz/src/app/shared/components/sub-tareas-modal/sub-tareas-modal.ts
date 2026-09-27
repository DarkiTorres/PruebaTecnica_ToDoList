import { Component, EventEmitter, inject, Input, Output } from '@angular/core';
import { Tarea } from '../../../models/tarea';
import { SubTarea } from '../../../models/subtarea';
import { SubTareaService } from '../../../core/services/subtarea.service';

@Component({
  selector: 'app-sub-tareas-modal',
  imports: [],
  templateUrl: './sub-tareas-modal.html',
  styleUrl: './sub-tareas-modal.css',
})
export class SubTareasModal {

  private readonly subTareaService = inject(SubTareaService);

  @Input() tarea: Tarea | null = null;
  @Input() subtareas: SubTarea[] | null = [];
  @Output() cerrar = new EventEmitter<void>();

  cerrarModal(): void {
    this.cerrar.emit();
  }
  
  cambiarEstadoSubTarea(subtarea: SubTarea, event: Event): void {
    const checkbox = event.target as HTMLInputElement;
    const estadoAnterior = subtarea.estaTerminada;
    const nuevoEstado = !estadoAnterior;

    subtarea.estaTerminada = nuevoEstado;

    this.subTareaService.completar(subtarea.id, nuevoEstado).subscribe(
      {
        next: () => {
          subtarea.estaTerminada = nuevoEstado;
        },
        error: error => {
          console.error('Error al cambiar estado de subtarea:',error);
          checkbox.checked = estadoAnterior;
          subtarea.estaTerminada = estadoAnterior;
        }
      }
    );
  }

}
