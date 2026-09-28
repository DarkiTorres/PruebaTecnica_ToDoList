import { Component, inject } from '@angular/core';
import { UsuarioService } from '../../core/services/usuario.service';
import { Usuario } from '../../models/usuario';
import { UsuarioStateService } from '../../core/services/usuario-state';
import { Router, RouterLink } from '@angular/router';

@Component({
  selector: 'app-header',
  imports: [RouterLink],
  templateUrl: './header.html',
  styleUrl: './header.css',
})
export class Header {
  private readonly usuarioService = inject(UsuarioService);
  private readonly usuarioState = inject(UsuarioStateService);
  private readonly router = inject(Router);

  usuarios: Usuario[] = [];

  menuAbierto = false;

  readonly usuarioActual = this.usuarioState.usuarioActual;

  ngOnInit(): void {
    this.usuarioService.obtenerTodos().subscribe({
      next: usuarios => {
        this.usuarios = usuarios;

        if (usuarios.length > 0) {
          this.usuarioState.establecerUsuario(usuarios[0]);
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
    this.usuarioState.establecerUsuario(usuario);
    this.menuAbierto = false;
  }

  esBitacora(): boolean {
    return this.router.url === '/bitacora';
  }

}
