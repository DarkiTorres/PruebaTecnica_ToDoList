import { Component, inject, signal } from '@angular/core';
import { TareaService } from '../../core/services/tarea.service';
import { UsuarioStateService } from '../../core/services/usuario-state';
import { Tarea } from '../../models/tarea';
import { Modal } from '../../shared/components/modal/modal';

@Component({
  selector: 'app-bitacora',
  imports: [Modal],
  templateUrl: './bitacora.html',
  styleUrl: './bitacora.css',
})
export class Bitacora {
  private readonly tareaService = inject(TareaService);
  private readonly usuarioState = inject(UsuarioStateService);

  tareas = signal<Tarea[]>([]);

  filtro = signal<'todas' | 'terminadas' | 'eliminadas'>(
    'todas'
  );

  readonly usuarioActual = this.usuarioState.usuarioActual;

  mostrarModal = false;
  tituloModal = '';
  mensajeModal = '';
  tareaAEliminar: Tarea | null = null;

  constructor() {
    this.cargarBitacora();
  }

  private cargarBitacora(): void {

    this.tareaService.obtenerBitacora().subscribe({
      next: tareas => {
        this.tareas.set(tareas);
      },

      error: error => {
        console.error(
          'Error al obtener bitácora:',
          error
        );
      }
    });
  }

  cambiarFiltro(
    filtro: 'todas' | 'terminadas' | 'eliminadas'
  ): void {
    this.filtro.set(filtro);
  }

  obtenerTareas(): Tarea[] {

  const filtroActual = this.filtro();

  switch (filtroActual) {

    case 'terminadas':
      return this.tareas().filter(
        tarea => tarea.estaTerminada
      );

    case 'eliminadas':
      return this.tareas().filter(
        tarea =>
          tarea.estaEliminada &&
          !tarea.estaTerminada
      );

    default:
      return this.tareas();
  }
}

  confirmarEliminarFisicamente(tarea: Tarea): void {
    this.tareaAEliminar = tarea;

    this.tituloModal = 'Eliminar definitivamente';
    this.mensajeModal =
      'Esta acción eliminará la tarea permanentemente. ¿Deseas continuar?';

    this.mostrarModal = true;
  }

  eliminarFisicamente(): void {
    if (!this.tareaAEliminar) {
      return;
    }

    const id = this.tareaAEliminar.id;

    this.tareaService.eliminarFisicamente(id).subscribe({
      next: () => {
        this.tareas.update(
          tareas =>
            tareas.filter(t => t.id !== id)
        );

        this.cerrarModal();
      },
      error: error => {
        console.error(
          'Error al eliminar físicamente la tarea:',
          error
        );

        this.cerrarModal();
      }
    });
  }

  cerrarModal(): void {
    this.mostrarModal = false;
    this.tareaAEliminar = null;
  }

  restaurarTarea(tarea: Tarea): void {
    const usuario = this.usuarioActual();

    if (!usuario) {
      return;
    }

    this.tareaService.restaurar(
      tarea.id,
      usuario.Id
    ).subscribe({

      next: () => {

        this.tareas.update(
          tareas =>
            tareas.filter(t => t.id !== tarea.id)
        );

      },

      error: error => {
        console.error(
          'Error al restaurar tarea:',
          error
        );
      }

    });
  }
}
