import { Routes } from '@angular/router';
import { TestApi } from './pages/test-api/test-api';

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
    }
];
