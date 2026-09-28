import { Component, inject, signal } from '@angular/core';
import { TareaService } from '../../core/services/tarea.service';
import { Tarea } from '../../models/tarea';

@Component({
  selector: 'app-test-api',
  imports: [],
  templateUrl: './test-api.html',
  styleUrl: './test-api.css',
})
export class TestApi {
  private readonly tareaService = inject(TareaService);

  tareas = signal<Tarea[]>([]);

  ngOnInit(): void {
    this.tareaService.obtenerTodas().subscribe({
      next: tareas => {
        console.log('Tarea recibidos: ', tareas);
        this.tareas.set(tareas);
      }, 
      error: error => {
        console.log('Error al consultar tareas:', error);
      }
    })
  }

}
