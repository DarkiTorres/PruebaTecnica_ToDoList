import { inject } from "@angular/core";
import { HttpClient } from "@angular/common/http";

export abstract class ApiService {
    protected readonly http = inject(HttpClient);
    protected readonly apiUrl = 'http://localhost:8080';
}