import { Component, inject } from '@angular/core';
import { UsuarioService } from '../../core/services/usuario.service';
import { Usuario } from '../../models/usuario';
import { UsuarioStateService } from '../../core/services/usuario-state';
import { Router, RouterLink } from '@angular/router';
import { Cache } from '../../core/services/cache';
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
  private readonly cacheService = inject(Cache);

  usuarios: Usuario[] = [];

  menuAbierto = false;

  readonly usuarioActual = this.usuarioState.usuarioActual;

  ngOnInit(): void {
    this.usuarioService.obtenerTodos().subscribe({
      next: usuarios => {
        this.usuarios = usuarios;

        this.cacheService.guardarUsuarios(usuarios);

        const usuarioGuardado = this.cacheService.obtenerUsuarioActual();

        const usuarioInicial = usuarios.find(usuario => usuario.Id === usuarioGuardado?.Id) ?? usuarios[0];

        if (usuarioInicial) {
          this.usuarioState.establecerUsuario(usuarioInicial);
          this.cacheService.guardarUsuarioActual(usuarioInicial);
        }

        // if (usuarios.length > 0) {
        //   this.usuarioState.establecerUsuario(usuarios[0]);
        // }
      },
      error: error => {
        console.error('Error al obtener usuarios:', error);

        const usuariosCache = this.cacheService.obtenerUsuarios();

        this.usuarios = usuariosCache;

        const usuarioCache = this.cacheService.obtenerUsuarioActual();

        if (usuarioCache) {
          this.usuarioState.establecerUsuario(usuarioCache);
        } else if (usuariosCache.length > 0) {
          this.usuarioState.establecerUsuario(usuariosCache[0]);
        }
      }
    });
  }

  cambiarMenu(): void {
    this.menuAbierto = !this.menuAbierto;
  }

  seleccionarUsuario(usuario: Usuario): void {
    this.usuarioState.establecerUsuario(usuario);

    this.cacheService.guardarUsuarioActual(usuario);

    this.menuAbierto = false;
  }

  esBitacora(): boolean {
    return this.router.url === '/bitacora';
  }

}
