import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Prioridad } from '../../../models/prioridad';
import { Usuario } from '../../../models/usuario';
import { CrearTareaForm } from '../../../models/crear_tarea_form';
import { Tarea } from '../../../models/tarea';

@Component({
  selector: 'app-nueva-tarea-modal',
  imports: [FormsModule],
  templateUrl: './nueva-tarea-modal.html',
  styleUrl: './nueva-tarea-modal.css',
})
export class NuevaTareaModal {
  @Input() prioridades: Prioridad[] = [];
  @Input() usuarios: Usuario[] = [];
  @Input() usuarioActual: Usuario | null = null;
  @Input() tareaEditar: Tarea | null = null;
  @Input() modoEdicion = false;

  @Output() cerrar = new EventEmitter<void>();
  @Output() crear = new EventEmitter<CrearTareaForm>();

  titulo = '';
  descripcion = '';
  prioridadId: number | null = null;
  asignadoA: number | null = null;
  fechaEntrega = '';

  ngOnChanges(): void {

    if (this.modoEdicion && this.tareaEditar) {

      console.log('Tarea para editar:', this.tareaEditar);
      console.log('AsignadoAId:', this.tareaEditar.asignadoAId);

      this.titulo = this.tareaEditar.titulo;

      this.descripcion = this.tareaEditar.descripcion ?? '';

      this.prioridadId = this.tareaEditar.prioridadId;

      this.asignadoA = this.tareaEditar.asignadoAId ?? null;

      this.fechaEntrega = this.formatearFecha(
        this.tareaEditar.fechaEntrega
      );

      return;
    }

    if (this.usuarioActual?.RolId === 2) {
      this.asignadoA = this.usuarioActual.Id;
    }
  }

  crearTarea(): void {
    if (this.prioridadId === null || this.asignadoA === null) {
      return;
    }

    this.crear.emit({
      titulo: this.titulo.trim(),
      descripcion: this.descripcion.trim() || null,
      prioridadId: this.prioridadId,
      fechaEntrega: this.fechaEntrega
        ? new Date(this.fechaEntrega).toISOString()
        : null,
      asignadoA: this.asignadoA
    });

  }

  cerrarModal(): void {
    this.cerrar.emit();
  }

  private formatearFecha(fecha: string | null | undefined): string {

    if (!fecha) {
      return '';
    }

    const date = new Date(fecha);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day}T${hours}:${minutes}`;
  }
}
