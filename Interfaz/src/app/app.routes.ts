import { Routes } from '@angular/router';
import { TestApi } from './pages/test-api/test-api';
import { Bitacora } from './pages/bitacora/bitacora';

export const routes: Routes = [
    {
        path: 'test-api',
        component: TestApi
    },
    {
        path: 'tareas',
        loadComponent: () =>
            import('./pages/tareas/tareas')
            .then(m => m.Tareas)
    },
    {
        path: '',
        redirectTo: 'tareas',
        pathMatch: 'full'
    },
    {
        path: 'bitacora',
        component: Bitacora
    }
];
