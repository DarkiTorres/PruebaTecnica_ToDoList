import { Component, effect, inject, signal } from '@angular/core';
import { TareaService } from '../../core/services/tarea.service';
import { UsuarioStateService } from '../../core/services/usuario-state';
import { Tarea } from '../../models/tarea';
import { PrioridadService } from '../../core/services/prioridad.service';
import { Prioridad } from '../../models/prioridad';
import { Modal } from '../../shared/components/modal/modal';
import { SubTareaService } from '../../core/services/subtarea.service';
import { SubTareasModal } from '../../shared/components/sub-tareas-modal/sub-tareas-modal';
import { SubTarea } from '../../models/subtarea';
import { NuevaTareaModal } from '../../shared/components/nueva-tarea-modal/nueva-tarea-modal';
import { CrearTareaForm } from '../../models/crear_tarea_form';
import { UsuarioService } from '../../core/services/usuario.service';
import { Usuario } from '../../models/usuario';
import { ActualizarTareaRequest } from '../../models/actualizar_tarea_request';

@Component({
  selector: 'app-tareas',
  imports: [Modal, SubTareasModal, NuevaTareaModal],
  templateUrl: './tareas.html',
  styleUrl: './tareas.css',
})
export class Tareas {
  private readonly tareaService = inject(TareaService);
  private readonly usuarioState = inject(UsuarioStateService);
  private readonly prioridadService = inject(PrioridadService);
  private readonly subTareaService = inject(SubTareaService);
  private readonly usuarioService = inject(UsuarioService);

  tareas = signal<Tarea[]>([]);
  usuarios = signal<Usuario[]>([]);
  prioridades = signal<Prioridad[]>([])

  paginaActual = signal(1);
  elemmentosPorPagina = 5;

  mostrarModal = false;
  tituloModal = '';
  mensajeModal = '';

  mostrarSubTareas = signal(false);

  tareaSeleccionada = signal<Tarea | null>(null);

  subTareasSeleccionadas = signal<SubTarea[]>([]);

  mostrarNuevaTarea = false;
  modoEdicion = false;

  readonly usuarioActual = this.usuarioState.usuarioActual;

  constructor() {

    this.cargarPrioridades();
    this.cargarUsuarios();

    effect(() => {
      const usuario = this.usuarioActual();

      console.log('Usuario cambio:', usuario);

      if (!usuario) {
        this.tareas.set([]);
        return;
      }

      this.cargarTareas(usuario.Id, usuario.RolId);
    })
  }

  get paginas(): number[] {
    return Array.from(
      { length: this.obtenerTotalPaginas() },
      (_, i) => i + 1
    );
  }

  verSubTareas(tarea: Tarea): void {

    console.log('Click en subtareas:', tarea.id);

    this.subTareaService
      .obtenerPorTarea(tarea.id)
      .subscribe({
        next: subtareas => {
          this.tareaSeleccionada.set(tarea)
          this.subTareasSeleccionadas.set(subtareas ?? []);

          this.mostrarSubTareas.set(true);
        },
        error: error => {
          console.error('Error al obtener subtareas:', error);
        }
      });
  }

  crearTarea(form: CrearTareaForm): void {
    const usuario = this.usuarioActual();
    
    if (!usuario) {
      return;
    }

    const request: CrearTareaRequest = {
      titulo: form.titulo,
      descripcion: form.descripcion,
      prioridadId: form.prioridadId,
      fechaEntrega: form.fechaEntrega,
      creadoPor: usuario.Id,
      asignadoA: form.asignadoA
    };


    this.tareaService.crear(request).subscribe({
      next: tarea => {
        this.tareas.update(
          tareas => [...tareas, tarea]
        );
        this.cerrarNuevaTarea();
      },
      error: error => {
        console.error('Error al crear tarea:', error);
      }
    });
  }

  guardarTarea(form: CrearTareaForm): void {
    if (this.modoEdicion) {
      this.actualizarTarea(form);
      return;
    }

    this.crearTarea(form);
  }

  actualizarTarea(form: CrearTareaForm): void {
    const tarea = this.tareaSeleccionada();
    const usuario = this.usuarioActual();

    if (!tarea || !usuario) {
      return;
    }

    const request: ActualizarTareaRequest = {
      titulo: form.titulo,
      descripcion: form.descripcion,
      prioridadId: form.prioridadId,
      fechaEntrega: form.fechaEntrega,
      modificadoPor: usuario.Id,
      asignadoA: form.asignadoA,
      subTareas: []
    };

    this.tareaService.actualizar(
      tarea.id,
      request
    ).subscribe({
      next: tareaActualizada => {

        if (tareaActualizada.estaEliminada) {
          this.tareas.update(
            tareas =>
              tareas.filter(t => t.id !== tareaActualizada.id)
          );

          const totalPaginas = this.obtenerTotalPaginas();

          if (this.paginaActual() > totalPaginas) {
            this.paginaActual.set(totalPaginas);
          }

          return;
        }

        this.tareas.update(
          tareas => 
            tareas.map(t =>
              t.id === tareaActualizada.id
              ? tareaActualizada 
              : t
            )
        );

        this.cerrarNuevaTarea();
      }, 
      error: error => {
        console.error('Error al actualizar tarea:', error);
      }
    });

  }

  //#region Paginacion
  obtenerTareasPagina(): Tarea[] {
    const inicio = (this.paginaActual() - 1) * this.elemmentosPorPagina;
    const fin = inicio + this.elemmentosPorPagina;

    return this.tareas().slice(inicio, fin);
  }

  obtenerTotalPaginas(): number {
    return Math.ceil(
      this.tareas().length / this.elemmentosPorPagina
    );
  }

  cambiarPagina(pagina: number): void {
    if (pagina < 1 || pagina > this.obtenerTotalPaginas()) {
      return;
    }
    this.paginaActual.set(pagina);
  }
  //#endregion

  eliminarTarea(id: number): void {

    const usuario = this.usuarioActual();

    if (!usuario){
      return;
    }

    this.tareaService.eliminar(id, usuario.Id).subscribe({
      next: () => {
        console.log('Tarea eliminada:', id);

        this.tareas.update(
          tareas => tareas.filter(tarea => tarea.id !== id)
        );

        const totalPaginas = this.obtenerTotalPaginas();

        if (this.paginaActual() > totalPaginas) {
          this.paginaActual.set(totalPaginas);
        }
      },
      error: error => {
        console.error('Error al eliminar tarea:', error);
      }
    });
  }

  cambiarEstadoTarea(tarea: Tarea, event: Event): void {

    const checkbox = event.target as HTMLInputElement;

    const estadoAnterior = tarea.estaTerminada;
    const nuevoEstado = !estadoAnterior;

    this.tareaService.completar(
      tarea.id,
      nuevoEstado
    ).subscribe({

      next: tareaActualizada => {

        // La tarea fue soft deleteada por el backend
        if (tareaActualizada.estaEliminada) {

          this.tareas.update(
            tareas =>
              tareas.filter(t => t.id !== tareaActualizada.id)
          );

          const totalPaginas = this.obtenerTotalPaginas();

          if (this.paginaActual() > totalPaginas) {
            this.paginaActual.set(
              Math.max(1, totalPaginas)
            );
          }

          return;
        }

        // La tarea sigue activa
        this.tareas.update(
          tareas =>
            tareas.map(t =>
              t.id === tareaActualizada.id
                ? tareaActualizada
                : t
            )
        );
      },

      error: error => {

        console.error(
          'Error al cambiar estado:',
          error
        );

        checkbox.checked = estadoAnterior;

        this.tareas.update(
          tareas =>
            tareas.map(t =>
              t.id === tarea.id
                ? {
                    ...t,
                    estaTerminada: estadoAnterior
                  }
                : t
            )
        );

        if (error.status === 400) {
          this.abrirModal(
            'Tarea no completada.',
            'La tarea tiene subtareas pendientes.'
          );
        }
      }

    });
  }

  marcarTareaPendiente(): void {
    const tarea = this.tareaSeleccionada();

    if (!tarea) {
      return;
    }

    this.tareas.update(
      tareas =>
        tareas.map(t =>
          t.id === tarea.id
            ? {
                ...t,
                estaTerminada: false
              }
            : t
        )
    );

    this.tareaSeleccionada.update(
      tareaActual =>
        tareaActual
          ? {
              ...tareaActual,
              estaTerminada: false
            }
          : tareaActual
    );
  }

  private cargarTareas(usuarioId: number, rolId: number): void {
    if (rolId === 1) {
      console.log('Líder → GET /tareas');

      this.tareaService.obtenerTodas().subscribe({
        next: tareas => {

          console.log('Tareas del líder:', tareas);

          this.tareas.set(tareas);

        },
        error: error => {
          console.error('Error al obtener tareas:', error);
        }
      });


    } else {
      console.log(`Colaborador → GET /usuarios/${usuarioId}/tareas`);

      this.tareaService.obtenerPorUsuario(usuarioId).subscribe({
        next: tareas => {

          console.log('Tareas del colaborador:', tareas);

          this.tareas.set(tareas);

        },
        error: error => {
          console.error(
            'Error al obtener tareas del usuario:',
            error
          );
        }
      });
    }
  }

  private cargarUsuarios(): void {
    this.usuarioService.obtenerTodos().subscribe(
      {
        next: usuarios => {
          this.usuarios.set(usuarios);
        },
        error: error => {
          console.error('Error al obtener usuarios:', error);
        }
      }
    );
  }

  //#region Prioridades
  obtenerDescripcionPrioridad(id: number): string {
    const prioridad = this.prioridades().find(
      prioridad => prioridad?.Id === id
    );

    return prioridad?.Descripcion ?? 'Sin prioridad';
  }

  obtenerClasePrioridad(id: number): string {
    const prioridad = this.prioridades().find(
      prioridad => prioridad.Id === id
    );

    if (!prioridad) {
      return 'otro';
    }

    switch (prioridad.Descripcion.toLowerCase()) {
      case 'urgente':
        return 'urgente';

      case 'alto':
        return 'alta';

      case 'medio':
        return 'media';

      case 'bajo':
        return 'baja';

      default:
        return 'otro';
    }
  }

  private cargarPrioridades(): void {
    this.prioridadService.obtenerTodas().subscribe({
      next: prioridades => {
        console.log('Prioridades recibidas:', prioridades);
        this.prioridades.set(prioridades)     
      },
      error: error => {
        console.error('Error al obtener prioridades:', error);
      }
    })
  }
  //#endregion

  //#region Modal
  abrirModal(titulo: string, mensaje: string): void {
    this.tituloModal = titulo;
    this.mensajeModal = mensaje;
    this.mostrarModal = true;
  }

  cerrarModal() : void {
    this.mostrarModal = false;
  }

  cerrarSubTareas() : void {
    this.mostrarSubTareas.set(false);
    this.tareaSeleccionada.set(null);
    this.subTareasSeleccionadas.set([]);
  }

  abrirNuevaTarea(): void {
    this.modoEdicion = false;
    this.tareaSeleccionada.set(null);
    this.mostrarNuevaTarea = true;
  }

  editarTarea(tarea: Tarea): void {
    this.modoEdicion = true;
    this.tareaSeleccionada.set(tarea);
    this.mostrarNuevaTarea = true;
  }

  cerrarNuevaTarea(): void {
    this.modoEdicion = false;
    this.mostrarNuevaTarea = false;
    this.tareaSeleccionada.set(null);
  }
  //#endregion
}
