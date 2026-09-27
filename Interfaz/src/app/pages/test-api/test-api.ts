import { Component, inject } from '@angular/core';
import { UsuarioService } from '../../core/services/usuario.service';
import { Usuario } from '../../models/usuario';

@Component({
  selector: 'app-test-api',
  imports: [],
  templateUrl: './test-api.html',
  styleUrl: './test-api.css',
})
export class TestApi {
  private readonly usuarioService = inject(UsuarioService);

  usuarios: Usuario[] = [];

  ngOnInit(): void {
    this.usuarioService.obtenerTodos().subscribe({
      next: usuarios => {
        console.log('Usuarios recibidos: ', usuarios);
        this.usuarios = usuarios;
      }, 
      error: error => {
        console.log('Error al consultar usuarios:', error);
      }
    })
  }

}
