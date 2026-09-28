import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-modal',
  imports: [],
  templateUrl: './modal.html',
  styleUrl: './modal.css',
})
export class Modal {
  @Input() titulo = '';
  @Input() mensaje = '';

  @Input() esConfirmacion = false;

  @Output() cerrar = new EventEmitter<void>();
  @Output() confirmar = new EventEmitter<void>();

  cerrarModal(): void {
    this.cerrar.emit();
  }

  confirmarModal(): void {
    this.confirmar.emit();
  }
}
