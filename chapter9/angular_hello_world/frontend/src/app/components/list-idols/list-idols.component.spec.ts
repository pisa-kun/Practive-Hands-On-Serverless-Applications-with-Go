import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListIdolsComponent } from './list-idols.component';

describe('ListIdolsComponent', () => {
  let component: ListIdolsComponent;
  let fixture: ComponentFixture<ListIdolsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListIdolsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ListIdolsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
