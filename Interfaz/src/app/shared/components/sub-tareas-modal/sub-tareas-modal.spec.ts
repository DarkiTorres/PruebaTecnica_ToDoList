import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubTareasModal } from './sub-tareas-modal';

describe('SubTareasModal', () => {
  let component: SubTareasModal;
  let fixture: ComponentFixture<SubTareasModal>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SubTareasModal],
    }).compileComponents();

    fixture = TestBed.createComponent(SubTareasModal);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
