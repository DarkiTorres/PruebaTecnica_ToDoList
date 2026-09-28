import { Injectable, signal } from '@angular/core';
import { Usuario } from '../../models/usuario';

@Injectable({
  providedIn: 'root',
})
export class UsuarioStateService {
  private readonly usuarioActualSignal = signal<Usuario | null>(null);

  readonly usuarioActual = this.usuarioActualSignal.asReadonly();

  establecerUsuario(usuario: Usuario): void {
    this.usuarioActualSignal.set(usuario);
  }
}
