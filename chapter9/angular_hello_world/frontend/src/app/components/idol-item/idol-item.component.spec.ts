import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IdolItemComponent } from './idol-item.component';

describe('IdolItemComponent', () => {
  let component: IdolItemComponent;
  let fixture: ComponentFixture<IdolItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IdolItemComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(IdolItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
