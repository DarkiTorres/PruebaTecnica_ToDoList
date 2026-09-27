import { Component, inject } from '@angular/core';
import { UsuarioService } from '../../core/services/usuario.service';
import { Usuario } from '../../models/usuario';

@Component({
  selector: 'app-header',
  imports: [],
  templateUrl: './header.html',
  styleUrl: './header.css',
})
export class Header {
  private readonly usuarioService = inject(UsuarioService);

  usuarios: Usuario[] = [];
  usuarioActual: Usuario | null = null;

  menuAbierto = false;

  ngOnInit(): void {
    this.usuarioService.obtenerTodos().subscribe({
      next: usuarios => {
        this.usuarios = usuarios;

        if (usuarios.length > 0) {
          this.usuarioActual = usuarios[0];
        }
      },
      error: error => {
        console.error('Error al obtener usuarios:', error);
      }
    });
  }

  cambiarMenu(): void {
    this.menuAbierto = !this.menuAbierto;
  }

  seleccionarUsuario(usuario: Usuario): void {
    this.usuarioActual = usuario;
    this.menuAbierto = false;
  }

}
