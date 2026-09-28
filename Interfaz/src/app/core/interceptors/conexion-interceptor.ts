import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { Conexion } from '../services/conexion';
import { catchError, tap, throwError } from 'rxjs';

export const conexionInterceptor: HttpInterceptorFn = (req, next) => {

  const conexion = inject(Conexion)

  return next(req).pipe(
    // tap(() => {
    //   conexion.marcarConectada();
    // }),
    catchError(error => {
      if (error.status === 0) {
        conexion.marcarDesconectada();
      }

      return throwError(() => error);
    })
  )
};
