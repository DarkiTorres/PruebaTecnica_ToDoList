import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NuevaTareaModal } from './nueva-tarea-modal';

describe('NuevaTareaModal', () => {
  let component: NuevaTareaModal;
  let fixture: ComponentFixture<NuevaTareaModal>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NuevaTareaModal],
    }).compileComponents();

    fixture = TestBed.createComponent(NuevaTareaModal);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
