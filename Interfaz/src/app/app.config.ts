import { ApplicationConfig, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideHttpClient, withInterceptors } from '@angular/common/http';

import { routes } from './app.routes';
import { conexionInterceptor } from './core/interceptors/conexion-interceptor';

export const appConfig: ApplicationConfig = {
  providers: [
    provideHttpClient(
      withInterceptors([
        conexionInterceptor
      ])
    ),
    provideBrowserGlobalErrorListeners(),
    provideRouter(routes)
  ]
};
