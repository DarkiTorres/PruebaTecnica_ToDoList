import { Component, EventEmitter, inject, Input, Output, signal } from '@angular/core';
import { Tarea } from '../../../models/tarea';
import { SubTarea } from '../../../models/subtarea';
import { SubTareaService } from '../../../core/services/subtarea.service';
import { FormsModule } from '@angular/forms';
import { Usuario } from '../../../models/usuario';
import { CrearSubTareaRequest } from '../../../models/crear_sub_tarea_request';

@Component({
  selector: 'app-sub-tareas-modal',
  imports: [FormsModule],
  templateUrl: './sub-tareas-modal.html',
  styleUrl: './sub-tareas-modal.css',
})
export class SubTareasModal {

  private readonly subTareaService = inject(SubTareaService);

  @Input() tarea: Tarea | null = null;
  @Input() subtareas: SubTarea[] | null = [];
  @Input() usuarioActual: Usuario | null = null;

  @Output() cerrar = new EventEmitter<void>();

  mostrarNuevaSubTarea = signal(false);
  tituloNuevaSubTarea = '';

  cerrarModal(): void {
    this.cerrar.emit();
  }

  abrirNuevaSubTarea(): void {
    this.tituloNuevaSubTarea = '';
    this.mostrarNuevaSubTarea.set(true);
  }

  cancelarNuevaSubTarea(): void {
    this.mostrarNuevaSubTarea.set(false);
    this.tituloNuevaSubTarea = '';
  }

  crearSubTarea(): void {
    if (!this.tarea || !this.usuarioActual) {
      return;
    }

    const titulo = this.tituloNuevaSubTarea.trim();

    if (!titulo) {
      return;
    }

    const request: CrearSubTareaRequest = {
      titulo,
      tareaId: this.tarea.id,
      creadoPor: this.usuarioActual.Id
    };

    this.subTareaService.crear(request).subscribe({
      next: subtarea => {
        this.subtareas = [
          ...(this.subtareas ?? []),
          subtarea
        ];

        this.mostrarNuevaSubTarea.set(false);
        this.tituloNuevaSubTarea = '';
      },
      error: error => {
        console.error('Error al crear subtarea:', error);
      }
    });
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
