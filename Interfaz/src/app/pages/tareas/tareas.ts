import { Component, effect, inject, signal } from '@angular/core';
import { TareaService } from '../../core/services/tarea.service';
import { UsuarioStateService } from '../../core/services/usuario-state';
import { Tarea } from '../../models/tarea';

@Component({
  selector: 'app-tareas',
  imports: [],
  templateUrl: './tareas.html',
  styleUrl: './tareas.css',
})
export class Tareas {
  private readonly tareaService = inject(TareaService);
  private readonly usuarioState = inject(UsuarioStateService);

  tareas = signal<Tarea[]>([]);

  readonly usuarioActual = this.usuarioState.usuarioActual;

  constructor() {
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

  // ngOnInit(): void {
  //   this.cargarTareas();
  // }

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
}
