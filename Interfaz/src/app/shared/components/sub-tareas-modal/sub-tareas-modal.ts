import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Tarea } from '../../../models/tarea';
import { SubTarea } from '../../../models/subtarea';

@Component({
  selector: 'app-sub-tareas-modal',
  imports: [],
  templateUrl: './sub-tareas-modal.html',
  styleUrl: './sub-tareas-modal.css',
})
export class SubTareasModal {
  @Input() tarea: Tarea | null = null;
  @Input() subtareas: SubTarea[] | null = [];
  @Output() cerrar = new EventEmitter<void>();

  cerrarModal(): void {
    this.cerrar.emit();
  }
  
}
