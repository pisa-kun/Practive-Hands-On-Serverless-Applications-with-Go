import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewIdolComponent } from './new-idol.component';

describe('NewIdolComponent', () => {
  let component: NewIdolComponent;
  let fixture: ComponentFixture<NewIdolComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewIdolComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewIdolComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
