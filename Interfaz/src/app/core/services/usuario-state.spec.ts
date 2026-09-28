import { TestBed } from '@angular/core/testing';

import { UsuarioState } from './usuario-state';

describe('UsuarioState', () => {
  let service: UsuarioState;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UsuarioState);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
