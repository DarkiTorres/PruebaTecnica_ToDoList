import { HttpBackend, HttpClient } from '@angular/common/http';
import { inject, Injectable, signal } from '@angular/core';
import { catchError, distinctUntilChanged, exhaustMap, map, of, timer } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class Conexion {
  private readonly http = new HttpClient(inject(HttpBackend));

  private readonly apiUrl = 'http://localhost:8080';

  readonly conectada = signal(false);

  constructor() {

    timer(0, 15000)
      .pipe(
        exhaustMap(() =>
          this.http
            .get(`${this.apiUrl}/health`)
            .pipe(
              map(() => true),
              catchError(() => of(false))
            )
        ),
        distinctUntilChanged()
      )
      .subscribe(conectada => {
        this.conectada.set(conectada);
      });
  }

  marcarDesconectada(): void {
    this.conectada.set(false);
  }

  marcarConectada(): void {
    this.conectada.set(true);
  }

}
